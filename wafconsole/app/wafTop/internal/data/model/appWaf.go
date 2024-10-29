package model

import "gorm.io/gorm"

const AppWafTableName = "app_waf"

type AppWaf struct {
	*gorm.Model
	Name     string `json:"name"  gorm:"type:varchar(255);not null;unique;comment:'应用名称'"`
	Url      string `json:"url"  gorm:"type:varchar(255);comment:'web程序地址'"`
	ServerID int64  `json:"server_id"  gorm:"type:bigint;not null;unique;comment:'web程序所在服务器id'"`
}

func (AppWaf) TableName() string {
	return AppWafTableName
}
