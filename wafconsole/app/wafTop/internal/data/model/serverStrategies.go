package model

import "gorm.io/gorm"

const ServersStrategiesTableName = "waf_servers_strategies"

type ServerStrategies struct {
	*gorm.Model
	ServerID   int64 `json:"server_id"  gorm:"type:int;not null;comment:'服务器ID'"`
	StrategyId int64 `json:"strategy_id"  gorm:"type:int;not null;comment:'策略ID'"`
}

func (ServerStrategies) TableName() string {
	return ServersStrategiesTableName
}
