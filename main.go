package main

import (
	"github.com/saicem/api/configs"
	_ "github.com/saicem/api/docs"
	router "github.com/saicem/api/routers"
	mysql "github.com/saicem/api/widgets/mysql-server"
	redis "github.com/saicem/api/widgets/redis-server"
	"log"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @Host localhost:9101
// @BathPath /
func main() {
	mysql.InitMySQL()
	redis.InitRedis()
	initRouter()
}

func initRouter() {
	r := router.SetupRouter()
	err := r.Run(":" + configs.ProjectPort)
	if err != nil {
		log.Println(err)
	}
}

// todo redis 多账户？？
// todo 统一 api 参数
