package main

import (
	"github.com/saicem/api/configs"
	db "github.com/saicem/api/dbManager"
	_ "github.com/saicem/api/docs"
	router "github.com/saicem/api/routers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @Host localhost:9101
// @BathPath /
func main() {

	db.InitDB()
	initRouter()
}

func initRouter() {
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + configs.ProjectPort)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Listening on port localhost:%s\n", configs.ProjectPort)
	}
}

// todo swagger UI 从注释中读取 但有些东西不能写死在注释 得写在配置文件里
