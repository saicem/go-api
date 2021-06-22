package global

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var Redis *redis.Pool
var Mysql *gorm.DB
