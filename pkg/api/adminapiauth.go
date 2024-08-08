package api

import "net/http"

// 管理员 API 权限 检查
//
//authorityId 权限ID
//
//path 路径
//
// method 方法
//
//返回时候有权限，错误返回 error
type AdminApiChecker func(authorityId string, path string, method string) (bool, error)

// 管理员权限检查
//
// 依赖 管理员 JWT 查件
type AdminApiAuth struct {
	Checker AdminApiChecker
}

func (l *AdminApiAuth) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的URI
		path := r.RequestURI
		// 获取请求方法
		method := r.Method
		// 获取用户的角色
		authorityId := r.Context().Value(Middle_Header_AuthorityId).(string)
		isok, err := l.Checker(authorityId, path, method)
		if err != nil {
			JsonResp(w, &BaseResp{
				Code: Error_Server,
				Msg:  err.Error(),
			})
			return
		}
		if !isok {
			JsonResp(w, &BaseResp{
				Code: Error_NoPrivilege,
				Msg:  "无访问权限",
			})
			return
		}
		next(w, r)
	}
}
