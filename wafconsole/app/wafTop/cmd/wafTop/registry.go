package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"go.etcd.io/etcd/client/v3" // 使用新的 etcd 客户端库
	"wafconsole/app/wafTop/internal/conf"
)

func initRegistryConf() *conf.Registry {
	registryCfg := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/registry.yaml", flagconf)),
		),
	)
	defer registryCfg.Close()

	if err := registryCfg.Load(); err != nil {
		panic(fmt.Sprintf("failed to load registry config: %v", err))
	}

	registryConfig := &conf.Registry{}
	if err := registryCfg.Scan(registryConfig); err != nil {
		panic(fmt.Sprintf("failed to scan registry config: %v", err))
	}

	return registryConfig
}

func initRegistry(conf *conf.Registry) registry.Registrar {
	switch conf.Type {
	case "consul":
		if conf.Consul == nil {
			panic("consul config is nil")
		}

		consulClient, err := api.NewClient(&api.Config{
			Address: conf.Consul.Address,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to create Consul client: %v", err))
		}

		consulRegistry := consul.New(consulClient, consul.WithHealthCheck(false))
		return consulRegistry
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
		panic(fmt.Sprintf("unknown registry driver: %s", conf.Type))
	}
}
