package biz

import (
	"github.com/google/wire"

	"wafconsole/app/mqconsume/internal/biz/notice"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(notice.NewUserNoticeUsecase)
