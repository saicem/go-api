package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/config"
	"github.com/saicem/api/middleware"
	"github.com/saicem/api/router/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(engine *gin.Engine) {
	engine.Use(middleware.LoggerToFile())
	// 初始化 swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", config.ProjectPort)
	// todo 统一路径风格
	v1.BasicController(engine.Group(""))

	v1.UserLogController(engine.Group("/user/log", Authentication))

	v1.LoginController(engine.Group("/login"))

}
