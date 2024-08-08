package api

import (
	"context"
	"net/http"

	"mall-pkg/jwt"
)

type UserNoErrJwt struct {
	Secret      string
	ExpiresTime int64
}

func (l *UserNoErrJwt) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 我们这里jwt鉴权取头部信息 authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := r.Header.Get(Req_Header_Token2)
		if token == "" {
			token = r.Header.Get(Req_Header_Token1)
		}
		if token == "" {
			next(w, r)
			return
		}

		claims, err := jwt.ParseUserToken(token, l.Secret)
		if err != nil {
			next(w, r)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), Context_Key_UID, claims.UserID))
		next(w, r)
	}
}
