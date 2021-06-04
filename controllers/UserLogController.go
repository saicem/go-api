package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/saicem/api/dbManager"
	"github.com/saicem/api/models"
	"net/http"
	"time"
)

func UserLogController(rg *gin.RouterGroup) {
	rg.GET("/retrieve", RetrieveUserLog)
	rg.POST("/upload", CreateUserLog)
}

// CreateUserLog 添加用户日志
// @Summary 添加用户日志
// @Description
// @Param json body models.PushJson true "json"
// @Router /user/log/upload [post]
// @Success 200 object models.ApiResponse
func CreateUserLog(c *gin.Context) {
	var pushJson models.PushJson
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, models.ApiResponse{Status: -1, Message: "无法解析JSON"})
	}
	err2 := json.Unmarshal(rawData, &pushJson)
	if err2 != nil {
		c.JSON(http.StatusOK, models.ApiResponse{Status: -1, Message: "无法解析JSON"})
	}

	conn := db.NewConn()

	uid := pushJson.Uid
	//deviceId := pushJson.DeviceId
	for _, userEvent := range pushJson.UserEvents {
		eventName := userEvent.EventName
		for _, eventInfo := range userEvent.EventInfos {
			db.InsertUserLog(&models.UserLog{
				Uid:        uid,
				EventName:  eventName,
				ActTime:    eventInfo.ActTime,
				AppendInfo: eventInfo.Append,
			}, conn)
		}
	}

	logCount := 0

	c.JSON(
		http.StatusOK,
		models.ApiResponse{Message: "添加成功", Data: string(logCount)},
	)
}

// RetrieveUserLog 获取用户日志
// @Summary 获取用户日志
// @Description
// @Param name query string true "事件名称"
// @Param start query string true "查询开始时间"
// @Param end query string true "查询结束时间"
// @Router /user/log/retrieve [get]
// @Success 200 object models.ApiResponse
func RetrieveUserLog(c *gin.Context) {
	eventName := c.Query("name")
	startTime, err1 := time.Parse("2006-01-02", c.Query("start"))
	endTime, err2 := time.Parse("2006-01-02", c.Query("end"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, models.ApiResponse{Status: -1, Message: "wrong format time, need YYYY-MM-DD"})
	}
	conn := db.NewConn()
	db.GetUserLog(eventName, startTime, endTime, conn)
	c.JSON(http.StatusOK, models.ApiResponse{Status: 0, Message: "ok", Data: fmt.Sprint(eventName, startTime, endTime)})
}
