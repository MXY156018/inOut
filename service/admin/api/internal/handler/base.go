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

type Base struct{}

func (l *Base) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine)
	{
		r.Post("/base/captcha", l.Captcha(ctx))
		r.Post("/base/login", l.Login(ctx))
	}

}

// 获取图形验证码
func (l *Base) Captcha(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewBase(r.Context(), ctx)
		resp := l.Captcha()
		httpx.OkJson(w, resp)
	}
}

// 登录
func (l *Base) Login(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewBase(r.Context(), ctx)
		resp := l.Login(req)
		httpx.OkJson(w, resp)
	}
}
