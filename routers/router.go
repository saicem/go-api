package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/controllers"
	"github.com/saicem/api/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.LoggerToFile())

	controllers.BasicController(r.Group(""))

	controllers.UserLogController(r.Group("/user/log"))

	return r
}
