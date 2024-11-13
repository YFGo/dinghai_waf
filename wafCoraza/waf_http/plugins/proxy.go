package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Proxy(wafDeny bool, realAddr string, req *http.Request, rw http.ResponseWriter, requestBody []byte) {
	if wafDeny && len(realAddr) != 0 { //允许放行
		targetURL, err := url.Parse(fmt.Sprintf("http://%s", realAddr))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		req.Body = io.NopCloser(strings.NewReader(string(requestBody))) /* 重置请求体 */
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(rw, req)
	} else {
		badMessage := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    string `json:"data"`
		}{
			Code:    403,
			Message: "Forbidden",
			Data:    "",
		}
		badMessageByte, _ := json.Marshal(badMessage)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(badMessageByte)
	}
}
