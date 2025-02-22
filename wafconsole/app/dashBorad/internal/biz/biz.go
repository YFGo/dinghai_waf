package biz

import (
	"github.com/google/wire"
	"wafconsole/app/dashBorad/internal/biz/viewLogic"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(viewLogic.NewDataViewUsecase)
