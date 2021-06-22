package mysql_server

import (
	"github.com/saicem/api/models"
	"time"
)

// InsertUserLog 插入单条用户日志
func InsertUserLog(userLog *models.UserLog) {
	db.Create(&userLog)
}

// InsertUserLogs 插入多条用户日志
func InsertUserLogs(userLogs *[]models.UserLog) {
	db.Create(&userLogs)
}

// GetUserLog 获取用户日志
func GetUserLog(objectName string, eventName string, startTime time.Time, endTime time.Time) []models.UserLog {
	var userLogs []models.UserLog
	db.Where("object_name = ? AND event_name = ? AND act_time >= ? AND act_time <= ?", objectName, eventName, startTime, endTime).Find(&userLogs)
	return userLogs
}

func GetUserLogByDay(objectName string, eventName string, startTime time.Time, endTime time.Time) []models.UserLogByDay {
	var result []models.UserLogByDay
	rawSql := "SELECT DISTINCT uid, cast(act_time AS date) as act_day FROM user_logs " +
		"WHERE object_name = ? AND event_name = ? AND act_time >= ? AND act_time <= ?"
	db.Raw(rawSql, objectName, eventName, startTime, endTime).Scan(&result)
	return result
}
