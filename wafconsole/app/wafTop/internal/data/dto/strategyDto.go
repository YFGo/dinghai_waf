package dto

import "wafconsole/app/wafTop/internal/data/model"

type StrategyDetailDto struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Status        uint8             `json:"status"`
	Action        uint8             `json:"action"`
	NextAction    uint8             `json:"next_action"`
	RuleGroupInfo []model.RuleGroup `json:"rule_group_info"`
}
