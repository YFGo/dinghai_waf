package model

import "gorm.io/gorm"

const TableNameRuleGroup = "waf_rule_group"

type RuleGroup struct {
	*gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);unique;not null;comment:'规则组名称'"`
	Description string `json:"description" gorm:"type:varchar(255);comment:'规则组描述'"`
	IsBuildin   bool   `json:"is_buildin" gorm:"type:tinyint;comment:'是否是内置规则组 1:是 -1否'"`
}

func (RuleGroup) TableName() string {
	return TableNameRuleGroup
}
