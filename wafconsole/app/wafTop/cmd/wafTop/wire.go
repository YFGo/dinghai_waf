//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/registry"
	"wafconsole/app/wafTop/internal/biz"
	"wafconsole/app/wafTop/internal/conf"
	"wafconsole/app/wafTop/internal/data"
	"wafconsole/app/wafTop/internal/server"
	"wafconsole/app/wafTop/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
//
//go:generate wire
func wireApp(*conf.Server, *conf.Bootstrap, log.Logger, registry.Registrar) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
