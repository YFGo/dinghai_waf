package data

import (
	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

func NewRabbitMQBroker(cfg *conf.Data, logger log.Logger) broker.Broker {
	l := log.NewHelper(log.With(logger, "module", "data", "rabbitmq"))
	b := rabbitmq.NewBroker(
		broker.WithAddress(cfg.Rabbitmq.Addrs...),
		broker.WithCodec(cfg.Rabbitmq.Codec),
		rabbitmq.WithExchangeName(cfg.Rabbitmq.Exchange),
		rabbitmq.WithDurableExchange(),
	)
	if b == nil {
		l.Errorf("rabbitmq broker is nil")
		return nil
	}

	_ = b.Init()

	if err := b.Connect(); err != nil {
		l.Errorf("rabbitmq connect error: %v", err)
		return nil
	}

	return b
}
