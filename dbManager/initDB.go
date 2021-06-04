package dbManager

import (
	"fmt"
	"github.com/saicem/api/models"
	"github.com/saicem/api/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *gorm.DB
var userLog *models.UserLog

func InitDB() {
	fmt.Println("InitDB...")
	db := NewConn()
	// 迁移 schema
	err := db.AutoMigrate(&models.UserLog{})
	if err != nil {
		panic(err)
	}
}

func NewConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(settings.DSN), &gorm.Config{})
	if err != nil {
		// todo 什么是panic
		panic("failed to connect dbManager")
	}
	fmt.Println("new conn...")
	return db
}
