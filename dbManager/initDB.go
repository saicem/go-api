package dbManager

import (
	"fmt"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *gorm.DB
//var userLog *models.UserLog

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
	config := configs.Get()
	db, err := gorm.Open(mysql.Open(config.MySQL.Log.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect dbManager")
	}
	fmt.Println("new conn...")
	return db
}
