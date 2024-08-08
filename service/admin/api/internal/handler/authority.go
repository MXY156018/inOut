package handler

import (
	"mall-admin/api/internal/logic"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-pkg/api"
	"mall-pkg/utils"

	// "mall-admin/api/internal/types"
	"mall-admin/model"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Authority struct{}

func (l *Authority) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/authority/createAuthority", l.Create(ctx))
		r.Post("/authority/deleteAuthority", l.Delete(ctx))
		r.Put("/authority/updateAuthority", l.Update(ctx))
		r.Post("/authority/copyAuthority", l.Copy(ctx))
		r.Post("/authority/getAuthorityList", l.GetList(ctx))
		r.Post("/authority/setDataAuthority", l.SetData(ctx))
	}

}

func (l *Authority) Create(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.Create(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Authority) Delete(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthorityIdReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Authority) Update(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.Update(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Authority) Copy(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthorityCopy
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.Copy(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Authority) GetList(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.PageQuery
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.GetAuthorityList(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Authority) SetData(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewAuthority(r.Context(), ctx)
		resp := l.SetDataAuthority(&req)
		httpx.OkJson(w, resp)
	}
}
