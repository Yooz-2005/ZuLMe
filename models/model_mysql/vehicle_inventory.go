package model_mysql

import (
	"ZuLMe/ZuLMe/Common/global"
	"time"

	"gorm.io/gorm"
)

// VehicleInventory 车辆库存表
type VehicleInventory struct {
	gorm.Model
	VehicleID uint      `json:"vehicle_id" gorm:"not null;comment:车辆ID"`
	StartDate time.Time `json:"start_date" gorm:"not null;comment:开始日期"`
	EndDate   time.Time `json:"end_date" gorm:"not null;comment:结束日期"`
	Status    int       `json:"status" gorm:"type:tinyint;default:1;comment:库存状态 1:可租用 2:已预订 3:租用中 4:维护中 5:不可用"`
	OrderID   uint      `json:"order_id" gorm:"default:0;comment:关联订单ID，0表示无订单"`
	Quantity  int       `json:"quantity" gorm:"default:1;comment:数量(通常为1，支持同型号多辆车)"`
	Notes     string    `json:"notes" gorm:"type:varchar(500);comment:备注"`
	CreatedBy uint      `json:"created_by" gorm:"comment:创建人ID"`
	UpdatedBy uint      `json:"updated_by" gorm:"comment:更新人ID"`
}

// 库存状态常量
const (
	InventoryStatusAvailable   = 1 // 可租用
	InventoryStatusReserved    = 2 // 已预订
	InventoryStatusRented      = 3 // 租用中
	InventoryStatusMaintenance = 4 // 维护中
	InventoryStatusUnavailable = 5 // 不可用
)

// TableName 设置表名
func (VehicleInventory) TableName() string {
	return "vehicle_inventories"
}

// Create 创建库存记录
func (vi *VehicleInventory) Create() error {
	return global.DB.Create(vi).Error
}

// Update 更新库存记录
func (vi *VehicleInventory) Update() error {
	return global.DB.Save(vi).Error
}

// Delete 删除库存记录
func (vi *VehicleInventory) Delete() error {
	return global.DB.Delete(vi).Error
}

// GetByID 根据ID获取库存记录
func (vi *VehicleInventory) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).First(vi).Error
}

// CheckAvailability 检查车辆在指定时间段的可用性
func (vi *VehicleInventory) CheckAvailability(vehicleID uint, startDate, endDate time.Time) (bool, error) {
	var count int64

	// 查询是否有冲突的预订或租用记录
	err := global.DB.Model(&VehicleInventory{}).Where(
		"vehicle_id = ? AND status IN (?, ?) AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
		vehicleID,
		InventoryStatusReserved, InventoryStatusRented,
		startDate, startDate,
		endDate, endDate,
		startDate, endDate,
	).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

// CreateReservation 创建预订记录（新流程：先预订后订单）
func (vi *VehicleInventory) CreateReservation(vehicleID uint, startDate, endDate time.Time, createdBy uint, notes string) (*VehicleInventory, error) {
	// 先检查可用性
	available, err := vi.CheckAvailability(vehicleID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gorm.ErrRecordNotFound // 可以自定义错误类型
	}

	// 创建预订记录（不需要order_id）
	reservation := &VehicleInventory{
		VehicleID: vehicleID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    InventoryStatusReserved,
		OrderID:   0, // 初始为0，创建订单时再关联
		Quantity:  1,
		Notes:     notes,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}

	err = reservation.Create()
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

// UpdateReservationOrderID 更新预订的订单ID（创建订单时调用）
func (vi *VehicleInventory) UpdateReservationOrderID(reservationID uint, orderID uint) error {
	return global.DB.Model(&VehicleInventory{}).
		Where("id = ?", reservationID).
		Update("order_id", orderID).Error
}

// UpdateReservationToRented 将预订状态更新为租用中
func (vi *VehicleInventory) UpdateReservationToRented(orderID uint) error {
	return global.DB.Model(&VehicleInventory{}).
		Where("order_id = ? AND status = ?", orderID, InventoryStatusReserved).
		Update("status", InventoryStatusRented).Error
}

// CompleteRental 完成租用，释放库存
func (vi *VehicleInventory) CompleteRental(orderID uint) error {
	return global.DB.Where("order_id = ? AND status = ?", orderID, InventoryStatusRented).
		Delete(&VehicleInventory{}).Error
}

// CancelReservation 取消预订
func (vi *VehicleInventory) CancelReservation(orderID uint) error {
	return global.DB.Where("order_id = ? AND status = ?", orderID, InventoryStatusReserved).
		Delete(&VehicleInventory{}).Error
}

// GetVehicleInventoryByDateRange 获取指定日期范围内的车辆库存状态
func (vi *VehicleInventory) GetVehicleInventoryByDateRange(vehicleID uint, startDate, endDate time.Time) ([]VehicleInventory, error) {
	var inventories []VehicleInventory

	err := global.DB.Where(
		"vehicle_id = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
		vehicleID,
		startDate, startDate,
		endDate, endDate,
		startDate, endDate,
	).Order("start_date ASC").Find(&inventories).Error

	return inventories, err
}

// GetAvailableVehicles 获取指定时间段内可用的车辆列表
func (vi *VehicleInventory) GetAvailableVehicles(startDate, endDate time.Time, vehicleIDs []uint) ([]uint, error) {
	var unavailableVehicleIDs []uint

	// 查询在指定时间段内不可用的车辆ID
	err := global.DB.Model(&VehicleInventory{}).
		Select("DISTINCT vehicle_id").
		Where(
			"status IN (?, ?, ?) AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?))",
			InventoryStatusReserved, InventoryStatusRented, InventoryStatusMaintenance,
			startDate, startDate,
			endDate, endDate,
			startDate, endDate,
		).
		Pluck("vehicle_id", &unavailableVehicleIDs).Error

	if err != nil {
		return nil, err
	}

	// 从输入的车辆ID列表中排除不可用的车辆
	var availableVehicleIDs []uint
	unavailableMap := make(map[uint]bool)
	for _, id := range unavailableVehicleIDs {
		unavailableMap[id] = true
	}

	for _, id := range vehicleIDs {
		if !unavailableMap[id] {
			availableVehicleIDs = append(availableVehicleIDs, id)
		}
	}

	return availableVehicleIDs, nil
}

// SetMaintenanceStatus 设置车辆维护状态
func (vi *VehicleInventory) SetMaintenanceStatus(vehicleID uint, startDate, endDate time.Time, notes string, createdBy uint) error {
	maintenance := &VehicleInventory{
		VehicleID: vehicleID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    InventoryStatusMaintenance,
		Notes:     notes,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}

	return maintenance.Create()
}

// GetMaintenanceSchedule 获取车辆维护计划
func (vi *VehicleInventory) GetMaintenanceSchedule(vehicleID uint) ([]VehicleInventory, error) {
	var maintenances []VehicleInventory

	err := global.DB.Where("vehicle_id = ? AND status = ? AND end_date >= ?",
		vehicleID, InventoryStatusMaintenance, time.Now()).
		Order("start_date ASC").Find(&maintenances).Error

	return maintenances, err
}

// GetInventoryStatistics 获取库存统计信息
func (vi *VehicleInventory) GetInventoryStatistics(merchantID uint) (map[string]int64, error) {
	stats := make(map[string]int64)

	// 获取商家的所有车辆ID
	var vehicleIDs []uint
	err := global.DB.Model(&Vehicle{}).Where("merchant_id = ?", merchantID).Pluck("id", &vehicleIDs).Error
	if err != nil {
		return nil, err
	}

	if len(vehicleIDs) == 0 {
		return stats, nil
	}

	now := time.Now()

	// 统计各种状态的车辆数量
	statuses := []int{InventoryStatusReserved, InventoryStatusRented, InventoryStatusMaintenance}
	statusNames := []string{"reserved", "rented", "maintenance"}

	for i, status := range statuses {
		var count int64
		err := global.DB.Model(&VehicleInventory{}).
			Where("vehicle_id IN ? AND status = ? AND start_date <= ? AND end_date >= ?",
				vehicleIDs, status, now, now).
			Count(&count).Error
		if err != nil {
			return nil, err
		}
		stats[statusNames[i]] = count
	}

	// 计算可用车辆数量
	stats["total"] = int64(len(vehicleIDs))
	stats["available"] = stats["total"] - stats["reserved"] - stats["rented"] - stats["maintenance"]

	return stats, nil
}

// BatchCreateReservations 批量创建预订
func (vi *VehicleInventory) BatchCreateReservations(reservations []VehicleInventory) error {
	// 使用事务确保数据一致性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, reservation := range reservations {
		// 检查每个预订的可用性
		available, err := vi.CheckAvailability(reservation.VehicleID, reservation.StartDate, reservation.EndDate)
		if err != nil {
			tx.Rollback()
			return err
		}
		if !available {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}

		// 创建预订记录
		if err := tx.Create(&reservation).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// BatchCancelReservations 批量取消预订
func (vi *VehicleInventory) BatchCancelReservations(orderIDs []uint) error {
	return global.DB.Where("order_id IN ? AND status = ?", orderIDs, InventoryStatusReserved).
		Delete(&VehicleInventory{}).Error
}

// GetInventoryReport 获取库存报表
func (vi *VehicleInventory) GetInventoryReport(merchantID uint, startDate, endDate time.Time) (map[string]interface{}, error) {
	report := make(map[string]interface{})

	// 获取商家的所有车辆ID
	var vehicleIDs []uint
	err := global.DB.Model(&Vehicle{}).Where("merchant_id = ?", merchantID).Pluck("id", &vehicleIDs).Error
	if err != nil {
		return nil, err
	}

	if len(vehicleIDs) == 0 {
		return report, nil
	}

	// 统计预订数量
	var reservationCount int64
	err = global.DB.Model(&VehicleInventory{}).
		Where("vehicle_id IN ? AND status = ? AND start_date >= ? AND end_date <= ?",
			vehicleIDs, InventoryStatusReserved, startDate, endDate).
		Count(&reservationCount).Error
	if err != nil {
		return nil, err
	}

	// 统计租用数量
	var rentalCount int64
	err = global.DB.Model(&VehicleInventory{}).
		Where("vehicle_id IN ? AND status = ? AND start_date >= ? AND end_date <= ?",
			vehicleIDs, InventoryStatusRented, startDate, endDate).
		Count(&rentalCount).Error
	if err != nil {
		return nil, err
	}

	// 统计维护数量
	var maintenanceCount int64
	err = global.DB.Model(&VehicleInventory{}).
		Where("vehicle_id IN ? AND status = ? AND start_date >= ? AND end_date <= ?",
			vehicleIDs, InventoryStatusMaintenance, startDate, endDate).
		Count(&maintenanceCount).Error
	if err != nil {
		return nil, err
	}

	// 计算利用率
	totalDays := int(endDate.Sub(startDate).Hours()/24) + 1
	totalCapacity := int64(len(vehicleIDs)) * int64(totalDays)
	usedCapacity := reservationCount + rentalCount + maintenanceCount
	utilizationRate := float64(usedCapacity) / float64(totalCapacity) * 100

	report["total_vehicles"] = int64(len(vehicleIDs))
	report["total_days"] = int64(totalDays)
	report["total_capacity"] = totalCapacity
	report["reservations"] = reservationCount
	report["rentals"] = rentalCount
	report["maintenances"] = maintenanceCount
	report["used_capacity"] = usedCapacity
	report["utilization_rate"] = utilizationRate

	return report, nil
}

// GetVehicleUtilization 获取单个车辆的利用率
func (vi *VehicleInventory) GetVehicleUtilization(vehicleID uint, startDate, endDate time.Time) (float64, error) {
	var usedDays int64

	err := global.DB.Model(&VehicleInventory{}).
		Where("vehicle_id = ? AND status IN (?, ?, ?) AND start_date >= ? AND end_date <= ?",
			vehicleID, InventoryStatusReserved, InventoryStatusRented, InventoryStatusMaintenance,
			startDate, endDate).
		Select("SUM(DATEDIFF(end_date, start_date) + 1)").
		Scan(&usedDays).Error

	if err != nil {
		return 0, err
	}

	totalDays := int(endDate.Sub(startDate).Hours()/24) + 1
	utilizationRate := float64(usedDays) / float64(totalDays) * 100

	return utilizationRate, nil
}
