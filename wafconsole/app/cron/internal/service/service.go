package service

import (
	"github.com/google/wire"
	"wafconsole/app/cron/internal/service/normalhttp"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewJobService, normalhttp.NewServiceNormalHttp)
