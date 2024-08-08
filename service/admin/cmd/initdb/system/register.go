package system

import (
	"mall-admin/cmd/initdb/logic"
)

func Register() {
	logic.RegisterInit(initOrderApi, &initApi{})
	logic.RegisterInit(initOrderUser, &initUser{})
	logic.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
	logic.RegisterInit(initOrderAuthority, &initAuthority{})
	logic.RegisterInit(initOrderCasbin, &initCasbin{})
	logic.RegisterInit(initOrderMenu, &initMenu{})
	logic.RegisterInit(initOrderMenu, &initAuthorityBtn{})
}
