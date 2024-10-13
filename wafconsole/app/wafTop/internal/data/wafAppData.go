package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type appWafRepo struct {
	data *Data
	log  *log.Helper
}

// NewAppWafRepo .
func NewAppWafRepo(data *Data, logger log.Logger) siteBiz.WafAppRepo {
	return &appWafRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g appWafRepo) Get(ctx context.Context, i int64) (model.AppWaf, error) {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) GetByName(ctx context.Context, s string) (model.AppWaf, error) {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) Create(ctx context.Context, t model.AppWaf) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) Update(ctx context.Context, i int64, t model.AppWaf) error {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) Delete(ctx context.Context, i int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (g appWafRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.AppWaf, error) {
	//TODO implement me
	panic("implement me")
}
