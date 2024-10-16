package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "wafconsole/api/wafTop/v1"
	"wafconsole/app/wafTop/internal/conf"
	rule "wafconsole/app/wafTop/internal/service/rule"
	site "wafconsole/app/wafTop/internal/service/site"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, wafApp *site.WafAppService, serverWaf *site.ServerService, buildRule *rule.BuildRuleService, ruleGroup *rule.RuleGroupService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterWafAppServer(srv, wafApp)
	v1.RegisterServerServer(srv, serverWaf)
	v1.RegisterBuildRuleServer(srv, buildRule)
	v1.RegisterRuleGroupServer(srv, ruleGroup)
	return srv
}