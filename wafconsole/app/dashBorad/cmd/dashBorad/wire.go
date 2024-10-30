//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"wafconsole/app/dashBorad/internal/biz"
	"wafconsole/app/dashBorad/internal/conf"
	"wafconsole/app/dashBorad/internal/data"
	"wafconsole/app/dashBorad/internal/server"
	"wafconsole/app/dashBorad/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
//
//go:generate wire
func wireApp(*conf.Server, *conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
