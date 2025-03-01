//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"wafconsole/app/mqconsume/internal/biz"
	"wafconsole/app/mqconsume/internal/conf"
	"wafconsole/app/mqconsume/internal/data"
	"wafconsole/app/mqconsume/internal/server"
	"wafconsole/app/mqconsume/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
//
//go:generate wire
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger, registry.Registrar) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
