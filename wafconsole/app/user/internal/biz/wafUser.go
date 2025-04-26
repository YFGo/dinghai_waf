package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/types"
	up "wafconsole/utils/plugin"
)

type WafUserRepo interface {
	iface.BaseRepo[types.UserInfo]
	LoginByEmailPassword(ctx context.Context, user types.UserInfo) (types.UserInfo, error)
	GetRedisValueByKey(ctx context.Context, key string) (string, error)
}

type WafUserUsecase struct {
	repo WafUserRepo
	log  *log.Helper
}

func NewWafUserUsecase(repo WafUserRepo, logger log.Logger) *WafUserUsecase {
	return &WafUserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// SignUp 用户注册
func (w *WafUserUsecase) SignUp(ctx context.Context, user types.UserInfo, inputCode string) error {
	// 判断验证码是否正确
	systemCode, err := w.repo.GetRedisValueByKey(ctx, user.Email)
	if err != nil {
		w.log.WithContext(ctx).Error("get sign code is error", err)
		return err
	}
	if systemCode != inputCode {
		return status.Error(codes.PermissionDenied, "code is error")
	}
	_, err = w.repo.Create(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "SignUp error : ", err)
		return err
	}
	return nil
}

// LoginByEmailPassword 用户登录
func (w *WafUserUsecase) LoginByEmailPassword(ctx context.Context, user types.UserInfo) (string, string, error) {
	userInfo, err := w.repo.LoginByEmailPassword(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", status.Error(codes.NotFound, "user not found")
		}
		slog.ErrorContext(ctx, "LoginByEmailPassword error : ", err)
		return "", "", status.Error(codes.Internal, err.Error())
	}
	userclaims := up.UserClaims{
		UserId:   uint64(userInfo.ID),
		Username: userInfo.UserName,
	}
	jwtUtils := up.InitNewJWTUtils()
	accessToken, refreshToken, _ := jwtUtils.GetToken(userclaims)
	return accessToken, refreshToken, nil
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
