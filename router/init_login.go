package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saicem/api/api/v1"
)

func InitLogin(rg *gin.RouterGroup) {
	LoginRouter := rg.Group("login")
	{
		LoginRouter.GET("/", v1.AdminLogin)
	}
}
