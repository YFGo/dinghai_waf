package ruleBiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/dto"
	"wafconsole/app/wafTop/internal/data/model"
)

type RuleGroupRepo interface {
	iface.BaseRepo[model.RuleGroup]
}

type RuleGroupUsecase struct {
	repo            RuleGroupRepo
	userRuleRepo    UserRuleRepo
	buildinRuleRepo BuildRuleRepo
}

func NewRuleGroupUsecase(repo RuleGroupRepo, userRuleRepo UserRuleRepo, buildinRuleRepo BuildRuleRepo) *RuleGroupUsecase {
	return &RuleGroupUsecase{
		repo:            repo,
		userRuleRepo:    userRuleRepo,
		buildinRuleRepo: buildinRuleRepo,
	}
}

// checkNameIsExist 检查规则组昵称是否存在
func (r *RuleGroupUsecase) checkNameIsExist(ctx context.Context, name string, id int64) bool {
	_, err := r.repo.GetByNameAndID(ctx, name, id)
	if errors.Is(err, gorm.ErrRecordNotFound) { //如果没有查询到记录 , 可以添加
		return true
	}
	return false
}

// GetRuleGroupDetail 获取规则组详情
func (r *RuleGroupUsecase) GetRuleGroupDetail(ctx context.Context, id int64) (*dto.RuleGroupDto, error) {
	// 1 . 首先获取规则组详情
	ruleGroupDetail, err := r.repo.Get(ctx, id)
	if err != nil { //如果err不是未查询到记录 , 则说明是其它错误
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.ErrorContext(ctx, "get rule group detail err : %v", err)
			return nil, err
		} else {
			// 规则组不存在
			return nil, nil
		}
	}
	ruleGroupInformations := &dto.RuleGroupDto{
		RuleGroup: ruleGroupDetail,
	}
	// 2 . 根据规则组id 查询所有的规则
	var ruleInfos []dto.RuleInfo
	switch ruleGroupDetail.IsBuildin {
	case 1: //内置规则组
		buildinRules, err := r.buildinRuleRepo.ListBuildinRulesByGroupId(ctx, int64(ruleGroupDetail.ID))
		if err != nil {
			slog.ErrorContext(ctx, "get buildin rules by group id err : %v", err)
			return nil, err
		}
		for _, buildinRule := range buildinRules {
			ruleInfo := dto.RuleInfo{
				ID:          int64(buildinRule.ID),
				Name:        buildinRule.Name,
				Description: buildinRule.Description,
				Status:      buildinRule.Status,
				RiskLevel:   buildinRule.RiskLevel,
			}
			ruleInfos = append(ruleInfos, ruleInfo)
		}
	case 2: //自定义规则组
		userRules, err := r.userRuleRepo.ListUserRulesByGroupId(int64(ruleGroupDetail.ID))
		if err != nil {
			slog.ErrorContext(ctx, "get user rules by group id err : %v", err)
			return nil, err
		}
		for _, userRule := range userRules {
			ruleInfo := dto.RuleInfo{
				ID:          int64(userRule.ID),
				Name:        userRule.Name,
				Description: userRule.Description,
				Status:      userRule.Status,
				RiskLevel:   userRule.RiskLevel,
				SecLangMod:  userRule.SeclangMod,
			}
			ruleInfos = append(ruleInfos, ruleInfo)
		}
	}
	ruleGroupInformations.RuleInfos = ruleInfos
	return ruleGroupInformations, nil
}

// CreateRuleGroup 新增规则组
func (r *RuleGroupUsecase) CreateRuleGroup(ctx context.Context, ruleGroup model.RuleGroup) error {
	if !r.checkNameIsExist(ctx, ruleGroup.Name, 0) { //昵称重复
		return nil
	}
	_, err := r.repo.Create(ctx, ruleGroup)
	if err != nil {
		slog.ErrorContext(ctx, "create rule group err : %v", err)
		return err
	}
	return nil
}

// UpdateRuleGroup 修改规则组
func (r *RuleGroupUsecase) UpdateRuleGroup(ctx context.Context, id int64, ruleGroup model.RuleGroup) error {
	if !r.checkNameIsExist(ctx, ruleGroup.Name, id) { //昵称重复
		return nil
	}
	if err := r.repo.Update(ctx, id, ruleGroup); err != nil {
		slog.ErrorContext(ctx, "update rule_group err: ", err)
		return err
	}
	return nil
}

// DeleteRuleGroup 删除规则组
func (r *RuleGroupUsecase) DeleteRuleGroup(ctx context.Context, ids []int64) error {
	affected, err := r.repo.Delete(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete rule_group err: ", err)
		return err
	}
	if int(affected) != len(ids) {
		slog.ErrorContext(ctx, "rule_group is not exists", ids)
	}
	return nil
}

// ListRuleGroupSearch 查询规则组列表
func (r *RuleGroupUsecase) ListRuleGroupSearch(ctx context.Context, name string, isBuildin int8, limit, offset int64) ([]model.RuleGroup, int64, error) {
	whereOptions := make([]iface.WhereOptionWithReturn, 0)
	if len(name) != 0 {
		whereOptions = append(whereOptions, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", "%"+name+"%")
		})
	}
	if isBuildin != 0 {
		whereOptions = append(whereOptions, func(db *gorm.DB) *gorm.DB {
			return db.Where("is_buildin = ?", isBuildin)
		})
	}
	total, err := r.repo.Count(ctx, whereOptions...) // 总数量
	if err != nil {
		slog.ErrorContext(ctx, "count rule_group err : %v", err)
		return nil, 0, err
	}
	listRuleGroup, err := r.repo.ListByWhere(ctx, limit, offset, whereOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "list rule_group err : %v", err)
		return nil, 0, err
	}
	return listRuleGroup, total, nil
}