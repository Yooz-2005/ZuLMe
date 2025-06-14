package model_mysql

import (
	"Common/global"

	"gorm.io/gorm"
)

// VehicleBrand 车辆品牌表
type VehicleBrand struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(50);not null;unique;comment:品牌名称"`
	EnglishName string `json:"english_name" gorm:"type:varchar(50);comment:英文名称"`
	Logo        string `json:"logo" gorm:"type:varchar(255);comment:品牌Logo URL"`
	Country     string `json:"country" gorm:"type:varchar(50);comment:品牌国家"`
	Description string `json:"description" gorm:"type:varchar(500);comment:品牌描述"`
	Status      int    `json:"status" gorm:"type:tinyint;default:1;comment:状态 1:启用 0:禁用"`
	Sort        int    `json:"sort" gorm:"type:int;default:0;comment:排序"`
	IsHot       int    `json:"is_hot" gorm:"type:tinyint;default:0;comment:是否热门 1:是 0:否"`
}

// TableName 设置表名
func (VehicleBrand) TableName() string {
	return "vehicle_brands"
}

// Create 创建车辆品牌
func (vb *VehicleBrand) Create() error {
	return global.DB.Create(vb).Error
}

// Update 更新车辆品牌
func (vb *VehicleBrand) Update() error {
	return global.DB.Save(vb).Error
}

// Delete 删除车辆品牌
func (vb *VehicleBrand) Delete() error {
	return global.DB.Delete(vb).Error
}

// GetByID 根据ID获取车辆品牌
func (vb *VehicleBrand) GetByID(id uint) error {
	return global.DB.Where("id = ? AND status = 1", id).First(vb).Error
}

// GetByName 根据名称获取车辆品牌
func (vb *VehicleBrand) GetByName(name string) error {
	return global.DB.Where("name = ? AND status = 1", name).First(vb).Error
}

// GetList 获取品牌列表
func (vb *VehicleBrand) GetList(page, pageSize int, isHot *int) ([]VehicleBrand, int64, error) {
	var brands []VehicleBrand
	var total int64

	query := global.DB.Model(&VehicleBrand{}).Where("status = 1")
	
	// 如果指定了是否热门
	if isHot != nil {
		query = query.Where("is_hot = ?", *isHot)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("sort DESC, id DESC").Offset(offset).Limit(pageSize).Find(&brands).Error; err != nil {
		return nil, 0, err
	}

	return brands, total, nil
}

// GetAll 获取所有启用的品牌
func (vb *VehicleBrand) GetAll() ([]VehicleBrand, error) {
	var brands []VehicleBrand
	err := global.DB.Where("status = 1").Order("sort DESC, name ASC").Find(&brands).Error
	return brands, err
}

// GetHotBrands 获取热门品牌
func (vb *VehicleBrand) GetHotBrands(limit int) ([]VehicleBrand, error) {
	var brands []VehicleBrand
	query := global.DB.Where("status = 1 AND is_hot = 1").Order("sort DESC, id DESC")
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	err := query.Find(&brands).Error
	return brands, err
}

// UpdateStatus 更新状态
func (vb *VehicleBrand) UpdateStatus(id uint, status int) error {
	return global.DB.Model(&VehicleBrand{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateSort 更新排序
func (vb *VehicleBrand) UpdateSort(id uint, sort int) error {
	return global.DB.Model(&VehicleBrand{}).Where("id = ?", id).Update("sort", sort).Error
}

// SetHot 设置热门状态
func (vb *VehicleBrand) SetHot(id uint, isHot int) error {
	return global.DB.Model(&VehicleBrand{}).Where("id = ?", id).Update("is_hot", isHot).Error
}

// CheckNameExists 检查品牌名称是否存在
func (vb *VehicleBrand) CheckNameExists(name string, excludeID uint) (bool, error) {
	var count int64
	query := global.DB.Model(&VehicleBrand{}).Where("name = ?", name)
	
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	
	err := query.Count(&count).Error
	return count > 0, err
}

// GetVehicleCountByBrand 获取品牌下的车辆数量
func (vb *VehicleBrand) GetVehicleCountByBrand(brandID uint) (int64, error) {
	var count int64
	err := global.DB.Model(&Vehicle{}).Where("brand_id = ?", brandID).Count(&count).Error
	return count, err
}
