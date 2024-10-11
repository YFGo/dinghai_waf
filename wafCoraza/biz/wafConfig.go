package biz

import (
	"fmt"
	coreruleset "github.com/corazawaf/coraza-coreruleset/v4"
	"github.com/corazawaf/coraza/v3"
	"github.com/jcchavezs/mergefs"
	"github.com/jcchavezs/mergefs/io"
	"log/slog"
	"strings"
	"wafCoraza/data/model"
	"wafCoraza/data/types"
)

// WafConfigRepo 加载waf配置
type WafConfigRepo interface {
	GetAppForStrategyIDs(appAddr string) ([]string, error)
	GetAllSeclangRules() ([]model.WAFStrategy, error)
}

type WafConfigUsercase struct {
	repo WafConfigRepo
	waf  map[string]coraza.WAF
}

func NewWafConfigUsercase(repo WafConfigRepo) *WafConfigUsercase {
	return &WafConfigUsercase{
		repo: repo,
		waf:  make(map[string]coraza.WAF),
	}
}

// CreateWaf 内核服务启动之初 , 根据策略创建waf实列 , 后续通过etcd通知 修改waf实列
func (w *WafConfigUsercase) CreateWaf() {
	wafStrategy, err := w.repo.GetAllSeclangRules()
	if err != nil {
		slog.Error("w.repo.GetAllSeclangRules failed : ", err)
		return
	}
	for _, strategy := range wafStrategy {
		cfg := coraza.NewWAFConfig()
		seclang := strategy.SeclangRules
		rightSeclang := w.disposeSeclang(seclang) //获取正确格式的seclang
		if strategy.ID == types.BUILDIN {         //如果是内置策略
			cfg = cfg.WithDirectives(rightSeclang).WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS))
		} else {
			cfg = cfg.WithDirectives(rightSeclang)
		}
		//创建waf
		waf, err := coraza.NewWAF(cfg)
		if err != nil {
			slog.Error("创建waf失败", err)
			panic(err)
		}
		w.waf[strategy.ID] = waf //每个策略均对应一个waf实列
	}
}

// GetAppWAF 根据域名和访问的端口 , 获取此web程序应用的策略的waf实列
func (w *WafConfigUsercase) GetAppWAF(host string) []coraza.WAF {
	wafs := make([]coraza.WAF, 0)
	strategyIDs, err := w.repo.GetAppForStrategyIDs(host)
	if err != nil {
		slog.Error("获取策略失败", err)
		return nil
	}
	//目前只有单节点etcd , 直接获取即可
	strategyIDsAll := strategyIDs[0]
	strategyIDsArr := strings.Split(strategyIDsAll, types.CutOFF)
	for _, strategyIDStr := range strategyIDsArr {
		fmt.Println(strategyIDStr)
		wafs = append(wafs, w.waf[strategyIDStr])
	}
	return wafs
}

// 处理从etcd中取出的seclang 安全规则 , 使其符合规范
func (w *WafConfigUsercase) disposeSeclang(seclang string) string {
	seclangArr := strings.Split(seclang, types.CutOFF)
	// 使用strings.Builder来构建最终的字符串
	var formattedSeclang strings.Builder
	for i, rule := range seclangArr {
		if i > 0 {
			formattedSeclang.WriteString("\n")
		}
		formattedSeclang.WriteString(rule)
	}

	return formattedSeclang.String()
}
