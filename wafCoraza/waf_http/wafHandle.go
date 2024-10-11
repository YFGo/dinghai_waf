package wafHttp

import (
	"github.com/corazawaf/coraza/v3/types"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"
	"wafCoraza/biz"
)

type WafHandleService struct {
	uc          *biz.AttackEventUsercase
	wafConfigUc *biz.WafConfigUsercase
}

func NewWafHandleService(uc *biz.AttackEventUsercase, wafConfigUc *biz.WafConfigUsercase) *WafHandleService {
	return &WafHandleService{
		uc:          uc,
		wafConfigUc: wafConfigUc,
	}
}

// ProxyHandler 创建反向代理服务器
func (w *WafHandleService) ProxyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		//根据访问的域名 获取收到保护的web程序所应用的策略 对应的WAF实列
		wafs := w.wafConfigUc.GetAppWAF(req.Host)
		//根据这些waf实列 , 校验请求是否可以放行
		for _, waf := range wafs {
			tx := waf.NewTransaction()
			defer func() {
				tx.ProcessLogging()
				if err := tx.Close(); err != nil {
					slog.Error("Error closing transaction: ", err)
				}
			}()
			// 获取真实的客户端 IP 地址和端口
			clientIP, clientPort, err := net.SplitHostPort(req.RemoteAddr)
			if err != nil {
				log.Printf("Error splitting RemoteAddr: %v", err)
			}
			port, _ := strconv.Atoi(clientPort)
			// 模拟网络连接，使用请求的远程地址和端口
			tx.ProcessConnection(clientIP, port, "152.136.50.60", 8888)
			// Request URI was /some-url?with=args
			tx.ProcessURI(req.RequestURI, req.Method, req.Proto)
			// 阶段1 处理请求头
			_, isAllow1 := w.WafParseHeader(tx, req, rw)
			// 阶段2 处理请求体
			_, isAllow2, requestBody := w.WafParseReqBody(tx, req, rw)
			attackMathRules := w.WafMatchRules(tx) //处理命中的规则
			if isAllow1 && isAllow2 {
			} else {
				//记录攻击日志
				w.uc.LogAttackEvent(attackMathRules, req, requestBody)
			}
		}
	}
}

// InitWAF 内核启动之初 , 初始化WAF
func (w *WafHandleService) InitWAF() {
	w.wafConfigUc.CreateWaf()
}

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
		http.Error(rw, "非法请求", 403)
		return itParse1, false
	}
	return itParse1, true
}

func (w *WafHandleService) WafParseReqBody(tx types.Transaction, req *http.Request, rw http.ResponseWriter) (*types.Interruption, bool, []byte) {
	requestBody, err := io.ReadAll(req.Body)
	defer func() {
		if err := req.Body.Close(); err != nil {
			slog.Error("Error closing request body: ", err)
		}
	}()
	if err != nil {
		slog.Error("LogAttackEvent Error reading request body: ", err)
		return nil, false, requestBody
	}
	// 将读取的数据写入请求体缓冲区
	if it, _, err := tx.WriteRequestBody(requestBody); it != nil || err != nil {
		if it != nil {
			return it, false, requestBody
		}
		slog.Error("Error writing request body: ", err)
		http.Error(rw, "Error writing request body", http.StatusBadRequest)
		return nil, false, requestBody
	}
	// 处理请求体阶段
	itReqBody, err := tx.ProcessRequestBody()
	if itReqBody != nil {
		http.Error(rw, itReqBody.Action, 403)
		return itReqBody, false, requestBody
	}
	return itReqBody, true, requestBody
}

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
