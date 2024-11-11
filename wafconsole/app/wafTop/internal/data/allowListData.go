package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/allow"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type allowlistRepo struct {
	data *Data
}

func NewAllowListRepo(data *Data) allow.ListAllowRepo {
	return &allowlistRepo{
		data: data,
	}
}

func (a allowlistRepo) Get(ctx context.Context, i int64) (model.AllowList, error) {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) GetByNameAndID(ctx context.Context, s string, i int64) (model.AllowList, error) {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) Create(ctx context.Context, t model.AllowList) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) Update(ctx context.Context, i int64, t model.AllowList) error {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a allowlistRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.AllowList, error) {
	//TODO implement me
	panic("implement me")
}
