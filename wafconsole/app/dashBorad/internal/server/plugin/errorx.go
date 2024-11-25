package plugin

import (
	"github.com/go-kratos/kratos/v2/errors"
)

const (
	ErrTokenIsFailed = 900
	ErrServer        = 901
)

func TokenErr() error {
	return errors.New(ErrTokenIsFailed, "", "token已失效,请重新登陆")
}

func ServerErr() error {
	return errors.New(ErrServer, "", "服务繁忙")
}
