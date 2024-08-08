package pkg

import (
	"mall-pkg/config"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ApiConfig struct {
	rest.RestConf
	// mysql 配置
	Mysql config.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// Zap 日志配置
	Zap config.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Captcha 配置
	Captcha config.Captcha `json:"captcha"`
	// JWT 配置
	JWT config.Jwt `json:"jwt"`
	// casbin 配置
	Casbin config.Casbin `json:"casbin"`
	//默认密码
	DefaultPassword string `json:"defaultPassword"`
	//Redis
	Redis config.Redis `json:"redis"`
}

type RpcConfig struct {
	zrpc.RpcServerConf
}
