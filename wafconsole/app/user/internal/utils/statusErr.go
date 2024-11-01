package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StatusErr(err error, code codes.Code) bool {
	if s, _ := status.FromError(err); s.Code() == code {
		return true
	}
	return false
}
