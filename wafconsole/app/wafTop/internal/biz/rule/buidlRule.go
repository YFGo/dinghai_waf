package ruleBiz

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type BuildRuleRepo interface {
	iface.BaseRepo[model.BuildinRule]
	ListBuildinRulesByGroupId(ctx context.Context, groupId int64) ([]model.BuildinRule, error)
}

type BuildRuleUsecase struct {
	repo BuildRuleRepo
}

func NewBuildRuleUsecase(repo BuildRuleRepo) *BuildRuleUsecase {
	return &BuildRuleUsecase{
		repo: repo,
	}
}

// GetBuildinRuleDetailById 根据id 获取内置规则详情
func (b *BuildRuleUsecase) GetBuildinRuleDetailById(ctx context.Context, id int64) (model.BuildinRule, error) {
	buildinRule, err := b.repo.Get(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "get buildin_rule by id failed: ", err)
		return buildinRule, err
	}
	return buildinRule, nil
}

func (b *BuildRuleUsecase) ListBuildinRules(ctx context.Context, name string, limit, offset int64) ([]model.BuildinRule, int64, error) {
	whereOpts := make([]iface.WhereOptionWithReturn, 0)
	if len(name) != 0 {
		whereOpts = append(whereOpts, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+name+"%")
		})
	}
	//获取总数量
	total, err := b.repo.Count(ctx, whereOpts...)
	if err != nil {
		slog.ErrorContext(ctx, "count buildin_rule failed: ", err)
		return nil, total, err
	}
	//获取内置规则
	buildinRules, err := b.repo.ListByWhere(ctx, limit, offset, whereOpts...)
	if err != nil {
		slog.ErrorContext(ctx, "list buildin_rule failed: ", err)
		return nil, total, err
	}
	return buildinRules, total, nil
}
