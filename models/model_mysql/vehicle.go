package model_mysql

import (
	"ZuLMe/ZuLMe/Common/global"
	"gorm.io/gorm"
)

// Vehicle 车辆表
type Vehicle struct {
	gorm.Model
	MerchantID  int64   `json:"merchant_id" gorm:"not null;comment:商家ID"`
	TypeID      int64   `json:"type_id" gorm:"not null;comment:车辆类型ID"`
	BrandID     int64   `json:"brand_id" gorm:"not null;comment:品牌ID"`
	Brand       string  `json:"brand" gorm:"type:varchar(50);not null;comment:品牌名称"` // 冗余字段，便于查询
	Style       string  `json:"style" gorm:"type:varchar(50);not null;comment:型号"`
	Year        int64   `json:"year" gorm:"type:int;not null;comment:年份"`
	Color       string  `json:"color" gorm:"type:varchar(20);comment:颜色"`
	Mileage     int64   `json:"mileage" gorm:"type:int;comment:里程数"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null;comment:价格"`
	Status      int64   `json:"status" gorm:"type:tinyint;default:1;comment:状态 1:上架 2:下架 3:可租 4:需预定"`
	Description string  `json:"description" gorm:"type:text;comment:描述"`
	Images      string  `json:"images" gorm:"type:text;comment:图片URL，多个用逗号分隔"`
	Location    string  `json:"location" gorm:"type:varchar(100);comment:车辆位置"`
	Contact     string  `json:"contact" gorm:"type:varchar(50);comment:联系方式"`
}

// 车辆状态常量
const (
	VehicleStatusOnline      = 1 // 上架
	VehicleStatusOffline     = 2 // 下架
	VehicleStatusAvailable   = 3 // 可租
	VehicleStatusReservation = 4 // 需预定
)

// TableName 设置表名
func (Vehicle) TableName() string {
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

// GetListByBrandID 根据品牌ID获取车辆列表
func (v *Vehicle) GetListByBrandID(brandID int64, page, pageSize int) ([]Vehicle, int64, error) {
	var vehicles []Vehicle
	var total int64

	query := global.DB.Model(&Vehicle{}).Where("brand_id = ? AND status = 1", brandID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&vehicles).Error; err != nil {
		return nil, 0, err
	}

	return vehicles, total, nil
}

// SearchByBrand 根据品牌名称搜索车辆
func (v *Vehicle) SearchByBrand(brandName string, page, pageSize int) ([]Vehicle, int64, error) {
	var vehicles []Vehicle
	var total int64

	query := global.DB.Model(&Vehicle{}).Where("brand LIKE ? AND status = 1", "%"+brandName+"%")

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&vehicles).Error; err != nil {
		return nil, 0, err
	}

	return vehicles, total, nil
}

// UpdateBrandInfo 更新车辆的品牌信息（当品牌表信息更新时同步）
func (v *Vehicle) UpdateBrandInfo(brandID int64, brandName string) error {
	return global.DB.Model(&Vehicle{}).Where("brand_id = ?", brandID).Update("brand", brandName).Error
}
