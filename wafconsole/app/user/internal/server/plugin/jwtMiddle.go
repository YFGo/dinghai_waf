package plugin

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

const (
	auth     = "Authorization"
	UserID   = "user_id"
	UserName = "user_name"
)

func JWTMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				authHeader := tr.RequestHeader().Get(auth)
				if authHeader == "" {
					return nil, TokenErr()
				}
				parts := strings.Split(authHeader, " ")
				if !(len(parts) == 2) && parts[0] == "Bearer" {
					return nil, TokenErr()
				}
				jwtUtil := InitNewJWTUtils()
				parseToken, isUpd, err := jwtUtil.ParseAccessToken(parts[1])
				if err != nil {
					return nil, TokenErr()
				}
				if isUpd {
					return nil, TokenErr()
				}
				ctx = context.WithValue(ctx, UserID, parseToken.UserId)
				ctx = context.WithValue(ctx, UserName, parseToken.Username)
			}
			return handler(ctx, req)
		}
	}
}
