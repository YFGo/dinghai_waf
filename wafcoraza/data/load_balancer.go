package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// RoundRobinBalancer 负载均衡器，实现轮询策略
type RoundRobinBalancer struct {
	addresses []string
	index     uint32
	mu        sync.RWMutex
}

func (b *RoundRobinBalancer) Update(addresses []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.addresses = addresses
	atomic.StoreUint32(&b.index, 0) // 重置索引
}

func (b *RoundRobinBalancer) Next() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if len(b.addresses) == 0 {
		return ""
	}
	next := atomic.AddUint32(&b.index, 1) - 1
	return b.addresses[next%uint32(len(b.addresses))]
}

// BalancerManager 管理所有URI对应的负载均衡器
type BalancerManager struct {
	etcdClient *clientv3.Client
	balancers  sync.Map // uri -> *RoundRobinBalancer
}

func (bm *BalancerManager) GetBalancer(uri string) (*RoundRobinBalancer, error) {
	// 检查是否已有负载均衡器
	if b, ok := bm.balancers.Load(uri); ok {
		return b.(*RoundRobinBalancer), nil
	}

	// 从etcd获取地址列表
	resp, err := bm.etcdClient.Get(context.Background(), uri)
	if err != nil {
		return nil, fmt.Errorf("etcd get error: %v", err)
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("no etcd key for URI: %s", uri)
	}

	var addresses []string
	if err := json.Unmarshal(resp.Kvs[0].Value, &addresses); err != nil {
		return nil, fmt.Errorf("failed to parse addresses: %v", err)
	}

	// 创建负载均衡器并保存
	balancer := &RoundRobinBalancer{addresses: addresses}
	bm.balancers.Store(uri, balancer)

	// 监听该URI的变化
	go bm.watchUpdates(uri)

	return balancer, nil
}

// watchUpdates 监听etcd中特定URI的变更
func (bm *BalancerManager) watchUpdates(uri string) {
	watchChan := bm.etcdClient.Watch(context.Background(), uri)
	for resp := range watchChan {
		for _, event := range resp.Events {
			if event.Type == clientv3.EventTypePut {
				var addresses []string
				if err := json.Unmarshal(event.Kv.Value, &addresses); err != nil {
					log.Printf("Failed to update addresses for %s: %v", uri, err)
					continue
				}
				if balancer, ok := bm.balancers.Load(uri); ok {
					balancer.(*RoundRobinBalancer).Update(addresses)
					log.Printf("Updated addresses for URI %s: %v", uri, addresses)
				}
			} else if event.Type == clientv3.EventTypeDelete {
				bm.balancers.Delete(uri)
				log.Printf("Removed balancer for URI %s", uri)
			}
		}
	}
}
