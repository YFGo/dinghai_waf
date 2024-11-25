package service

import (
	"github.com/google/wire"
	"wafconsole/app/dashBorad/internal/service/view"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(view.NewDataViewService)
