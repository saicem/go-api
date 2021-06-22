package mysql_server

import (
	"github.com/saicem/api/models"
)

// SearchAdminUser 查找管理用户中是否存在该账户并验证
func SearchAdminUser(userName string, password string) bool {
	var admin models.AdminUser
	if res := db.Where("user_name = ? AND password = ?", userName, password).First(&admin); res.Error != nil {
		return false
	}
	return true
}
