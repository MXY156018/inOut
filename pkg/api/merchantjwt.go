package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"mall-pkg/jwt"
)

const (
	// 商户权限(固定100)
	Merchant_AuthorityId = "100"
	Merchant_MainId      = 7
	Express_Fee          = 10
)

type MerchantJwt struct {
	Secret      string
	ExpiresTime int64
}

func (l *MerchantJwt) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 我们这里jwt鉴权取头部信息 authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := r.Header.Get(Req_Header_Token2)
		if token == "" {
			token = r.Header.Get(Req_Header_Token1)
		}
		if token == "" {
			respInvalidToken(w, Error_InvalidToken, "无效token")
			return
		}
		claims, err := jwt.ParseMerchantToken(token, l.Secret)
		if err != nil {
			if err == jwt.TokenExpired {
				respInvalidToken(w, Error_TokenExpire, "token超时")
			} else {
				respInvalidToken(w, Error_InvalidToken, "无效token")
			}
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + l.ExpiresTime
			newToken, err := jwt.GetMerchantToken(l.Secret, *claims)
			if err != nil {
				r.Header.Set(Resp_Header_NewToken, newToken)
				r.Header.Set(Resp_Header_NewTokenExpire, strconv.FormatInt(claims.ExpiresAt, 10))
			}
		}

		r = r.WithContext(context.WithValue(r.Context(), Context_Key_UID, claims.MchID))
		r = r.WithContext(context.WithValue(r.Context(), Context_Key_Privilege, claims.Privilege))
		r = r.WithContext(context.WithValue(r.Context(), Context_Key_MerchantLogin, claims.IsLogin))
		next(w, r)
	}
}
