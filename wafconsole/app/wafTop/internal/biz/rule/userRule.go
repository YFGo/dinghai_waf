package ruleBiz

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
	"strings"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/dto"
	"wafconsole/app/wafTop/internal/data/model"
)

const (
	IPSameMod = "SecRule REMOTE_ADDR \"METHOD IP_MOD\" \"id:1,block,msg:'IP address is not allowed.'\""
	Same      = "等于"
	NotSame   = "不等于"
	Like      = "Like"

	RulePrefix = "rule_"
)

type UserRuleRepo interface {
	iface.BaseRepo[model.UserRule]
	ListUserRulesByGroupId(groupId int64) ([]model.UserRule, error)
	CreateRuleInfoToEtcd(ctx context.Context, ruleInfoKey, ruleInfoValue string) error
	DeleteRuleInfoToEtcd(ctx context.Context, ruleInfoKey string) error
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

func (u *UserRuleUsecase) createRuleToEtcd(ctx context.Context, ruleID int64, userRule model.UserRule) error {
	userRuleInfo := dto.UserRuleInfo{
		ID:        ruleID,
		RiskLevel: userRule.RiskLevel,
		Seclang:   userRule.ModSecurity,
	}
	ruleInfoKey := RulePrefix + strconv.Itoa(int(userRuleInfo.ID))
	ruleInfoValue, err := json.Marshal(&userRuleInfo)
	if err != nil {
		slog.ErrorContext(ctx, "marshal user_rule_info is failed: ", err)
		return err
	}
	ruleGroupKey := "rule_group_" + strconv.Itoa(int(userRule.GroupId))
	//根据规则组id查询 规则组信息
	ruleGroup, err := u.ruleGroupRepo.GetRuleGroupEtcd(ctx, ruleGroupKey)
	if err != nil {
		slog.ErrorContext(ctx, "get rule_group_info from etcd is failed: ", err)
		return err
	}
	// 将规则组信息转换为json
	var ruleGroupDto dto.RuleGroupEtcd
	err = json.Unmarshal([]byte(ruleGroup), &ruleGroupDto)
	if err != nil {
		slog.ErrorContext(ctx, "unmarshal rule_group_info is failed: ", err)
		return err
	}
	// 将新的信息追加到此规则组中
	ruleGroupDto.RuleIdList = append(ruleGroupDto.RuleIdList, ruleID)
	newRuleInfoDto, err := json.Marshal(&ruleGroupDto)
	if err != nil {
		slog.ErrorContext(ctx, "marshal rule_group_info is failed: ", err)
		return err
	}
	// 更新数据
	if err = u.ruleGroupRepo.CreateRuleGroupInfoToEtcd(ctx, ruleGroupKey, string(newRuleInfoDto)); err != nil {
		slog.ErrorContext(ctx, "create rule_group_info to etcd is failed: ", err)
		return err
	}
	// 需要更新规则组的键值对
	if err = u.repo.CreateRuleInfoToEtcd(ctx, ruleInfoKey, string(ruleInfoValue)); err != nil {
		slog.ErrorContext(ctx, "create rule_info to etcd is failed: ", err)
		return err
	}
	return nil
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
	userRuleID, err := u.repo.Create(ctx, userRule)
	if err != nil {
		slog.ErrorContext(ctx, "create user_rule is failed: ", err)
		return err
	}
	err = u.createRuleToEtcd(ctx, userRuleID, userRule)
	if err != nil {
		slog.ErrorContext(ctx, "create rule_info to etcd is failed: ", err)
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
	// 更新规则信息存入etcd
	err := u.createRuleToEtcd(ctx, id, userRule)
	if err != nil {
		slog.ErrorContext(ctx, "create rule_info to etcd is failed: ", err)
		return err
	}
	return nil
}

// DeleteUserRule 删除用户自定义规则
func (u *UserRuleUsecase) DeleteUserRule(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		ruleInfoKey := RulePrefix + strconv.Itoa(int(id))
		if err := u.repo.DeleteRuleInfoToEtcd(ctx, ruleInfoKey); err != nil { // 删除etcd中的规则信息
			slog.ErrorContext(ctx, "delete rule_info to etcd is failed: ", err)
			return err
		}
		// 获取规则组id
		ruleGroupInfo, err := u.repo.Get(ctx, id)
		if err != nil {
			slog.ErrorContext(ctx, "get rule_group_id is failed: ", err)
			return err
		}
		ruleGroupKey := "rule_group_" + strconv.Itoa(int(ruleGroupInfo.GroupId))
		// 获取规则组信息
		ruleGroupInfoEtcd, err := u.ruleGroupRepo.GetRuleGroupEtcd(ctx, ruleGroupKey)
		if err != nil {
			slog.ErrorContext(ctx, "get rule_group_info from etcd is failed: ", err)
			return err
		}
		// 将规则组信息转换为json
		var ruleGroupDto dto.RuleGroupEtcd
		err = json.Unmarshal([]byte(ruleGroupInfoEtcd), &ruleGroupDto)
		if err != nil {
			slog.ErrorContext(ctx, "unmarshal rule_group_info is failed: ", err)
			return err
		}
		// 删除规则组中的此规则id
		var newRuleIDList []int64
		for _, ruleID := range ruleGroupDto.RuleIdList {
			if ruleID != id {
				newRuleIDList = append(newRuleIDList, ruleID)
			}
		}
		ruleGroupDto.RuleIdList = newRuleIDList
		newRuleInfoDto, err := json.Marshal(&ruleGroupDto)
		if err != nil {
			slog.ErrorContext(ctx, "marshal rule_group_info is failed: ", err)
			return err
		}
		// 更新数据
		if err = u.ruleGroupRepo.CreateRuleGroupInfoToEtcd(ctx, ruleGroupKey, string(newRuleInfoDto)); err != nil { // 更新规则组信息
			slog.ErrorContext(ctx, "create rule_group_info to etcd is failed: ", err)
			return err
		}
	}
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
