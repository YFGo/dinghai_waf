package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wafcoraza/appinit"
)

func main() {

	cleanup, wafService, file := appinit.InitApp() //初始化数据层面
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
