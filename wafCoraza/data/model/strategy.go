package model

import "github.com/corazawaf/coraza/v3"

// WAFStrategy WAF策略
type WAFStrategy struct {
	ID           string `json:"id"`
	SeclangRules string `json:"seclang_rules"`
}

type WafConfig struct {
	ID              int     `json:"id"`
	Action          uint8   `json:"action"`
	NextAction      uint8   `json:"next_action"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	RuleGroupIdList []int64 `json:"rule_group_id_list"`
}

type RuleGroup struct {
	ID         int     `json:"id"`
	IsBuildin  uint8   `json:"is_buildin"`
	RuleIDList []int64 `json:"rule_id_list"`
}

type Rule struct {
	ID        int    `json:"id"`
	RiskLevel uint8  `json:"risk_level"`
	Seclang   string `json:"seclang"`
}

type ModifyStrategyDTO struct {
	IsBuildin uint8  `json:"is_buildin"`
	RuleName  string `json:"rule_name"`
	Seclang   string `json:"seclang"`
}

type CorazaWaf struct {
	WAF         coraza.WAF `json:"waf"`
	Action      uint8      `json:"action"`
	NextAction  uint8      `json:"next_action"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}
