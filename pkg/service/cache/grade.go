package cache

import (
	"encoding/json"
	"fmt"
	"mall-pkg/service/model"
	"strconv"

	"github.com/go-redis/redis"
)

// 会员等级定义信息

type Grade struct{}

// 从 redis 缓存中获取
//
//是否存在,数据,错误
func (l *Grade) Get() (bool, []model.Grade, error) {
	res := Ctx.Redis.Get(Ctx.Key(Key_GradeInfo))
	if res.Err() == redis.Nil {
		return false, nil, nil
	}
	if res.Err() != nil {
		return false, nil, res.Err()
	}

	var data []model.Grade
	err := json.Unmarshal([]byte(res.Val()), &data)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil
}

func (l *Grade) DefaultGrade() (int, error) {
	ok, data, err := l.Get()
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("等级信息不存在")
	}
	for i := 0; i < len(data); i++{
		item := &data[i]
		if item.IsDefault != 0 {
			return int(item.GradeId), nil
		}
	}

	return 0, fmt.Errorf("未配置默认等级")
}

// 更新缓存
//
//data 数据
func (l *Grade) Update(data []model.Grade) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return Ctx.Redis.Set(Ctx.Key(Key_GradeInfo), string(str), 0).Err()
}

// 删除缓存
func (l *Grade) Delete() error {
	return Ctx.Redis.Del(Ctx.Key(Key_GradeInfo)).Err()
}

// 获取用户级别
//
// uid 用户UID
//
//1. 用户级别不存在返回默认级别
//
func (l *Grade) GetUserGrade(uid int) (int, error) {
	if uid > 0 {
		key := Ctx.Keys(Key_UserGrade, uid)
		res := Ctx.Redis.Get(key)
		err := res.Err()
		if err == redis.Nil {
			// 不存在
		} else if err != nil {
			return 0, err
		} else {
			grade, err := strconv.ParseInt(res.Val(), 10, 32)
			if err != nil {
				return 0, err
			}
			return int(grade), nil
		}
	}

	isExists, data, err := l.Get()
	if err != nil {
		return 0, err
	}
	if !isExists {
		return 0, fmt.Errorf("等级缓存信息不存在")
	}
	for i := 0; i < len(data); i++ {
		item := &data[i]
		if item.IsDefault != 0 {
			return int(item.GradeId), nil
		}
	}

	return 0, fmt.Errorf("未配置默认级别")
}

// 设置用户等级
//
//uid 用户uid
//
//gradeId 用户级别
func (l *Grade) SetUserGrade(uid int, gradeId int) error {
	key := Ctx.Keys(Key_UserGrade, uid)
	res := Ctx.Redis.Set(key, fmt.Sprintf("%d", gradeId), 0)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
