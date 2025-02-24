package main

import (
	"context"
	"errors"
	"gopkg.in/ini.v1"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wafcoraza/biz"
	"wafcoraza/data"
	wafHttp "wafcoraza/waf_http"
)

func initApp() (func(), *wafHttp.WafHandleService, *ini.File) {
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
	// 初始化定时任务
	// 开启定时任务
	//开启定时任务
	timeTask := attackUsercase.StartTimeTask()
	timeTask()
	attackHttp := wafHttp.NewWafHandleService(attackUsercase, wafConfigUsercase, wafAllowUsercase)

	// 在服务启动之处 , 创建存储攻击日志的csv文件
	return cleanup, attackHttp, file
}

func main() {

	cleanup, wafService, file := initApp() //初始化数据层面
	// 初始化waf 实列
	wafService.InitWAF()
	// 配置热更新waf实列

	wafService.WatchEtcdService()

	// 设置 HTTP 处理函数
	http.HandleFunc("/", wafService.ProxyHandler())

	// 监听并在  上启动服务器
	slog.Info("Starting HTTP server on :" + file.Section("app").Key("port").String())
	httpServer := &http.Server{Addr: ":" + file.Section("app").Key("port").String()}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("ListenAndServe", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	<-quit
	// 停止 HTTP 服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error("HTTP server shutdown: ", err)
	}
	cleanup()
}
