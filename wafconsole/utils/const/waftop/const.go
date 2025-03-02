package waftop

// 数据库迁移相关
const (
	WafTop           = "waf_top"
	MigratePathMysql = "wafconsole/app/wafTop/migrations/mysql"
	MigratePathCk    = "wafconsole/app/wafTop/migrations/clickhouse"
)

// gRPC服务调用名
const (
	WafTopRpc = "/microservices/wafTop/wafTop"
)

const (
	NormalHttpTopic = "normal_http"
)
