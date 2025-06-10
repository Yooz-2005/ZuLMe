package model_mysql

import "gorm.io/gorm"

// Merchant 商家模型
type Merchant struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Phone    string `gorm:"type:varchar(20);unique;not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"` // 存储哈希后的密码
	Status   int    `gorm:"type:int;default:0"`         // 状态：0-未审核，1-审核通过，2-审核失败
}
