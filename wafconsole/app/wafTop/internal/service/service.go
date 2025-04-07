package service

import (
	"github.com/google/wire"

	allow "wafconsole/app/wafTop/internal/service/allow"
	sn "wafconsole/app/wafTop/internal/service/normalhttp"
	rule "wafconsole/app/wafTop/internal/service/rule"
	site "wafconsole/app/wafTop/internal/service/site"
	strategy "wafconsole/app/wafTop/internal/service/strategy"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(site.NewWafAppService, site.NewServerService,
	rule.NewBuildRuleService, rule.NewRuleGroupService, rule.NewUserRuleService,
	strategy.NewStrategyService, allow.NewAllowListService, sn.NewNormalHttpService, site.NewServerPrincipalService)
