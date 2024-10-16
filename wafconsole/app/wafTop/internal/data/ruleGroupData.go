package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	ruleBiz "wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"
)

type ruleGroupRepo struct {
	data *Data
}

func NewRuleGroupRepo(data *Data) ruleBiz.RuleGroupRepo {
	return &ruleGroupRepo{data: data}
}

// Get 获取规则组详情
func (r ruleGroupRepo) Get(ctx context.Context, id int64) (model.RuleGroup, error) {
	var ruleGroupInfo model.RuleGroup
	err := r.data.db.Where("id = ?", id).First(&ruleGroupInfo).Error
	return ruleGroupInfo, err
}

// GetByNameAndID 通过昵称查询规则组信息
func (r ruleGroupRepo) GetByNameAndID(ctx context.Context, name string, id int64) (model.RuleGroup, error) {
	var ruleGroupInfo model.RuleGroup
	if id == 0 {
		err := r.data.db.Where("name = ?", name).First(&ruleGroupInfo).Error
		return ruleGroupInfo, err
	}
	err := r.data.db.Where("name = ? AND id != ?", name, id).First(&ruleGroupInfo).Error
	return ruleGroupInfo, err
}

// Create 创建规则组
func (r ruleGroupRepo) Create(ctx context.Context, ruleGroup model.RuleGroup) (int64, error) {
	err := r.data.db.Create(&ruleGroup).Error
	return int64(ruleGroup.ID), err
}

// Update 更新规则组信息
func (r ruleGroupRepo) Update(ctx context.Context, id int64, ruleGroup model.RuleGroup) error {
	err := r.data.db.Where("id = ?", id).Updates(&ruleGroup).Error
	return err
}

// Delete 删除规则组信息
func (r ruleGroupRepo) Delete(ctx context.Context, ids []int64) (int64, error) {
	res := r.data.db.Where("id in (?)", ids).Delete(&model.RuleGroup{})
	return res.RowsAffected, res.Error
}

// Count 统计符合条件的数据数量
func (r ruleGroupRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	mysqlDB := r.data.db.Table(model.TableNameRuleGroup)
	for _, opt := range withReturn {
		mysqlDB = opt(mysqlDB)
	}
	var total int64
	err := mysqlDB.Count(&total).Error
	return total, err
}

// ListByWhere 查询规则组列表
func (r ruleGroupRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.RuleGroup, error) {
	mysqlDB := r.data.db.Table(model.TableNameRuleGroup)
	for _, opt := range opts {
		mysqlDB = opt(mysqlDB)
	}
	var ruleGroupList []model.RuleGroup
	err := mysqlDB.Limit(int(limit)).Offset(int(offset)).Find(&ruleGroupList).Error
	return ruleGroupList, err
}