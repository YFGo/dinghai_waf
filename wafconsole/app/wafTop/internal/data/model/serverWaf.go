package model

import "gorm.io/gorm"

const ServerWafTableName = "server_waf"

type ServerWaf struct {
	*gorm.Model
	Name          string  `json:"name" gorm:"type:varchar(255);not null;unique;comment:'服务器名称'"`
	Host          string  `json:"host"  gorm:"type:varchar(255);not null;unique;comment:'服务器域名'"`
	IP            string  `json:"ip"  gorm:"type:varchar(255);not null;comment:'服务器IP地址'"`
	Port          int     `json:"port" gorm:"type:int;not null;comment:'服务器端口'"`
	TLS           uint8   `json:"tls"  gorm:"type:tinyint;not null;default:2;comment:'是否开启TLS 2否 1是'"`
	Cert          string  `json:"cert" gorm:"type:varchar(255);comment:'TLS证书'"`
	Key           string  `json:"key" gorm:"type:varchar(255);comment:'TLS密钥'"`
	ServerGroupID int64   `json:"server_group_id"  gorm:"type:bigint;not null;default:0;comment:'服务器组ID-预留字段'"`
	StrategiesID  []int64 `json:"strategies_id" gorm:"-"` //关联的策略id
	AllowListID   []int64 `json:"allow_list_id" gorm:"-"` //白名单id
}

func (ServerWaf) TableName() string {
	return ServerWafTableName
}
