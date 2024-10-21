package biz

import (
	"encoding/json"
	coreruleset "github.com/corazawaf/coraza-coreruleset/v4"
	"github.com/corazawaf/coraza/v3"
	"github.com/jcchavezs/mergefs"
	"github.com/jcchavezs/mergefs/io"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"strings"
	"sync"
	"wafCoraza/data/model"
	"wafCoraza/data/types"
)

// WafConfigRepo 加载waf配置
type WafConfigRepo interface {
	GetAppForStrategyIDs(appAddr string) ([]string, error)
	GetAllSeclangRules() ([]model.WAFStrategy, error)
	GetRealAddr(domain string) (string, error)
	AirUpdateStrategy() clientv3.WatchChan
	GetAllRuleGroup() ([]model.RuleGroup, error)
	GetAllRule() ([]model.Rule, error)
}

type WafConfigUsercase struct {
	repo         WafConfigRepo
	waf          map[string]*model.CorazaWaf
	ruleGroupMap map[int]model.RuleGroup
	ruleMap      map[int]model.Rule
	mu           sync.Mutex
}

func NewWafConfigUsercase(repo WafConfigRepo) *WafConfigUsercase {
	return &WafConfigUsercase{
		repo:         repo,
		waf:          make(map[string]*model.CorazaWaf),
		ruleGroupMap: make(map[int]model.RuleGroup),
		ruleMap:      make(map[int]model.Rule),
	}
}

// 处理从etcd中取出的seclang 安全规则 , 使其符合规范
func (w *WafConfigUsercase) disposeSeclang(seclang string) string {
	seclangArr := strings.Split(seclang, types.SeclangCutOFF)
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

// CreateWaf 内核服务启动之初 , 根据策略创建waf实列 , 后续通过etcd通知 修改waf实列
func (w *WafConfigUsercase) CreateWaf() {
	wafStrategy, err := w.repo.GetAllSeclangRules() // 获取策略信息
	if err != nil {
		slog.Error("w.repo.GetAllSeclangRules failed : ", err)
		return
	}
	ruleGroupList, err := w.repo.GetAllRuleGroup()
	if err != nil {
		slog.Error("w.repo.GetAllRuleGroup failed : ", err)
		return
	}
	ruleList, err := w.repo.GetAllRule()
	if err != nil {
		slog.Error("w.repo.GetAllRule failed : ", err)
		return
	}

	w.wafRuleGroup(ruleGroupList)
	w.wafRule(ruleList)
	w.wafConfig(wafStrategy)
}

// GetAppWAF 根据域名和访问的端口 , 获取此web程序应用的策略的waf实列
func (w *WafConfigUsercase) GetAppWAF(host string) []*model.CorazaWaf {
	wafs := make([]*model.CorazaWaf, 0)
	strategyIDs, err := w.repo.GetAppForStrategyIDs(host) //获取策略ID
	if err != nil || len(strategyIDs) == 0 {
		slog.Error("get strategy failed: ", err, "strategyIDs: ", strategyIDs)
		return nil
	}
	//目前只有单节点etcd , 直接获取即可
	strategyIDsAll := strategyIDs[0]
	strategyIDsArr := strings.Split(strategyIDsAll, types.CutOFF)
	for _, strategyIDStr := range strategyIDsArr {
		if wafValue, ok := w.waf[strategyIDStr]; ok {
			wafs = append(wafs, wafValue)
		}
	}
	return wafs
}

// GetRealAddr 根据请求地址 , 获取真正的后端请求地址
func (w *WafConfigUsercase) GetRealAddr(host string) (string, error) {
	realAddr, err := w.repo.GetRealAddr(host)
	if err != nil {
		slog.Error("GetRealAddr is failed: ", err)
		return realAddr, nil
	}
	return realAddr, nil
}

func (w *WafConfigUsercase) WatchStrategy() {
	watchChan := w.repo.AirUpdateStrategy()
	for wresp := range watchChan {
		slog.Info("watcher", "key", wresp.Header.Revision)
		for _, ev := range wresp.Events {
			keyArr := strings.Split(string(ev.Kv.Key), types.CutOFF) //获取策略id
			strategyTemp := model.WAFStrategy{
				ID:           keyArr[1],
				SeclangRules: string(ev.Kv.Value),
			}
			// 判断不同的情况 新增 , 修改 , 删除
			wafValue, ok := w.waf[strategyTemp.ID]
			if len(ev.Kv.Value) == 0 && !ok { // 避免一些因策略格式错误而创建失败的waf实列  在监听到删除操作时,错误的进入新增操作
				continue
			}
			if ok && wafValue != nil { //如果key存在 , 并且 waf实列不为空 根据 value判断是修改   还是删除
				switch len(strategyTemp.SeclangRules) {
				case 0: //删除
					w.mu.Lock()
					delete(w.waf, strategyTemp.ID)
					w.mu.Unlock()
				default: //修改
					w.wafConfig([]model.WAFStrategy{strategyTemp})
				}
			} else { //如果key不存在 , 新增
				w.wafConfig([]model.WAFStrategy{strategyTemp})
			}
		}
	}
}

// wafRuleGroup 处理规则组信息
func (w *WafConfigUsercase) wafRuleGroup(ruleGroupList []model.RuleGroup) {
	for _, ruleGroup := range ruleGroupList {
		w.ruleGroupMap[ruleGroup.ID] = ruleGroup
	}
}

// 处理规则信息
func (w *WafConfigUsercase) wafRule(ruleList []model.Rule) {
	for _, rule := range ruleList {
		w.ruleMap[rule.ID] = rule
	}
}

func (w *WafConfigUsercase) wafConfig(wafStrategy []model.WAFStrategy) {
	for _, strategy := range wafStrategy {
		cfg := coraza.NewWAFConfig()
		seclang := strategy.SeclangRules
		var wafConfigInfo model.WafConfig
		err := json.Unmarshal([]byte(seclang), &wafConfigInfo)
		if err != nil {
			slog.Error("waf_config json unmarshal is failed: ", err)
			continue
		}
		var (
			userRuleSeclang    string
			buildinRuleSeclang string
		)
		for _, ruleGroupID := range wafConfigInfo.RuleGroupIdList {
			//根据规则组id查询规则组信息
			ruleGroup, ok := w.ruleGroupMap[int(ruleGroupID)]
			if !ok {
				slog.Error("ruleGroup is not exist: ", ruleGroupID)
				continue
			}
			for _, ruleID := range ruleGroup.RuleIDList {
				//根据规则id查询规则信息
				rule, ok := w.ruleMap[int(ruleID)]
				if !ok {
					slog.Error("rule is not exist: ", ruleID)
					continue
				}
				if ruleGroup.IsBuildin == types.BUILDIN {
					buildinRuleSeclang = buildinRuleSeclang + types.SeclangCutOFF + rule.Seclang
				} else {
					userRuleSeclang = userRuleSeclang + types.SeclangCutOFF + rule.Seclang
				}
			}

		}
		buildinRuleSeclang = w.disposeSeclang(buildinRuleSeclang)
		userRuleSeclang = w.disposeSeclang(userRuleSeclang)
		cfg = cfg.WithDirectives(buildinRuleSeclang).WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS))
		cfg = cfg.WithDirectives(userRuleSeclang)
		waf, err := coraza.NewWAF(cfg)
		if err != nil {
			slog.Error("创建waf失败", err)
			return
		}
		w.mu.Lock()
		w.waf[strategy.ID] = &model.CorazaWaf{ //每个策略均对应一个waf实列
			WAF:         waf,
			Action:      wafConfigInfo.Action,
			NextAction:  wafConfigInfo.NextAction,
			Name:        wafConfigInfo.Name,
			Description: wafConfigInfo.Description,
		}
		w.mu.Unlock()
	}
}
