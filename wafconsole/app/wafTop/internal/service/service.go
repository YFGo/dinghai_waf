package service

import (
	"github.com/google/wire"
	rule "wafconsole/app/wafTop/internal/service/rule"
	site "wafconsole/app/wafTop/internal/service/site"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(site.NewWafAppService, site.NewServerService, rule.NewBuildRuleService, rule.NewRuleGroupService, rule.NewUserRuleService)
