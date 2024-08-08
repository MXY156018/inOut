package types

import (
	"mall-admin/model"
)

type SysAuthorityResp struct {
	Authority *model.SysAuthority `json:"authority"`
}

type SysAuthorityCopy struct {
	Authority      *model.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"oldAuthorityId"` // 旧角色ID
}

type AuthorityIdReq struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}
