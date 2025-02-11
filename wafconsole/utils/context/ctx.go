package utils

import (
	"context"
	"time"
)

func GetAppCtx(ctx context.Context) (appCtx context.Context) {
	// 定义上下文，设置超时时间
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	appCtx = ctx

	return appCtx
}
