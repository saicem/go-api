package redis_server

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/api/configs"
)

var pool *redis.Pool

func InitRedis() {
	// todo redis 未能连接的处理
	config := configs.Get()
	pool = &redis.Pool{
		MaxIdle:   3, /*最大的空闲连接数*/
		MaxActive: 8, /*最大的激活连接数*/
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", "127.0.0.1:8888", redis.DialPassword("密码"))
			c, err := redis.Dial("tcp", config.Redis.Addr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func Get() redis.Conn {
	return pool.Get()
}
