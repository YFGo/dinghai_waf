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
	"wafconsole/app/user/internal/server/plugin"
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
	userclaims := plugin.UserClaims{
		UserId:   uint64(userInfo.ID),
		Username: userInfo.UserName,
	}
	jwtUtils := plugin.InitNewJWTUtils()
	accessToken, refreshToken, _ := jwtUtils.GetToken(userclaims)
	return accessToken, refreshToken, nil
}
