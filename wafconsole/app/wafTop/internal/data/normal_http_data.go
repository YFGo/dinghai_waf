package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/biz/normalhttp"
	"wafconsole/app/wafTop/internal/data/model"
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

// CreateNormalHttpBatch 批量创建正常http请求日志
func (n normalHttpRepo) CreateNormalHttpBatch(ctx context.Context,
	normalHttpInfo []model.NormalHttpModel) error {
	err := n.data.clickHouse.Table(model.NormalHttpTableName).Create(&normalHttpInfo).Error
	if err != nil {
		n.log.WithContext(ctx).Errorf("create normal http info failed: %v", err)
		return err
	}
	return nil
}
