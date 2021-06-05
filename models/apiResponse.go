package models

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/models/ApiCode"
)

type ApiResponse struct {
	Status  ApiCode.ApiCode `json:"status"`
	Message string          `json:"message"`
	Data    string          `json:"data"`
}

func (r ApiResponse) ToJson() gin.H {
	return gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
	}
}
