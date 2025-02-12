package utils

import (
	"context"
	"log/slog"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

const (
	auth        = "Authorization"
	UserIDMid   = "user_id"
	UserNameMid = "user_name"
)

func JWTMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				authHeader := tr.RequestHeader().Get(auth)
				if authHeader == "" {
					slog.ErrorContext(ctx, "auth header is empty")
					return nil, TokenErr()
				}
				parts := strings.Split(authHeader, " ")
				if len(parts) != 2 && parts[0] != "Bearer" {
					slog.ErrorContext(ctx, "parts is empty")
					return nil, TokenErr()
				}
				jwtUtil := InitNewJWTUtils()
				parseToken, isUpd, err := jwtUtil.ParseAccessToken(parts[1])
				if err != nil {
					slog.ErrorContext(ctx, "parts[1] is err", err, "parts[1]", parts[1])
					return nil, TokenErr()
				}
				if isUpd {
					slog.ErrorContext(ctx, "token is expired")
					return nil, TokenErr()
				}
				ctx = context.WithValue(ctx, UserIDMid, parseToken.UserId)
				ctx = context.WithValue(ctx, UserNameMid, parseToken.Username)
			}
			return handler(ctx, req)
		}
	}
}
