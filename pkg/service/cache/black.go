package cache

import (
	"github.com/go-redis/redis"
)

// 会员等级定义信息

type Black struct{}

//  是否在黑名单中
//
//是否存在,数据,错误
func (l *Black) Get(uid int) (bool, error) {
	res := Ctx.Redis.Get(Ctx.Keys(Key_UserBlack, uid))
	if res.Err() == redis.Nil {
		return false, nil
	}
	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}

// 更新缓存
//
//data 数据
func (l *Black) Set(uid int, isBlack bool) error {
	key := Ctx.Keys(Key_UserBlack, uid)
	if !isBlack {
		Ctx.Redis.Del(key)
		return nil
	}

	return Ctx.Redis.Set(key, "1", 0).Err()
}

//  是否在黑名单中[商户]
//
//是否存在,数据,错误
func (l *Black) GetMerchant(uid int) (bool, error) {
	res := Ctx.Redis.Get(Ctx.Keys(Key_MerchantBlack, uid))
	if res.Err() == redis.Nil {
		return false, nil
	}
	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}

// 更新缓存 [商户]
//
//data 数据
func (l *Black) SetMerchant(uid int, isBlack bool) error {
	key := Ctx.Keys(Key_MerchantBlack, uid)
	if !isBlack {
		Ctx.Redis.Del(key)
		return nil
	}

	return Ctx.Redis.Set(key, "1", 0).Err()
}
