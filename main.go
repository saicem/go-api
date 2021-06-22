package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/configs"
	_ "github.com/saicem/api/docs"
	"github.com/saicem/api/router"
	mysql "github.com/saicem/api/widgets/mysql_server"
	redis "github.com/saicem/api/widgets/redis_server"
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
