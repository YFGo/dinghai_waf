package biz

import (
	"github.com/google/wire"

	"wafconsole/app/cron/internal/biz/normalhttp"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(normalhttp.NewUsercaseNormalHttp)
