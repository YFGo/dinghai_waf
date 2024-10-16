package model

import "gorm.io/gorm"

const StrategyTableName = "waf_strategy"

type Strategy struct {
	*gorm.Model
	Name        string `json:"name"  gorm:"type:varchar(255);not null;unique;comment:'策略名称'"`
	Description string `json:"description" gorm:"type:varchar(255);comment:'策略描述'"`
	Kind        uint8  `json:"kind" gorm:"type:tinyint;comment:'策略类别-预留字段'"`
	Status      uint8  `json:"status"  gorm:"type:tinyint;not null;default:1;comment:'2:disable 1:enable'"`      //策略状态
	AlertLevel  uint8  `json:"alert_level"  gorm:"type:tinyint;not null;default:1;comment:'2:disable 1:enable'"` //策略的风险等级 预留字段
	Action      uint8  `json:"action"  gorm:"type:tinyint;not null;default:1;comment:'2:拦截 1:记录'"`               //命中时的行为
	NextAction  uint8  `json:"next_action" gorm:"type:tinyint;not null;default:1;comment:'2:拦截 1:记录'"`           //命中后的行为
}

func (Strategy) TableName() string {
	return StrategyTableName
}
