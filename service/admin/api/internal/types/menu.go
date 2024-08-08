package types

import (
	"mall-admin/model"
)

type SysMenusResp struct {
	Menus []model.SysMenu `json:"menus"`
}

type SysBaseMenusResp struct {
	Menus []model.SysBaseMenu `json:"menus"`
}

// 添加菜单权限
type AddMenuAuthorityReq struct {
	Menus       []model.SysBaseMenu `json:"menus"`
	AuthorityId string              `json:"authorityId"` // 角色ID
}

type SysBaseMenuResp struct {
	Menu *model.SysBaseMenu `json:"menu"`
}

// 默认菜单
var DefaultMenu = []model.SysBaseMenu{
	{
		GormBaseModel: model.GormBaseModel{ID: 1},
		ParentId:      "0",
		Path:          "dashboard",
		Name:          "dashboard",
		Component:     "view/dashboard/index.vue",
		Sort:          1,
		Meta: model.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	},
}
