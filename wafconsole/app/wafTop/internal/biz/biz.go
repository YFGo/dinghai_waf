package biz

import (
	"github.com/google/wire"
	"wafconsole/app/wafTop/internal/biz/rule"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(siteBiz.NewGreeterUsecase, siteBiz.NewServerUsecase, ruleBiz.NewBuildRuleUsecase, ruleBiz.NewRuleGroupUsecase, ruleBiz.NewUserRuleUsecase)
