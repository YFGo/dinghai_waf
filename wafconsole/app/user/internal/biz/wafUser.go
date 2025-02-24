package biz

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/model"
	up "wafconsole/utils/plugin"
)

type WafUserRepo interface {
	iface.BaseRepo[model.UserInfo]
	LoginByEmailPassword(ctx context.Context, user model.UserInfo) (model.UserInfo, error)
}

type WafUserUsecase struct {
	repo WafUserRepo
}

func NewWafUserUsecase(repo WafUserRepo) *WafUserUsecase {
	return &WafUserUsecase{repo: repo}
}

// SignUp 用户注册
func (w *WafUserUsecase) SignUp(ctx context.Context, user model.UserInfo) error {
	_, err := w.repo.Create(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "SignUp error : ", err)
		return err
	}
	return nil
}

// LoginByEmailPassword 用户登录
func (w *WafUserUsecase) LoginByEmailPassword(ctx context.Context, user model.UserInfo) (string, string, error) {
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
func (w *WafUserUsecase) GetUserInfoByID(ctx context.Context, id int64) (model.UserInfo, error) {
	userInfo, err := w.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserInfo{}, status.Error(codes.NotFound, "user not found")
		}
		slog.ErrorContext(ctx, "GetUserInfo error : ", err)
		return model.UserInfo{}, err
	}
	return userInfo, nil
}
