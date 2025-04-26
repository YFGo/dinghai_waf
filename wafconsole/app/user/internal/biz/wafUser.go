package biz

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log/slog"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/types"
	up "wafconsole/utils/plugin"
)

type WafUserRepo interface {
	iface.BaseRepo[types.UserInfo]
	LoginByEmailPassword(ctx context.Context, user types.UserInfo) (types.UserInfo, error)
	GetRedisValueByKey(ctx context.Context, key string) (string, error)
	GetUserInfoByEmail(ctx context.Context, email string) (types.UserInfo, error)
}

type WafUserUsecase struct {
	repo          WafUserRepo
	wafCommonRepo WafUserCommonRepo
	log           *log.Helper
}

func NewWafUserUsecase(repo WafUserRepo, wafCommonRepo WafUserCommonRepo, logger log.Logger) *WafUserUsecase {
	return &WafUserUsecase{
		repo:          repo,
		wafCommonRepo: wafCommonRepo,
		log:           log.NewHelper(logger),
	}
}

// 校验验证码是否正确
func (w *WafUserUsecase) checkEmailCodeIsRight(ctx context.Context, userEmail, inputCode string) error {
	// 判断验证码是否正确
	systemCode, err := w.repo.GetRedisValueByKey(ctx, userEmail)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return status.Error(codes.NotFound, "code is not exist")
		}
		w.log.WithContext(ctx).Error("get sign code is error", err)
		return err
	}
	if systemCode != inputCode {
		return status.Error(codes.PermissionDenied, "code is error")
	}
	// 删除验证码 , 防止被重复使用
	if err := w.wafCommonRepo.DeleteKvFromRs(ctx, userEmail); err != nil {
		return err
	}
	return nil
}

// SignUp 用户注册
func (w *WafUserUsecase) SignUp(ctx context.Context, user types.UserInfo, inputCode string) error {
	err := w.checkEmailCodeIsRight(ctx, user.Email, inputCode)
	if err != nil {
		return err
	}
	_, err = w.repo.Create(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "SignUp error : ", err)
		return err
	}
	return nil
}

// LoginByEmailPassword 用户登录
func (w *WafUserUsecase) LoginByEmailPassword(ctx context.Context, account, accountCheck string, loginMethod uint8) (string, string, string, uint, error) {
	var (
		userInfo types.UserInfo
		err      error
	)
	switch loginMethod {
	case 1, 3: //邮箱密码 + 电话密码
		// 区分登录信息
		loginInfo := types.UserInfo{}
		if loginMethod == 1 {
			loginInfo.Email = account
		} else {
			loginInfo.Phone = account
		}
		loginInfo.Password = accountCheck
		// 校验登录
		userInfo, err = w.repo.LoginByEmailPassword(ctx, loginInfo)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", "", "", 0, status.Error(codes.NotFound, "user not found")
			}
			w.log.WithContext(ctx).Error("LoginByEmailPassword error : ", err)
			return "", "", "", 0, status.Error(codes.Internal, err.Error())
		}
	case 2: // 邮箱验证码登录
		err := w.checkEmailCodeIsRight(ctx, account, accountCheck)
		if err != nil {
			return "", "", "", 0, err
		}
		// 验证码通过 , 根据邮箱获取用户信息
		userInfo, err = w.repo.GetUserInfoByEmail(ctx, account)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", "", "", 0, status.Error(codes.NotFound, "user not found")
			}
			w.log.WithContext(ctx).Error("GetUserInfoByEmail error : ", err)
			return "", "", "", 0, status.Error(codes.Internal, err.Error())
		}
	}

	userclaims := up.UserClaims{
		UserId:   uint64(userInfo.ID),
		Username: userInfo.UserName,
	}
	jwtUtils := up.InitNewJWTUtils()
	accessToken, refreshToken, _ := jwtUtils.GetToken(userclaims)
	return accessToken, refreshToken, userInfo.AvatarAddr, userInfo.ID, nil
}

// GetUserInfoByID 获取用户信息
func (w *WafUserUsecase) GetUserInfoByID(ctx context.Context, id int64) (types.UserInfo, error) {
	userInfo, err := w.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return types.UserInfo{}, status.Error(codes.NotFound, "user not found")
		}
		slog.ErrorContext(ctx, "GetUserInfo error : ", err)
		return types.UserInfo{}, err
	}
	return userInfo, nil
}
