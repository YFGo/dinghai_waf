package biz

import (
	"github.com/google/wire"
	"wafconsole/app/wafTop/internal/biz/allow"
	"wafconsole/app/wafTop/internal/biz/rule"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	strategyBiz "wafconsole/app/wafTop/internal/biz/strategy"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(siteBiz.NewGreeterUsecase, siteBiz.NewServerUsecase, ruleBiz.NewBuildRuleUsecase, ruleBiz.NewRuleGroupUsecase, ruleBiz.NewUserRuleUsecase, strategyBiz.NewWafStrategyUsecase, allow.NewListAllowUsecase)
