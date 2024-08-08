package model

type SysAuthorityBtn struct {
	AuthorityId      string         `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}

func (l *SysAuthorityBtn) TableName() string {
	return "sys_authority_btns"
}
