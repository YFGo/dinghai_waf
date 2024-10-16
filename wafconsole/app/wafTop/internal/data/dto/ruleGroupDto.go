package dto

import "wafconsole/app/wafTop/internal/data/model"

type RuleGroupDto struct {
	RuleGroup model.RuleGroup `json:"rule_group"`
	RuleInfos []RuleInfo      `json:"rule_infos"`
}

type RuleInfo struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Status      uint8            `json:"status"`
	RiskLevel   uint8            `json:"risk_level"`
	SecLangMod  model.SeclangMod `json:"sec_lang_mod"`
}
