package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-pkg/api"
	"strconv"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

const menuRoot = "0"

type Menu struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewMenu(ctx context.Context, svcCtx *svc.ServiceContext) Menu {
	return Menu{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

// 根据权限获取菜单树
func (l *Menu) getMenuTree(authorityId string) (map[string][]model.SysMenu, error) {
	// var menus []model.SysMenu
	// treeMap := make(map[string][]model.SysMenu)
	// // TODO authority_menus?
	// err := l.sCtx.DB.Table("authority_menu").Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&menus).Error
	// if err != nil {
	// 	return nil, err
	// }

	// for _, v := range menus {
	// 	treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	// }
	// return treeMap, nil

	var allMenus []model.SysMenu
	var baseMenu []model.SysBaseMenu
	var btns []model.SysAuthorityBtn
	treeMap := make(map[string][]model.SysMenu)
	var sysAuthorityMenus []model.SysAuthorityMenu
	err := l.sCtx.DB.Where("sys_authority_authority_id = ?", authorityId).Find(&sysAuthorityMenus).Error
	if err != nil {
		return nil, err
	}

	var menuIds []string

	for i := range sysAuthorityMenus {
		menuIds = append(menuIds, sysAuthorityMenus[i].MenuId)
	}

	err = l.sCtx.DB.Where("id in (?)", menuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return nil, err
	}

	for i := range baseMenu {
		allMenus = append(allMenus, model.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = l.sCtx.DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return nil, err
	}

	var btnMap = make(map[uint]map[string]string)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]string)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// 获取子菜单
func (l *Menu) getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		l.getChildrenList(&menu.Children[i], treeMap)
	}
}

// 获取菜单树
func (l *Menu) GetMenu() *api.BaseResp {
	var authorityId = l.ctx.Value(api.Middle_Header_AuthorityId).(string)
	resp := &api.BaseResp{}
	tree, err := l.getMenuTree(authorityId)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("get menu", zap.Error(err))
		return resp
	}
	menus := tree[menuRoot]
	for i := 0; i < len(menus); i++ {
		l.getChildrenList(&menus[i], tree)
	}
	resp.Data = &types.SysMenusResp{
		Menus: menus,
	}
	return resp
}

func (l *Menu) getBaseMenuTree() (map[string][]model.SysBaseMenu, error) {
	var menus []model.SysBaseMenu
	tree := make(map[string][]model.SysBaseMenu)
	err := l.sCtx.DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Preload("Parameters").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	for _, v := range menus {
		tree[v.ParentId] = append(tree[v.ParentId], v)
	}
	return tree, nil
}

func (l *Menu) getBaseChildrenList(menu *model.SysBaseMenu, tree map[string][]model.SysBaseMenu) {
	menu.Children = tree[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		l.getBaseChildrenList(&menu.Children[i], tree)
	}
}

//获取基础menu列表
func (l *Menu) GetBaseMenuList() *api.BaseResp {
	resp := &api.BaseResp{}
	tree, err := l.getBaseMenuTree()
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("get base menu", zap.Error(err))
		return resp
	}

	menus := tree[menuRoot]
	for i := 0; i < len(menus); i++ {
		l.getBaseChildrenList(&menus[i], tree)
	}

	resp.Data = &api.PageResp{
		List: menus,
	}

	return resp
}

// 新增菜单
func (l *Menu) AddBaseMenu(req *model.SysBaseMenu) *api.BaseResp {
	resp := &api.BaseResp{}
	var count int64
	err := l.sCtx.DB.Model(req).Where("name = ?", req.Name).Count(&count).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("get base menu", zap.Error(err))
		return resp
	}
	if count > 0 {
		resp.Code = api.Error_Admin_MenuExists
		resp.Msg = "存在重复name，请修改name"
		return resp
	}
	err = l.sCtx.DB.Create(req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("create base menu", zap.Error(err))
		return resp
	}

	return resp
}

// 获取基础菜单树
func (l *Menu) GetBaseMenuTree() *api.BaseResp {
	resp := &api.BaseResp{}
	tree, err := l.getBaseMenuTree()
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("get base menu", zap.Error(err))
		return resp
	}

	menus := tree[menuRoot]
	for i := 0; i < len(menus); i++ {
		l.getBaseChildrenList(&menus[i], tree)
	}

	resp.Data = &types.SysBaseMenusResp{
		Menus: menus,
	}
	return resp
}

// 添加菜单权限
func (l *Menu) AddMenuAuthority(req *types.AddMenuAuthorityReq) *api.BaseResp {
	resp := &api.BaseResp{}
	err := addMenuAuthority(l.sCtx.DB, req.AuthorityId, req.Menus)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}
	return resp
}

// 获取角色菜单
func (l *Menu) GetMenuAuthority(req *types.AuthorityIdReq) *api.BaseResp {
	resp := &api.BaseResp{}
	menus, err := getMenuAuthority(l.sCtx.DB, req.AuthorityId)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}
	resp.Data = &types.SysMenusResp{
		Menus: menus,
	}

	return resp
}

// 删除菜单
func (l *Menu) DeleteBaseMenu(req *api.IDReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var menu model.SysBaseMenu
	err := l.sCtx.DB.Preload("Parameters").Where("parent_id = ?", req.ID).First(menu).Error
	if err == nil {
		resp.Code = api.Error_Admin_MenuHasChild
		resp.Msg = "此菜单存在子菜单不可删除"
		return resp
	}
	err = l.sCtx.DB.Preload("SysAuthoritys").Where("id = ?", req.ID).First(&menu).Delete(&menu).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}
	err = l.sCtx.DB.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", req.ID).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}
	if len(menu.SysAuthoritys) > 0 {
		err = l.sCtx.DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = err.Error()
			l.sCtx.Log.Error("menu", zap.Error(err))
			return resp
		}
	}

	return resp
}

//更新菜单
func (l *Menu) UpdateBaseMenu(req *model.SysBaseMenu) *api.BaseResp {
	resp := &api.BaseResp{}
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = req.Meta.KeepAlive
	upDateMap["close_tab"] = req.Meta.CloseTab
	upDateMap["default_menu"] = req.Meta.DefaultMenu
	upDateMap["parent_id"] = req.ParentId
	upDateMap["path"] = req.Path
	upDateMap["name"] = req.Name
	upDateMap["hidden"] = req.Hidden
	upDateMap["component"] = req.Component
	upDateMap["title"] = req.Meta.Title
	upDateMap["icon"] = req.Meta.Icon
	upDateMap["sort"] = req.Sort

	err := l.sCtx.DB.Where("id = ?", req.ID).Find(&oldMenu).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}

	if oldMenu.Name != req.Name {
		var count int64
		err := l.sCtx.DB.Model(req).Where("id <> ? AND name = ?", req.ID, req.Name).Count(&count).Error
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = err.Error()
			l.sCtx.Log.Error("menu", zap.Error(err))
			return resp
		}
		if count > 0 {
			resp.Code = api.Error_Admin_MenuExists
			resp.Msg = "存在相同name修改失败"
			return resp
		}
	}

	err = l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Unscoped().Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", req.ID).Error
		if merr != nil {
			return merr
		}
		if len(req.Parameters) > 0 {
			for k := range req.Parameters {
				req.Parameters[k].SysBaseMenuID = req.ID
			}
			merr = tx.Create(&req.Parameters).Error
			if merr != nil {
				return merr
			}
		}

		merr = tx.Model(oldMenu).Updates(upDateMap).Error
		if merr != nil {
			return merr
		}
		return nil
	})

	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}

	return resp
}

//根据id获取菜单
func (l *Menu) GetBaseMenuById(req *api.IDReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var menu model.SysBaseMenu
	err := l.sCtx.DB.Preload("Parameters").Where("id = ?", req.ID).First(&menu).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("menu", zap.Error(err))
		return resp
	}
	resp.Data = &types.SysBaseMenuResp{
		Menu: &menu,
	}
	return resp
}
