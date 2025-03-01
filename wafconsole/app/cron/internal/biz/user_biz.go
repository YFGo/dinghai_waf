package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserAccountInfo struct {
	Id      int64  // 用户ID
	Account string // 用户账号
	Passwd  string // 用户密码
}

// UserRepo is a Greater repo.
type UserRepo interface {
	ListByPage(ctx context.Context, page int64, pageSize int64) ([]*UserAccountInfo, error)
	Count(ctx context.Context) (int64, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase new a User usecase.
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

// List list users.
func (u *UserUseCase) List(ctx context.Context, page int64, pageSize int64) ([]*UserAccountInfo, error) {
	return u.repo.ListByPage(ctx, page, pageSize)
}

// Count count users.
func (u *UserUseCase) Count(ctx context.Context) (int64, error) {
	return u.repo.Count(ctx)
}
