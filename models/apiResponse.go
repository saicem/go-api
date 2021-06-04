package models

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (r ApiResponse) ToJson() gin.H {
	return gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
	}
}
