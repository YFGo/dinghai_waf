package model

import "gorm.io/gorm"

const StrategyTableName = "waf_strategy"

type Strategy struct {
	*gorm.Model
	Name        string `json:"name"  gorm:"type:varchar(255);not null;unique;comment:'策略名称'"`
	Description string `json:"description" gorm:"type:varchar(255);comment:'策略描述'"`
	Kind        uint8  `json:"kind" gorm:"type:tinyint;comment:'策略类别-预留字段'"`
	IsBuildin   bool   `json:"is_buildin" gorm:"type:tinyint;comment:'是否是内置策略'"`
}

func (Strategy) TableName() string {
	return StrategyTableName
}
