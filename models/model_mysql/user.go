package model_mysql

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;not null;comment:'用户名'"`
	Password string `gorm:"type:varchar(20);not null;comment:'密码'"`
	Email    string `gorm:"type:varchar(50);not null;comment:'邮箱'"`
	Phone    string `gorm:"type:varchar(20);not null;comment:'手机号'"`
}
