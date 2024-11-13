package data

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"wafCoraza/biz"
	"wafCoraza/data/model"
	"wafCoraza/data/types"
)

type wafAllowListRepo struct {
	data *Data
}

func NewWafAllowListRepo(data *Data) biz.WafAllowListRepo {
	return &wafAllowListRepo{data: data}
}

// GetValueInfoByKey 根据key 获取对应的value
func (w wafAllowListRepo) GetValueInfoByKey(ctx context.Context, key string) (string, error) {
	var res string
	valueResp, err := w.data.etcdClient.KV.Get(ctx, key)
	if err != nil {
		return res, err
	}
	for _, kv := range valueResp.Kvs {
		res = string(kv.Value)
	}
	return res, nil
}

// GetAllowListByPrefix 根据前缀 , 获取所有白名单信息
func (w wafAllowListRepo) GetAllowListByPrefix(ctx context.Context, prefix string) ([]model.AllowAction, error) {
	resp, err := w.data.etcdClient.KV.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var allowList []model.AllowAction
	for _, kv := range resp.Kvs {
		keyArr := strings.Split(string(kv.Key), types.CutOFF)
		allow := model.AllowAction{
			Goal:    keyArr[1],
			Content: string(kv.Value),
		}
		allowList = append(allowList, allow)
	}
	return allowList, nil
}

// WatchAllowList 监视白名单的变化
func (w wafAllowListRepo) WatchAllowList(ctx context.Context, prefix string) clientv3.WatchChan {
	return w.data.etcdClient.Watch(ctx, prefix, clientv3.WithPrefix())
}
