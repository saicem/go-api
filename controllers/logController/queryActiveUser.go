package logController

import (
	"github.com/gin-gonic/gin"
	"github.com/saicem/api/dbManager"
	"github.com/saicem/api/models"
	"github.com/saicem/api/models/api"
	"github.com/saicem/api/models/api/code"
	"net/http"
	"strconv"
	"time"
)

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
		c.JSON(http.StatusOK, api.Response{Status: code.ERROR, Message: "参数读取失败"})
	} else if totalDaySpan > 60 || // 参数有效性判断
		totalDaySpan < activeDaySpan ||
		startTime.Unix() > endTime.Unix() ||
		endTime.Unix() > time.Now().Unix() {
		c.JSON(http.StatusOK, api.Response{Status: code.ERROR, Message: "参数无效"})
	} else {
		// 获取事件
		queryResult := dbManager.GetUserLogByDay(objectName, eventName, startTime.AddDate(0, 0, -totalDaySpan), endTime)
		res := compute(queryResult, totalDaySpan, activeDaySpan)
		c.JSON(http.StatusOK, api.Response{Status: code.OK, Message: "成功获取", Data: res})
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