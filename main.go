package main

import (
	"github.com/saicem/api/config"
	_ "github.com/saicem/api/docs"
	"github.com/saicem/api/initialize"
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
	engine := initialize.Routers()
	err := engine.Run(":" + config.ProjectPort)
	if err != nil {
		log.Println(err)
	}
}

// todo redis 多账户？？
