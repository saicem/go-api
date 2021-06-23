package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/config"
	"github.com/saicem/api/middleware"
	"github.com/saicem/api/router"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	engine := gin.New()
	// 设置静态文件
	// Router.StaticFS()
	//gin.SetMode(gin.ReleaseMode)

	engine.Use(middleware.LoggerToFile())
	// 初始化 swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("open swagger UI http://localhost:%s/swagger/index.html\n", config.ProjectPort)

	PrivateGroup := engine.Group("")
	router.InitBasic(PrivateGroup)
	router.InitLog(PrivateGroup)
	router.InitLogin(PrivateGroup)

	return engine
}
