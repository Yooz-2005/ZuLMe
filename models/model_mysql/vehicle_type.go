package model_mysql

import (
	"ZuLMe/ZuLMe/Common/global"
	"gorm.io/gorm"
)

// VehicleType 车辆类型表
type VehicleType struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(50);not null;comment:类型名称"`
	Description string `json:"description" gorm:"type:varchar(200);comment:类型描述"`
	Status      int    `json:"status" gorm:"type:tinyint;default:1;comment:状态 1:启用 0:禁用"`
	Sort        int    `json:"sort" gorm:"type:int;default:0;comment:排序"`
}

// TableName 设置表名
func (VehicleType) TableName() string {
	return "vehicle_types"
}

// Create 创建车辆类型
func (vt *VehicleType) Create() error {
	return global.DB.Create(vt).Error
}

// Update 更新车辆类型
func (vt *VehicleType) Update() error {
	return global.DB.Save(vt).Error
}

// Delete 删除车辆类型
func (vt *VehicleType) Delete() error {
	return global.DB.Delete(vt).Error
}

// GetByID 根据ID获取车辆类型
func (vt *VehicleType) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).First(vt).Error
}
