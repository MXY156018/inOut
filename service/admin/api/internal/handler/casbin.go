package handler

import (
	"mall-admin/api/internal/logic"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-pkg/api"
	"mall-pkg/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Casbin struct{}

func (l *Casbin) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/casbin/updateCasbin", l.Update(ctx))
		r.Post("/casbin/getPolicyPathByAuthorityId", l.GetPolicyPathByAuthorityId(ctx))
	}

}

func (l *Casbin) Update(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CasbinInReceiveReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewCasbin(r.Context(), ctx)
		resp := l.Update(&req)
		httpx.OkJson(w, resp)
	}
}

func (l *Casbin) GetPolicyPathByAuthorityId(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthorityIdReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewCasbin(r.Context(), ctx)
		resp := l.GetPolicyPathByAuthorityId(&req)
		httpx.OkJson(w, resp)
	}
}
