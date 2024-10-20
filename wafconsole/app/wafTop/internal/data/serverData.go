package data

import (
	"context"
	"gorm.io/gorm"
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
	if err != nil {
		return server, err
	}
	return server, nil
}

// GetByNameAndID 根据名称获取服务器信息
func (s serverRepo) GetByNameAndID(ctx context.Context, serverName string, id int64) (model.ServerWaf, error) {
	var server model.ServerWaf
	if id == 0 {
		err := s.data.db.Where("name = ?", serverName).First(&server).Error
		return server, err
	}
	err := s.data.db.Where("name = ? AND id != ?", serverName, id).First(&server).Error
	return server, err
}

// Create 新增服务器
func (s serverRepo) Create(ctx context.Context, serverInfo model.ServerWaf) (int64, error) {
	var serverStrategies []model.ServerStrategies
	err := s.data.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&serverInfo).Error // 插入主表数据
		if err != nil {
			return err
		}
		for _, strategyId := range serverInfo.StrategiesID {
			serverStrategy := model.ServerStrategies{
				ServerID:   int64(serverInfo.ID),
				StrategyId: strategyId,
			}
			serverStrategies = append(serverStrategies, serverStrategy)
		}
		err = tx.Create(&serverStrategies).Error // 插入关联表数据
		if err != nil {
			return err
		}
		return nil
	})
	return int64(serverInfo.ID), err
}

func (s serverRepo) Update(ctx context.Context, id int64, serverInfo model.ServerWaf) error {
	var serverStrategies []model.ServerStrategies
	serverWaf := model.ServerWaf{
		Name: serverInfo.Name,
		Host: serverInfo.Host,
		IP:   serverInfo.IP,
		Port: serverInfo.Port,
	}
	err := s.data.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Updates(&serverWaf).Error
		if err != nil {
			return err
		}
		err = tx.Where("server_id = ?", id).Unscoped().Delete(&model.ServerStrategies{}).Error //删除关联表数据
		if err != nil {
			return err
		}
		for _, strategyId := range serverInfo.StrategiesID {
			serverStrategy := model.ServerStrategies{
				ServerID:   int64(serverInfo.ID),
				StrategyId: strategyId,
			}
			serverStrategies = append(serverStrategies, serverStrategy)
		}
		err = tx.Create(&serverStrategies).Error // 插入关联表数据
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s serverRepo) Delete(ctx context.Context, serverIds []int64) (int64, error) {
	err := s.data.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("server_id in (?)", serverIds).Unscoped().Delete(&model.ServerStrategies{}).Error //删除关联表数据
		if err != nil {
			return err
		}
		err = tx.Where("id in (?)", serverIds).Delete(&model.ServerWaf{}).Error //删除主表数据")
		if err != nil {
			return err
		}
		return nil
	})
	return 0, err
}

func (s serverRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	mysqlDB := s.data.db.Table(model.ServerWafTableName)
	for _, opt := range withReturn {
		opt(mysqlDB)
	}
	var total int64
	err := mysqlDB.Count(&total).Error
	return total, err
}

func (s serverRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.ServerWaf, error) {
	var serverList []model.ServerWaf
	mysqlDB := s.data.db.Table(model.ServerWafTableName)
	for _, opt := range opts {
		opt(mysqlDB)
	}
	err := mysqlDB.Limit(int(limit)).Offset(int(offset)).Find(&serverList).Error
	return serverList, err
}

func (s serverRepo) GetServerStrategiesID(ctx context.Context, id int64) ([]int64, error) {
	var strategyIds []int64
	err := s.data.db.Table(model.ServersStrategiesTableName).Select("strategy_id").Where("server_id = ?", id).Find(&strategyIds).Error
	if err != nil {
		return nil, err
	}
	return strategyIds, nil
}

// SaveServerToEtcd 将web程序应用的策略 , 及真实地址存入etcd
func (s serverRepo) SaveServerToEtcd(ctx context.Context, serverStrategiesKey, serverRealAddrKey, serverStrategies, serverRealAddr string) error {
	_, err := s.data.etcd.KV.Put(ctx, serverStrategiesKey, serverStrategies)
	if err != nil {
		return err
	}
	_, err = s.data.etcd.KV.Put(ctx, serverRealAddrKey, serverRealAddr)
	return err
}
