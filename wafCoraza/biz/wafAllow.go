package biz

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"strings"
	"wafCoraza/data/model"
	"wafCoraza/data/types"
)

type WafAllowListRepo interface {
	GetValueInfoByKey(ctx context.Context, key string) (string, error)
	GetAllowListByPrefix(ctx context.Context, prefix string) ([]model.AllowAction, error)
	WatchAllowList(ctx context.Context, prefix string) clientv3.WatchChan
}

type WafAllowListUsecase struct {
	repo         WafAllowListRepo
	allowListMap map[string]model.Allow //key 为白名单id , value为白名单信息
}

func NewWafAllowListUsecase(repo WafAllowListRepo) *WafAllowListUsecase {
	return &WafAllowListUsecase{
		repo:         repo,
		allowListMap: map[string]model.Allow{},
	}
}

// InitAllowList 初始化白名单
func (w *WafAllowListUsecase) InitAllowList(ctx context.Context) {
	allowList, err := w.repo.GetAllowListByPrefix(ctx, types.AllowPrefix)
	if err != nil {
		slog.ErrorContext(ctx, "InitAllowList err : %v", err)
		return
	}
	for _, allowValue := range allowList {
		var allow model.Allow
		if err = json.Unmarshal([]byte(allowValue.Content), &allow); err != nil {
			slog.ErrorContext(ctx, "InitAllowList err : %v", err)
			continue
		}
		w.allowListMap[allowValue.Goal] = allow
	}

}

// GetAllowInfo 获取站点应用的白名单id列表
func (w *WafAllowListUsecase) GetAllowInfo(ctx context.Context, host string) ([]string, error) {
	allowIdValue, err := w.repo.GetValueInfoByKey(ctx, host+types.AllowSuffix)
	if err != nil {
		slog.ErrorContext(ctx, "get_allow_info from etcd is failed: ", err)
		return nil, err
	}
	allowIdList := strings.Split(allowIdValue, types.CutOFF)
	return allowIdList, nil
}

// GetAllowsDetail 根据id 在map集合中查询白名单详细信息
func (w *WafAllowListUsecase) GetAllowsDetail(ctx context.Context, allowIdList []string) []model.Allow {
	var allowList []model.Allow
	for _, allowId := range allowIdList {
		allow, ok := w.allowListMap[allowId]
		if ok {
			allowList = append(allowList, allow)
		}
	}
	return allowList
}

// WatchAllowChange 监听白名单的变化
func (w *WafAllowListUsecase) WatchAllowChange(ctx context.Context) {
	watchChan := w.repo.WatchAllowList(ctx, types.AllowPrefix)
	for {
		select {
		case watchResp := <-watchChan:
			for _, event := range watchResp.Events {
				keyArr := strings.Split(string(event.Kv.Key), types.CutOFF)
				switch event.Type {
				case clientv3.EventTypePut: // 更新或者新增
					var allow model.Allow
					if err := json.Unmarshal(event.Kv.Value, &allow); err != nil {
						slog.ErrorContext(ctx, "watch_allow_change err : %v", err)
						continue
					}
					w.allowListMap[keyArr[1]] = allow
				case clientv3.EventTypeDelete: //删除
					delete(w.allowListMap, keyArr[1])
				}
			}
		}
	}
}
