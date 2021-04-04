package dao

import (
	"ChatRoom/Gin/cache"
)

/*   mysql */
type MySqlSmsDao struct {
	ISmsDao
}

//  redis
type RedisSmsDao struct {
}

func (rsd *RedisSmsDao) Deposite(content string) (err error) {
	// 将数据存入mesList[userId]中
	err = cache.RedisSet("sms", content)
	if err != nil {
		return err
	}
	// 退出
	return
}

func (rsd *RedisSmsDao) Withdraw() (string, error) {
	// 将数据存入mesList[userId]中
	sms, err := cache.RedisGet("sms")
	if err != nil {
		return "", err
	}

	// 退出
	return sms, err
}

// end redis */
