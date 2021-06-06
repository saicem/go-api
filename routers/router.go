package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/controllers"
	"github.com/saicem/api/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {

	r := gin.New()

	r.Use(middleware.LoggerToFile())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.BasicController(r.Group(""))

	controllers.UserLogController(r.Group("/user/log"))

	return r
}
