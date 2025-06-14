package model_mysql

import (
	"Common/global"
	"errors"

	"fmt"
	"gorm.io/gorm"
)

// Vehicle 车辆表
type Vehicle struct {
	gorm.Model
	MerchantID  int64   `json:"merchant_id" gorm:"not null;comment:商家ID"`
	TypeID      int64   `json:"type_id" gorm:"not null;comment:车辆类型ID"`
	Brand       string  `json:"brand" gorm:"type:varchar(50);not null;comment:品牌"`
	Style       string  `json:"style" gorm:"type:varchar(50);not null;comment:型号"`
	Year        int64   `json:"year" gorm:"type:int;not null;comment:年份"`
	Color       string  `json:"color" gorm:"type:varchar(20);comment:颜色"`
	Mileage     int64   `json:"mileage" gorm:"type:int;comment:里程数"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null;comment:价格"`
	Status      int64   `json:"status" gorm:"type:tinyint;default:1;comment:状态 1:上架 0:下架"`
	Description string  `json:"description" gorm:"type:text;comment:描述"`
	Images      string  `json:"images" gorm:"type:text;comment:图片URL，多个用逗号分隔"`
	Location    string  `json:"location" gorm:"type:varchar(100);comment:车辆位置"`
	Contact     string  `json:"contact" gorm:"type:varchar(50);comment:联系方式"`
}

// TableName 设置表名
func (v *Vehicle) TableName() string {
	return "vehicles"
}

// Create 创建车辆
func (v *Vehicle) Create() error {
	return global.DB.Create(v).Error
}

// Update 更新车辆
func (v *Vehicle) Update() error {
	return global.DB.Save(v).Error
}

// Delete 删除车辆
func (v *Vehicle) Delete() error {
	return global.DB.Delete(v).Error
}

// GetByID 根据ID获取车辆
func (v *Vehicle) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).Limit(1).Find(v).Error
}

// CheckVehicleAvailable 检查车辆是否存在且状态为上架
func (v *Vehicle) CheckVehicleAvailable(vehicleID int64) (*Vehicle, error) {
	vehicle := &Vehicle{}
	err := global.DB.Where("id = ? AND status = ?", vehicleID, 1).Limit(1).Find(vehicle).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("车辆不存在或已下架")
		}
		return nil, fmt.Errorf("查询车辆失败: %w", err)
	}
	return vehicle, nil
}
