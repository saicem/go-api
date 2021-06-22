package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/configs"
	_ "github.com/saicem/api/docs"
	"github.com/saicem/api/initialize"
	"github.com/saicem/api/router"
	"log"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @Host localhost:9101
// @BathPath /
func main() {
	initialize.InitMySQL()
	initialize.Redis()
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.InitRouter(engine)
	err := engine.Run(":" + configs.ProjectPort)
	if err != nil {
		log.Println(err)
	}
}

// todo redis 多账户？？
// todo 统一 api 参数
