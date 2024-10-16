//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/conf"
	"wafconsole/app/user/internal/data"
	"wafconsole/app/user/internal/server"
	"wafconsole/app/user/internal/service"

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