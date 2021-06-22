package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/controllers"
	"github.com/saicem/api/controllers/logController"
	"github.com/saicem/api/controllers/loginController"
	"github.com/saicem/api/middleware"
	redis "github.com/saicem/api/widgets/redis-server"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", configs.ProjectPort)

	controllers.BasicController(r.Group(""))

	logController.UserLogController(r.Group("/user/log", Authentication))

	loginController.LoginController(r.Group("/login"))

	return r
}

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
