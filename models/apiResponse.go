package models

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

//type ApiCode int
//
//const (
//	Ok          ApiCode = 0
//	Error       ApiCode = -1
//	WrongFormat ApiCode = -2
//)

func (r ApiResponse) ToJson() gin.H {
	return gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
	}
}
