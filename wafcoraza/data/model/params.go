package model

import (
	"github.com/corazawaf/coraza/v3"
	"time"
)

// WAFStrategy WAF策略
type WAFStrategy struct {
	ID           string `json:"id"`
	SeclangRules string `json:"seclang_rules"`
}

type WafConfig struct {
	ID              int     `json:"id"`
	Action          uint8   `json:"action"`
	NextAction      uint8   `json:"next_action"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	RuleGroupIdList []int64 `json:"rule_group_id_list"`
}

type RuleGroup struct {
	ID         int     `json:"id"`
	IsBuildin  uint8   `json:"is_buildin"`
	RuleIDList []int64 `json:"rule_id_list"`
}

type Rule struct {
	ID        int    `json:"id"`
	RiskLevel uint8  `json:"risk_level"`
	Seclang   string `json:"seclang"`
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

type Allow struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AllowAction struct {
	Goal    string `json:"goal"`
	Content string `json:"content"`
}

// ListNormalHttp 正常请求列表
type ListNormalHttp struct {
	ID            string    `json:"id"`
	IP            string    `json:"ip"`
	RequestURI    string    `json:"request_uri"`
	RequestTime   time.Time `json:"request_time"`
	RequestMethod string    `json:"request_method"`
	Protocol      string    `json:"protocol"`
	RequestBody   string    `json:"request_body"`
}
