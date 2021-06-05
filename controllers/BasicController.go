package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/models/api"
	"net/http"
)

func BasicController(rg *gin.RouterGroup) {
	rg.GET("/ping", ping)
}

// ping
// @Summary ping
// @Description 连接测试
// @Router /ping [get]
// @Success 200 object models.ApiResponse
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, api.Response{
		Status:  0,
		Message: "pong",
	})
}
