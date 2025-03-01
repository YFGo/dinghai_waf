package data

import (
	"wafconsole/app/cron/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

// initRegistry init the registry
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
	default:
		panic("unknown registry driver")
	}
}
