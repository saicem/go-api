package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/models"
	"net/http"
)

func BasicController(rg *gin.RouterGroup) {
	rg.GET("/ping", Ping)
}

// Ping
// @Summary ping
// @Description 连接测试
// @Router /ping [get]
// @Success 200 object models.ApiResponse
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.ApiResponse{
		Status:  0,
		Message: "pong",
		Data:    "this is data",
	})
}
