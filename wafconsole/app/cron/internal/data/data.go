package data

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"

	"wafconsole/api/grpc_user"
	"wafconsole/app/cron/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewRabbitMQBroker, NewUserRpcClient, NewUserRpcRepo, NewMqPushRepo)

// Data .
type Data struct {
	log *log.Helper

	rabbitmqBroker broker.Broker

	guClient grpc_user.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rabbitmqBroker broker.Broker, uc grpc_user.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log:            l,
		rabbitmqBroker: rabbitmqBroker,
		guClient:       uc,
	}, cleanup, nil
}

func NewUserRpcClient(r registry.Discovery) grpc_user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///wafconsole.grpc_user"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			// 异常恢复
			recovery.Recovery(),
			//熔断器
			circuitbreaker.Client(),
			//链路追踪
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return grpc_user.NewUserClient(conn)
}
