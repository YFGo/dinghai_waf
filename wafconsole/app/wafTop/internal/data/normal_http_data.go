package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"

	"wafconsole/app/wafTop/internal/biz/normalhttp"
)

type normalHttpRepo struct {
	data *Data
	log  *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	return &normalHttpRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (n normalHttpRepo) Get(ctx context.Context, i int64) (model.NormalHttpModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) GetByNameAndID(ctx context.Context, s string, i int64) (model.NormalHttpModel, error) {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) Create(ctx context.Context, t model.NormalHttpModel) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) Update(ctx context.Context, i int64, t model.NormalHttpModel) error {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (n normalHttpRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.NormalHttpModel, error) {
	//TODO implement me
	panic("implement me")
}
