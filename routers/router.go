package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/controllers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	controllers.BasicController(r.Group(""))

	controllers.UserLogController(r.Group("/user/log"))

	return r
}
