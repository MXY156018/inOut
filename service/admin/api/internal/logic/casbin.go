package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-pkg/api"

	"go.uber.org/zap"
)

type Casbin struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewCasbin(ctx context.Context, svcCtx *svc.ServiceContext) Casbin {
	return Casbin{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

func (l *Casbin) Update(req *types.CasbinInReceiveReq) *api.BaseResp {
	resp := &api.BaseResp{}

	err := l.sCtx.Casbin.UpdateCasbin(req.AuthorityId, req.CasbinInfos)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("casbin", zap.Error(err))
		return resp
	}

	return resp
}

func (l *Casbin) GetPolicyPathByAuthorityId(req *types.AuthorityIdReq) *api.BaseResp {
	resp := &api.BaseResp{}

	paths := l.sCtx.Casbin.GetPolicyPathByAuthorityId(req.AuthorityId)
	resp.Data = &types.PolicyPathResp{
		Paths: paths,
	}

	return resp
}
