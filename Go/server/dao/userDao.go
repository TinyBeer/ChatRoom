package dao

import (
	"ChatRoom/Go/common/userinfo"
	"ChatRoom/Go/server/cache"
	"encoding/json"
	"fmt"
)

type RedisUserDao struct {
	IUserDao
}

func (rud *RedisUserDao) Update() {
	panic("not implemented") // TODO: Implement
}

func (rud *RedisUserDao) Insert(id int, pwd string, name string) error {
	_, err := rud.GetUserByID(id)
	if err != ERROR_USER_NOTEXIST {
		return err
	}

	user := userinfo.User{
		UserID:   id,
		UserPwd:  string(pwd),
		UserName: name,
	}

	// 该用户ID可用
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// 入库
	err = cache.RedisHSet("users", user.UserID, string(data))
	if err != nil {
		fmt.Println("用户信息入库失败")
		return err
	}
	return nil
}

func (rud *RedisUserDao) GetUserByID(id int) (user *userinfo.User, err error) {
	// 通过给定的id 去redis查询用户
	res, err := cache.RedisHGetStr("users", id)
	fmt.Println(res, err)
	if err != nil {
		// 发生错误
		if err == cache.ErrNil {
			// 没有找到对应id
			err = ERROR_USER_NOTEXIST
		}
		return nil, err
	}
	// fmt.Println(res)
	user = &userinfo.User{}
	// 无错误  将res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("josn.Unmarshal failed, err=", err.Error())
		return
	}
	return
}

func (rud *RedisUserDao) Delete(id int, pwd string) error {
	return cache.RedisDel("users", id)
}

func (rud *RedisUserDao) IsExist(id int) bool {
	if _, err := rud.GetUserByID(id); err != nil {
		return false
	}
	return true
}
