package dao

import (
	"ChatRoom/Go/common/userinfo"
	"ChatRoom/Go/server/cache"
	"ChatRoom/Go/server/database"
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

func Init(mode int) {

	switch mode {
	case MODE_O_REDIS:
		cache.InitPool()
		MyUserDao = &RedisUserDao{}
		MySmsDao = &RedisSmsDao{}
	case MODE_REDIS_MYSQL:
		cache.InitPool()
		database.Init()
		MyUserDao = &RedisUserDao{}
		MySmsDao = &MySqlSmsDao{}
	default:
		panic("init with unkown mode")
	}

}

type IUserDao interface {
	Insert(int, string, string) error        // 添加用户
	Delete(int, string) error                // 删除用户  需要提供ID和密码
	GetUserByID(int) (*userinfo.User, error) // 根据ID获取用户
	IsExist(int) bool                        // 查看用户是否存在

	Update() // 更新用户信息
}

type ISmsDao interface {
	DepositeByID(int, string) error
	WithdrawByID(int) ([]string, error)
}
