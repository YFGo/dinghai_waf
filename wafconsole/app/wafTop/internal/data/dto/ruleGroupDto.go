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
	RiskLevel   uint8            `json:"risk_level"`
	SecLangMod  model.SeclangMod `json:"sec_lang_mod"`
}

type RuleGroupInfo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsBuildin   uint8  `json:"is_buildin"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type RuleGroupEtcd struct {
	ID         int64   `json:"id"`
	IsBuildin  uint8   `json:"is_buildin"`
	RuleIdList []int64 `json:"rule_id_list"`
}
