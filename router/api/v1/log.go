package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/global"
	"github.com/saicem/api/models"
	"github.com/saicem/api/models/request"
	"github.com/saicem/api/models/response"
	"net/http"
	"strconv"
	"time"
)

func UserLogController(rg *gin.RouterGroup) {
	rg.GET("/retrieve", retrieveUserLog)
	rg.POST("/upload", uploadUserLogs)
	rg.GET("/query/active", queryActiveUser)
}

// uploadUserLogs 添加用户日志
// @Summary 添加用户日志
// @Description
// @Param json body models.PushJson true "json"
// @Router /user/log/upload [post]
// @Success 200 object api.Response
func uploadUserLogs(c *gin.Context) {
	var pushJson request.PushJson
	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "无法解析JSON"})
		return
	}
	err2 := json.Unmarshal(rawData, &pushJson)
	if err2 != nil {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "无法解析JSON"})
		return
	}

	var userLogs []models.UserLog
	logCount := 0
	uid := pushJson.Uid
	deviceId := pushJson.DeviceId
	for _, userObject := range pushJson.UserObjects {
		objectName := userObject.ObjectName
		userEvents := userObject.UserEvents
		for _, userEvent := range userEvents {
			eventName := userEvent.EventName
			eventInfos := userEvent.EventInfos
			for _, eventInfo := range eventInfos {
				userLogs = append(userLogs, models.UserLog{
					Uid:        uid,
					ObjectName: objectName,
					EventName:  eventName,
					ActTime:    eventInfo.ActTime,
					AppendInfo: eventInfo.Append,
				})
				logCount++
			}
		}
	}
	InsertUserLogs(&userLogs)
	InsertUserLog(&models.UserLog{
		Uid:        uid,
		ObjectName: "system",
		EventName:  "upload",
		ActTime:    time.Now(),
		AppendInfo: deviceId,
	})

	c.JSON(
		http.StatusOK,
		response.Response{Status: response.OK, Message: "添加成功", Data: logCount},
	)
}

// retrieveUserLog 获取用户日志
// @Summary 获取用户日志
// @Description
// @Param object_name query string true "对象名称"
// @Param event_name query string true "事件名称"
// @Param start_time query string true "查询开始时间 YYYY-MM-DD"
// @Param end_time query string true "查询结束时间 YYYY-MM-DD"
// @Router /user/log/retrieve [get]
// @Success 200 object api.Response
func retrieveUserLog(c *gin.Context) {
	// 获取参数
	objectName := c.Query("object_name")
	eventName := c.Query("event_name")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")
	// 获取北京时区
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	// 解析时间
	startTime, err1 := time.ParseInLocation("2006-01-02", startTimeStr, beijing)
	endTime, err2 := time.ParseInLocation("2006-01-02", endTimeStr, beijing)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "wrong format time, need YYYY-MM-DD"})
	}
	// 获取数据
	userLogs := GetUserLog(objectName, eventName, startTime, endTime)
	c.JSON(http.StatusOK, response.Response{Status: response.OK, Message: "ok", Data: fmt.Sprint(userLogs)})
}

// queryActiveUser 获取活跃（留存）用户数量
// @Summary 获取活跃（留存）用户数量
// @Description 查询 start_time 到 end_time 范围 total_day_span 内有 active_day_span 天活跃的每日数据
// @Description 计算方式为 计算 前 total_day_span 天内有 active_day_span 天有活跃记录 视为活跃
// @Param object_name query string true "对象名称"
// @Param event_name query string true "事件名称"
// @Param start_time query string true "查询开始时间 YYYY-MM-DD"
// @Param end_time query string true "查询结束时间 YYYY-MM-DD"
// @Param total_day_span query uint true "查询天数范围"
// @Param active_day_span query uint true "活跃天数（需小于查询天数）"
// @Router /user/log/query/active [get]
// @Success 200 object api.Response
func queryActiveUser(c *gin.Context) {
	// 获取参数
	objectName := c.Query("object_name")
	eventName := c.Query("event_name")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")
	totalDaySpanStr := c.Query("total_day_span")
	activeDaySpanStr := c.Query("active_day_span")
	// 获取北京时区
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	// 转换参数
	startTime, err := time.ParseInLocation("2006-01-02", startTimeStr, beijing)
	endTime, err := time.ParseInLocation("2006-01-02", endTimeStr, beijing)
	totalDaySpan, err := strconv.Atoi(totalDaySpanStr)
	activeDaySpan, err := strconv.Atoi(activeDaySpanStr)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "参数读取失败"})
	} else if totalDaySpan > 60 || // 参数有效性判断
		totalDaySpan < activeDaySpan ||
		startTime.Unix() > endTime.Unix() ||
		endTime.Unix() > time.Now().Unix() {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "参数无效"})
	} else {
		// 获取事件
		queryResult := GetUserLogByDay(objectName, eventName, startTime.AddDate(0, 0, -totalDaySpan), endTime)
		res := compute(queryResult, totalDaySpan, activeDaySpan)
		c.JSON(http.StatusOK, response.Response{Status: response.OK, Message: "成功获取", Data: res})
	}
}

// ids 的数量
func compute(ids []models.UserLogByDay, totalDaySpan int, activeDaySpan int) *map[time.Time]int {
	res := map[time.Time]int{}
	countLs := models.CountList{}
	countLs.New()
	i := 0
	pre := time.Time{}
	dayCount := 0
	for index, id := range ids {
		if id.ActDay != pre {
			pre = id.ActDay
			dayCount++
			if dayCount >= totalDaySpan {
				i = index
				break
			}
		}
		countLs.Add(id.UID)
	}

	pre = time.Time{}
	iPre := 0
	for ; i < len(ids); i++ {
		// 不是同一天 则进行一次计算 并删除最前面一天的所有数据
		if ids[i].ActDay != pre {
			pre = ids[i].ActDay
			res[pre] = countLs.Compute(activeDaySpan)
			for true {
				countLs.Del(ids[iPre].UID)
				iPre++
				if ids[iPre-1].ActDay != ids[iPre].ActDay {
					break
				}
			}
		}
		countLs.Add(ids[i].UID)
	}
	res[pre.AddDate(0, 0, 1)] = countLs.Compute(activeDaySpan)
	return &res
}

// InsertUserLog 插入单条用户日志
func InsertUserLog(userLog *models.UserLog) {
	global.Mysql.Create(&userLog)
}

// InsertUserLogs 插入多条用户日志
func InsertUserLogs(userLogs *[]models.UserLog) {
	global.Mysql.Create(&userLogs)
}

// GetUserLog 获取用户日志
func GetUserLog(objectName string, eventName string, startTime time.Time, endTime time.Time) []models.UserLog {
	var userLogs []models.UserLog
	global.Mysql.Where("object_name = ? AND event_name = ? AND act_time >= ? AND act_time <= ?", objectName, eventName, startTime, endTime).Find(&userLogs)
	return userLogs
}

func GetUserLogByDay(objectName string, eventName string, startTime time.Time, endTime time.Time) []models.UserLogByDay {
	var result []models.UserLogByDay
	rawSql := "SELECT DISTINCT uid, cast(act_time AS date) as act_day FROM user_logs " +
		"WHERE object_name = ? AND event_name = ? AND act_time >= ? AND act_time <= ?"
	global.Mysql.Raw(rawSql, objectName, eventName, startTime, endTime).Scan(&result)
	return result
}
