package appinit

import (
	"gopkg.in/ini.v1"
	"wafcoraza/biz"
	"wafcoraza/data"
	wafHttp "wafcoraza/waf_http"
)

func InitApp() (func(), *wafHttp.WafHandleService, *ini.File) {
	file := data.NewConfFile()
	// 链接数据
	dataDB, cleanup := data.NewData(file)

	// 初始化app
	attackRepo := data.NewSaveAttackEventRepo(dataDB)
	attackUsercase := biz.NewAttackEventUsercase(attackRepo)
	// waf
	loadWafRepo := data.NewLoadWAFConfigRepo(dataDB)
	wafConfigUsercase := biz.NewWafConfigUsercase(loadWafRepo)
	// allow
	wafAllowRepo := data.NewWafAllowListRepo(dataDB)
	wafAllowUsercase := biz.NewWafAllowListUsecase(wafAllowRepo)
	// normal http
	normalHttpRepo := data.NewNormalHttpRepo(dataDB)
	normalHttpUsercase := biz.NewNormalHttpUsercase(normalHttpRepo)

	attackHttp := wafHttp.NewWafHandleService(attackUsercase, wafConfigUsercase, wafAllowUsercase, normalHttpUsercase)

	// 在服务启动之处 , 创建存储攻击日志的csv文件
	return cleanup, attackHttp, file
}
