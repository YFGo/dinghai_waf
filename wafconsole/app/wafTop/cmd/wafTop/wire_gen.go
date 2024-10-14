// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/conf"
	"wafconsole/app/wafTop/internal/data"
	"wafconsole/app/wafTop/internal/server"
	"wafconsole/app/wafTop/internal/service/site"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
//
//go:generate wire
func wireApp(confServer *conf.Server, bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confServer, bootstrap)
	if err != nil {
		return nil, nil, err
	}
	wafAppRepo := data.NewAppWafRepo(dataData, logger)
	wafAppUsecase := siteBiz.NewGreeterUsecase(wafAppRepo, logger)
	wafAppService := service.NewWafAppService(wafAppUsecase)
	serverRepo := data.NewServerRepo(dataData)
	serverUsecase := siteBiz.NewServerUsecase(serverRepo)
	serverService := service.NewServerService(serverUsecase)
	grpcServer := server.NewGRPCServer(confServer, wafAppService, serverService, logger)
	httpServer := server.NewHTTPServer(confServer, wafAppService, serverService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
