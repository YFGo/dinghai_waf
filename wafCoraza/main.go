package main

import (
	"context"
	"errors"
	coreruleset "github.com/corazawaf/coraza-coreruleset/v4"
	"github.com/corazawaf/coraza/v3"
	"github.com/jcchavezs/mergefs"
	"github.com/jcchavezs/mergefs/io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wafCoraza/biz"
	"wafCoraza/data"
	wafHttp "wafCoraza/waf_http"
)

func initApp() (func(), *wafHttp.WafHandleService) {
	// 链接数据
	dataDB, cleanup := data.NewData()

	// 初始化app
	attackRepo := data.NewSaveAttackEventRepo(dataDB)
	attackUsercase := biz.NewAttackEventUsercase(attackRepo)
	//开启定时任务
	timeTask := attackUsercase.StartTimeTask()
	timeTask()
	attackHttp := wafHttp.NewWafHandleService(attackUsercase)

	// 在服务启动之处 , 创建存储攻击日志的csv文件
	return cleanup, attackHttp
}

func initCoraza() coraza.WAF {
	cfg := coraza.NewWAFConfig().
		WithDirectives(`
			Include wafCoraza/ruleset/coraza.conf
			Include wafCoraza/ruleset/coreruleset/crs-setup.conf.example
			Include wafCoraza/ruleset/coreruleset/rules/*.conf
 		`).
		WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS))
	waf, err := coraza.NewWAF(cfg)
	if err != nil {
		panic(err)
	}
	return waf

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic: ", err)
		}
	}()
	waf := initCoraza()
	cleanup, wafService := initApp() //初始化数据层面
	// 设置 HTTP 处理函数
	http.HandleFunc("/", wafService.WAFHandle(waf))

	// 监听并在 0.0.0.0:8888 上启动服务器
	slog.Info("Starting HTTP server on :8887")
	httpServer := &http.Server{Addr: ":8887"}
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
