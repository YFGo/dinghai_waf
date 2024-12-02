package plugins

import (
	"context"
	"log/slog"
	"wafCoraza/biz"
	"wafCoraza/data/types"
)

// AllowHandle 白名单处理事务
func AllowHandle(allowUc *biz.WafAllowListUsecase, ctx context.Context, host string, reqURI string, reqIP string) (error, bool) {
	allowIdList, err := allowUc.GetAllowInfo(ctx, host) // 获取此站点应用的白名单
	if err != nil {
		slog.ErrorContext(ctx, "allow_handle get_allow_info is failed", err)
		return err, false
	}
	allowDetails := allowUc.GetAllowsDetail(ctx, allowIdList, host) //根据白名单id 获取白名单信息
	for _, allowInfo := range allowDetails {
		switch allowInfo.Key {
		case types.URI:
			if allowInfo.Value == reqURI {
				return nil, true
			}
		case types.IP:
			if allowInfo.Value == reqIP {
				return nil, true
			}
		}
	}
	return nil, false
}
