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

// Get 根据id获取服务器信息
func (s serverRepo) Get(ctx context.Context, id int64) (model.ServerWaf, error) {
	var server model.ServerWaf
	err := s.data.db.Where("id = ?", id).First(&server).Error
	return server, err
}

// GetByNameAndID 根据名称获取服务器信息
func (s serverRepo) GetByNameAndID(ctx context.Context, serverName string, id int64) (model.ServerWaf, error) {
	var server model.ServerWaf
	err := s.data.db.Where("server_name = ?", serverName).First(&server).Error
	return server, err
}

// Create 新增服务器
func (s serverRepo) Create(ctx context.Context, serverInfo model.ServerWaf) (int64, error) {
	err := s.data.db.Create(&serverInfo).Error
	return int64(serverInfo.ID), err
}

func (s serverRepo) Update(ctx context.Context, i int64, t model.ServerWaf) error {
	//TODO implement me
	panic("implement me")
}

func (s serverRepo) Delete(ctx context.Context, i []int64) (int64, error) {
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
