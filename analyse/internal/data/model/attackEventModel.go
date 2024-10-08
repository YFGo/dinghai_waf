package model

import "time"

// AttackEvent 定义攻击事件的结构
type AttackEvent struct {
	RuleId        int       `json:"rule_id"`
	Port          int       `json:"port"`
	Timestamp     time.Time `json:"timestamp"`
	IP            string    `json:"ip"`
	ID            string    `json:"id"`
	RequestMethod string    `json:"request_method"`
	RequestURI    string    `json:"request_uri"`
	Action        string    `json:"action"`
	NextAction    string    `json:"next_action"`
	Message       string    `json:"message"`
	Protocol      string    `json:"protocol"`
	RuleName      string    `json:"rule_name"`
	RuleDesc      string    `json:"rule_desc"`
	Request       string    `json:"request"` //请求报文
}
