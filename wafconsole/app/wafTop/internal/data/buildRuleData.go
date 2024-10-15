package data

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"
)

type buildRuleRepo struct {
	data *Data
}

func NewBuildRuleRepo(data *Data) ruleBiz.BuildRuleRepo {
	return &buildRuleRepo{data: data}
}

// Get 根据id 获取内置规则详情
func (b buildRuleRepo) Get(ctx context.Context, id int64) (model.BuildinRule, error) {
	var buildinRuleInfo model.BuildinRule
	err := b.data.db.Where("id = ?", id).First(&buildinRuleInfo).Error
	return buildinRuleInfo, err
}

func (b buildRuleRepo) GetByName(ctx context.Context, s string) (model.BuildinRule, error) {
	//TODO implement me
	panic("implement me")
}

func (b buildRuleRepo) Create(ctx context.Context, t model.BuildinRule) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b buildRuleRepo) Update(ctx context.Context, i int64, t model.BuildinRule) error {
	//TODO implement me
	panic("implement me")
}

func (b buildRuleRepo) Delete(ctx context.Context, i int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

// Count 获取符合条件的总数量
func (b buildRuleRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	var total int64
	err := b.data.db.Table(model.TableNameBuildinRule).Where(withReturn).Count(&total).Error
	return total, err
}

// ListByWhere 根据条件获取内置规则
func (b buildRuleRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.BuildinRule, error) {
	var buildinRules []model.BuildinRule
	err := b.data.db.Where(opts).Offset(int(offset)).Limit(int(limit)).Find(&buildinRules).Error
	return buildinRules, err
}
