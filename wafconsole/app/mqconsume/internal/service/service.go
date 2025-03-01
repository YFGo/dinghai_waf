package service

import (
	"github.com/google/wire"

	"wafconsole/app/mqconsume/internal/service/notice"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(notice.NewNoticeMsgService)
