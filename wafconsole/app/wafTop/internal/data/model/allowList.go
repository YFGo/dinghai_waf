package model

import "gorm.io/gorm"

const AllowListTableName = "waf_allow_list"

type AllowList struct {
	*gorm.Model
	Name        string `json:"name"  gorm:"type:varchar(255);not null;comment:'白名单昵称'"`
	Description string `json:"description" gorm:"type:varchar(255);not null;comment:'白名单描述'"`
	Key         string `json:"key" gorm:"type:varchar(255);not null;comment:'白名单匹配方式[IP , URI]'"`
	Value       string `json:"value" gorm:"type:varchar(255);not null;comment:'白名单匹配值'"`
}

func (AllowList) TableName() string {
	return AllowListTableName
}
