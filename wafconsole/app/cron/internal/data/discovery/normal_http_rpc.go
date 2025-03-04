package discovery

import (
	"context"
	"wafconsole/utils/const/waftop"

	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "wafconsole/api/wafTop/v1"
)

// NewNormalHttpRpc 正常流量服务
func NewNormalHttpRpc(r registry.Discovery) v1.NormalHttpClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(waftop.WafTopRpc),
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
	return v1.NewNormalHttpClient(conn)
}
