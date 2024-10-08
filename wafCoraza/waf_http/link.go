package wafHttp

import "strings"

func getLinkUrl(uri string) string {
	//Link配置
	var link string
	switch {
	case strings.Contains(uri, "user"):
		link = "http://" + "47.93.85.12:8887" + uri
	case strings.Contains(uri, "tag"):
		link = "http://" + "47.93.85.12:8888" + uri
	case strings.Contains(uri, "chat"):
		link = "http://" + "47.93.85.12:8889" + uri
	}
	return link
}
