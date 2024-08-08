package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type InviteCode struct {
}

//生成随机邀请码

// 六位数字加字母

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

func (i *InviteCode) GenInviteCode(uid int) string {
	var u = fmt.Sprintf("%d", uid)
	if uid <= 9 {
		u = "0" + u
	}
	b := make([]rune, 4)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	var str = string(b) + u

	return str
}

func (i *InviteCode) KeyToUid(key string) (bool, int64, error) {
	var k = Ctx.Keys(Key_UserInviteCode, key)
	res := Ctx.Redis.Get(k)
	if res.Err() != redis.Nil {
		var uidstr = res.Val()
		if uidstr == "" {
			return false, 0, nil
		}
		uid, _ := strconv.ParseInt(uidstr, 10, 32)
		return true, uid, nil
	}
	if res.Err() != nil {
		return false, 0, res.Err()
	}
	return false, 0, nil
}

func (i *InviteCode) SaveToRedis(key string, uid int) error {
	var k = Ctx.Keys(Key_UserInviteCode, key)
	res := Ctx.Redis.Set(k, fmt.Sprintf("%d", uid), 0)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}