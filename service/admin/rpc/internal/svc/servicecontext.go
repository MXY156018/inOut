package svc

import (
	"mall-admin/pkg"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Config pkg.RpcConfig
	Log *zap.Logger
}

func NewServiceContext(c pkg.RpcConfig, log *zap.Logger) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Log: log,
	}
}
