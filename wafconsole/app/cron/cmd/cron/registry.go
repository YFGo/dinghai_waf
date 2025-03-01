package main

import (
	"fmt"

	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/hashicorp/consul/api"
)

// initRegistryConf init the registry config
func initRegistryConf() *conf.Registry {
	registry := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/registry.yaml", flagconf)),
		),
	)
	defer registry.Close()

	if err := registry.Load(); err != nil {
		panic(err)
	}
	registryConfig := &conf.Registry{}
	if err := registry.Scan(registryConfig); err != nil {
		panic(err)
	}
	return registryConfig
}

// initRegistry init the registry
func initRegistry(conf *conf.Registry) registry.Registrar {
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
