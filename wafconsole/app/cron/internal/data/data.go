package data

import (
	"github.com/tx7do/kratos-transport/broker"
	"wafconsole/app/cron/internal/data/discovery"
	"wafconsole/app/cron/internal/data/mq"

	v1 "wafconsole/api/wafTop/v1"

	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, discovery.NewDiscovery, mq.NewKafkaBroker,
	discovery.NewSiteServerRpc, NewNormalHttpRepo)

// Data .
type Data struct {
	log *log.Helper

	kafkaClient broker.Broker

	siteServerRpc v1.ServerClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, kafkaClient broker.Broker,
	siteServerRpc v1.ServerClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log:           l,
		kafkaClient:   kafkaClient,
		siteServerRpc: siteServerRpc,
	}, cleanup, nil
}
