package dbManager

import (
	"github.com/saicem/api/models"
	"time"
)

// InsertUserLog 插入单条用户日志
func InsertUserLog(userLog *models.UserLog) {
	db := NewConn()
	db.Create(&userLog)
}

// InsertUserLogs 插入多条用户日志
func InsertUserLogs(userLogs *[]models.UserLog) {
	db := NewConn()
	db.Create(&userLogs)
}

// GetUserLog 获取用户日志
func GetUserLog(eventName string, fromTime time.Time, endTime time.Time) []models.UserLog {
	db := NewConn()
	var userLogs []models.UserLog
	db.Where("event_name = ? AND act_time >= ? AND act_time <= ?", eventName, fromTime, endTime).Find(&userLogs)
	return userLogs
}
