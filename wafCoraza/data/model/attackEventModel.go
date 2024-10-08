package model

import "time"

const AttackEventLogTopic = "attack_events"

type AttackEventKey struct {
	AttackEvents []AttackEvent `json:"attack_events"`
}

// AttackEvent 定义攻击事件的结构
type AttackEvent struct {
	RuleId        int       `json:"rule_id" csv:"rule_id"`
	Port          int       `json:"port" csv:"port"`
	Timestamp     time.Time `json:"timestamp" csv:"timestamp"`
	IP            string    `json:"ip" csv:"ip"`
	ID            string    `json:"id" csv:"id"`
	RequestMethod string    `json:"request_method" csv:"request_method"`
	RequestURI    string    `json:"request_uri" csv:"request_uri"`
	Action        string    `json:"action" csv:"action"`
	NextAction    string    `json:"next_action" csv:"next_action"`
	Message       string    `json:"message" csv:"message"`
	Protocol      string    `json:"protocol" csv:"protocol"`
	RuleName      string    `json:"rule_name" csv:"rule_name"`
	RuleDesc      string    `json:"rule_desc" csv:"rule_desc"`
	Request       string    `json:"request" csv:"request"` //请求报文
}
