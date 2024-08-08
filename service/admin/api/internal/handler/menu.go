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

type Menu struct{}

// 注册处理函数
func (l *Menu) Router(engine *rest.Server, ctx *svc.ServiceContext) {
	r := utils.NewRouter(engine).Use(ctx.MiddleJwt, ctx.MiddleCasbin)
	{
		r.Post("/menu/getMenu", l.GetMenu(ctx))
		r.Post("/menu/getMenuList", l.GetMenuList(ctx))
		r.Post("/menu/addBaseMenu", l.AddBaseMenu(ctx))
		r.Post("/menu/getBaseMenuTree", l.GetBaseMenuTree(ctx))
		r.Post("/menu/addMenuAuthority", l.AddMenuAuthority(ctx))
		r.Post("/menu/getMenuAuthority", l.GetMenuAuthority(ctx))
		r.Post("/menu/deleteBaseMenu", l.DeleteBaseMenu(ctx))
		r.Post("/menu/updateBaseMenu", l.UpdateBaseMenu(ctx))
		r.Post("/menu/getBaseMenuById", l.GetBaseMenuById(ctx))
	}
}

// 获取菜单
func (l *Menu) GetMenu(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.GetMenu()
		httpx.OkJson(w, resp)
	}
}

// 基础menu列表
func (l *Menu) GetMenuList(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.GetBaseMenuList()
		httpx.OkJson(w, resp)
	}
}

//新增菜单
func (l *Menu) AddBaseMenu(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysBaseMenu
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.AddBaseMenu(&req)
		httpx.OkJson(w, resp)
	}
}

//获取用户动态路由
func (l *Menu) GetBaseMenuTree(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.GetBaseMenuTree()
		httpx.OkJson(w, resp)
	}
}

//增加menu和角色关联关系
func (l *Menu) AddMenuAuthority(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddMenuAuthorityReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.AddMenuAuthority(&req)
		httpx.OkJson(w, resp)
	}
}

//获取指定角色menu
func (l *Menu) GetMenuAuthority(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthorityIdReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.GetMenuAuthority(&req)
		httpx.OkJson(w, resp)
	}
}

//删除菜单
func (l *Menu) DeleteBaseMenu(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.DeleteBaseMenu(&req)
		httpx.OkJson(w, resp)
	}
}

//更新菜单
func (l *Menu) UpdateBaseMenu(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SysBaseMenu
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.UpdateBaseMenu(&req)
		httpx.OkJson(w, resp)
	}
}

//根据id获取菜单
func (l *Menu) GetBaseMenuById(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.IDReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.OkJson(w, api.NewInvalidParameter(err.Error()))
			return
		}
		l := logic.NewMenu(r.Context(), ctx)
		resp := l.GetBaseMenuById(&req)
		httpx.OkJson(w, resp)
	}
}
