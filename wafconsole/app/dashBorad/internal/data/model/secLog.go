package model

import "time"

const AttackEventsTopic = "attack_events"

type SecLog struct {
	LogID           string    `db:"log_id"`            // 日志唯一id
	URI             string    `db:"uri"`               // uri
	TransactionID   string    `db:"transaction_id"`    // 请求id
	Ctime           time.Time `db:"ctime"`             // 发生时间
	Protocol        string    `db:"protocol"`          // 协议类型
	Request         string    `db:"request"`           // 请求报文
	RequestMethod   string    `db:"request_method"`    // 请求方法
	Response        string    `db:"response"`          // 响应报文
	ClientIP        string    `db:"client_ip"`         // 源ip
	RuleName        string    `db:"rule_name"`         // 规则名称
	RuleDesc        string    `db:"rule_desc"`         // 规则描述
	RuleGroupName   string    `db:"rule_group_name"`   // 特征组名称
	ServerGroupName string    `db:"server_group_name"` // 服务器组名称
	ServerName      string    `db:"server_name"`       // 服务器名称
	AppName         string    `db:"app_name"`          // 应用程序名称
	Strategy        string    `db:"strategy"`          // 策略名称
	Action          string    `db:"action"`            // 动作
	NextAction      string    `db:"next_action"`       // 后续动作
	ClientPort      int32     `db:"client_port"`       // 源端口
	AlertLevel      int32     `db:"alert_level"`       // 告警级别
	ServerGroupID   int32     `db:"server_group_id"`   // 服务器组id
	RuleGroupID     int32     `db:"rule_group_id"`     // 特征组id
	RuleID          int       `db:"rule_id"`           //命中的规则id
}

func (SecLog) TableName() string {
	return `sec_log`
}
