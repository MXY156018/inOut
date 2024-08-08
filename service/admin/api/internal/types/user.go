package types

import (
	"mall-admin/model"
)

const DefaultAuthrityId = "8888"

type ChangeUserInfoReq struct {
	ID           uint                 `gorm:"primarykey"`                                                                                    // 主键ID
	NickName     string               `json:"nickName,optional" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string               `json:"phone,optional"  gorm:"comment:用户手机号"`                                                          // 用户角色ID
	AuthorityIds []string             `json:"authorityIds,optional" gorm:"-"`                                                                // 角色ID
	Email        string               `json:"email,optional"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string               `json:"headerImg,optional" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string               `json:"sideMode,optional"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Authorities  []model.SysAuthority `json:"-,optional" gorm:"many2many:sys_user_authority;"`
}

// 更改个人密码
type ChangePasswordReq struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type SetUserAuthReq struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

type SetUserAuthoritiesReq struct {
	ID           uint
	AuthorityIds string `json:"authorityIds"` // 角色ID
}

type SysuserResp struct {
	UserInfo *model.SysUser `json:"userInfo"`
}

type RegisterReq struct {
	Username     string `json:"userName"`
	Password     string `json:"passWord"`
	NickName     string `json:"nickName" gorm:"default:'QMPlusUser'"`
	AuthorityId  string `json:"authorityId" gorm:"default:888"`
	AuthorityIds string `json:"authorityIds"`
	Platform     string `json:"platform"`
	MerchantId   int    `json:"merchant_id"`
}

type SetSelfInfoReq struct { // 主键ID
	NickName  string `json:"nickName,optional" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone     string `json:"phone,optional"  gorm:"comment:用户手机号"`                                                          // 用户角色ID                                                           // 角色ID
	Email     string `json:"email,optional"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg string `json:"headerImg,optional" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode  string `json:"sideMode,optional"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
}
