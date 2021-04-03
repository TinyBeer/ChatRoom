package dao

import (
	"ChatRoom/Go/server/cache"
	"log"
	"strconv"
)

type RedisSmsDao struct {
}

func (rsd *RedisSmsDao) DepositeByID(id int, content string) (err error) {
	// 将数据存入mesList[userId]中
	err = cache.RedisLpush("smsList"+strconv.Itoa(id), content)
	if err != nil {
		return err
	}
	// 退出
	return
}

func (rsd *RedisSmsDao) WithdrawByID(id int) (dataSlice []string, err error) {
	// 将数据存入mesList[userId]中
	dataSlice, err = cache.RedisGetList("smsList" + strconv.Itoa(id))
	log.Println(dataSlice)
	if err != nil {
		return
	}

	// 如果留言数量不为零
	if len(dataSlice) != 0 {
		err = cache.RedisDel("smsList" + strconv.Itoa(id))
		if err != nil {
			log.Println(err.Error())
		}
	}

	// 退出
	return
}
