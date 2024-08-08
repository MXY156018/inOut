package logic

import (
	"context"

	"mall-admin/pkg"
	"mall-admin/rpc/admin"
	"mall-admin/rpc/internal/svc"
	"mall-pkg/api"
)

type ApiAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiAuthLogic {
	return &ApiAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// api 授权
func (l *ApiAuthLogic) ApiAuth(in *admin.ApiAuthReq) (*admin.ApiAuthResp, error) {
	// todo: add your logic here and delete this line
	resp := &admin.ApiAuthResp{}
	// 获取请求的URI
	obj := in.Path
	// 获取请求方法
	act := in.Method
	// 获取用户的角色
	sub := in.AuthorityId
	e := pkg.Casbin.Casbin()
	success, err := e.Enforce(sub, obj, act)
	if err != nil {
		resp.Code = int32(api.Error_Server)
		resp.Message = err.Error()
		return resp, nil
	}
	if !success {
		resp.Code = int32(api.Error_NoPrivilege)
		resp.Message = "无权限"
	}
	return resp, nil
}
