package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-pkg/api"
	"mall-pkg/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Api struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewApi(ctx context.Context, svcCtx *svc.ServiceContext) Api {
	return Api{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

// 新建API
func (l *Api) Create(req *model.SysApi) *api.BaseResp {
	resp := &api.BaseResp{}
	var count int64
	err := l.sCtx.DB.Model(&req).Where("path = ? AND method = ?", req.Path, req.Method).Count(&count).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	if count > 0 {
		resp.Code = api.Error_Admin_ApiNotExists
		resp.Msg = "API已经存在"
		return resp
	}

	err = l.sCtx.DB.Create(&req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}

	return resp
}

//删除基础api
func (l *Api) Delete(req *model.SysApi) *api.BaseResp {
	resp := &api.BaseResp{}
	err := utils.Verify(*req, IdVerify)
	if err != nil {
		resp.Code = api.Error_Parameter
		resp.Msg = err.Error()
		return resp
	}

	err = l.sCtx.DB.Table("sys_apis").Where("id=?", req.ID).Update("deleted_at", time.Now()).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}

	l.sCtx.Casbin.ClearCasbin(1, req.Path, req.Method)
	return resp
}

// 查询 API 信息
func (l *Api) SearchApi(req *types.ApiSearchReq) *api.BaseResp {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	resp := &api.BaseResp{}
	db := l.sCtx.DB.WithContext(context.Background())
	db = db.Table("sys_apis")
	var apiList []model.SysApi

	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Description != "" {
		db = db.Where("description LIKE ?", "%"+req.Description+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group LIKE  ?", "%"+req.ApiGroup+"%")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	db = db.Limit(limit).Offset(offset)
	if req.OrderKey != "" {
		var OrderStr string
		if req.Desc {
			OrderStr = req.OrderKey + " desc"
		} else {
			OrderStr = req.OrderKey
		}
		err = db.Order(OrderStr).Find(&apiList).Error
	} else {
		err = db.Order("api_group").Find(&apiList).Error
	}
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	resp.Data = &api.PageResp{
		List:  apiList,
		Total: total,
	}

	return resp
}

// 根据 API 获取 API信息
func (l *Api) GetById(req *api.IDReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var data model.SysApi
	err := l.sCtx.DB.Where("id = ?", req.ID).First(&data).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	resp.Data = &types.SysApiResp{
		Api: &data,
	}

	return resp
}

//根据id更新api
func (l *Api) Update(req *model.SysApi) *api.BaseResp {
	resp := &api.BaseResp{}
	var oldA model.SysApi
	err := l.sCtx.DB.Where("id = ?", req.ID).First(&oldA).Error
	if err == gorm.ErrRecordNotFound {
		resp.Code = api.Error_Admin_ApiNotExists
		resp.Msg = "API不存在"
		return resp
	}

	if oldA.Path != req.Path || oldA.Method != req.Method {
		var count int64
		err2 := l.sCtx.DB.Model(&model.SysApi{}).Where("path = ? AND method = ?", req.Path, req.Method).Count(&count).Error
		if err2 != nil {
			resp.Code = api.Error_Server
			resp.Msg = err2.Error()
			l.sCtx.Log.Error("api", zap.Error(err2))
			return resp
		}
		if count > 0 {
			resp.Code = api.Error_Admin_ApiExists
			resp.Msg = "存在相同api路径"
			return resp
		}
	}

	err = l.sCtx.Casbin.UpdateCasbinApi(oldA.Path, req.Path, req.Method, req.Method)
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	err = l.sCtx.DB.Save(&req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	return resp
}

// 获取所有API
func (l *Api) GetAll() *api.BaseResp {
	resp := &api.BaseResp{}
	var data []model.SysApi
	err := l.sCtx.DB.Find(&data).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	resp.Data = types.SysApiListResp{
		Apis: data,
	}
	return resp
}

// 删除 API
func (l *Api) DeleteByIds(req *api.IdsReq) *api.BaseResp {
	resp := &api.BaseResp{}
	err := l.sCtx.DB.Delete(&model.SysApi{}, "id in ? ", req.Ids).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("api", zap.Error(err))
		return resp
	}
	return resp
}
