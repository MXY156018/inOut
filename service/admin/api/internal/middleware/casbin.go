package middleware

import (
	"encoding/json"
	"mall-admin/pkg"
	"mall-pkg/api"
	"mall-pkg/jwt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 拦截器
func Casbin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := r.Header.Get(api.Middle_Header_Claims)
		var waitUse jwt.AdminClaims
		err := json.Unmarshal([]byte(str), &waitUse)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 获取请求的URI
		obj := r.RequestURI
		// 获取请求方法
		act := r.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := pkg.Casbin.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			next(w, r)
		} else {
			res := api.BaseResp{
				Code: api.Error_NoPrivilege,
				Msg:  "无权限",
			}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
	}
}
