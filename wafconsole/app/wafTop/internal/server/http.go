package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"log/slog"
	v1 "wafconsole/api/wafTop/v1"
	"wafconsole/app/wafTop/internal/conf"
	"wafconsole/app/wafTop/internal/server/plugin"
	rule "wafconsole/app/wafTop/internal/service/rule"
	site "wafconsole/app/wafTop/internal/service/site"
	strategy "wafconsole/app/wafTop/internal/service/strategy"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, appWafTop *site.WafAppService, serverWaf *site.ServerService, buildRule *rule.BuildRuleService, ruleGroup *rule.RuleGroupService, userRule *rule.UserRuleService, strategyHttp *strategy.StrategyService, logger log.Logger) *http.Server {
	protoValidate, err := plugin.NewValidate()
	if err != nil {
		slog.Error("protoValidate", err)
		return nil
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			protoValidate.ValidateUnaryServerInterceptor(), //参数校验
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	v1.RegisterWafAppHTTPServer(srv, appWafTop)
	v1.RegisterServerHTTPServer(srv, serverWaf)
	v1.RegisterBuildRuleHTTPServer(srv, buildRule)
	v1.RegisterRuleGroupHTTPServer(srv, ruleGroup)
	v1.RegisterUserRuleHTTPServer(srv, userRule)
	v1.RegisterStrategyHTTPServer(srv, strategyHttp)
	return srv
}
