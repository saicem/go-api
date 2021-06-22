package models

type AdminUser struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
