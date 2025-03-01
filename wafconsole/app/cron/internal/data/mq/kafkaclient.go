package mq

import (
	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"
)

func NewKafkaBroker(cfg *conf.Data, logger log.Logger) broker.Broker {
	l := log.NewHelper(log.With(logger, "module", "data", "kafka")) // 6. 修改日志标识

	// 2. 使用kafka.NewBroker
	b := kafka.NewBroker(
		broker.WithAddress(cfg.Kafka.Addrs...), // 假设配置结构已改为Kafka
		broker.WithCodec(cfg.Kafka.Codec),
	)

	if b == nil {
		l.Errorf("kafka broker is nil")
		return nil
	}

	_ = b.Init()

	if err := b.Connect(); err != nil {
		l.Errorf("kafka connect error: %v", err) // 6. 修改错误信息
		return nil
	}

	return b
}
