package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/saicem/api/dbManager"
	"github.com/saicem/api/models"
	"github.com/saicem/api/models/api"
	"github.com/saicem/api/models/api/code"
	"net/http"
	"time"
)

func UserLogController(rg *gin.RouterGroup) {
	rg.GET("/retrieve", retrieveUserLog)
	rg.POST("/upload", uploadUserLogs)
}

// uploadUserLogs 添加用户日志
// @Summary 添加用户日志
// @Description
// @Param json body models.PushJson true "json"
// @Router /user/log/upload [post]
// @Success 200 object api.Response
func uploadUserLogs(c *gin.Context) {
	var pushJson models.PushJson
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, api.Response{Status: code.ERROR, Message: "无法解析JSON"})
		return
	}
	err2 := json.Unmarshal(rawData, &pushJson)
	if err2 != nil {
		c.JSON(http.StatusOK, api.Response{Status: code.ERROR, Message: "无法解析JSON"})
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
		api.Response{Status: code.OK, Message: "添加成功", Data: string(rune(logCount))},
	)
}

// retrieveUserLog 获取用户日志
// @Summary 获取用户日志
// @Description
// @Param name query string true "事件名称"
// @Param start query string true "查询开始时间 YYYY-MM-DD"
// @Param end query string true "查询结束时间 YYYY-MM-DD"
// @Router /user/log/retrieve [get]
// @Success 200 object api.Response
func retrieveUserLog(c *gin.Context) {
	// 获取参数
	eventName := c.Query("name")
	startTimeStr := c.Query("start")
	endTimeStr := c.Query("end")
	// 获取北京时区
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	// 解析时间
	startTime, err1 := time.ParseInLocation("2006-01-02", startTimeStr, beijing)
	endTime, err2 := time.ParseInLocation("2006-01-02", endTimeStr, beijing)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, api.Response{Status: code.ERROR, Message: "wrong format time, need YYYY-MM-DD"})
	}
	// 获取数据
	// todo 对每个功能单独写查询
	userLogs := db.GetUserLog(eventName, startTime, endTime)
	c.JSON(http.StatusOK, api.Response{Status: code.OK, Message: "ok", Data: fmt.Sprint(userLogs)})
}
