package dto

type UserRuleInfo struct {
	ID        int64  `json:"id"`
	RiskLevel uint8  `json:"risk_level"`
	Seclang   string `json:"seclang"`
}
