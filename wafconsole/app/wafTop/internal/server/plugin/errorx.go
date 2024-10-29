package plugin

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

const (
	// general error type
	ErrUnknownType    = 900
	ErrParamsInvalid  = 901
	ErrRecordNotFound = 910
	ErrUnexpectedDB   = 911

	// 策略类逻辑错误 3000 ~ 4000
	ErrStrategy               = 3000
	ErrStrategyNotFound       = 3001
	ErrStrategyConfig         = 3100
	ErrStrategyConfigNotFound = 3101

	// 站点类逻辑错误 5000 ~ 6000
	ErrSiteApp            = 5000
	ErrSiteAppNotFound    = 5001
	ErrSiteServer         = 5100
	ErrSiteServerNotFound = 5101
	ErrSiteGroup          = 5200
	ErrSiteGroupNotFound  = 5201

	// 黑白名单逻辑错误 7000 ~ 8000
	ErrListPattern       = 7000
	ErrListApplyNotFound = 7001
	ErrListIsExist       = 7002
	ErrListNotFound      = 7003
	ErrListCreateFailed  = 7004
)

func UnknownTypeErr(e error) error {
	return e
}

func ReqParamErr(e error) error {
	return e
}

func StrategyErr(e error) error {
	return e
}
func StrategyNotFound() error {
	return errors.New(ErrStrategyNotFound, "", "该记录不存在！")
}
func StrategyConfigErr(e error) error {
	return e
}
func StrategyConfigNotFound() error {
	return errors.New(ErrStrategyConfigNotFound, "", "该记录不存在！")
}
func SiteAppErr(e error) error {
	return e
}
func SiteAppNotFound() error {
	return errors.New(ErrSiteAppNotFound, "", "该记录不存在！")
}
func SiteServerErr(e error) error {
	return e
}
func SiteServerNotFound() error {
	return errors.New(ErrSiteServerNotFound, "", "该记录不存在！")
}
func SiteGroupErr(e error) error {
	return e
}
func SiteGroupNotFound() error {
	return errors.New(ErrSiteGroupNotFound, "", "该记录不存在！")
}

func ListPatternErr() error {
	return errors.New(ErrListPattern, "", "该规则不合法！")
}

func ListApplyNotFoundErr() error {
	return errors.New(ErrListApplyNotFound, "", "应用对象不存在！")
}
func ListErr(err error) error {
	return err
}

func ListIsExistErr() error {
	return errors.New(ErrListIsExist, "", "名称重复，请重新填写！")
}

func ListNotFoundErr() error {
	return errors.New(ErrListNotFound, "", "该记录不存在！")
}

func ListCreateFailedErr() error {
	return errors.New(ErrListCreateFailed, "", "创建失败")
}

func ParamsInvalid() error {
	return errors.New(ErrParamsInvalid, "", "参数错误")
}

func RecordNotFound() error {
	return errors.New(ErrRecordNotFound, "", "获取不到记录")
}

func UnexpectedDB(origin error) error {
	return errors.New(ErrUnexpectedDB, origin.Error(), "数据库错误")
}

func UnAuthErr() error {
	return errors.New(http.StatusUnauthorized, "", "未登录")
}
func AuthTokenErr() error {
	return errors.New(http.StatusUnauthorized, "", "token已失效,请重新登陆")
}
func UnOrgErr() error {
	return errors.New(http.StatusUnauthorized, "没有权限访问", "没有此组织权限")
}
func UnAuthAccess() error {
	return errors.New(http.StatusUnauthorized, "没有权限访问", "当前暂无权限，如有需求请联系管理员")
}
