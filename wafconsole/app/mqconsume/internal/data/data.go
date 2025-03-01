package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	"wafconsole/app/mqconsume/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewNoticeUserRepo)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: l,
	}, cleanup, nil
}

// NewUserRpcClient . 新建用户rpc客户端
func NewUserRpcClient(r registry.Discovery) grpc_user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos-project.grpc_user"),
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
