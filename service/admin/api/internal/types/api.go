package types

import (
	"mall-admin/model"
	"mall-pkg/api"
)

type SysApiResp struct {
	Api *model.SysApi `json:"api"`
}

type SysApiListResp struct {
	Apis []model.SysApi `json:"apis"`
}

type ApiSearchReq struct {
	model.SysApi
	api.PageQuery
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
