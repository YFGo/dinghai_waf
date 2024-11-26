package dto

type AttackDayCount struct {
	AttackCount       int `json:"attack_count"`        //异常攻击数
	AttackIPCount     int `json:"attack_ip_count"`     //异常IP数
	AttackYesterday   int `json:"attack_yesterday"`    // 异常攻击数 昨日
	AttackIPYesterday int `json:"attack_ip_yesterday"` // 异常IP数 昨日
}
