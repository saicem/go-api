package mysql_server

import (
	"fmt"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//var userLog *models.UserLog

func InitMySQL() {
	fmt.Println("InitMySQL...")
	NewConn()
	// 迁移 schema
	err := db.AutoMigrate(&models.UserLog{})
	if err != nil {
		panic(err)
	}
}

func NewConn() {
	config := configs.Get()
	var err error
	db, err = gorm.Open(mysql.Open(config.MySQL.Log.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql-server")
	}
	fmt.Println("new conn...")
}
