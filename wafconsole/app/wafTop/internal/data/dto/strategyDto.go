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

type StrategyEtcdInfo struct {
	ID                    int64               `json:"id"`
	Action                uint8               `json:"action"`
	NextAction            uint8               `json:"next_action"`
	Name                  string              `json:"name"`
	Description           string              `json:"description"`
	ModifyStrategyDtoList []ModifyStrategyDto `json:"modify_strategy_dto_list"`
}

type ModifyStrategyDto struct {
	IsBuilding uint8  `json:"is_building"`
	RuleName   string `json:"rule_name"`
	Seclang    string `json:"seclang"`
}
