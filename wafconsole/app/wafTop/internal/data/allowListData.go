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

func (a allowlistRepo) Get(ctx context.Context, id int64) (model.AllowList, error) {
	var allowList model.AllowList
	err := a.data.db.Where("id = ?", id).First(&allowList).Error
	return allowList, err
}

func (a allowlistRepo) GetByNameAndID(ctx context.Context, name string, id int64) (model.AllowList, error) {
	var allowList model.AllowList
	if id != 0 {
		err := a.data.db.Where("name = ? AND id != ?", name, id).First(&allowList).Error
		return allowList, err
	}
	err := a.data.db.Where("name = ?", name).First(&allowList).Error
	return allowList, err
}

func (a allowlistRepo) Create(ctx context.Context, allowList model.AllowList) (int64, error) {
	err := a.data.db.Create(&allowList).Error
	return int64(allowList.ID), err
}

func (a allowlistRepo) Update(ctx context.Context, id int64, allowList model.AllowList) error {
	err := a.data.db.Where("id = ?", id).Updates(&allowList).Error
	return err
}

func (a allowlistRepo) Delete(ctx context.Context, ids []int64) (int64, error) {
	affectRows := a.data.db.Where("id IN (?)", ids).Delete(&model.AllowList{}).RowsAffected
	return affectRows, nil
}

func (a allowlistRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	var total int64
	mysqlDB := a.data.db.Table(model.AllowListTableName)
	for _, opt := range withReturn {
		mysqlDB = opt(mysqlDB)
	}
	err := mysqlDB.Count(&total).Error
	return total, err
}

func (a allowlistRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.AllowList, error) {
	var allowLists []model.AllowList
	mysqlDB := a.data.db.Table(model.AllowListTableName)
	for _, opt := range opts {
		mysqlDB = opt(mysqlDB)
	}
	err := mysqlDB.Limit(int(limit)).Offset(int(offset)).Find(&allowLists).Error
	return allowLists, err
}
