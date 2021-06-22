package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/middleware"
	v12 "github.com/saicem/api/router/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(engine *gin.Engine) {
	engine.Use(middleware.LoggerToFile())
	// 初始化 swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", configs.ProjectPort)
	// todo 统一路径风格
	v12.BasicController(engine.Group(""))

	v12.UserLogController(engine.Group("/user/log", Authentication))

	v12.LoginController(engine.Group("/login"))

}
