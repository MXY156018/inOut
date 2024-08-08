package api

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

func JsonResp(w http.ResponseWriter, data interface{}) {
	resp, _ := json.Marshal(data)
	w.Write(resp)
}

// 获取客户端IP
//
//r 请求
func ClientIp(r *http.Request) (string, error) {
	realIp := r.Header.Get("X-real-ip")
	if realIp == "" {
		realIp = r.RemoteAddr
	}
	var host string
	var err error
	if strings.Contains(realIp, ":") {
		host, _, err = net.SplitHostPort(realIp)
	} else {
		host = realIp
	}

	return host, err
}
