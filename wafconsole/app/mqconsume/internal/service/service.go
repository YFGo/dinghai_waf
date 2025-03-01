package service

import (
	"github.com/google/wire"

	"wafconsole/app/mqconsume/internal/service/usermsg"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(usermsg.NewNoticeMsgService)
