package model

import (
	"gorm.io/gorm"
	"log/slog"
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

// AfterCreate 在建表操作完成之后
func (BuildinRule) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Exec("INSERT INTO waf_buildin_rule (id , created_at, updated_at, name, description, group_id, seclang  , risk_level) VALUES (90001 , NOW() , NOW() , 'coraza基础配置', 'waf内核的基础配置文件' , 1 , 'Include wafCoraza/ruleset/coraza.conf' , 1 , 4),(90002 , NOW() , NOW() , 'coraza的内置防护规则1' , 'coraza内置防护规则的文件' , 1 , 'Include wafCoraza/ruleset/coreruleset/crs-setup.conf.example' , 1 , 4),(9003 , NOW() , NOW() , 'coraza的内置防护规则2' , 'coraza内置防护规则文件' , 1 , 'Include wafCoraza/ruleset/coreruleset/rules/*.conf' , 1 , 4);").Error
	if err != nil {
		slog.Error("buildinRule AfterCreate err:", err.Error())
		return err
	}
	return nil

}
