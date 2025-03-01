package discovery

import (
	"fmt"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"go.etcd.io/etcd/client/v3"

	"wafconsole/app/cron/internal/conf"
)

// NewDiscovery 初始化注册中心
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	switch conf.Type {
	case "consul":
		if conf.Consul == nil {
			panic("consul config is nil")
		}
		// 读取consul配置
		cli, err := api.NewClient(&api.Config{Address: conf.Consul.Address})
		if err != nil {
			panic(err)
		}
		// 创建consul注册中心
		return consul.New(cli, consul.WithHealthCheck(false))
	case "etcd":
		if conf.Etcd == nil {
			panic("etcd config is nil")
		}
		etcdCfg := clientv3.Config{
			Endpoints: []string{conf.Etcd.Address},
		}

		etcdClient, err := clientv3.New(etcdCfg)
		if err != nil {
			panic(fmt.Sprintf("failed to create Etcd client: %v", err))
		}

		etcdRegistry := etcd.New(etcdClient)
		return etcdRegistry
	default:
		panic("unknown registry driver")
	}
}
