package wafHttp

import (
	"fmt"
	"github.com/corazawaf/coraza/v3/types"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
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

// InitWAF 内核启动之初 , 初始化WAF
func (w *WafHandleService) InitWAF() {
	w.wafConfigUc.CreateWaf()
}

// ProxyHandler 创建反向代理服务器
func (w *WafHandleService) ProxyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		//根据访问的域名 获取收到保护的web程序所应用的策略 对应的WAF实列
		wafs := w.wafConfigUc.GetAppWAF(req.Host)
		//根据这些waf实列 , 校验请求是否可以放行 , 只要存在一个waf实列拦截了请求 , 就不再检测
		var tx types.Transaction
		defer func() {
			tx.ProcessLogging()
			if err := tx.Close(); err != nil {
				slog.Error("Error closing transaction: ", err)
			}
		}()
		clientIP, clientPort, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			log.Printf("Error splitting RemoteAddr: %v", err)
		}
		port, _ := strconv.Atoi(clientPort)
		var isAllow bool
		for _, waf := range wafs {
			tx = waf.NewTransaction()
			tx.ProcessConnection(clientIP, port, "152.136.50.60", 8888) // 模拟网络连接，使用请求的远程地址和端口
			tx.ProcessURI(req.RequestURI, req.Method, req.Proto)        // Request URI was /some-url?with=args
			_, reqHeaderIsLegal := w.WafParseHeader(tx, req, rw)
			_, reqBodyIsLegal, requestBody := w.WafParseReqBody(tx, req, rw)
			attackMathRules := w.WafMatchRules(tx)  //处理命中的规则
			if reqHeaderIsLegal && reqBodyIsLegal { // 此waf实列检测 请求头和请求体的检测均无问题
				isAllow = true
			} else {
				isAllow = false
				w.uc.LogAttackEvent(attackMathRules, req, requestBody) //记录攻击日志
				break
			}
		}
		if isAllow { //允许放行
			readAddr, err := w.wafConfigUc.GetRealAddr(req.Host)
			targetURL, err := url.Parse(fmt.Sprintf("http://%s", readAddr))
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			proxy := httputil.NewSingleHostReverseProxy(targetURL)
			proxy.ServeHTTP(rw, req)
		} else {
			http.Error(rw, "非法访问", 403)
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
		return itParse1, false
	}
	return itParse1, true
}

// WafParseReqBody 处理请求体
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
		return nil, false, requestBody
	}
	// 处理请求体阶段
	itReqBody, err := tx.ProcessRequestBody()
	if itReqBody != nil {
		return itReqBody, false, requestBody
	}
	return itReqBody, true, requestBody
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
