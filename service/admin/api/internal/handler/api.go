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

type Api struct{}

func (l *Api) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/api/createApi", l.Create(ctx))
		r.Post("/api/deleteApi", l.Delete(ctx))
		r.Post("/api/getApiList", l.SearchApi(ctx))
		r.Post("/api/getApiById", l.GetById(ctx))
		r.Post("/api/updateApi", l.Update(ctx))
		r.Post("/api/getAllApis", l.GetAll(ctx))
		r.Delete("/api/deleteApisByIds", l.DeleteByIds(ctx))
	}

}

//新增api
func (l *Api) Create(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysApi
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.Create(&req)
		httpx.OkJson(w, resp)
	}
}

// 删除api
func (l *Api) Delete(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysApi
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.Delete(&req)
		httpx.OkJson(w, resp)
	}
}

//获取apilist
func (l *Api) SearchApi(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiSearchReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.SearchApi(&req)
		httpx.OkJson(w, resp)
	}
}

//获取单条Api消息
func (l *Api) GetById(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.GetById(&req)
		httpx.OkJson(w, resp)
	}
}

//更新api
func (l *Api) Update(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysApi
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.Update(&req)
		httpx.OkJson(w, resp)
	}
}

//获取所有api
func (l *Api) GetAll(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewApi(r.Context(), ctx)
		resp := l.GetAll()
		httpx.OkJson(w, resp)
	}
}

//删除选中Api
func (l *Api) DeleteByIds(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IdsReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewApi(r.Context(), ctx)
		resp := l.DeleteByIds(&req)
		httpx.OkJson(w, resp)
	}
}
