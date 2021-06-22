package mysql_server

import (
	"fmt"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL() {
	fmt.Println("InitMySQL...")
	NewConn()
	migrate()
}

// migrate 数据库迁移
func migrate() {
	if err := db.AutoMigrate(
		&models.UserLog{},
		&models.AdminUser{},
	); err != nil {
		panic(err)
	}
}

func NewConn() {
	config := configs.Get()
	var err error
	db, err = gorm.Open(mysql.Open(config.MySQL.Log.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect iwut-server")
	}
	fmt.Println("new conn...")
}
