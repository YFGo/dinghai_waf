package hooks

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
)

// InitEtcd 在每一次平台服务启动时 , 初始化etcd中的值
func InitEtcd(etcd *clientv3.Client) {
	if etcd == nil {
		return
	}
	// 定义上下文
	ctx := context.Background()

	// 定义要插入的键值对
	keys := []string{"rule_90001", "rule_90002", "rule_90003"}
	values := []string{`{"id":90001,"risk_level":4,"seclang":"Include wafCoraza/ruleset/coraza.conf"}`, `{"id":90002,"risk_level":4,"seclang":"Include wafCoraza/ruleset/coreruleset/crs-setup.conf.example"}`, `{"id":90003,"risk_level":4,"seclang":"Include wafCoraza/ruleset/coreruleset/rules/*.conf"}`}
	// 遍历键值对
	for i, key := range keys {
		// 获取键的当前值
		resp, err := etcd.KV.Get(ctx, key)
		if err != nil {
			slog.Error("etcd.KV.Get error: %v", err)
		}
		// 检查键是否存在
		if len(resp.Kvs) == 0 {
			// 键不存在，插入键值对
			_, err = etcd.KV.Put(ctx, key, values[i])
			if err != nil {
				slog.Error("etcd.KV.Get error: %v", err)
			}
		}
	}
}
