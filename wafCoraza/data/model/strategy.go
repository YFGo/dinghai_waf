package model

// WAFStrategy WAF策略
type WAFStrategy struct {
	ID           string `json:"id"`
	SeclangRules string `json:"seclang_rules"`
}
