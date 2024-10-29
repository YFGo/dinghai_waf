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

func (g appWafRepo) Get(ctx context.Context, id int64) (model.AppWaf, error) {
	var appWaf model.AppWaf
	err := g.data.db.Where("id = ?", id).First(&appWaf).Error
	return appWaf, err
}

func (g appWafRepo) GetByNameAndID(ctx context.Context, name string, id int64) (model.AppWaf, error) {
	var appWaf model.AppWaf
	if id == 0 {
		err := g.data.db.Where("name = ?", name).First(&appWaf).Error
		return appWaf, err
	}
	err := g.data.db.Where("name = ? and id != ?", name, id).First(&appWaf).Error
	return appWaf, err
}

func (g appWafRepo) Create(ctx context.Context, appInfo model.AppWaf) (int64, error) {
	err := g.data.db.Create(&appInfo).Error
	return int64(appInfo.ID), err
}

func (g appWafRepo) Update(ctx context.Context, id int64, appInfo model.AppWaf) error {
	err := g.data.db.Where("id = ?", id).Updates(&appInfo).Error
	return err
}

func (g appWafRepo) Delete(ctx context.Context, i []int64) (int64, error) {
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

func (g appWafRepo) GetAppWafByServerId(ctx context.Context, serverId int64) (appInfo model.AppWaf, err error) {
	err = g.data.db.Where("server_id = ?", serverId).First(&appInfo).Error
	return
}
