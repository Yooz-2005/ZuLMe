package model_mysql

import (
	"Common/global"
	"time"
)

type UserProfile struct {
	Id             int32      `gorm:"column:id;type:int UNSIGNED;comment:id;primaryKey;not null;" json:"id"`                                       // id
	UserId         int32      `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                                               // 用户id
	RealName       string     `gorm:"column:real_name;type:varchar(10);comment:真实姓名;not null;" json:"real_name"`                                   // 真实姓名
	IdType         string     `gorm:"column:id_type;type:varchar(255);comment:证件类型1:身份证2:台湾居民来往大陆通行证3:港澳居民来往大陆通行证4:外籍护照;not null;" json:"id_type"` // 证件类型1:身份证2:台湾居民来往大陆通行证3:港澳居民来往大陆通行证4:外籍护照
	IdNumber       string     `gorm:"column:id_number;type:varchar(30);comment:证件号码;not null;" json:"id_number"`                                   // 证件号码
	IdExpireDate   *time.Time `gorm:"column:id_expire_date;type:datetime;comment:有效期;default:NULL;" json:"id_expire_date"`                         // 有效期
	Phone          string     `gorm:"column:phone;type:varchar(255);comment:手机号;not null;" json:"phone"`                                           // 手机号
	Email          string     `gorm:"column:email;type:varchar(255);comment:邮箱;not null;" json:"email"`                                            // 邮箱
	Province       string     `gorm:"column:province;type:varchar(255);comment:通讯地址-省;not null;" json:"province"`                                  // 通讯地址-省
	City           string     `gorm:"column:city;type:varchar(255);comment:通讯地址-市;not null;" json:"city"`                                          // 通讯地址-市
	District       string     `gorm:"column:district;type:varchar(255);comment:通讯地址-区;not null;" json:"district"`                                  // 通讯地址-区
	EmergencyName  string     `gorm:"column:emergency_name;type:varchar(255);comment:紧急联系人姓名;not null;" json:"emergency_name"`                     // 紧急联系人姓名
	EmergencyPhone string     `gorm:"column:emergency_phone;type:varchar(255);comment:紧急联系人电话;not null;" json:"emergency_phone"`                   // 紧急联系人电话
	CreatedAt      time.Time  `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`                                 // 创建时间
	UpdatedAt      time.Time  `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                                 // 修改时间
	DeletedAt      time.Time  `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                             // 删除时间
}

func (up *UserProfile) TableName() string {
	return "user_profile"
}

// todo添加用户默认个人信息
func (up *UserProfile) CreateUserProfile() error {
	return global.DB.Create(&up).Error
}

// todo 修改个人信息 使用 map 更新用户资料，允许更新零值字段
func (up *UserProfile) UpdateUserProfileByMap(userId int64, updates map[string]interface{}) error {
	err := global.DB.Model(&UserProfile{}).Where("user_id = ?", userId).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

// todo 修改用户手机号
func (up *UserProfile) UpdateUserPhoneByUserId(userId int64, phone string) error {
	err := global.DB.Model(&UserProfile{}).Where("user_id =?", userId).Update("phone", phone).Error
	if err != nil {
		return err
	}
	return nil
}
