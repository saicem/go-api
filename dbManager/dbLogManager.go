package dbManager

import (
	"fmt"
	"github.com/saicem/api/models"
	"time"
)

// InsertUserLog 插入单条用户日志
func InsertUserLog(userLog *models.UserLog) {
	db := NewConn()
	db.Create(&userLog)
	fmt.Println("插入单条用户日志")
}

// InsertUserLogs 插入多条用户日志
func InsertUserLogs(userLogs *[]models.UserLog) {
	db := NewConn()
	db.Create(&userLogs)
	fmt.Println("插入多条用户日志")
}

// GetUserLog 获取用户日志
// todo 还没写
func GetUserLog(eventName string, fromTime time.Time, endTime time.Time) *models.UserLog {
	db := NewConn()
	var userLogs []models.UserLog
	db.Where("EventName = ? AND CreatedAt >= ? AND CreatedAt <= ?", eventName, fromTime, endTime).Find(&userLogs)
	for _, userLog := range userLogs {
		println(fmt.Sprint(userLog))
	}
	return nil
}
