package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-pkg/api"
	mjwt "mall-pkg/jwt"

	// "mall-pkg/service/cache"
	"mall-pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 用户管理相关功能
type User struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewUser(ctx context.Context, svcCtx *svc.ServiceContext) User {
	return User{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

// 用户注册
func (l *User) Register(req *types.RegisterReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var count int64
	err := l.sCtx.DB.Model(&model.SysUser{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("获取用户信息", zap.Error(err))
		return resp
	}
	if count > 0 {
		resp.Code = api.Error_Admin_UserExists
		resp.Msg = "用户名已注册"
		return resp
	}
	var authorities []model.SysAuthority
	authorities = append(authorities, model.SysAuthority{
		AuthorityId: req.AuthorityIds,
	})
	encPwd := utils.BcryptHash(req.Password)
	user := &model.SysUser{
		Username:    req.Username,
		NickName:    req.NickName,
		Password:    encPwd,
		AuthorityId: req.AuthorityIds,
		Authorities: authorities,
		Platform:    req.Platform,
		MerchantId:  req.MerchantId,
	}
	if user.AuthorityId == "" {
		user.AuthorityId = types.DefaultAuthrityId
	}
	user.UUID = uuid.NewV4()
	err = l.sCtx.DB.Omit("id").Create(user).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("创建用户", zap.Error(err))
		return resp
	}
	//设置商户-用户账号映射
	// cache.Ctx.MerchantId.Set(int(user.ID), req.MerchantId)
	resp.Data = user
	return resp
}

// 设置个人信息
func (l *User) SetSelfInfo(req *types.SetSelfInfoReq) *api.BaseResp {
	resp := &api.BaseResp{}
	//管理员 UID
	uid := l.ctx.Value(api.Context_Key_UID).(int)
	err := l.sCtx.DB.Updates(&model.SysUser{
		GormBaseModel: model.GormBaseModel{ID: uint(uid)},
		NickName:      req.NickName,
		HeaderImg:     req.HeaderImg,
		SideMode:      req.SideMode,
	}).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("重置密码", zap.Error(err))
		return resp
	}
	return resp
}

// 重置密码
func (l *User) ResetPassword(req *api.IDReq) *api.BaseResp {
	resp := &api.BaseResp{}
	encPwd := utils.BcryptHash(l.sCtx.Config.DefaultPassword)
	err := l.sCtx.DB.Model(&model.SysUser{}).Where("id = ?", req.ID).Update("password", encPwd).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("重置密码", zap.Error(err))
		return resp
	}
	return resp
}

// 更改个人密码
func (l *User) ChangePassword(req *types.ChangePasswordReq) *api.BaseResp {
	resp := &api.BaseResp{}
	var user model.SysUser
	err := l.sCtx.DB.Where("username = ?", req.Username).Select("id,password").First(&user).Error
	if err == gorm.ErrRecordNotFound {
		resp.Code = api.Error_Admin_UserNotExists
		resp.Msg = "用户不存在"
		return resp
	}
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("更改密码错误", zap.Error(err))
		return resp
	}

	isOk := utils.BcryptCheck(req.Password, user.Password)
	if !isOk {
		resp.Code = api.Error_Admin_Password
		resp.Msg = "密码错误"
		return resp
	}

	encPwd := utils.BcryptHash(req.NewPassword)
	err = l.sCtx.DB.Model(&user).Where("id = ?", user.ID).Update("password", encPwd).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("更改密码错误", zap.Error(err))
		return resp
	}
	return resp
}

// 获取用户列表
func (l *User) GetUserList(req *api.PageQuery) *api.BaseResp {
	resp := &api.BaseResp{}
	var count int64
	var users []model.SysUser
	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize
	err := l.sCtx.DB.Model(users).Count(&count).Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&users).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("更改密码错误", zap.Error(err))
		return resp
	}
	resp.Data = &api.PageResp{
		List:  users,
		Total: count,
	}
	return resp
}

// 设置个人权限(管理员)
func (l *User) SetSelfAuthority(req *types.SetUserAuthReq, r *http.Request) *api.BaseResp {
	uid := l.ctx.Value(api.Context_Key_UID).(int)
	resp := &api.BaseResp{}
	err := l.sCtx.DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", uid, req.AuthorityId).Error
	if err == gorm.ErrRecordNotFound {
		resp.Code = api.Error_Admin_UserNotExists
		resp.Msg = "用户不存在"
		return resp
	}
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("设置管理员权限错误", zap.Error(err))
		return resp
	}
	err = l.sCtx.DB.Model(&model.SysUser{}).Where("id=?", uid).Update("authority_id", req.AuthorityId).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("更改密码错误", zap.Error(err))
		return resp
	}

	expire := l.sCtx.Config.JWT.ExpiresTime
	claims := mjwt.AdminClaims{
		UserID:      int(uid),
		AuthorityId: req.AuthorityId,
		BufferTime:  l.sCtx.Config.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + expire,
		},
	}
	token, err := mjwt.GetAdminToken(l.sCtx.Config.JWT.Secret, claims)
	if err != nil {
		l.sCtx.Log.Error("生成Token", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		return resp
	}
	r.Header.Set(api.Resp_Header_NewToken, token)
	r.Header.Set(api.Resp_Header_NewTokenExpire, strconv.FormatInt(claims.ExpiresAt, 10))
	return resp
}

// 删除用户
func (l *User) DeleteUser(req *api.IDReq) *api.BaseResp {
	resp := &api.BaseResp{}
	uid := l.ctx.Value(api.Context_Key_UID).(int)
	if uid == req.ID {
		resp.Code = api.Error_Admin_DelSelf
		resp.Msg = "不能删除自己"
		return resp
	}
	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Where("id=?", req.ID).Delete(&model.SysUser{}).Error
		if merr != nil {
			return merr
		}
		merr = tx.Table("sys_user_authority").Where("sys_user_id = ?", req.ID).Delete(&model.SysUserAuthority{}).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("删除用户", zap.Error(err))
		return resp
	}
	return resp
}

// 设置用户信息
func (l *User) SetUserInfo(req *model.SysUser) *api.BaseResp {
	resp := &api.BaseResp{}
	err := l.sCtx.DB.Updates(req).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("删除用户", zap.Error(err))
		return resp
	}
	return resp
}

// 设置用户权限组
func (l *User) SetUserAuthorities(req *types.SetUserAuthoritiesReq) *api.BaseResp {
	resp := &api.BaseResp{}

	var useAuthority model.SysUserAuthority
	useAuthority.SysUserId = req.ID
	useAuthority.SysAuthorityAuthorityId = req.AuthorityIds
	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Table("sys_user_authority").Delete(&[]model.SysUserAuthority{}, "sys_user_id = ?", req.ID).Error
		if merr != nil {
			return merr
		}
		merr = tx.Table("sys_user_authority").Create(&useAuthority).Error
		if merr != nil {
			return merr
		}
		merr = tx.Table("sys_users").Where("id = ?", req.ID).Update("authority_id", req.AuthorityIds).Error
		return merr
	})
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("删除用户", zap.Error(err))
		return resp
	}
	return resp
}

// 获取用户信息
func (l *User) GetUserInfo() *api.BaseResp {
	resp := &api.BaseResp{}
	var user model.SysUser
	uid := l.ctx.Value(api.Context_Key_UID).(int)
	err := l.sCtx.DB.Preload("Authorities").Preload("Authority").First(&user, "id = ?", uid).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("删除用户", zap.Error(err))
		return resp
	}
	resp.Data = &types.SysuserResp{
		UserInfo: &user,
	}
	return resp
}
