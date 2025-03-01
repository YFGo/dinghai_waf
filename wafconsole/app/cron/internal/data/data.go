package data

import (
	"github.com/tx7do/kratos-transport/broker"
	"wafconsole/app/cron/internal/data/discovery"

	v1 "wafconsole/api/wafTop/v1"

	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, discovery.NewDiscovery, NewRabbitMQBroker,
	discovery.NewSiteServerRpc, NewUserRpcRepo, NewMqPushRepo)

// Data .
type Data struct {
	log *log.Helper

	rabbitmqBroker broker.Broker

	siteServerRpc v1.ServerClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rabbitmqBroker broker.Broker,
	siteServerRpc v1.ServerClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log:            l,
		rabbitmqBroker: rabbitmqBroker,
		siteServerRpc:  siteServerRpc,
	}, cleanup, nil
}
