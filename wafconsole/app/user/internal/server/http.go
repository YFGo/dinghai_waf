package server

import (
	"context"
	"log/slog"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "wafconsole/api/user/v1"
	"wafconsole/app/user/internal/conf"
	is "wafconsole/app/user/internal/service"
	"wafconsole/app/user/internal/service/ginuser"
	"wafconsole/utils/const/user"
	up "wafconsole/utils/plugin"
	utils "wafconsole/utils/resp"
)

// NewWhiteListMatcher 初始化白名单接口
func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.user.v1.WafUser/CreateWafUser"] = struct{}{}
	whiteList["/api.user.v1.WafUser/Login"] = struct{}{}
	whiteList["/api.user.v1.Common/CreateNewToken"] = struct{}{}
	whiteList["/api.user.v1.Common/GetCaptcha"] = struct{}{}
	whiteList["/api.user.v1.Common/VerifyCaptcha"] = struct{}{}
	whiteList["/api.user.v1.Common/SendCode"] = struct{}{}
	whiteList["/api.user.v1.Common/SseConnect"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *is.WafUserService, commonService *is.CommonService, logger log.Logger) *http.Server {
	protoValidate, err := up.NewValidate()
	if err != nil {
		slog.Error("protoValidate", err)
		return nil
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			up.MiddlewareCors(),
			tracing.Server(),
			protoValidate.ValidateUnaryServerInterceptor(), //参数校验
			selector.Server(
				up.JWTMiddleware(), //token验证
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.ResponseEncoder(utils.ResponseEncoder),
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
	v1.RegisterWafUserHTTPServer(srv, userService)
	v1.RegisterCommonHTTPServer(srv, commonService)
	// 集成sse接口
	srv.HandlePrefix(user.GinRouterSse, &ginuser.SSEHandlerWrapper{})
	return srv
}
