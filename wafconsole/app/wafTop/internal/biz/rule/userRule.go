package ruleBiz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log/slog"
	"strings"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

const (
	IPSameMod = "SecRule REMOTE_ADDR \"METHOD IP_MOD\" \"id:1,block,msg:'IP address is not allowed.'\""
	Same      = "等于"
	NotSame   = "不等于"
	Like      = "Like"
)

type UserRuleRepo interface {
	iface.BaseRepo[model.UserRule]
	ListUserRulesByGroupId(groupId int64) ([]model.UserRule, error)
}

type UserRuleUsecase struct {
	repo          UserRuleRepo
	ruleGroupRepo RuleGroupRepo
}

func NewUserRuleUsecase(repo UserRuleRepo, ruleGroupRepo RuleGroupRepo) *UserRuleUsecase {
	return &UserRuleUsecase{
		repo:          repo,
		ruleGroupRepo: ruleGroupRepo,
	}
}

// checkUserRuleIsExist 检查用户自定义规则是否已经存在
func (u *UserRuleUsecase) checkUserRuleIsExist(ctx context.Context, ruleId int64, name string) bool {
	_, err := u.repo.GetByNameAndID(ctx, name, ruleId)
	if errors.Is(err, gorm.ErrRecordNotFound) { //如果没有查询到记录 , 可以添加
		return true
	}
	return false
}

// checkRuleGroupIsExist 检查所选的规则组是否存在
func (u *UserRuleUsecase) checkRuleGroupIsExist(ctx context.Context, groupId int64) bool {
	_, err := u.ruleGroupRepo.Get(ctx, groupId)
	if errors.Is(err, gorm.ErrRecordNotFound) { //如果没有查询到记录 , 规则组不存在
		return false
	}
	return true
}

// CreateUserRule 创建用户自定义规则
func (u *UserRuleUsecase) CreateUserRule(ctx context.Context, userRule model.UserRule) error {
	if !u.checkUserRuleIsExist(ctx, 0, userRule.Name) {
		return errors.New("用户自定义规则已经存在")
	}
	if !u.checkRuleGroupIsExist(ctx, userRule.GroupId) {
		return errors.New("规则组不存在")
	}
	// 处理用户自定义规则
	seclang := u.disposeUserRule(ctx, userRule.SeclangMod)
	userRule.ModSecurity = seclang
	_, err := u.repo.Create(ctx, userRule)
	if err != nil {
		slog.ErrorContext(ctx, "create user_rule is failed: ", err)
		return err
	}
	return nil
}

// UpdateUserRule 修改用户自定义规则
func (u *UserRuleUsecase) UpdateUserRule(ctx context.Context, id int64, userRule model.UserRule) error {
	if !u.checkUserRuleIsExist(ctx, id, userRule.Name) {
		return errors.New("用户自定义规则已经存在")
	}
	if !u.checkRuleGroupIsExist(ctx, userRule.GroupId) {
		return errors.New("规则组不存在")
	}
	seclang := u.disposeUserRule(ctx, userRule.SeclangMod)
	userRule.ModSecurity = seclang
	if err := u.repo.Update(ctx, id, userRule); err != nil {
		slog.ErrorContext(ctx, "update user_rule is failed: ", err)
		return err
	}
	return nil
}

// DeleteUserRule 删除用户自定义规则
func (u *UserRuleUsecase) DeleteUserRule(ctx context.Context, ids []int64) error {
	_, err := u.repo.Delete(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete user_rule is failed: ", err)
		return err
	}
	return nil
}

func (u *UserRuleUsecase) disposeUserRule(ctx context.Context, userRule model.SeclangMod) string {
	var (
		res    string
		method string
	)
	switch userRule.MatchGoal {
	case "IP":
		if userRule.MatchAction == Same {
			method = "@eq"
		} else if userRule.MatchAction == NotSame {
			method = "!@eq"
		}
		res = strings.ReplaceAll(IPSameMod, "METHOD", method)
		res = strings.ReplaceAll(res, "IP_MOD", userRule.MatchContent)
	}
	return res
}
