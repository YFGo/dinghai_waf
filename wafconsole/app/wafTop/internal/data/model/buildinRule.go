package model

import (
	"gorm.io/gorm"
)

const TableNameBuildinRule = "waf_buildin_rule"

type BuildinRule struct {
	*gorm.Model
	Name        string `json:"name"  gorm:"type:varchar(255);not null;unique;comment:'内置规则名称'"`
	Description string `json:"description" gorm:"type:varchar(255);comment:'内置规则描述'"`
	RiskLevel   uint8  `json:"risk_level" gorm:"type:tinyint;not null;default:1;comment:'风险等级'"` //4个等级
	GroupId     int64  `json:"group_id"  gorm:"type:bigint;not null;comment:'分组id'"`
	Seclang     string `json:"seclang"  gorm:"type:varchar(255);not null;comment:'安全规则'"`
}

func (BuildinRule) TableName() string {
	return TableNameBuildinRule
}
