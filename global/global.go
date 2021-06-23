package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/api/config"
	"gorm.io/gorm"
)

var Redis *redis.Pool
var Mysql *gorm.DB
var Config *config.Config

func init() {
	Config = config.Get()
}
