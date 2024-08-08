package api

import (
	"net/http"
)

// 跨域中间件
type CorsMiddleware struct {
	Headers string
	Origin  string
	Methods string
}

func NewCorsMiddleware() *CorsMiddleware {
	c := &CorsMiddleware{}
	c.Origin = "*"
	c.Methods = "POST, GET, OPTIONS, PUT, DELETE"
	c.Headers = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, User-Agent, x-token, token,x-user-id"
	return c
}

func (m *CorsMiddleware) setCors(w http.ResponseWriter) {
	//跨域请求头
	w.Header().Set("Access-Control-Allow-Origin", m.Origin)
	w.Header().Set("Access-Control-Allow-Methods", m.Methods)
	w.Header().Set("Access-Control-Allow-Headers", m.Headers)
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.setCors(w)
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, User-Agent, x-token, token,x-user-id")
		next(w, r)
	}
}

// http.Handle 接口实现
func (m *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if r.Method != "OPTIONS" || origin == "" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//跨域请求头
	m.setCors(w)
}
