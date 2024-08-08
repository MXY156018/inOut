package types

import (
	"mall-admin/model"
)

type CaptchaResp struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}

// 登录请求
type LoginReq struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
	Platform  string `json:"platform"`  //账号平台
}

// 登录回复
type LoginResp struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
