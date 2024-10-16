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
	//TODO implement me
	panic("implement me")
}

func (u userRuleRepo) Create(ctx context.Context, t model.UserRule) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRuleRepo) Update(ctx context.Context, i int64, t model.UserRule) error {
	//TODO implement me
	panic("implement me")
}

func (u userRuleRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
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
