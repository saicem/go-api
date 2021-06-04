package dbManager

import (
	"fmt"
	"github.com/saicem/api/models"
	"gorm.io/gorm"
	"time"
)

// InsertUserLog 插入单条用户日志
func InsertUserLog(userLog *models.UserLog, db *gorm.DB) {
	fmt.Println("插入用户日志")
	db.Create(&userLog)
}

// todo 插入多条用户日志 并最后添加"添加日志的"日志

// GetUserLog 获取用户日志
func GetUserLog(eventName string, fromTime time.Time, endTime time.Time, db *gorm.DB) *models.UserLog {
	var userLogs []models.UserLog
	db.Where("EventName = ? AND CreatedAt >= ? AND CreatedAt <= ?", eventName, fromTime, endTime).Find(&userLogs)
	for _, log := range userLogs {
		println(log)
	}
	return nil
}
