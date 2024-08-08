package handler

import (
	"mall-admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

type register interface {
	Router(engine *rest.Server, ctx *svc.ServiceContext)
}

func esayRegister(engine *rest.Server, ctx *svc.ServiceContext, r register) {
	r.Router(engine, ctx)
}

// 注册处理函数
func Register(engine *rest.Server, ctx *svc.ServiceContext) {
	esayRegister(engine, ctx, &Base{})
	esayRegister(engine, ctx, &Api{})
	esayRegister(engine, ctx, &Authority{})
	esayRegister(engine, ctx, &User{})
	esayRegister(engine, ctx, &Menu{})
	esayRegister(engine, ctx, &Casbin{})
	esayRegister(engine, ctx, &TODO{})
	esayRegister(engine, ctx, &Out{})
	esayRegister(engine, ctx, &In{})
}
