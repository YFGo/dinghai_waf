package data

import (
	"context"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/model"
)

type wafUserRepo struct {
	data *Data
}

func NewWafUserRepo(data *Data) biz.WafUserRepo {
	return &wafUserRepo{data: data}
}

func (w wafUserRepo) Get(ctx context.Context, i int64) (model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) GetByNameAndID(ctx context.Context, s string, i int64) (model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Create(ctx context.Context, t model.UserInfo) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Update(ctx context.Context, i int64, t model.UserInfo) error {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}
