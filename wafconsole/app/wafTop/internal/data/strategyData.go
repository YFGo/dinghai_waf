package data

import (
	"context"
	"gorm.io/gorm"
	"wafconsole/app/wafTop/internal/biz/iface"
	strategyBiz "wafconsole/app/wafTop/internal/biz/strategy"
	"wafconsole/app/wafTop/internal/data/model"
)

type wafStrategyRepo struct {
	data *Data
}

func NewWafStrategyRepo(data *Data) strategyBiz.WafStrategyRepo {
	return &wafStrategyRepo{
		data: data,
	}
}

// Get 获取策略详情
func (w wafStrategyRepo) Get(ctx context.Context, id int64) (model.Strategy, error) {
	var strategyInfo model.Strategy
	err := w.data.db.Where("id = ?", id).First(&strategyInfo).Error
	return strategyInfo, err
}

// GetByNameAndID 根据昵称和id获取策略信息
func (w wafStrategyRepo) GetByNameAndID(ctx context.Context, name string, id int64) (model.Strategy, error) {
	var strategyInfo model.Strategy
	if id == 0 {
		err := w.data.db.Where("name = ?", name).First(&strategyInfo).Error
		return strategyInfo, err
	}
	err := w.data.db.Where("name = ? and id != ?", name, id).First(&strategyInfo).Error
	return strategyInfo, err
}

// Create 用户新增策略
func (w wafStrategyRepo) Create(ctx context.Context, strategyInfo model.Strategy) (int64, error) {
	strategy := model.Strategy{
		Name:        strategyInfo.Name,
		Description: strategyInfo.Description,
		Status:      strategyInfo.Status,
		Action:      strategyInfo.Action,
		NextAction:  strategyInfo.NextAction,
	}
	err := w.data.db.Transaction(func(tx *gorm.DB) error {
		// 1 . 首先插入策略表数据
		if err := tx.Table(model.StrategyTableName).Create(&strategy).Error; err != nil {
			return err
		}
		// 2 . 插入策略表和规则组表的关联表中的数据
		for _, config := range strategyInfo.StrategyConfig {
			configInfo := model.StrategyConfig{
				StrategyId:  int64(strategy.ID),
				RuleGroupID: config.RuleGroupID,
			}
			if err := tx.Table(model.StrategyConfigTableName).Create(&configInfo).Error; err != nil {
				return err
			}
		}

		return nil //返回nil , 提交事务
	})
	return int64(strategy.ID), err
}

func (w wafStrategyRepo) Update(ctx context.Context, id int64, strategyInfo model.Strategy) error {
	strategy := model.Strategy{
		Name:        strategyInfo.Name,
		Description: strategyInfo.Description,
		Status:      strategyInfo.Status,
		Action:      strategyInfo.Action,
		NextAction:  strategyInfo.NextAction,
	}
	err := w.data.db.Transaction(func(tx *gorm.DB) error {
		// 1 . 首先插入策略表数据
		if err := tx.Table(model.StrategyTableName).Where("id = ?", id).Updates(&strategy).Error; err != nil {
			return err
		}
		// 2 . 删除中间表中的数据
		if err := tx.Where("strategy_id = ?", id).Unscoped().Delete(&model.StrategyConfig{}).Error; err != nil {
			return err
		}
		// 3 . 插入新的数据
		for _, config := range strategyInfo.StrategyConfig {
			configInfo := model.StrategyConfig{
				StrategyId:  id,
				RuleGroupID: config.RuleGroupID,
			}
			if err := tx.Table(model.StrategyConfigTableName).Create(&configInfo).Error; err != nil {
				return err
			}
		}
		return nil //返回nil , 提交事务
	})
	return err
}

// Delete 删除策略
func (w wafStrategyRepo) Delete(ctx context.Context, ids []int64) (int64, error) {
	err := w.data.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("strategy_id in (?)", ids).Unscoped().Delete(&model.StrategyConfig{}).Error; err != nil { //中间表的数据物理删除
			return err
		}
		if err := tx.Where("id in (?)", ids).Delete(&model.Strategy{}).Error; err != nil {
			return err
		}
		return nil
	})
	return 0, err
}

// Count 统计符合条件的数量
func (w wafStrategyRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	mysqlDB := w.data.db.Table(model.StrategyTableName)
	for _, opt := range withReturn {
		opt(mysqlDB)
	}
	var total int64
	err := mysqlDB.Count(&total).Error
	return total, err
}

// ListByWhere 批量查询
func (w wafStrategyRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.Strategy, error) {
	mysqlDB := w.data.db.Table(model.StrategyTableName)
	for _, opt := range opts {
		opt(mysqlDB)
	}
	var strategyInfos []model.Strategy
	err := mysqlDB.Limit(int(limit)).Offset(int(offset)).Find(&strategyInfos).Error
	return strategyInfos, err
}

func (w wafStrategyRepo) CreateStrategyForEtcd(ctx context.Context, strategyKey, strategyValue string) error {
	_, err := w.data.etcd.KV.Put(ctx, strategyKey, strategyValue)
	return err
}
