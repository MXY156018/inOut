package handler

import (
	"mall-admin/api/internal/svc"
	"mall-pkg/api"
	"mall-pkg/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 待完成
// 1. 管理员退出, token 无效
// - 方案1. 将无效的token缓存到 redis， 认证服务里传 token

type TODO struct{}

func (l *TODO) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/jwt/jsonInBlacklist", l.Logout(ctx))
	}

}

// 使当前token无效
func (l *TODO) Logout(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &api.BaseResp{}
		httpx.OkJson(w, resp)
	}
}
