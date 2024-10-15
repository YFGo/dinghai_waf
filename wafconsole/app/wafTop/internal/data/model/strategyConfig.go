package model

import "gorm.io/gorm"

const StrategyConfigTableName = "waf_strategy_config"

type StrategyConfig struct {
	*gorm.Model
	StrategyId  int64 `json:"strategy_id"  gorm:"type:int;not null;comment:'策略id'"`
	RuleGroupID int64 `json:"rule_group_id"  gorm:"type:int;not null;comment:'规则组id'"`
}

func (StrategyConfig) TableName() string {
	return StrategyConfigTableName
}
