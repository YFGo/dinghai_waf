package ruleBiz

import (
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type UserRuleRepo interface {
	iface.BaseRepo[model.UserRule]
	ListUserRulesByGroupId(groupId int64) ([]model.UserRule, error)
}

type UserRuleUsecase struct {
	repo UserRuleRepo
}

func NewUserRuleUsecase(repo UserRuleRepo) *UserRuleUsecase {
	return &UserRuleUsecase{repo: repo}
}
