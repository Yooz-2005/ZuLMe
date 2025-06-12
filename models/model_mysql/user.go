package model_mysql

import (
	"Common/global"
	"time"
)

type User struct {
	Id        int64     `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Phone     string    `gorm:"column:phone;type:varchar(20);comment:''''手机号'''';default:NULL;" json:"phone"` // ''''手机号''''
}

func (u *User) TableName() string {
	return "users"
}

// todo注册
func (u *User) Register() error {
	return global.DB.Create(&u).Error
}

// todo登录
func (u *User) Login(phone string) error {
	return global.DB.Where("phone = ?", phone).First(&u).Error
}
