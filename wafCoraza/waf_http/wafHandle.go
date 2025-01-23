package wafHttp

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"
	"wafCoraza/biz"
	constType "wafCoraza/data/types"
	"wafCoraza/waf_http/plugins"

	"github.com/corazawaf/coraza/v3/types"
)

type WafHandleService struct {
	uc          *biz.AttackEventUsercase
	wafConfigUc *biz.WafConfigUsercase
	wafAllowUc  *biz.WafAllowListUsecase
}

func NewWafHandleService(uc *biz.AttackEventUsercase, wafConfigUc *biz.WafConfigUsercase, wafAllowUc *biz.WafAllowListUsecase) *WafHandleService {
	return &WafHandleService{
		uc:          uc,
		wafConfigUc: wafConfigUc,
		wafAllowUc:  wafAllowUc,
	}
}

// InitWAF 内核启动之初 , 初始化WAF
func (w *WafHandleService) InitWAF() {
	w.wafConfigUc.InitWAF()                          // 初始化waf内核
	w.wafAllowUc.InitAllowList(context.Background()) // 初始化白名单
}

// ProxyHandler 创建反向代理服务器
func (w *WafHandleService) ProxyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// 读取请求体
		requestBody, err := io.ReadAll(req.Body)
		defer func() {
			if err := req.Body.Close(); err != nil {
				slog.Error("Error closing request body: ", err)
			}
		}()
		if err != nil {
			slog.Error("LogAttackEvent Error reading request body: ", err)
			return
		}
		hostBaseUrl := strings.Split(req.URL.Path, "/")            // 提取请求地址中的关键字段
		wafs := w.wafConfigUc.GetAppWAF(hostBaseUrl[1])            //根据访问的域名 获取收到保护的web程序所应用的策略 对应的WAF实列
		realAddr, err := w.wafConfigUc.GetRealAddr(hostBaseUrl[1]) //  获取真实的后端地址
		if err != nil {
			slog.Error("get realAddr error", err, "hostBaseUrl", hostBaseUrl[1])
		}
		newPath := strings.Join(hostBaseUrl[2:], "/") // 重构请求路径
		req.URL.Path = newPath
		var tx types.Transaction
		defer func() {
			if tx != nil {
				tx.ProcessLogging()
				if err := tx.Close(); err != nil {
					slog.Error("Error closing transaction: ", err)
				}
			}
		}()
		clientIP, clientPort, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			slog.Error("Error parsing remote address: ", err, "remote_addr : ", req.RemoteAddr)
		}
		serverIP, serverPortStr, err := net.SplitHostPort(realAddr)
		if err != nil {
			slog.Error("get server_ip , server_port is failed: ", err, "real_addr", realAddr, "req_path", req.URL.Path)
		}

		serverPort, _ := strconv.Atoi(serverPortStr)
		port, _ := strconv.Atoi(clientPort)
		err, isAllow := plugins.AllowHandle(w.wafAllowUc, context.Background(), hostBaseUrl[1], newPath, clientIP) //百名单检测
		if err != nil {
			slog.Error("allow plugins is failed: ", err)
		}
		switch isAllow {
		case true:
			plugins.Proxy(true, realAddr, req, rw, requestBody) // 处理结果
		case false: // 不存在于白名单中
			//根据这些waf实列 , 校验请求是否可以放行 , 只要存在一个waf实列拦截了请求 , 就不再检测
			for _, waf := range wafs {
				isAllow = true
				tx = waf.WAF.NewTransaction()
				tx.ProcessConnection(clientIP, port, serverIP, serverPort) // 模拟网络连接，使用请求的远程地址和端口
				tx.ProcessURI(req.RequestURI, req.Method, req.Proto)       // Request URI was /some-url?with=args
				_, reqHeaderIsLegal := w.WafParseHeader(tx, req, rw)
				if !reqHeaderIsLegal {
					isAllow = w.commonRes(waf.Action, waf.NextAction, tx, req, requestBody)
					break
				}
				_, reqBodyIsLegal := w.WafParseReqBody(tx, requestBody)
				if !reqBodyIsLegal {
					isAllow = w.commonRes(waf.Action, waf.NextAction, tx, req, requestBody)
					break
				}
			}
			plugins.Proxy(isAllow, realAddr, req, rw, requestBody) // 处理结果
		}
	}
}

// WafParseHeader 处理请求头
func (w *WafHandleService) WafParseHeader(tx types.Transaction, req *http.Request, rw http.ResponseWriter) (*types.Interruption, bool) {
	// 添加请求头
	for name, values := range req.Header {
		for _, value := range values {
			tx.AddRequestHeader(name, value)
		}
	}
	// 处理请求头
	itParse1 := tx.ProcessRequestHeaders()
	if itParse1 != nil {
		switch itParse1.Action {
		case constType.WafDeny: //访问行为被禁止
			return itParse1, false
		default:
			return itParse1, false
		}
	}
	return itParse1, true
}

// WafParseReqBody 处理请求体
func (w *WafHandleService) WafParseReqBody(tx types.Transaction, requestBody []byte) (*types.Interruption, bool) {
	// 将读取的数据写入请求体缓冲区
	if it, _, err := tx.WriteRequestBody(requestBody); it != nil || err != nil {
		if it != nil {
			return it, false
		}
		slog.Error("Error writing request body: ", err)
		return nil, false
	}
	// 处理请求体阶段
	itReqBody, err := tx.ProcessRequestBody()
	if err != nil {
		return itReqBody, false
	}
	if itReqBody != nil { //处理结果
		switch itReqBody.Action {
		case constType.WafDeny: //访问行为被禁止
			return itReqBody, false
		default:
			return itReqBody, false
		}
	}
	return itReqBody, true
}

// WafMatchRules 获取命中的规则
func (w *WafHandleService) WafMatchRules(tx types.Transaction) []types.MatchedRule {
	// 获取恶意请求匹配到的规则
	attackMatchRule := make([]types.MatchedRule, 0)
	// 获取所有匹配的规则
	matchRules := tx.MatchedRules()
	for _, rule := range matchRules {
		if !strings.Contains(rule.Rule().Raw(), "SecRule TX") && !strings.Contains(rule.Rule().Raw(), "SecAction") && !strings.Contains(rule.Rule().Raw(), "SecRule &TX") {
			attackMatchRule = append(attackMatchRule, rule)
		}
	}
	return attackMatchRule
}

// WatchEtcdService 监听etcd键值对的变化
func (w *WafHandleService) WatchEtcdService() {
	go func() {
		w.wafConfigUc.WatchStrategy()
	}()
	go func() {
		w.wafConfigUc.WatchRuleGroup()
	}()
	go func() {
		w.wafConfigUc.WatchRule()
	}()
	go func() {
		w.wafAllowUc.WatchAllowChange(context.Background())
	}()
}

func (w *WafHandleService) commonRes(action, nextAction uint8, tx types.Transaction, req *http.Request, requestBody []byte) bool {
	attackMathRules := w.WafMatchRules(tx) //处理命中的规则
	var res bool
	w.uc.LogAttackEvent(attackMathRules, req, requestBody, action, nextAction) //记录攻击日志
	if nextAction == 1 {                                                       //不再拦截
		res = true
	} else {
		res = false
	}
	return res
}
