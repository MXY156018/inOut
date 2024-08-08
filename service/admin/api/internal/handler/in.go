package handler

import (
	"mall-admin/api/internal/logic"
	"mall-admin/api/internal/svc"
	"mall-admin/types"
	"mall-pkg/api"
	"mall-pkg/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type In struct{}

// 注册处理函数
func (l *In) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine)
	{
		r.Post("/in/add", l.InAdd(ctx))
		r.Post("/in/get", l.InGet(ctx))
		r.Post("/in/del", l.InDel(ctx))
		r.Post("/in/getType", l.InGetType(ctx))
		r.Post("/in/getTypeDetail", l.InGetTypeDetail(ctx))
	}
}

//增加记录
func (l *In) InAdd(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InRecordReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewIn(r.Context(), ctx)
		resp := l.AddInRecord(&req)
		httpx.OkJson(w, resp)
	}
}

// 获取记录
func (l *In) InGet(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRecordsReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewIn(r.Context(), ctx)
		resp := l.GetInRecords(&req)
		httpx.OkJson(w, resp)
	}
}

// 删除记录
func (l *In) InDel(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewIn(r.Context(), ctx)
		resp := l.DelInRecords(&req)
		httpx.OkJson(w, resp)
	}
}
func (l *In) InGetType(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIn(r.Context(), ctx)
		resp := l.InGetType()
		httpx.OkJson(w, resp)
	}
}

func (l *In) InGetTypeDetail(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TypeDetailReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewIn(r.Context(), ctx)
		resp := l.InGetTypeDetail(&req)
		httpx.OkJson(w, resp)
	}
}
