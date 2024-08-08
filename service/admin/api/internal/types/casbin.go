package types

import (
	"mall-admin/pkg"
)

type CasbinInReceiveReq struct {
	AuthorityId string           `json:"authorityId"` // 权限id
	CasbinInfos []pkg.CasbinInfo `json:"casbinInfos"`
}

type PolicyPathResp struct {
	Paths []pkg.CasbinInfo `json:"paths"`
}
