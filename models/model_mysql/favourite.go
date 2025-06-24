package model_mysql

import (
	"Common/global"
	"time"
)

type Favourite struct {
	Id            int32     `gorm:"column:id;type:int UNSIGNED;comment:id;primaryKey;not null;" json:"id"`            // id
	UserId        int32     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                    // 用户id
	VehicleId     int32     `gorm:"column:vehicle_id;type:int;comment:车辆id;not null;" json:"vehicle_id"`              // 车辆id
	VehicleName   string    `gorm:"column:vehicle_name;type:text;comment:车辆名称;not null;" json:"vehicle_name"`         // 车辆名称
	Image         string    `gorm:"column:image;type:varchar(255);comment:图片;not null;" json:"image"`                 // 图片
	FavouriteTime time.Time `gorm:"column:favourite_time;type:datetime;comment:收藏时间;not null;" json:"favourite_time"` // 收藏时间
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`      // 创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`      // 修改时间
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`  // 删除时间
}

func (f *Favourite) TableName() string {
	return "favourite"
}

// 查看车辆是否收藏
func (f *Favourite) IsCollectVehicle(vehicleId, userId int64) bool {
	// 只检查未删除的收藏记录
	err := global.DB.Where("vehicle_id = ? AND user_id = ? AND deleted_at IS NULL", vehicleId, userId).First(f).Error
	if err != nil {
		return false
	}
	return true
}

// 收藏车辆
func (f *Favourite) CollectVehicle() error {
	err := global.DB.Create(f).Error
	if err != nil {
		return err
	}
	return nil
}

// 取消收藏车辆
func (f *Favourite) CancelCollectVehicle(vehicleId, userId int64) error {
	// 关键点：使用 Update 而非 Delete
	err := global.DB.Model(&Favourite{}).
		Where("vehicle_id = ? AND user_id = ?", vehicleId, userId).
		Update("deleted_at", time.Now()).Error // 手动更新标记
	if err != nil {
		return err
	}
	return nil
}

// 查看用户收藏的车辆
func (f *Favourite) GetUserCollectVehicle(userId int64) ([]*Favourite, error) {
	var favourite []*Favourite
	// 只查询未删除的收藏记录
	err := global.DB.Where("user_id = ? AND deleted_at IS NULL", userId).Find(&favourite).Error
	if err != nil {
		return nil, err
	}
	return favourite, nil
}
