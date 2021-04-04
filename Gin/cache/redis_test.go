package cache

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func TestRedis(t *testing.T) {
	err := RedisHSet("users", "100", "hello")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestMain(m *testing.M) {
	redisPool = &redis.Pool{
		MaxIdle:     4,                      // 最大空闲数
		MaxActive:   8,                      // 最大连接数
		IdleTimeout: 300 * time.Millisecond, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err == nil {
				_, err = conn.Do("select", 1)
			}
			return conn, err
		},
	}
	if _, err := redisPool.Get().Do("ping"); err != nil {
		panic(err)
	}

	m.Run()

	redisPool.Close()
}
