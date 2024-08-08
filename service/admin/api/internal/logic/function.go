package logic

import (
	"errors"
	"mall-pkg/api"
	"strconv"

	"mall-admin/model"
	"mall-admin/pkg"

	"gorm.io/gorm"
)

// 添加菜单权限
func addMenuAuthority(db *gorm.DB, authorityId string, menus []model.SysBaseMenu) error {

	var s model.SysAuthority
	err := db.Preload("SysBaseMenus").First(&s, "authority_id = ?", authorityId).Error
	if err != nil {
		return err
	}

	err = db.Model(&s).Association("SysBaseMenus").Replace(&menus)
	if err != nil {
		return err
	}

	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: menus []system.SysMenu, err error
func getMenuAuthority(db *gorm.DB, authorityId string) ([]model.SysMenu, error) {
	var baseMenu []model.SysBaseMenu
	var SysAuthorityMenus []model.SysAuthorityMenu
	err := db.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return nil, err
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}
	var menus []model.SysMenu
	err = db.Where("id in (?) ", MenuIds).Order("sort").Find(&baseMenu).Error
	for i := range baseMenu {
		menus = append(menus, model.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}
	// sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	// err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return menus, err
	//?
	// var menus []model.SysMenu
	// err := db.Table("authority_menu").Where("authority_id = ? ", authorityId).Order("sort").Find(&menus).Error
	// return menus, err
}

//删除角色
func deleteAuthority(db *gorm.DB, authorityId string) (int, error) {
	var count int64
	err := db.Model(&model.SysAuthority{}).Where("parent_id = ?", authorityId).Count(&count).Error
	if err != nil {
		return api.Error_Server, err
	}
	if count > 0 {
		return api.Error_Admin_AuthorityHasChild, errors.New("此角色存在子角色不允许删除")
	}
	var auth model.SysAuthority
	mdb := db.Preload("SysBaseMenus").Where("authority_id = ?", authorityId).First(&auth)
	if mdb.Error != nil {
		return api.Error_Server, err
	}
	err = mdb.Unscoped().Delete(auth).Error
	if err != nil {
		return api.Error_Server, err
	}
	if len(auth.SysBaseMenus) > 0 {
		err = db.Model(&auth).Association("SysBaseMenus").Delete(&auth.SysBaseMenus)
		if err != nil {
			return api.Error_Server, err
		}
	}
	err = db.Table("sys_user_authority").Delete(&model.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	if err != nil {
		return api.Error_Server, err
	}
	pkg.Casbin.ClearCasbin(0, auth.AuthorityId)
	return api.Error_OK, nil
}
