package model_mysql

import (
	"Common/global"
	"gorm.io/gorm"
)

// Merchant 商家模型
type Merchant struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100);not null"`
	Phone        string  `gorm:"type:varchar(20);unique;not null"`
	Email        string  `gorm:"type:varchar(100);unique;not null"`
	Password     string  `gorm:"type:varchar(255);not null"` // 存储哈希后的密码
	Status       int     `gorm:"type:int;default:0"`         // 状态：0-未审核，1-审核通过，2-审核失败
	Location     string  `gorm:"type:varchar(255)"`          // 网点地址
	BusinessTime string  `gorm:"type:varchar(255)"`          // 营业时间
	Longitude    float64 `gorm:"type:decimal(10,7)"`         // 经度
	Latitude     float64 `gorm:"type:decimal(10,7)"`         // 纬度
}

// GetByID 根据ID获取商家信息
func (m *Merchant) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).First(m).Error
}
