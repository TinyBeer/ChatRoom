package dao

import (
	"ChatRoom/Gin/cache"
	"ChatRoom/Gin/model"
	"errors"
)

const (
	MODE_O_REDIS = iota
	MODE_REDIS_MYSQL
)

var (
	ERR_USER_NOTEXIST = errors.New("用户不存在")
	ERR_USER_EXIST    = errors.New("用户已经存在")
	ERR_USER_PWD      = errors.New("密码不正确")
)

// 服务器启动后初始化一个全局的UserDao
var (
	MyUserDao IUserDao
	MySmsDao  ISmsDao
)

func InitDao() {
	cache.InitPool()
	MyUserDao = &RedisUserDao{}
	MySmsDao = &RedisSmsDao{}
}

type IUserDao interface {
	Insert(string, string) error                 // 添加用户
	Delete(string) error                         // 删除用户  需要提供ID和密码
	GetUserByID(string) (*model.UserInfo, error) // 根据ID获取用户

	Update() // 更新用户信息
}

type ISmsDao interface {
	Deposite(string) error
	Withdraw() (string, error)
}
