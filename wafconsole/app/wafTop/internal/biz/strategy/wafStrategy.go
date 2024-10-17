package strategyBiz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	ruleBiz "wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"
)

type WafStrategyRepo interface {
	iface.BaseRepo[model.Strategy]
}

type WafStrategyUsecase struct {
	repo          WafStrategyRepo
	ruleGroupRepo ruleBiz.RuleGroupRepo
}

func NewWafStrategyUsecase(repo WafStrategyRepo) *WafStrategyUsecase {
	return &WafStrategyUsecase{repo: repo}
}

func (w *WafStrategyUsecase) checkStrategyIsExist(ctx context.Context, id int64, name string) bool {
	_, err := w.repo.GetByNameAndID(ctx, name, id)
	if errors.Is(err, gorm.ErrRecordNotFound) { //策略不存在 , 可以插入数据
		return true
	}
	return false
}

func (w *WafStrategyUsecase) checkRuleGroupIsExist(ctx context.Context, groupId int64) bool {
	_, err := w.ruleGroupRepo.Get(ctx, groupId)
	if errors.Is(err, gorm.ErrRecordNotFound) { //规则不存在 , 禁止插入数据
		return false
	}
	return true
}

// CreateStrategy 新增策略
func (w *WafStrategyUsecase) CreateStrategy(ctx context.Context, strategy model.Strategy) error {
	if !w.checkStrategyIsExist(ctx, 0, strategy.Name) {
		return errors.New("策略已存在")
	}
	for _, strategyConfig := range strategy.StrategyConfig {
		if !w.checkRuleGroupIsExist(ctx, strategyConfig.RuleGroupID) {
			return errors.New("规则组不存在")
		}
	}
	_, err := w.repo.Create(ctx, strategy)
	if err != nil {
		slog.ErrorContext(ctx, "create strategy failed: ", err)
		return err
	}
	return nil
}

// UpdateStrategy 修改策略
func (w *WafStrategyUsecase) UpdateStrategy(ctx context.Context, id int64, strategy model.Strategy) error {
	if !w.checkStrategyIsExist(ctx, 0, strategy.Name) {
		return errors.New("策略已存在")
	}
	for _, strategyConfig := range strategy.StrategyConfig {
		if !w.checkRuleGroupIsExist(ctx, strategyConfig.RuleGroupID) {
			return errors.New("规则组不存在")
		}
	}
	if err := w.repo.Update(ctx, id, strategy); err != nil {
		slog.ErrorContext(ctx, "update strategy failed: ", err)
		return err
	}
	return nil
}

// DeleteStrategy 删除策略
func (w *WafStrategyUsecase) DeleteStrategy(ctx context.Context, ids []int64) error {
	_, err := w.repo.Delete(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete strategy failed: ", err)
		return err
	}
	return nil
}

// GetStrategyDetail 查询策略详情
func (w *WafStrategyUsecase) GetStrategyDetail(ctx context.Context, id int64) {

}
