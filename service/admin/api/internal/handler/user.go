package handler

import (
	"mall-admin/api/internal/logic"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-pkg/api"
	"mall-pkg/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type User struct{}

// 注册处理函数
func (l *User) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/user/admin_register", l.RegisterUser(ctx))
		r.Post("/user/changePassword", l.ChangePassword(ctx))
		r.Post("/user/getUserList", l.GetUserList(ctx))
		r.Put("/user/setSelfInfo", l.SetSelfInfo(ctx))
		r.Post("/user/resetPassword", l.ResetPassword(ctx))
		r.Post("/user/setUserAuthority", l.SetUserAuthority(ctx))
		r.Delete("/user/deleteUser", l.DeleteUser(ctx))
		r.Put("/user/setUserInfo", l.SetUserInfo(ctx))
		r.Post("/user/setUserAuthorities", l.SetUserAuthoritys(ctx))
		r.Get("/user/getUserInfo", l.GetUserInfo(ctx))
	}
}

// 用户注册
func (l *User) RegisterUser(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.Register(&req)
		httpx.OkJson(w, resp)
	}
}

// 设置个人信息
func (l *User) SetSelfInfo(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSelfInfoReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.SetSelfInfo(&req)
		httpx.OkJson(w, resp)
	}
}

// 重置用户密码
func (l *User) ResetPassword(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.ResetPassword(&req)
		httpx.OkJson(w, resp)
	}
}

// 更改用户密码
func (l *User) ChangePassword(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.ChangePassword(&req)
		httpx.OkJson(w, resp)
	}
}

// 获取用户列表
func (l *User) GetUserList(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.PageQuery
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.GetUserList(&req)
		httpx.OkJson(w, resp)
	}
}

// 设置用户权限
func (l *User) SetUserAuthority(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserAuthReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.SetSelfAuthority(&req, r)
		httpx.OkJson(w, resp)
	}
}

// 删除用户
func (l *User) DeleteUser(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.DeleteUser(&req)
		httpx.OkJson(w, resp)
	}
}

// 设置用户信息
func (l *User) SetUserInfo(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysUser
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.SetUserInfo(&req)
		httpx.OkJson(w, resp)
	}
}

// 批量设置用户权限
func (l *User) SetUserAuthoritys(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserAuthoritiesReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewUser(r.Context(), ctx)
		resp := l.SetUserAuthorities(&req)
		httpx.OkJson(w, resp)
	}
}

// 获取用户信息
func (l *User) GetUserInfo(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req model.SysApi
		// if err := utils.Bind(r, &req); err != nil {
		// 	httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
		// 	return
		// }
		l := logic.NewUser(r.Context(), ctx)
		resp := l.GetUserInfo()
		httpx.OkJson(w, resp)
	}
}
