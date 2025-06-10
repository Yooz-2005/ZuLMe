package model_mysql

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Username  string    `gorm:"column:username;type:varchar(20);comment:''用户名'';default:NULL;" json:"username"` // ''用户名''
	Phone     string    `gorm:"column:phone;type:varchar(20);comment:''手机号'';default:NULL;" json:"phone"`       // ''手机号''
}

func (User) TableName() string {
	return "users"
}
