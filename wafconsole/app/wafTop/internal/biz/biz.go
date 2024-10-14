package biz

import (
	"github.com/google/wire"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(siteBiz.NewGreeterUsecase, siteBiz.NewServerUsecase)
