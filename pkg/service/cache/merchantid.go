package cache

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/go-redis/redis"
// )

// // 会员等级定义信息

// type MerchantId struct{}

// //	根据管理员ID 获取 商户ID
// //
// // 是否存在,数据,错误
// func (l *MerchantId) Get(uid int) (int, error) {
// 	key := Ctx.Keys(Key_AdminMerchantId, uid)
// 	res := Ctx.Redis.Get(key)
// 	if res.Err() == redis.Nil {
// 		return 0, nil
// 	}
// 	if res.Err() != nil {
// 		return 0, res.Err()
// 	}
// 	id, err := strconv.ParseInt(res.Val(), 10, 64)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(id), nil
// }

// // 更新缓存
// //
// // 设置管理员商户ID
// func (l *MerchantId) Set(uid int, merchantId int) error {
// 	key := Ctx.Keys(Key_AdminMerchantId, uid)
// 	if merchantId <= 0 {
// 		res := Ctx.Redis.Del(key)
// 		if res.Err() != nil {
// 			return res.Err()
// 		}
// 		return nil
// 	} else {
// 		res := Ctx.Redis.Set(key, fmt.Sprintf("%d", merchantId), 0)
// 		if res.Err() != nil {
// 			return res.Err()
// 		}
// 		return nil
// 	}
// }
