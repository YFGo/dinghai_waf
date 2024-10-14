package service

import (
	"github.com/google/wire"
	service "wafconsole/app/wafTop/internal/service/site"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(service.NewWafAppService, service.NewServerService)
