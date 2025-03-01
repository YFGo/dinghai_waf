package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/cron/internal/biz"
)

type userRpcRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRpcRepo . Create a new user rpc repo.
func NewUserRpcRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRpcRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// ListByPage . List users by page.
func (u *userRpcRepo) ListByPage(ctx context.Context, page int64, pageSize int64) ([]*biz.UserAccountInfo, error) {
	return nil, nil
}

// Count . Count users.
func (u *userRpcRepo) Count(ctx context.Context) (int64, error) {
	return 0, nil
}
