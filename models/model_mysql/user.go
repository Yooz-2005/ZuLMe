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
	Phone     string    `gorm:"column:phone;type:varchar(20);comment:''手机号'';default:NULL;" json:"phone"`       // ''手机号''
	Location  string    `gorm:"column:location;type:varchar(255);comment:''地址'';default:NULL;" json:"location"` // ''地址''
	LatAndLon string    `gorm:"column:lat_and_lon;type:varchar(255);comment:''经纬度'';default:NULL;" json:"lat_and_lon"`
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

// todo修改手机号
func (u *User) UpdateUserByPhone(id int64, phone string) error {
	return global.DB.Model(&User{}).Where("id =?", id).Update("phone", phone).Error
}

// CheckPhoneExistExcludingUser 检查手机号码是否已经被注册，排除指定用户ID
func (u *User) CheckPhoneExistExcludingUser(phone string) (bool, error) {
	var count int64
	err := global.DB.Model(&User{}).Where("phone = ?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
