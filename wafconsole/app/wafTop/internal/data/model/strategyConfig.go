package model

import "gorm.io/gorm"

const StrategyConfigTableName = "waf_strategy_config"

type StrategyConfig struct {
	*gorm.Model
	StrategyId  int64 `json:"strategy_id"  gorm:"type:int;not null;comment:'策略id'"`
	RuleGroupID int64 `json:"rule_group_id"  gorm:"type:int;not null;comment:'规则组id'"`
	Status      uint8 `json:"status"  gorm:"type:tinyint;not null;default:1;comment:'-1:disable 1:enable'"`      //策略状态
	AlertLevel  uint8 `json:"alert_level"  gorm:"type:tinyint;not null;default:1;comment:'-1:disable 1:enable'"` //策略的风险等级 预留字段
	Action      uint8 `json:"action"  gorm:"type:tinyint;not null;default:1;comment:'-1:拦截 1:记录'"`               //命中时的行为
	NextAction  uint8 `json:"next_action" gorm:"type:tinyint;not null;default:1;comment:'-1:拦截 1:记录'"`           //命中后的行为
}

func (StrategyConfig) TableName() string {
	return StrategyConfigTableName
}
