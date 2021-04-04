package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var (
	redisPool *redis.Pool
	ErrNil    = errors.New("nil returned")
)

func InitPool() {
	address := fmt.Sprintf("%s:%d", viper.GetString("cache.host"), viper.GetInt("cache.port"))
	redisPool = &redis.Pool{
		MaxIdle:     viper.GetInt("cache.maxIdle"),                             // 最大空闲数
		MaxActive:   viper.GetInt("cache.maxActive"),                           // 最大连接数
		IdleTimeout: viper.GetDuration("cache.idleTimeout") * time.Millisecond, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			conn, err := redis.Dial("tcp", address)
			if err == nil {
				_, err = conn.Do("select", 1)
			}
			return conn, err
		},
	}
	if _, err := redisPool.Get().Do("ping"); err != nil {
		panic(err)
	}
}

func RedisDel(args ...interface{}) error {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()

	_, err := conn.Do("del", args...)
	return err
}

func RedisGetList(key string) ([]string, error) {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()

	return redis.Strings(conn.Do("lrange", key, 0, -1))
}

func RedisLpush(args ...interface{}) error {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()

	_, err := conn.Do("lpush", args...)
	return err

}

func RedisSet(args ...interface{}) error {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()
	_, err := redis.String(conn.Do("Set", args...))

	if err == redis.ErrNil {
		return ErrNil
	}
	return err
}

func RedisGet(args ...interface{}) (string, error) {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()
	res, err := redis.String(conn.Do("Get", args...))

	if err == redis.ErrNil {
		return "", ErrNil
	}
	return res, err
}

func RedisHGetStr(args ...interface{}) (string, error) {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()
	res, err := redis.String(conn.Do("HGet", args...))

	if err == redis.ErrNil {
		return "", ErrNil
	}
	return res, err
}

func RedisHSet(args ...interface{}) error {
	// 从连接池取出连接
	conn := redisPool.Get()
	// 延时关闭连接
	defer conn.Close()
	_, err := conn.Do("HSet", args...)
	return err
}
