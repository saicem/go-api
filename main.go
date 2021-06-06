package main

import (
	"github.com/saicem/api/configs"
	db "github.com/saicem/api/dbManager"
	_ "github.com/saicem/api/docs"
	router "github.com/saicem/api/routers"
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
	err := r.Run(":" + configs.ProjectPort)
	if err != nil {
		log.Println(err)
	}
}
