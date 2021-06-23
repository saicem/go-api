package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/api/config"
	"gorm.io/gorm"
)

var (
	Redis  *redis.Pool
	Mysql  *gorm.DB
	Config *config.Config
)

func init() {
	Config = config.Get()
}
