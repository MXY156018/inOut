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

type Out struct{}

// 注册处理函数
func (l *Out) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine)
	{
		r.Post("/out/add", l.OutAdd(ctx))
		r.Post("/out/get", l.OutGet(ctx))
		r.Post("/out/del", l.OutDel(ctx))
	}
}

//增加记录
func (l *Out) OutAdd(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OutRecordReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewOut(r.Context(), ctx)
		resp := l.AddOutRecord(&req)
		httpx.OkJson(w, resp)
	}
}

// 获取记录
func (l *Out) OutGet(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRecordsReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewOut(r.Context(), ctx)
		resp := l.GetOutRecords(&req)
		httpx.OkJson(w, resp)
	}
}

// 删除记录
func (l *Out) OutDel(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewOut(r.Context(), ctx)
		resp := l.DelOutRecords(&req)
		httpx.OkJson(w, resp)
	}
}
