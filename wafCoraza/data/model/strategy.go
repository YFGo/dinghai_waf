package model

import "github.com/corazawaf/coraza/v3"

// WAFStrategy WAF策略
type WAFStrategy struct {
	ID           string `json:"id"`
	SeclangRules string `json:"seclang_rules"`
}

type WafConfig struct {
	ID                    int                 `json:"id"`
	Action                uint8               `json:"action"`
	NextAction            uint8               `json:"next_action"`
	Name                  string              `json:"name"`
	Description           string              `json:"description"`
	ModifyStrategyDTOList []ModifyStrategyDTO `json:"modify_strategy_dto_list"`
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
