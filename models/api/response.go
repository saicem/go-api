package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/models/api/code"
)

type Response struct {
	Status  code.ApiCode `json:"status"`
	Message string       `json:"message"`
	Data    string       `json:"data"`
}

func (r Response) ToJson() gin.H {
	return gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
	}
}
