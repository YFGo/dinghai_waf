package biz

import (
	"encoding/json"
	coreruleset "github.com/corazawaf/coraza-coreruleset/v4"
	"github.com/corazawaf/coraza/v3"
	"github.com/jcchavezs/mergefs"
	"github.com/jcchavezs/mergefs/io"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"strconv"
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
	AirUpdateRuleGroup() clientv3.WatchChan
	AirUpdateRule() clientv3.WatchChan
	GetStrategyListById(strategyKey string) ([]model.WAFStrategy, error)
	GetCommonByKey(key string) (string, error)
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

// InitWAF 内核服务启动之初 , 根据策略创建waf实列 , 后续通过etcd通知 修改waf实列
func (w *WafConfigUsercase) InitWAF() {
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

	w.wafRule(ruleList)
	w.wafRuleGroup(ruleGroupList)
	w.wafConfig(wafStrategy)
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

// GetAppWAF 根据域名和访问的端口 , 获取此web程序应用的策略的waf实列
func (w *WafConfigUsercase) GetAppWAF(host string) []*model.CorazaWaf {
	wafs := make([]*model.CorazaWaf, 0)
	strategyIDs, err := w.repo.GetAppForStrategyIDs(host + types.StrategySuffix) //获取策略ID
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

// WatchStrategy 监视策略
func (w *WafConfigUsercase) WatchStrategy() {
	watchChan := w.repo.AirUpdateStrategy()
	for {
		select {
		case watchResp := <-watchChan:
			for _, event := range watchResp.Events {
				keyArr := strings.Split(string(event.Kv.Key), types.CutOFF) //获取策略id
				switch event.Type {                                         //判断此次更新是什么类型的操作
				case clientv3.EventTypePut:
					strategyTemp := model.WAFStrategy{
						ID:           keyArr[1],
						SeclangRules: string(event.Kv.Value),
					}
					w.wafConfig([]model.WAFStrategy{strategyTemp})
				case clientv3.EventTypeDelete:
					delete(w.waf, keyArr[1])
				}
			}
		}
	}
}

// WatchRuleGroup 监视规则组
func (w *WafConfigUsercase) WatchRuleGroup() {
	watchChan := w.repo.AirUpdateRuleGroup()
	for {
		select {
		case watchResp := <-watchChan:
			for _, event := range watchResp.Events {
				keyArr := strings.Split(string(event.Kv.Key), types.CutOFF) //获取规则组id
				groupID, err := strconv.Atoi(keyArr[1])
				if err != nil {
					slog.Error("strconv group_id is failed: ", err)
					continue
				}
				switch event.Type { //判断此次更新是什么类型的操作
				case clientv3.EventTypePut:
					var ruleGroup model.RuleGroup
					err = json.Unmarshal(event.Kv.Value, &ruleGroup)
					if err != nil {
						slog.Error("ruleGroup json unmarshal is failed: ", err)
						continue
					}
					w.ruleGroupMap[ruleGroup.ID] = ruleGroup //更新规则组信息
				case clientv3.EventTypeDelete:
					delete(w.ruleGroupMap, groupID)
				}
			}
		}
	}
}

func (w *WafConfigUsercase) WatchRule() {
	watchChan := w.repo.AirUpdateRule()
	for {
		select {
		case watchResp := <-watchChan:
			for _, event := range watchResp.Events {
				keyArr := strings.Split(string(event.Kv.Key), types.CutOFF) //获取规则id
				ruleID, err := strconv.Atoi(keyArr[1])
				if err != nil {
					slog.Error("strconv rule_id is failed: ", err)
					continue
				}
				switch event.Type { //判断此次更新是什么类型的操作
				case clientv3.EventTypePut:
					var rule model.Rule
					err = json.Unmarshal(event.Kv.Value, &rule)
					if err != nil {
						slog.Error("rule json unmarshal is failed: ", err)
						continue
					}
					w.ruleMap[rule.ID] = rule //更新规则信息
				case clientv3.EventTypeDelete:
					delete(w.ruleMap, ruleID)
				}
			}
		}
	}
}

func (w *WafConfigUsercase) wafConfig(wafStrategy []model.WAFStrategy) {
	w.mu.Lock()
	for _, strategy := range wafStrategy {
		cfg := coraza.NewWAFConfig()
		seclang := strategy.SeclangRules
		var wafConfigInfo model.WafConfig
		err := json.Unmarshal([]byte(seclang), &wafConfigInfo) // 解析策略信息
		if err != nil {
			slog.Error("waf_config json unmarshal is failed: ", err)
			continue
		}
		var (
			userRuleSeclang    string
			buildinRuleSeclang string
		)
		for _, ruleGroupID := range wafConfigInfo.RuleGroupIdList { // 遍历此策略包含的规则组ID
			ruleGroup, ok := w.ruleGroupMap[int(ruleGroupID)] //根据规则组id查询规则组信息
			if !ok {
				slog.Error("ruleGroup is not exist: ", ruleGroupID)
				continue
			}
			for _, ruleID := range ruleGroup.RuleIDList {
				rule, ok := w.ruleMap[int(ruleID)] //根据规则id查询规则信息
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
			slog.Error("创建waf失败", err, "user_rule_seclang: ", userRuleSeclang)
			return
		}
		w.waf[strategy.ID] = &model.CorazaWaf{ //每个策略均对应一个waf实列
			WAF:         waf,
			Action:      wafConfigInfo.Action,
			NextAction:  wafConfigInfo.NextAction,
			Name:        wafConfigInfo.Name,
			Description: wafConfigInfo.Description,
		}
	}
	w.mu.Unlock()
}
