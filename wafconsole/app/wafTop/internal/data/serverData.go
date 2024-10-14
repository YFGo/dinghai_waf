package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"
)

type serverRepo struct {
	data *Data
}

func NewServerRepo(data *Data) siteBiz.ServerRepo {
	return &serverRepo{
		data: data,
	}
}

func (s serverRepo) Get(ctx context.Context, i int64) (model.ServerWaf, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) GetByName(ctx context.Context, s2 string) (model.ServerWaf, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) Create(ctx context.Context, t model.ServerWaf) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) Update(ctx context.Context, i int64, t model.ServerWaf) error {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) Delete(ctx context.Context, i int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.ServerWaf, error) {
	//TODO implement me
	panic("implement me")
}
