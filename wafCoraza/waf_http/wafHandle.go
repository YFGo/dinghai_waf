package wafHttp

import (
	"bytes"
	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/types"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"wafCoraza/biz"
)

type WafHandleService struct {
	uc *biz.AttackEventUsercase
}

func NewWafHandleService(uc *biz.AttackEventUsercase) *WafHandleService {
	return &WafHandleService{uc: uc}
}

// WAFHandle HTTP 请求处理函数
func (w *WafHandleService) WAFHandle(waf coraza.WAF) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

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
			w.ForwardHandler(rw, req, requestBody)
		} else {
			//记录攻击日志
			w.uc.LogAttackEvent(attackMathRules, req, requestBody)
		}
	}
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

func (w *WafHandleService) ForwardHandler(rw http.ResponseWriter, r *http.Request, reqBody []byte) {
	// 解析请求的目标地址
	linkURL := getLinkUrl(r.RequestURI)
	targetURL, err := url.Parse(linkURL)
	if err != nil {
		// 处理错误
		return
	}
	// 构建新的请求对象
	forwardReq, err := http.NewRequest(r.Method, targetURL.String(), bytes.NewReader(reqBody))
	forwardReq.Header = r.Header
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 发送转发请求并获取响应
	resp, err := client.Do(forwardReq)
	if err != nil {
		// 处理错误
		return
	}
	defer resp.Body.Close()

	// 将目标服务器的响应返回给客户端
	for key, values := range resp.Header {
		for _, value := range values {
			rw.Header().Add(key, value)
		}
	}
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}
