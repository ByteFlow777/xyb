package request

import (
	"net/http"
)

type Header struct {
	Key string
	Val string
}

type Headers []Header

func GenHeaders(cookie string) Headers {
	userAgent := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 " +
		"NetType/WIFI MiniProgramEnv/Windows WindowsWechat"

	if cookie != "" {
		cookie = "JSESSIONID=" + cookie
	}

	// 需要 form 格式请求
	contentType := "application/x-www-form-urlencoded"

	Headers := Headers{
		Header{"User-Agent", userAgent},
		Header{"Connection", "keep-alive"},
		Header{"Cookie", cookie},
		Header{"content-type", contentType},
	}
	return Headers
}

func AddHeadersToReq(req *http.Request, cookie string) {
	headers := GenHeaders(cookie)
	for _, header := range headers {
		req.Header[header.Key] = []string{header.Val}
	}
}
