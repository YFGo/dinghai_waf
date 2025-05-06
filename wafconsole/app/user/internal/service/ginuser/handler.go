package ginuser

import (
	"fmt"
	"net/http"
	"time"
)

// SSEHandlerWrapper 自定义结构体，实现 http.Handler 接口
type SSEHandlerWrapper struct{}

func (s *SSEHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	SSEHandler(w, r)
}

// SSEHandler 处理 SSE 请求
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 确保连接不会被代理关闭
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 发送自定义事件
	events := []struct {
		EventType string
		Data      string
	}{
		{"message", "This is a normal message"},
		{"alert", "This is an alert message"},
		{"notification", "This is a notification message"},
	}

	for _, event := range events {
		fmt.Fprintf(w, "event: %s\n", event.EventType)
		fmt.Fprintf(w, "data: %s\n\n", event.Data)
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
}
