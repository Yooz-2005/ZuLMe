package logic

import (
	"ZuLMe/ZuLMe/Common/global"
	vehicle "ZuLMe/ZuLMe/Srv/vehicle_srv/proto_vehicle"
	"ZuLMe/ZuLMe/models/model_mysql"
	"context"
	"fmt"
	"time"
)

// CheckVehicleAvailability 检查车辆可用性
func CheckVehicleAvailability(ctx context.Context, req *vehicle.CheckAvailabilityRequest) (*vehicle.CheckAvailabilityResponse, error) {
	// 参数验证
	if req.VehicleId <= 0 {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	if req.StartDate == "" || req.EndDate == "" {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "开始日期和结束日期不能为空",
		}, nil
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 验证日期逻辑
	if endDate.Before(startDate) {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "结束日期不能早于开始日期",
		}, nil
	}

	if startDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return &vehicle.CheckAvailabilityResponse{
			Code:    400,
			Message: "开始日期不能早于今天",
		}, nil
	}

	// 检查车辆是否存在
	vehicleModel := &model_mysql.Vehicle{}
	if err := vehicleModel.GetByID(uint(req.VehicleId)); err != nil {
		return &vehicle.CheckAvailabilityResponse{
			Code:    404,
			Message: "车辆不存在",
		}, nil
	}

	// 检查车辆状态
	if vehicleModel.Status != 1 {
		return &vehicle.CheckAvailabilityResponse{
			Code:        200,
			Message:     "车辆当前不可用",
			IsAvailable: false,
		}, nil
	}

	// 检查库存可用性
	inventoryModel := &model_mysql.VehicleInventory{}
	available, err := inventoryModel.CheckAvailability(uint(req.VehicleId), startDate, endDate)
	if err != nil {
		return &vehicle.CheckAvailabilityResponse{
			Code:    500,
			Message: "检查库存失败",
		}, err
	}

	return &vehicle.CheckAvailabilityResponse{
		Code:        200,
		Message:     "检查成功",
		IsAvailable: available,
	}, nil
}

// CreateReservation 创建预订（新流程：先预订后订单）
func CreateReservation(ctx context.Context, req *vehicle.CreateReservationRequest) (*vehicle.CreateReservationResponse, error) {
	// 参数验证
	if req.VehicleId <= 0 {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	if req.UserId <= 0 {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 验证日期逻辑
	if endDate.Before(startDate) {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "结束日期不能早于开始日期",
		}, nil
	}

	// 1. 检查车辆是否存在
	var vehicleModel model_mysql.Vehicle
	if err := vehicleModel.GetByID(uint(req.VehicleId)); err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    404,
			Message: "车辆不存在",
		}, nil
	}

	// 2. 检查车辆状态是否可预订
	// 车辆状态：1:上架 2:下架 3:可租 4:需预定
	if vehicleModel.Status == model_mysql.VehicleStatusOffline { // 下架状态
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "车辆已下架，暂不可预订",
		}, nil
	}

	// 3. 检查车辆是否处于维护中
	inventoryModel := &model_mysql.VehicleInventory{}
	maintenances, err := inventoryModel.GetMaintenanceSchedule(uint(req.VehicleId))
	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    500,
			Message: "检查车辆维护状态失败",
		}, err
	}

	// 检查预订时间段是否与维护时间冲突
	for _, maintenance := range maintenances {
		if (startDate.Before(maintenance.EndDate) || startDate.Equal(maintenance.EndDate)) &&
			(endDate.After(maintenance.StartDate) || endDate.Equal(maintenance.StartDate)) {
			return &vehicle.CreateReservationResponse{
				Code: 400,
				Message: fmt.Sprintf("车辆在 %s 至 %s 期间处于维护中，无法预订",
					maintenance.StartDate.Format("2006-01-02"),
					maintenance.EndDate.Format("2006-01-02")),
			}, nil
		}
	}

	// 4. 检查车辆是否已被预订（时间冲突检查）
	available, err := inventoryModel.CheckAvailability(uint(req.VehicleId), startDate, endDate)
	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    500,
			Message: "检查车辆可用性失败",
		}, err
	}

	if !available {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: "车辆在指定时间段已被预订，请选择其他时间",
		}, nil
	}

	// 5. 创建预订
	reservation, err := inventoryModel.CreateReservation(
		uint(req.VehicleId),
		startDate,
		endDate,
		uint(req.UserId),
		req.Notes,
	)

	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    500,
			Message: "创建预订失败",
		}, err
	}

	return &vehicle.CreateReservationResponse{
		Code:          200,
		Message:       "预订创建成功",
		ReservationId: int64(reservation.ID),
	}, nil
}

// UpdateReservationStatus 更新预订状态
func UpdateReservationStatus(ctx context.Context, req *vehicle.UpdateReservationStatusRequest) (*vehicle.UpdateReservationStatusResponse, error) {
	if req.OrderId <= 0 {
		return &vehicle.UpdateReservationStatusResponse{
			Code:    400,
			Message: "订单ID不能为空",
		}, nil
	}

	inventoryModel := &model_mysql.VehicleInventory{}
	var err error

	switch req.Status {
	case "rented":
		// 更新为租用中
		err = inventoryModel.UpdateReservationToRented(uint(req.OrderId))
	case "completed":
		// 完成租用
		err = inventoryModel.CompleteRental(uint(req.OrderId))
	case "cancelled":
		// 取消预订
		err = inventoryModel.CancelReservation(uint(req.OrderId))
	default:
		return &vehicle.UpdateReservationStatusResponse{
			Code:    400,
			Message: "无效的状态",
		}, nil
	}

	if err != nil {
		return &vehicle.UpdateReservationStatusResponse{
			Code:    500,
			Message: "更新状态失败",
		}, err
	}

	return &vehicle.UpdateReservationStatusResponse{
		Code:    200,
		Message: "状态更新成功",
	}, nil
}

// GetAvailableVehicles 获取可用车辆列表
func GetAvailableVehicles(ctx context.Context, req *vehicle.GetAvailableVehiclesRequest) (*vehicle.GetAvailableVehiclesResponse, error) {
	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 获取所有符合条件的车辆ID
	var allVehicles []model_mysql.Vehicle
	var vehicleIDs []uint

	// 构建查询条件
	query := global.DB.Model(&model_mysql.Vehicle{}).Where("status = 1")

	if req.MerchantId > 0 {
		query = query.Where("merchant_id = ?", req.MerchantId)
	}
	if req.TypeId > 0 {
		query = query.Where("type_id = ?", req.TypeId)
	}
	if req.BrandId > 0 {
		query = query.Where("brand_id = ?", req.BrandId)
	}
	if req.PriceMin > 0 {
		query = query.Where("price >= ?", req.PriceMin)
	}
	if req.PriceMax > 0 {
		query = query.Where("price <= ?", req.PriceMax)
	}

	if err := query.Find(&allVehicles).Error; err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    500,
			Message: "查询车辆失败",
		}, err
	}

	// 提取车辆ID
	for _, v := range allVehicles {
		vehicleIDs = append(vehicleIDs, v.ID)
	}

	// 检查库存可用性
	inventoryModel := &model_mysql.VehicleInventory{}
	availableVehicleIDs, err := inventoryModel.GetAvailableVehicles(startDate, endDate, vehicleIDs)
	if err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    500,
			Message: "检查库存失败",
		}, err
	}

	// 过滤出可用的车辆
	var availableVehicles []*vehicle.VehicleInfo
	availableMap := make(map[uint]bool)
	for _, id := range availableVehicleIDs {
		availableMap[id] = true
	}

	for _, v := range allVehicles {
		if availableMap[v.ID] {
			vehicleInfo := &vehicle.VehicleInfo{
				Id:          int64(v.ID),
				MerchantId:  v.MerchantID,
				TypeId:      v.TypeID,
				BrandId:     v.BrandID,
				Brand:       v.Brand,
				Style:       v.Style,
				Year:        v.Year,
				Color:       v.Color,
				Mileage:     v.Mileage,
				Price:       v.Price,
				Status:      v.Status,
				Description: v.Description,
				Images:      v.Images,
				Location:    v.Location,
				Contact:     v.Contact,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			}
			availableVehicles = append(availableVehicles, vehicleInfo)
		}
	}

	return &vehicle.GetAvailableVehiclesResponse{
		Code:     200,
		Message:  "获取成功",
		Vehicles: availableVehicles,
		Total:    int64(len(availableVehicles)),
	}, nil
}

// GetInventoryStatistics 获取库存统计
func GetInventoryStatistics(ctx context.Context, req *vehicle.GetInventoryStatsRequest) (*vehicle.GetInventoryStatsResponse, error) {
	if req.MerchantId <= 0 {
		return &vehicle.GetInventoryStatsResponse{
			Code:    400,
			Message: "商家ID不能为空",
		}, nil
	}

	inventoryModel := &model_mysql.VehicleInventory{}
	stats, err := inventoryModel.GetInventoryStatistics(uint(req.MerchantId))
	if err != nil {
		return &vehicle.GetInventoryStatsResponse{
			Code:    500,
			Message: "获取统计信息失败",
		}, err
	}

	return &vehicle.GetInventoryStatsResponse{
		Code:        200,
		Message:     "获取成功",
		Total:       stats["total"],
		Available:   stats["available"],
		Reserved:    stats["reserved"],
		Rented:      stats["rented"],
		Maintenance: stats["maintenance"],
	}, nil
}

// SetVehicleMaintenance 设置车辆维护状态
func SetVehicleMaintenance(ctx context.Context, req *vehicle.SetMaintenanceRequest) (*vehicle.SetMaintenanceResponse, error) {
	// 参数验证
	if req.VehicleId <= 0 {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	if req.StartDate == "" || req.EndDate == "" {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "开始日期和结束日期不能为空",
		}, nil
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 验证日期逻辑
	if endDate.Before(startDate) {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "结束日期不能早于开始日期",
		}, nil
	}

	// 检查车辆是否存在
	vehicleModel := &model_mysql.Vehicle{}
	if err := vehicleModel.GetByID(uint(req.VehicleId)); err != nil {
		return &vehicle.SetMaintenanceResponse{
			Code:    404,
			Message: "车辆不存在",
		}, nil
	}

	// 检查时间段是否有冲突
	inventoryModel := &model_mysql.VehicleInventory{}
	available, err := inventoryModel.CheckAvailability(uint(req.VehicleId), startDate, endDate)
	if err != nil {
		return &vehicle.SetMaintenanceResponse{
			Code:    500,
			Message: "检查时间冲突失败",
		}, err
	}

	if !available {
		return &vehicle.SetMaintenanceResponse{
			Code:    400,
			Message: "该时间段已有预订或维护安排",
		}, nil
	}

	// 设置维护状态
	err = inventoryModel.SetMaintenanceStatus(
		uint(req.VehicleId),
		startDate,
		endDate,
		req.Notes,
		uint(req.CreatedBy),
	)

	if err != nil {
		return &vehicle.SetMaintenanceResponse{
			Code:    500,
			Message: "设置维护状态失败",
		}, err
	}

	return &vehicle.SetMaintenanceResponse{
		Code:    200,
		Message: "维护计划设置成功",
	}, nil
}

// GetMaintenanceSchedule 获取车辆维护计划
func GetMaintenanceSchedule(ctx context.Context, req *vehicle.GetMaintenanceScheduleRequest) (*vehicle.GetMaintenanceScheduleResponse, error) {
	if req.VehicleId <= 0 {
		return &vehicle.GetMaintenanceScheduleResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	inventoryModel := &model_mysql.VehicleInventory{}
	maintenances, err := inventoryModel.GetMaintenanceSchedule(uint(req.VehicleId))
	if err != nil {
		return &vehicle.GetMaintenanceScheduleResponse{
			Code:    500,
			Message: "获取维护计划失败",
		}, err
	}

	var schedules []*vehicle.MaintenanceInfo
	for _, m := range maintenances {
		schedules = append(schedules, &vehicle.MaintenanceInfo{
			Id:        int64(m.ID),
			VehicleId: int64(m.VehicleID),
			StartDate: m.StartDate.Format("2006-01-02"),
			EndDate:   m.EndDate.Format("2006-01-02"),
			Notes:     m.Notes,
			CreatedAt: m.CreatedAt.Format(time.RFC3339),
		})
	}

	return &vehicle.GetMaintenanceScheduleResponse{
		Code:         200,
		Message:      "获取成功",
		Maintenances: schedules,
		Total:        int64(len(schedules)),
	}, nil
}

// GetInventoryCalendar 获取库存日历
func GetInventoryCalendar(ctx context.Context, req *vehicle.GetInventoryCalendarRequest) (*vehicle.GetInventoryCalendarResponse, error) {
	if req.VehicleId <= 0 {
		return &vehicle.GetInventoryCalendarResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.GetInventoryCalendarResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.GetInventoryCalendarResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 获取指定日期范围内的库存记录
	inventoryModel := &model_mysql.VehicleInventory{}
	inventories, err := inventoryModel.GetVehicleInventoryByDateRange(uint(req.VehicleId), startDate, endDate)
	if err != nil {
		return &vehicle.GetInventoryCalendarResponse{
			Code:    500,
			Message: "获取库存日历失败",
		}, err
	}

	// 生成日历数据
	var calendar []*vehicle.InventoryCalendarItem
	currentDate := startDate

	for !currentDate.After(endDate) {
		dateStr := currentDate.Format("2006-01-02")
		status := int32(1) // 默认可用
		var orderID int64 = 0
		var notes string

		// 检查当前日期是否有库存记录
		for _, inv := range inventories {
			if !currentDate.Before(inv.StartDate) && !currentDate.After(inv.EndDate) {
				status = int32(inv.Status)
				orderID = int64(inv.OrderID)
				notes = inv.Notes
				break
			}
		}

		calendar = append(calendar, &vehicle.InventoryCalendarItem{
			Date:    dateStr,
			Status:  status,
			OrderId: orderID,
			Notes:   notes,
		})

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return &vehicle.GetInventoryCalendarResponse{
		Code:     200,
		Message:  "获取成功",
		Calendar: calendar,
	}, nil
}

// GetInventoryReport 获取库存报表
func GetInventoryReport(ctx context.Context, req *vehicle.GetInventoryReportRequest) (*vehicle.GetInventoryReportResponse, error) {
	if req.MerchantId <= 0 {
		return &vehicle.GetInventoryReportResponse{
			Code:    400,
			Message: "商家ID不能为空",
		}, nil
	}

	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return &vehicle.GetInventoryReportResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return &vehicle.GetInventoryReportResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 验证日期逻辑
	if endDate.Before(startDate) {
		return &vehicle.GetInventoryReportResponse{
			Code:    400,
			Message: "结束日期不能早于开始日期",
		}, nil
	}

	// 获取库存报表
	inventoryModel := &model_mysql.VehicleInventory{}
	report, err := inventoryModel.GetInventoryReport(uint(req.MerchantId), startDate, endDate)
	if err != nil {
		return &vehicle.GetInventoryReportResponse{
			Code:    500,
			Message: "获取库存报表失败",
		}, err
	}

	return &vehicle.GetInventoryReportResponse{
		Code:            200,
		Message:         "获取成功",
		TotalVehicles:   report["total_vehicles"].(int64),
		TotalDays:       report["total_days"].(int64),
		TotalCapacity:   report["total_capacity"].(int64),
		Reservations:    report["reservations"].(int64),
		Rentals:         report["rentals"].(int64),
		Maintenances:    report["maintenances"].(int64),
		UsedCapacity:    report["used_capacity"].(int64),
		UtilizationRate: report["utilization_rate"].(float64),
	}, nil
}
