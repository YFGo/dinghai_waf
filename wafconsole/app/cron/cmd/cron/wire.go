//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"wafconsole/app/cron/internal/biz"
	"wafconsole/app/cron/internal/conf"
	"wafconsole/app/cron/internal/data"
	"wafconsole/app/cron/internal/server"
	"wafconsole/app/cron/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger, registry.Registrar) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
