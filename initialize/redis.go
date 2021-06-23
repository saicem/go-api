package initialize

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/api/global"
)

func Redis() {
	// todo redis 验证是否成功连接 redis
	global.Redis = &redis.Pool{
		MaxIdle:   3, /*最大的空闲连接数*/
		MaxActive: 8, /*最大的激活连接数*/
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", "127.0.0.1:8888", redis.DialPassword("密码"))
			c, err := redis.Dial("tcp", global.Config.Redis.Addr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
