package main

import (
	db "github.com/saicem/api/dbManager"
	_ "github.com/saicem/api/docs"
	router "github.com/saicem/api/routers"
	"github.com/saicem/api/settings"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

// @title iwut api go
// @version 0.0.2
// @description iwut 用户日志记录
// @BathPath /
// @Host localhost:8080
func main() {
	db.InitDB()
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":" + settings.PORT); err != nil {
		log.Println(err)
	} else {
		log.Printf("Listening on port %s", settings.PORT)
	}
}

// todo swagger UI 从注释中读取 但有些东西不能写死在注释 得写在配置文件里
// todo 从JSON文件读取配置 而不是config.go
