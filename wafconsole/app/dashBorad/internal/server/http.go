package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "wafconsole/api/dashBorad/v1"
	"wafconsole/app/dashBorad/internal/conf"
	"wafconsole/app/dashBorad/internal/service/view"
	up "wafconsole/utils/plugin"
	utils "wafconsole/utils/resp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, dataView *view.DataViewService, logger log.Logger) *http.Server {
	protoValidate, err := up.NewValidate()
	if err != nil {
		panic(err)
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			protoValidate.ValidateUnaryServerInterceptor(),
			up.MiddlewareCors(),
			up.JWTMiddleware(),
		),
		http.ResponseEncoder(utils.ResponseEncoder),
		http.ErrorEncoder(utils.ErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterDataViewHTTPServer(srv, dataView)
	return srv
}
