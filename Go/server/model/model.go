package model

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var RPool *redis.Pool

func InitDao(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	InitPool(address, maxIdle, maxActive, idleTimeout)
	InitUserDao()
}

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	RPool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲数
		MaxActive:   maxActive,   // 最大连接数
		IdleTimeout: idleTimeout, // 最大空闲事件
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			return redis.Dial("tcp", address)
		},
	}
}
