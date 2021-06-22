package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/api/global"
	"github.com/saicem/api/initialize"
	"github.com/saicem/api/models"
	"github.com/saicem/api/models/response"
	"math/rand"
	"net/http"
	"time"
)

func LoginController(rg *gin.RouterGroup) {
	rg.GET("/", adminLogin)
}

// adminLogin 管理员登录
// @Summary 管理员登录
// @Description
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Router /login/ [get]
// @Success 200 object api.Response
func adminLogin(c *gin.Context) {
	// 获取参数
	userName := c.Query("username")
	password := c.Query("password")
	if isValid := SearchAdminUser(userName, password); isValid {
		// todo 其他 sessionId 策略
		// todo 参数待确定 & 加入 config
		sessionId := RandString(50)
		maxAge := 60 * 10
		domain := "localhost"
		c.SetCookie("SESSIONID", sessionId, maxAge, "/", domain, false, true)
		r := initialize.GetRedis()
		if _, err := r.Do("SET", sessionId, 1, "EX", maxAge); err != nil {
			panic("发送不了？？")
		}
		defer func(r redis.Conn) {
			err := r.Close()
			if err != nil {
				panic("关不掉？？")
			}
		}(r)
		c.JSON(http.StatusOK, response.Response{Status: response.OK, Message: "登录成功"})
	} else {
		c.JSON(http.StatusOK, response.Response{Status: response.ERROR, Message: "未通过验证"})
	}
}

func SearchAdminUser(userName string, password string) bool {
	var admin models.AdminUser
	if res := global.Mysql.Where("user_name = ? AND password = ?", userName, password).First(&admin); res.Error != nil {
		return false
	}
	return true
}

func RandString(length int) string {
	str := "0123456789QWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
