package mysql_server

import (
	"fmt"
	"github.com/saicem/api/configs"
	"github.com/saicem/api/models/iwut"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbLog *gorm.DB
var dbAdmin *gorm.DB

func InitMySQL() {
	fmt.Println("InitMySQL...")
	NewConn()
	migrate()
}

// migrate 数据库迁移
func migrate() {
	if err := dbLog.AutoMigrate(&iwut.UserLog{}); err != nil {
		panic(err)
	}
	if err := dbAdmin.AutoMigrate(&iwut.AdminUser{}); err != nil {
		panic(err)
	}
}

func NewConn() {
	config := configs.Get()
	var err error
	dbLog, err = gorm.Open(mysql.Open(config.MySQL.Log.Dsn), &gorm.Config{})
	dbAdmin, err = gorm.Open(mysql.Open(config.MySQL.Log.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect iwut-server")
	}
	fmt.Println("new conn...")
}
