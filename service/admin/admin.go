package main

import (
	"flag"
	"log"

	apiw "mall-admin/api/wrap"
	"mall-admin/pkg"
	mlog "mall-pkg/log"

	"github.com/zeromicro/go-zero/core/conf"
)

// 配置
type Config struct {
	// API 配置
	Api pkg.ApiConfig `json:"api"`
	// RPC 配置
	Rpc pkg.RpcConfig `json:"rpc"`
}

var configFile = flag.String("f", "etc/mall_admin.yaml", "the api config file")

// 管理员 核心模块
// RPC 和 API 在同一个进程
// 核心功能主要提供基础的功能，如权限/菜单配置等
// 核心功能不依赖任何服务，并 提供  管理员 API RPC 鉴权接口
// RPC 和 API 放同一进程主要是基于 API 接口鉴权 考虑
func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)

	// 日志
	logInst := mlog.NewZap(c.Api.Zap)
	// RPC 服务
	// go func() {
	// 	time.Sleep(time.Second)
	// 	err := rpcw.Start(c.Rpc, logInst)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// }()

	err := apiw.Start(c.Api, logInst)
	if err != nil {
		log.Panic(err)
	}
}
