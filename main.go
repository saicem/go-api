package main

import (
	"github.com/saicem/api/configs"
	db "github.com/saicem/api/dbManager"
	_ "github.com/saicem/api/docs"
	"github.com/saicem/api/middleware"
	router "github.com/saicem/api/routers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
	"path"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @Host localhost:9101
// @BathPath /
func main() {
	db.InitDB()
	initFiles()
	initRouter()
}

func initFiles() {
	logFilePath := configs.LogFilePath
	logFileName := configs.LogFileName
	fileName := path.Join(logFilePath, logFileName)

	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		panic(err)
	}
	if exists, err := middleware.IsPathExists(fileName); err != nil {
		panic(err)
	} else if !exists {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
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
