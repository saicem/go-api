package router

import (
	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
	redis "github.com/saicem/api/widgets/redis_server"
	"net/http"
)

func Authentication(c *gin.Context) {
	if sessionId, err := c.Cookie("SESSIONID"); err != nil {
		//c.SetCookie("sessionId", "asd", 10, "/", "localhost", false, true)
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		if isValid := SearchSession(sessionId); !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
	return
}

func SearchSession(sessionId string) bool {
	// todo 不能整个redis全给存这个 需要优化存储策略

	r := redis.Get()
	defer func(r redis2.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	if _, err := r.Do("GET", sessionId); err == nil {
		return true
	}
	return false
}
