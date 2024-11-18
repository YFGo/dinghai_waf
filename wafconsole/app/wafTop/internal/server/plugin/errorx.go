package plugin

import (
	"github.com/go-kratos/kratos/v2/errors"
)

const (
	ErrTokenIsFailed = 900
	ErrServer        = 901

	// 6500 ~ 7000 规则错误
	ErrRuleGroupIsExist = 6500
	ErrRuleGroupIsUsing = 6501

	// 7000 ~ 7500 策略错误
	ErrStrategyNotFound = 7000

	// 7500 - 8000 站点错误
	ErrServerExist  = 7500
	ErrServerChoose = 7501

	// 8500~9000 白名单错误
	ErrAllowExist    = 8500
	ErrAllowNotFound = 8501
)

func TokenErr() error {
	return errors.New(ErrTokenIsFailed, "", "token已失效,请重新登陆")
}

func ServerErr() error {
	return errors.New(ErrServer, "", "服务繁忙")
}

func AllowExistErr() error {
	return errors.New(ErrAllowExist, "", "白名单已存在")
}

func ServerExistErr() error {
	return errors.New(ErrServerExist, "", "服务器站点已存在")
}

func ServerChooseErr(err error) error {
	return errors.New(ErrServerChoose, "", err.Error())
}

func RuleGroupIsExistErr() error {
	return errors.New(ErrRuleGroupIsExist, "", "规则组已存在")
}

func RuleGroupIsUsingErr() error {
	return errors.New(ErrRuleGroupIsUsing, "", "规则组下存在未被删除的规则")
}
