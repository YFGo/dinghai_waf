package model

import "gorm.io/gorm"

const ServerAllowTableName = "waf_server_allow"

type ServerAllow struct {
	*gorm.Model
	AllowID  int64 `json:"allow_id" gorm:"type:int;not null;comment:'白名单id'"`
	ServerID int64 `json:"server_id" gorm:"type:int;not null;comment:'server_id'"`
}

func (ServerAllow) TableName() string {
	return ServerAllowTableName
}
