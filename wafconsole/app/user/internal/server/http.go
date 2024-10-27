package server

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	stdhttp "net/http"
	v1 "wafconsole/api/user/v1"
	"wafconsole/app/user/internal/conf"
	"wafconsole/app/user/internal/server/plugin"
	"wafconsole/app/user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Reply struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v any) error {
	data, err := json.Marshal(&Reply{
		Code: 0,
		Msg:  "",
		Data: v,
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return nil
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	e := errors.FromError(err)
	code := int(e.Code)
	rep := &Reply{
		Code: e.Code,
		Msg:  e.Message,
		Data: e.Metadata,
	}

	body, err := json.Marshal(rep)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	//参数错误 , 单独处理
	if codes.InvalidArgument == status.Code(e) {
		w.WriteHeader(stdhttp.StatusOK)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if 0 < code && code <= 600 {
		w.WriteHeader(code)
	} else {
		w.WriteHeader(stdhttp.StatusOK)
	}
	w.Write(body)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.WafUserService, logger log.Logger) *http.Server {
	protoValidate, err := plugin.NewValidate()
	if err != nil {
		slog.Error("protoValidate", err)
		return nil
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			protoValidate.ValidateUnaryServerInterceptor(), //参数校验
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(c.Jwt.Key), nil
			}),
		),
		http.ErrorEncoder(ErrorEncoder),
		http.ResponseEncoder(ResponseEncoder),
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
	return srv
}
