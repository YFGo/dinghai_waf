package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/api/grpc_user"
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
	userList, err := u.data.guClient.UserList(ctx, &grpc_user.UserListRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	var list []*biz.UserAccountInfo
	for _, v := range userList.List {
		list = append(list, &biz.UserAccountInfo{
			Id:      v.Id,
			Account: v.Account,
			Passwd:  v.Password,
		})
	}
	return list, nil
}

// Count . Count users.
func (u *userRpcRepo) Count(ctx context.Context) (int64, error) {
	rpcResp, err := u.data.guClient.UserList(ctx, &grpc_user.UserListRequest{})
	if err != nil {
		return 0, err
	}
	return rpcResp.Total, nil
}
