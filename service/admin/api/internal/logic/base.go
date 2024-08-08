package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	types "mall-admin/api/internal/types"
	"mall-admin/model"
	"mall-pkg/api"
	mjwt "mall-pkg/jwt"
	"mall-pkg/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaStore = base64Captcha.DefaultMemStore

type Base struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewBase(ctx context.Context, svcCtx *svc.ServiceContext) Base {
	return Base{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

// 获取图形验证码
func (l *Base) Captcha() *api.BaseResp {
	resp := &api.BaseResp{}
	cfg := &l.sCtx.Config.Captcha
	driver := base64Captcha.NewDriverDigit(cfg.ImgHeight, cfg.ImgWidth, cfg.KeyLong, cfg.MaxSkew, cfg.DotCount)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = err.Error()
		l.sCtx.Log.Error("", zap.Error(err))
		return resp
	}
	var cptData types.CaptchaResp
	cptData.CaptchaId = id
	cptData.PicPath = b64s
	resp.Data = &cptData
	return resp
}

// 登录
func (l *Base) Login(req types.LoginReq) *api.BaseResp {
	resp := &api.BaseResp{}
	isok := captchaStore.Verify(req.CaptchaId, req.Captcha, true)
	if !isok {
		resp.Code = api.Error_CaptchaInvalid
		resp.Msg = "验证码错误"
		return resp
	}

	var user model.SysUser
	// MYSQL 注入问题?
	err := l.sCtx.DB.Where("username = ?", req.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err != nil {
		resp.Code = api.Error_AccountOrPassword
		resp.Msg = err.Error()
		l.sCtx.Log.Error("", zap.Error(err))
		return resp
	}
	if !utils.BcryptCheck(req.Password, user.Password) {
		resp.Code = api.Error_AccountOrPassword
		resp.Msg = "用户名或者密码错误"
		return resp
	}
	expire := l.sCtx.Config.JWT.ExpiresTime
	claims := mjwt.AdminClaims{
		UserID:      int(user.ID),
		AuthorityId: user.AuthorityId,
		BufferTime:  l.sCtx.Config.JWT.BufferTime,
		MerchantId:  user.MerchantId,
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
	resp.Data = &types.LoginResp{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}

	return resp
}
