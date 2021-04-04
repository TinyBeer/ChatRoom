package dao

import (
	"ChatRoom/Gin/cache"
	"ChatRoom/Gin/model"
	"encoding/json"
	"fmt"
)

type RedisUserDao struct {
	IUserDao
}

func (rud *RedisUserDao) Update() {
	panic("not implemented") // TODO: Implement
}

func (rud *RedisUserDao) Insert(id string, name string) error {
	_, err := rud.GetUserByID(id)
	if err != ERR_USER_NOTEXIST {
		return err
	}

	user := model.UserInfo{
		ID:   id,
		Name: name,
	}

	// 该用户ID可用
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// 入库
	err = cache.RedisHSet("users", id, string(data))
	if err != nil {
		fmt.Println("用户信息入库失败")
		return err
	}
	return nil
}

func (rud *RedisUserDao) GetUserByID(id string) (*model.UserInfo, error) {
	// 通过给定的id 去redis查询用户
	res, err := cache.RedisHGetStr("users", id)
	if err != nil {
		// 发生错误
		if err == cache.ErrNil {
			// 没有找到对应id
			err = ERR_USER_NOTEXIST
		}
		return nil, err
	}
	user := &model.UserInfo{}
	// 无错误  将res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("josn.Unmarshal failed, err=", err.Error())
		return nil, err
	}
	return user, nil
}

func (rud *RedisUserDao) Delete(id string) error {
	return cache.RedisDel("users", id)
}
