package cache

import (
	"fmt"
	"mall-pkg/utils"

	"github.com/go-redis/redis"
)

type Context struct {
	// redis 客户端
	Redis *redis.Client

	// redis 所有的key的前缀，默认无
	Prefix string

	// 等级定义信息
	Grade *Grade
	// 黑名单
	Black *Black
	// 商户ID
	// MerchantId *MerchantId
	// 邀请码
	InviteCode *InviteCode
}

var Ctx = new(Context)

func init() {
	Ctx.Grade = &Grade{}
	Ctx.Black = &Black{}
	Ctx.InviteCode = &InviteCode{}
	// Ctx.MerchantId = &MerchantId{}
}

// 获取 k 值
func (l *Context) Key(k string) string {
	if Ctx.Prefix == "" {
		return k
	}
	return fmt.Sprintf("%s:%s", Ctx.Prefix, k)
}

func (l *Context) Keys(args ...interface{}) string {
	if Ctx.Prefix == "" {
		return utils.RedisFormatKey(args...)
	}
	return Ctx.Prefix + ":" + utils.RedisFormatKey(args...)
}
