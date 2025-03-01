package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"

	"wafconsole/app/mqconsume/internal/conf"
	"wafconsole/app/mqconsume/internal/service/notice"
)

// NewRabbitMQServer create a rabbitmq server.
func NewRabbitMQServer(cfg *conf.Data, _ log.Logger, mtc *notice.MsgNoticeService) *rabbitmq.
	Server {
	ctx := context.Background()

	srv := rabbitmq.NewServer(
		rabbitmq.WithGlobalTracerProvider(),
		rabbitmq.WithGlobalPropagator(),
		rabbitmq.WithCodec("json"),
		rabbitmq.WithAddress(cfg.Rabbitmq.Addrs),
		rabbitmq.WithExchange(cfg.Rabbitmq.Exchange, cfg.Rabbitmq.DurableExchange),
	)
	// 注册订阅者
	_ = rabbitmq.RegisterSubscriber(srv, ctx, cfg.Rabbitmq.Routing, mtc.SaveUser,
		broker.WithQueueName("user"))

	return srv
}

// NewKafkaServer 创建kafka服务
func NewKafkaServer(cfg *conf.Data, _ log.Logger, mtc *notice.MsgNoticeService) *kafka.Server {
	ctx := context.Background()
	srv := kafka.NewServer(
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
		kafka.WithCodec("json"),
		kafka.WithAddress(cfg.Kafka.Brokers),
	)

	// 注册订阅者
	_ = kafka.RegisterSubscriber(srv, ctx,
		cfg.Kafka.Topic, // topic 参数
		cfg.Kafka.Group, // queue 参数（消费者组）
		false,           // disableAutoAck（自动提交偏移量）
		mtc.SaveUser,    // 消息处理函数
	)

	return srv
}
