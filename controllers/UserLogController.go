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
	rg.POST("/upload", UploadUserLogs)
}

// UploadUserLogs 添加用户日志
// @Summary 添加用户日志
// @Description
// @Param json body models.PushJson true "json"
// @Router /user/log/upload [post]
// @Success 200 object models.ApiResponse
func UploadUserLogs(c *gin.Context) {
	var pushJson models.PushJson
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, models.ApiResponse{Status: -1, Message: "无法解析JSON"})
		return
	}
	err2 := json.Unmarshal(rawData, &pushJson)
	if err2 != nil {
		c.JSON(http.StatusOK, models.ApiResponse{Status: -1, Message: "无法解析JSON"})
		return
	}

	var userLogs []models.UserLog
	logCount := 0
	uid := pushJson.Uid
	//deviceId := pushJson.DeviceId
	for _, userEvent := range pushJson.UserEvents {
		eventName := userEvent.EventName
		for _, eventInfo := range userEvent.EventInfos {
			userLogs = append(userLogs, models.UserLog{
				Uid:        uid,
				EventName:  eventName,
				ActTime:    eventInfo.ActTime,
				AppendInfo: eventInfo.Append,
			})
			logCount++
		}
	}
	db.InsertUserLogs(&userLogs)
	db.InsertUserLog(&models.UserLog{
		Uid:       uid,
		EventName: "upload",
		ActTime:   time.Now(),
	})

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
	db.GetUserLog(eventName, startTime, endTime)
	c.JSON(http.StatusOK, models.ApiResponse{Status: 0, Message: "ok", Data: fmt.Sprint(eventName, startTime, endTime)})
}
