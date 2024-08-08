package model

type SysUserAuthority struct {
	SysUserId               uint   `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId string `gorm:"column:sys_authority_authority_id"`
}
