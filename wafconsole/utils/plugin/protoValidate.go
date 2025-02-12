package utils

import (
	"context"
	"errors"
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Validate struct {
	v *protovalidate.Validator
}

func NewValidate() (*Validate, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	return &Validate{v: v}, nil
}

func (v *Validate) ValidateUnaryServerInterceptor() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			switch req.(type) {
			case proto.Message:
				if err := v.v.Validate(req.(proto.Message)); err != nil {
					var valErr *protovalidate.ValidationError
					if ok := errors.As(err, &valErr); ok && len(valErr.ToProto().GetViolations()) > 0 {
						return nil, status.Error(codes.InvalidArgument, *valErr.ToProto().GetViolations()[0].Message)
					}
					return nil, status.Error(codes.InvalidArgument, err.Error())
				}
			}
			return handler(ctx, req)
		}
	}
}
