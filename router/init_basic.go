package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/api/api/v1"
)

func InitBasic(rg *gin.RouterGroup) {
	BasicRouter := rg.Group("")
	{
		BasicRouter.GET("/ping", v1.Ping)
	}
}
