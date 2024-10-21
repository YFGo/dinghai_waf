package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	ruleBiz "wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"
)

type userRuleRepo struct {
	data *Data
}

func NewUserRuleRepo(data *Data) ruleBiz.UserRuleRepo {
	return &userRuleRepo{
		data: data,
	}
}

func (u userRuleRepo) Get(ctx context.Context, i int64) (model.UserRule, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRuleRepo) GetByNameAndID(ctx context.Context, name string, id int64) (model.UserRule, error) {
	var userRule model.UserRule
	if id == 0 {
		err := u.data.db.Where("name = ?", name).First(&userRule).Error
		return userRule, err
	}
	err := u.data.db.Where("id != ? AND name = ?", id, name).First(&userRule).Error
	return userRule, err
}

// Create 新增自定义规则
func (u userRuleRepo) Create(ctx context.Context, userRule model.UserRule) (int64, error) {
	err := u.data.db.Create(&userRule).Error
	return int64(userRule.ID), err
}

func (u userRuleRepo) Update(ctx context.Context, id int64, userRule model.UserRule) error {
	err := u.data.db.Where("id = ?", id).Updates(&userRule).Error
	return err
}

func (u userRuleRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	affectRow := u.data.db.Where("id in (?)", int64s).Delete(&model.UserRule{}).RowsAffected
	return affectRow, nil
}

func (u userRuleRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRuleRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.UserRule, error) {
	//TODO implement me
	panic("implement me")
}

// ListUserRulesByGroupId 根据groupID 获取所有的规则
func (u userRuleRepo) ListUserRulesByGroupId(groupId int64) ([]model.UserRule, error) {
	var userRules []model.UserRule
	err := u.data.db.Where("group_id = ?", groupId).Find(&userRules).Error
	return userRules, err
}

func (u userRuleRepo) CreateRuleInfoToEtcd(ctx context.Context, ruleInfoKey, ruleInfoValue string) error {
	_, err := u.data.etcd.KV.Put(ctx, ruleInfoKey, ruleInfoValue)
	return err
}

func (u userRuleRepo) DeleteRuleInfoToEtcd(ctx context.Context, ruleInfoKey string) error {
	_, err := u.data.etcd.KV.Delete(ctx, ruleInfoKey)
	return err
}
