package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/api/api/v1"
	"github.com/saicem/api/middleware"
)

func InitLog(rg *gin.RouterGroup) {
	LogRouter := rg.Group("user/log", middleware.Authentication)
	{
		LogRouter.GET("/retrieve", v1.RetrieveUserLog)
		LogRouter.POST("/upload", v1.UploadUserLogs)
		LogRouter.GET("/query/active", v1.QueryActiveUser)
	}
}
