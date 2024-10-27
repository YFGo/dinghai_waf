package biz

import (
	"context"
	"log/slog"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/model"
)

type WafUserRepo interface {
	iface.BaseRepo[model.UserInfo]
	LoginByEmailPassword(ctx context.Context, user model.UserInfo) error
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
func (w *WafUserUsecase) LoginByEmailPassword(ctx context.Context, user model.UserInfo) error {
	err := w.repo.LoginByEmailPassword(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "LoginByEmailPassword error : ", err)
		return err
	}
	return nil
}
