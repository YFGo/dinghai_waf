package plugin

import (
	"github.com/go-kratos/kratos/v2/errors"
)

const (
	ErrTokenIsFailed = 900 + iota
	ErrServer
)

func TokenErr() error {
	return errors.New(ErrTokenIsFailed, "", "token已失效,请重新登陆")
}

func ServerEr() error {
	return errors.New(ErrServer, "", "服务繁忙")
}
