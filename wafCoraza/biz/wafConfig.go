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
	repo                 WafConfigRepo
	waf                  map[string]*model.CorazaWaf
	ruleGroupMap         map[int]model.RuleGroup
	ruleMap              map[int]model.Rule
	ruleGroupStrategyMap map[int64][]string //key为规则组id , value为策略id
	ruleToGroupMap       map[int64]int64    //key为规则id , value为规则组id
	mu                   sync.Mutex
}

func NewWafConfigUsercase(repo WafConfigRepo) *WafConfigUsercase {
	return &WafConfigUsercase{
		repo:                 repo,
		waf:                  make(map[string]*model.CorazaWaf),
		ruleGroupMap:         make(map[int]model.RuleGroup),
		ruleMap:              make(map[int]model.Rule),
		ruleGroupStrategyMap: make(map[int64][]string),
		ruleToGroupMap:       make(map[int64]int64),
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

	w.wafRule(ruleList)
	w.wafRuleGroup(ruleGroupList)
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

// WatchStrategy 监视策略
func (w *WafConfigUsercase) WatchStrategy() {
	watchChan := w.repo.AirUpdateStrategy()
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			keyArr := strings.Split(string(ev.Kv.Key), types.CutOFF) //获取策略id
			strategyTemp := model.WAFStrategy{
				ID:           keyArr[1],
				SeclangRules: string(ev.Kv.Value),
			}
			slog.Info("Key changed", "Key", ev.Kv.Key, "Value", ev.Kv.Value)
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

func (w *WafConfigUsercase) WatchRuleGroup() {
	watchChan := w.repo.AirUpdateRuleGroup()
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			var ruleGroup model.RuleGroup
			ruleGroupKeyArr := strings.Split(string(ev.Kv.Key), types.CutOFF)
			ruleGroup.ID, _ = strconv.Atoi(ruleGroupKeyArr[1]) //
			strategyListId, ok := w.ruleGroupStrategyMap[int64(ruleGroup.ID)]
			if !ok { //新增规则组的操作

			}
			if len(ev.Kv.Value) == 0 { //执行删除操作
				delete(w.ruleGroupMap, ruleGroup.ID)
				delete(w.ruleGroupStrategyMap, int64(ruleGroup.ID))
			} else {
				err := json.Unmarshal(ev.Kv.Value, &ruleGroup)
				if err != nil {
					slog.Error("ruleGroup json unmarshal is failed: ", err)
					continue
				}
				w.ruleGroupMap[ruleGroup.ID] = ruleGroup //更新规则组信息
				//更新相应的规则信息
				for _, ruleID := range ruleGroup.RuleIDList {
					if _, ok := w.ruleMap[int(ruleID)]; !ok { //此规则是新增的规则
						ruleValue, err := w.repo.GetCommonByKey(types.RulePrefix + strconv.Itoa(int(ruleID)))
						if err != nil {
							slog.Error("rule is not exist: ", ruleID)
							continue
						}
						var rule model.Rule
						err = json.Unmarshal([]byte(ruleValue), &rule)
						if err != nil {
							slog.Error(" WatchRuleGroup rule json unmarshal is failed: ", err)
							continue
						}
						w.ruleMap[rule.ID] = rule //更新规则信息
					}
				}
			}
			for _, strategyId := range strategyListId {
				//获取这些策略信息
				strategyList, err := w.repo.GetStrategyListById(types.StrategyKey + strategyId)
				if err != nil {
					slog.Error("GetStrategyListById is failed: ", err)
					continue
				}
				w.wafConfig(strategyList)
			}
		}
	}
}

func (w *WafConfigUsercase) WatchRule() {
	watchChan := w.repo.AirUpdateRule()
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			var rule model.Rule
			ruleKeyArr := strings.Split(string(ev.Kv.Key), types.CutOFF)
			rule.ID, _ = strconv.Atoi(ruleKeyArr[1]) // 获取到规则ID
			groupID, ok := w.ruleToGroupMap[int64(rule.ID)]
			if len(ev.Kv.Value) == 0 { //执行删除操作
				delete(w.ruleMap, rule.ID)
				delete(w.ruleToGroupMap, int64(rule.ID))
			} else {
				err := json.Unmarshal(ev.Kv.Value, &rule)
				if err != nil {
					slog.Error("rule json unmarshal is failed: ", err)
					continue
				}
				w.ruleMap[rule.ID] = rule //更新规则信息
			}
			//根据groupID , 获取 策略id列表
			strategyListId, ok := w.ruleGroupStrategyMap[groupID]
			if !ok { //如果不存在 , 表示当前操作为新增操作
				// 查询此规则绑定的规则组
			}
			for _, strategyId := range strategyListId {
				//获取这些策略信息
				strategyList, err := w.repo.GetStrategyListById(types.StrategyKey + strategyId)
				if err != nil {
					slog.Error("GetStrategyListById is failed: ", err)
					continue
				}
				w.wafConfig(strategyList)
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
		w.mu.Lock()
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
			isFlag             bool // 是否使用了内置规则
		)
		for _, ruleGroupID := range wafConfigInfo.RuleGroupIdList { // 遍历此策略包含的规则组ID
			w.ruleGroupStrategyMap[ruleGroupID] = append(w.ruleGroupStrategyMap[ruleGroupID], strategy.ID) //记录规则组id与策略id的关系 , 映射的map集合
			ruleGroup, ok := w.ruleGroupMap[int(ruleGroupID)]                                              //根据规则组id查询规则组信息
			if !ok {
				slog.Error("ruleGroup is not exist: ", ruleGroupID)
				continue
			}
			for _, ruleID := range ruleGroup.RuleIDList {
				rule, ok := w.ruleMap[int(ruleID)]     //根据规则id查询规则信息
				w.ruleToGroupMap[ruleID] = ruleGroupID //映射规则和规则组的关系 , 一个规则只能属于一个规则组
				if !ok {
					slog.Error("rule is not exist: ", ruleID)
					continue
				}
				if ruleGroup.IsBuildin == types.BUILDIN {
					isFlag = true
					buildinRuleSeclang = buildinRuleSeclang + types.SeclangCutOFF + rule.Seclang
				} else {
					userRuleSeclang = userRuleSeclang + types.SeclangCutOFF + rule.Seclang
				}
			}

		}
		buildinRuleSeclang = w.disposeSeclang(buildinRuleSeclang)
		userRuleSeclang = w.disposeSeclang(userRuleSeclang)
		cfg = cfg.WithDirectives(buildinRuleSeclang).WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS))
		if !isFlag {
			cfg = cfg.WithDirectives(`Include wafCoraza/ruleset/coraza.conf`).WithRootFS(mergefs.Merge(coreruleset.FS, io.OSFS))
		}
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
		w.mu.Unlock()
	}
}
