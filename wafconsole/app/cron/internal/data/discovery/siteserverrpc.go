package discovery

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "wafconsole/api/wafTop/v1"
)

// NewSiteServerRpc 站点服务
func NewSiteServerRpc(r registry.Discovery) v1.ServerClient {
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
	return v1.NewServerClient(conn)
}
