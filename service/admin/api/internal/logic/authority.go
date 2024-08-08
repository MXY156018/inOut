package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-admin/pkg"
	"mall-pkg/api"
	"strconv"

	"go.uber.org/zap"
)

type Authority struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewAuthority(ctx context.Context, svcCtx *svc.ServiceContext) Authority {
	return Authority{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

// 新建
func (l *Authority) Create(req *model.SysAuthority) *api.BaseResp {
	resp := &api.BaseResp{}

	var count int64
	err := l.sCtx.DB.Model(&req).Where("authority_id = ?", req.AuthorityId).Count(&count).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	if count > 0 {
		resp.Code = api.Error_Admin_AuthorityExists
		resp.Msg = "角色权限已经存在"
		return resp
	}

	err = l.sCtx.DB.Create(&req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}

	// 默认菜单
	req.SysBaseMenus = types.DefaultMenu
	err = addMenuAuthority(l.sCtx.DB, req.AuthorityId, req.SysBaseMenus)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}
	// 默认 API
	err = l.sCtx.Casbin.UpdateCasbin(req.AuthorityId, pkg.DefaultCasbin)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}

	resp.Data = &types.SysAuthorityResp{
		Authority: req,
	}
	return resp
}

// 删除角色
func (l *Authority) Delete(req *types.AuthorityIdReq) *api.BaseResp {
	resp := &api.BaseResp{}
	code, err := deleteAuthority(l.sCtx.DB, req.AuthorityId)
	resp.Code = code
	if err != nil {
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
	}
	return resp
}

//更改一个角色
func (l *Authority) Update(req *model.SysAuthority) *api.BaseResp {
	resp := &api.BaseResp{}

	err := l.sCtx.DB.Where("authority_id = ?", req.AuthorityId).First(&model.SysAuthority{}).Updates(req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	resp.Data = &types.SysAuthorityResp{
		Authority: req,
	}

	return resp
}

//复制角色权限
func (l *Authority) Copy(req *types.SysAuthorityCopy) *api.BaseResp {
	resp := &api.BaseResp{}

	authorityId := req.Authority.AuthorityId
	var count int64
	err := l.sCtx.DB.Model(&model.SysAuthority{}).Where("authority_id = ?", authorityId).Count(&count).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}
	if count > 0 {
		resp.Code = api.Error_Admin_AuthorityExists
		resp.Msg = "存在相同角色id"
		return resp
	}

	// 获取菜单
	menus, err := getMenuAuthority(l.sCtx.DB, req.OldAuthorityId)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}

	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	req.Authority.SysBaseMenus = baseMenu
	req.Authority.Children = []model.SysAuthority{}
	err = l.sCtx.DB.Create(&req.Authority).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}

	paths := l.sCtx.Casbin.GetPolicyPathByAuthorityId(req.OldAuthorityId)
	err = l.sCtx.Casbin.UpdateCasbin(req.Authority.AuthorityId, paths)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}
	return resp
}

func (l *Authority) findChildrenAuthority(authority *model.SysAuthority) error {
	err := l.sCtx.DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(authority.Children); i++ {
		err = l.findChildrenAuthority(&authority.Children[i])
		if err != nil {
			return err
		}
	}
	return nil
}

//分页获取数据
func (l *Authority) GetAuthorityList(req *api.PageQuery) *api.BaseResp {
	resp := &api.BaseResp{}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var authority []model.SysAuthority
	err := l.sCtx.DB.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}
	for i := 0; i < len(authority); i++ {
		err = l.findChildrenAuthority(&authority[i])
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = err.Error()
			l.sCtx.Log.Error("authority", zap.Error(err))
			return resp
		}
	}
	resp.Code = api.Error_OK
	resp.Data = api.PageResp{
		List:     authority,
		Total:    0,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp
}

func (l *Authority) SetDataAuthority(req *model.SysAuthority) *api.BaseResp {
	resp := &api.BaseResp{}
	var s model.SysAuthority
	err := l.sCtx.DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", req.AuthorityId).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}

	err = l.sCtx.DB.Model(&s).Association("DataAuthorityId").Replace(&req.DataAuthorityId)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}

	return resp
}

func (l *Authority) GetAuthorityInfo(req *types.AuthorityIdReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var s model.SysAuthority
	err := l.sCtx.DB.Preload("DataAuthorityId").Where("authority_id = ?", req.AuthorityId).First(&s).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("authority", zap.Error(err))
		return resp
	}

	resp.Data = &s
	return resp
}
