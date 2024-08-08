package svc

import (
	"mall-admin/api/internal/middleware"
	"mall-admin/pkg"
	"mall-pkg/api"
	"mall-pkg/db"

	// "mall-pkg/service/cache"

	"github.com/go-redis/redis"

	// "github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/go-zero/rest"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// API 上下文
type ServiceContext struct {
	// 配置
	Config pkg.ApiConfig
	// 日志
	Log *zap.Logger
	//数据库
	DB *gorm.DB
	// Redis
	Redis *redis.Client

	Casbin *pkg.CasbinService
	// casbin 中间件
	MiddleCasbin rest.Middleware
	// JWT 中间件
	MiddleJwt rest.Middleware
}

func NewServiceContext(c pkg.ApiConfig, log *zap.Logger) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Log:    log,
	}
}

// 初始化
func (l *ServiceContext) Init() error {
	err := l.initMysql()
	if err != nil {
		return err
	}
	// casbin 配置
	pkg.Casbin.DB = l.DB
	pkg.Casbin.ModelPath = l.Config.Casbin.ModelPath
	l.Casbin = pkg.Casbin

	err = l.initMiddleware()
	if err != nil {
		return err
	}
	err = l.initRedis()
	return err
}

// 初始化数据库
func (l *ServiceContext) initMysql() error {
	mdb, err := db.NewGorm(l.Config.Mysql, l.Log)
	if err != nil {
		return err
	}
	l.DB = mdb
	return nil
}

// 初始化Redis
func (l *ServiceContext) initRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     l.Config.Redis.Addr,
		Password: l.Config.Redis.Password, // no password set
		DB:       l.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	l.Redis = client
	// cache.Ctx.Redis = client
	return nil
}

//初始化中间件
func (l *ServiceContext) initMiddleware() error {
	var jwtMiddle = api.AdminJwt{
		Secret:      l.Config.JWT.Secret,
		ExpiresTime: l.Config.JWT.ExpiresTime,
	}
	l.MiddleJwt = jwtMiddle.Middleware
	l.MiddleCasbin = middleware.Casbin
	return nil
}
