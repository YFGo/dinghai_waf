package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	consulapi "github.com/hashicorp/consul/api"
	"wafconsole/app/dashBorad/internal/conf"
)

// initConfig init the config
func initConfig() config.Config {
	dirverC := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/config_driver.yaml", flagconf)),
		),
	)
	defer dirverC.Close()

	if err := dirverC.Load(); err != nil {
		panic(err)
	}
	driverConfig := conf.ConfDriver{}
	if err := dirverC.Scan(&driverConfig); err != nil {
		panic(err)
	}
	switch driverConfig.Type {
	case "file":
		// 读取本地文件
		return config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
	case "consul":
		if driverConfig.Consul == nil {
			panic("consul config is nil")
		}
		// 读取consul配置
		consulClient, err := consulapi.NewClient(&consulapi.Config{
			Address:    driverConfig.Consul.Address,
			Scheme:     driverConfig.Consul.Scheme,
			PathPrefix: driverConfig.Consul.PathPrefix,
			Token:      driverConfig.Consul.Token,
		})
		if err != nil {
			panic(err)
		}
		cs, err := consul.New(consulClient, consul.WithPath(driverConfig.Consul.Path))
		if err != nil {
			panic(err)
		}
		return config.New(config.WithSource(cs))
	case "etcd":
		panic("not support")
	default:
		panic("unknown config driver")
	}
}
