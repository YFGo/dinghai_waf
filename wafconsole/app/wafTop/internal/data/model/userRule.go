package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"log/slog"
)

const UserRuleTableName = "waf_user_rule"

type UserRule struct {
	*gorm.Model
	Name        string     `json:"name"  gorm:"type:varchar(255);not null;unique;comment:'规则名称'"`
	Description string     `json:"description" gorm:"type:varchar(255);comment:'规则描述'"`
	RiskLevel   uint8      `json:"risk_level" gorm:"type:tinyint;not null;default:1;comment:'风险等级'"` //4个等级
	Status      uint8      `json:"status"  gorm:"type:tinyint;not null;default:1;comment:'状态'"`      //1:启用 -1:禁用`
	GroupId     int64      `json:"group_id"  gorm:"type:bigint;not null;comment:'分组id'"`
	Seclang     string     `json:"seclang"  gorm:"type:varchar(255);not null;comment:'安全规则'"`
	SeclangMod  SeclangMod `json:"seclang_mod" gorm:"-"`
	ModSecurity string     `json:"mod_security" gorm:"type:varchar(255);not null;comment:'安全规则语言'"`
}

type SeclangMod struct {
	MatchGoal    string `json:"match_goal"`    //匹配目标
	MatchAction  string `json:"match_action"`  //匹配方式
	MatchContent string `json:"match_content"` //匹配内容
}

func (u *UserRule) TableName() string {
	return UserRuleTableName
}

// BeforeCreate 在创建数据之前
func (u *UserRule) BeforeCreate(tx *gorm.DB) (err error) {
	// 将SeclangMod转换为 JSON字符串
	seclang, err := json.Marshal(u.SeclangMod)
	if err != nil {
		slog.Error("user rule json Marshaling failed", err)
		return err
	}
	u.Seclang = string(seclang)
	return nil
}

// AfterFind 在查询之后
func (u *UserRule) AfterFind(tx *gorm.DB) (err error) {
	var seclangMod SeclangMod
	err = json.Unmarshal([]byte(u.Seclang), &seclangMod)
	if err != nil {
		slog.Error("user rule json Unmarshaling failed", err)
		return err
	}
	u.SeclangMod = seclangMod
	return nil
}
