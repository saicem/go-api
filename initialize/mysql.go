package initialize

import (
	"github.com/saicem/api/global"
	"github.com/saicem/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitMySQL() {
	log.Println("InitMySQL...")
	NewConn()
	migrate()
}

// migrate 数据库迁移
func migrate() {
	err := global.Mysql.AutoMigrate(
		&models.UserLog{},
		&models.AdminUser{},
	)
	if err != nil {
		// todo 更好的日志记录方式
		log.Println("迁移数据库失败")
		os.Exit(0)
	}
}

func NewConn() {
	config := global.Config
	if db, err := gorm.Open(mysql.Open(config.MySQL.Dsn()), &gorm.Config{}); err != nil {
		panic("failed to connect iwut-server")
	} else {
		global.Mysql = db
		log.Println("new conn...")
	}
}
