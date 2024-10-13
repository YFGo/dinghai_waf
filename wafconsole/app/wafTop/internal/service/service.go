package service

import (
	"github.com/google/wire"
	"wafconsole/app/wafTop/internal/service/site"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(site.NewWafAppService)
