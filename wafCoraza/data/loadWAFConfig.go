package data

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"wafCoraza/biz"
	"wafCoraza/data/model"
	"wafCoraza/data/types"
)

type loadWAFConfigRepo struct {
	data *Data
}

func NewLoadWAFConfigRepo(data *Data) biz.WafConfigRepo {
	return &loadWAFConfigRepo{data: data}
}

// GetAppForStrategyIDs 获取app应用的策略
func (l loadWAFConfigRepo) GetAppForStrategyIDs(appAddr string) ([]string, error) {
	strategyResp, err := l.data.etcdClient.Get(context.Background(), appAddr)
	if err != nil {
		return nil, err
	}
	var strategyIDsStr []string
	for _, kv := range strategyResp.Kvs {
		strategyIDsStr = append(strategyIDsStr, string(kv.Value))
	}
	return strategyIDsStr, nil
}

// GetAllSeclangRules 获取所有策略对应的安全规则
func (l loadWAFConfigRepo) GetAllSeclangRules() ([]model.WAFStrategy, error) {
	seclangRulesResp, err := l.data.etcdClient.KV.Get(context.Background(), types.StrategyKey, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var seclangRules []model.WAFStrategy
	for _, kv := range seclangRulesResp.Kvs {
		keyLoad := strings.Split(string(kv.Key), types.CutOFF)
		seclangRule := model.WAFStrategy{
			ID:           keyLoad[1],
			SeclangRules: string(kv.Value),
		}
		seclangRules = append(seclangRules, seclangRule)
	}
	return seclangRules, nil
}

// GetRealAddr 根据etcd中的域名+端口 , 获取其真实访问的地址
func (l loadWAFConfigRepo) GetRealAddr(domain string) (string, error) {
	var realAddr string
	realAddrResp, err := l.data.etcdClient.KV.Get(context.Background(), domain+types.RealAddr)
	if err != nil {
		return realAddr, err
	}
	for _, kv := range realAddrResp.Kvs {
		realAddr = string(kv.Value)
	}
	return realAddr, nil
}

// AirUpdateStrategy 热更新策略
func (l loadWAFConfigRepo) AirUpdateStrategy() clientv3.WatchChan {
	return l.data.etcdClient.Watch(context.Background(), types.StrategyKey, clientv3.WithPrefix())
}
