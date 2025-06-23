package logic

import (
	"Common/global"
	"context"
	"fmt"
	"models/model_mysql"
	"time"
	vehicle "vehicle_srv/proto_vehicle"

	"gorm.io/gorm"
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

	// 幂等性检查：检查用户是否有未支付的订单
	var orderModel model_mysql.Orders
	hasUnpaidOrder, unpaidOrder, err := orderModel.CheckUserHasUnpaidOrder(uint(req.UserId))
	if err != nil {
		return &vehicle.CreateReservationResponse{
			Code:    500,
			Message: "检查用户订单状态失败",
		}, err
	}

	if hasUnpaidOrder {
		return &vehicle.CreateReservationResponse{
			Code:    400,
			Message: fmt.Sprintf("您有未完成支付的订单（订单号：%s），请先完成支付后再进行新的预订", unpaidOrder.OrderSn),
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

// CancelReservation 取消预订
func CancelReservation(ctx context.Context, req *vehicle.CancelReservationRequest) (*vehicle.CancelReservationResponse, error) {
	if req.ReservationId == "" {
		return &vehicle.CancelReservationResponse{
			Code:    400,
			Message: "预订ID不能为空",
		}, nil
	}

	// 从预订ID中提取库存ID（格式：RES123 -> 123）
	var inventoryID uint
	if _, err := fmt.Sscanf(req.ReservationId, "RES%d", &inventoryID); err != nil {
		return &vehicle.CancelReservationResponse{
			Code:    400,
			Message: "无效的预订ID格式",
		}, nil
	}

	// 查找预订记录
	var inventory model_mysql.VehicleInventory
	if err := global.DB.Where("id = ?", inventoryID).First(&inventory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &vehicle.CancelReservationResponse{
				Code:    404,
				Message: "预订记录不存在",
			}, nil
		}
		return &vehicle.CancelReservationResponse{
			Code:    500,
			Message: "查询预订记录失败",
		}, err
	}

	// 检查预订状态，只有已预订状态才能取消
	if inventory.Status != model_mysql.InventoryStatusReserved {
		return &vehicle.CancelReservationResponse{
			Code:    400,
			Message: "只能取消已预订状态的预订",
		}, nil
	}

	// 如果已经有关联订单且订单已支付，不允许取消
	if inventory.OrderID > 0 {
		var paymentStatus int
		if err := global.DB.Table("orders").Select("payment_status").Where("id = ?", inventory.OrderID).Scan(&paymentStatus).Error; err == nil && paymentStatus == 2 {
			return &vehicle.CancelReservationResponse{
				Code:    400,
				Message: "订单已支付，无法取消预订",
			}, nil
		}
	}

	// 更新预订状态为不可用（表示已取消）
	inventory.Status = model_mysql.InventoryStatusUnavailable
	inventory.Notes = "用户取消预订"

	if err := global.DB.Save(&inventory).Error; err != nil {
		return &vehicle.CancelReservationResponse{
			Code:    500,
			Message: "取消预订失败",
		}, err
	}

	return &vehicle.CancelReservationResponse{
		Code:    200,
		Message: "预订已成功取消",
	}, nil
}

// GetAvailableVehicles 获取可用车辆列表
func GetAvailableVehicles(ctx context.Context, req *vehicle.GetAvailableVehiclesRequest) (*vehicle.GetAvailableVehiclesResponse, error) {
	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.GetStartDate())
	if err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    400,
			Message: "开始日期格式错误",
		}, nil
	}

	endDate, err := time.Parse("2006-01-02", req.GetEndDate())
	if err != nil {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:    400,
			Message: "结束日期格式错误",
		}, nil
	}

	// 获取所有符合条件的车辆ID
	var allVehicles []model_mysql.Vehicle
	var vehicleIDs []uint

	// 构建基础车辆查询条件（只查询上架的车辆）
	query := global.DB.Model(&model_mysql.Vehicle{}).Where("status = 1") // 只查询上架的车辆

	if req.GetMerchantId() > 0 {
		query = query.Where("merchant_id = ?", req.GetMerchantId())
	}
	if req.GetTypeId() > 0 {
		query = query.Where("type_id = ?", req.GetTypeId())
	}
	if req.GetBrandId() > 0 {
		query = query.Where("brand_id = ?", req.GetBrandId())
	}
	if req.GetPriceMin() > 0 {
		query = query.Where("price >= ?", req.GetPriceMin())
	}
	if req.GetPriceMax() > 0 {
		query = query.Where("price <= ?", req.GetPriceMax())
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

	if len(vehicleIDs) == 0 {
		return &vehicle.GetAvailableVehiclesResponse{
			Code:     200,
			Message:  "获取成功",
			Vehicles: []*vehicle.VehicleInfo{},
			Total:    0,
		}, nil
	}

	// 根据库存状态筛选车辆
	var filteredVehicleIDs []uint

	if req.GetStatus() == -1 {
		// 查询所有状态的车辆（包括有库存记录和无库存记录的）
		filteredVehicleIDs = vehicleIDs
	} else if req.GetStatus() > 0 {
		// 查询指定库存状态的车辆
		err := global.DB.Model(&model_mysql.VehicleInventory{}).
			Select("DISTINCT vehicle_id").
			Where("vehicle_id IN ? AND status = ? AND start_date <= ? AND end_date >= ?",
				vehicleIDs, req.GetStatus(), endDate, startDate).
			Pluck("vehicle_id", &filteredVehicleIDs).Error
		if err != nil {
			return &vehicle.GetAvailableVehiclesResponse{
				Code:    500,
				Message: "查询库存状态失败",
			}, err
		}
	} else {
		// 默认查询可用车辆（没有库存记录或库存状态为可租用的车辆）
		var unavailableVehicleIDs []uint
		err := global.DB.Model(&model_mysql.VehicleInventory{}).
			Select("DISTINCT vehicle_id").
			Where("vehicle_id IN ? AND status IN (?, ?, ?, ?) AND start_date <= ? AND end_date >= ?",
				vehicleIDs,
				model_mysql.InventoryStatusReserved,
				model_mysql.InventoryStatusRented,
				model_mysql.InventoryStatusMaintenance,
				model_mysql.InventoryStatusUnavailable,
				endDate, startDate).
			Pluck("vehicle_id", &unavailableVehicleIDs).Error
		if err != nil {
			return &vehicle.GetAvailableVehiclesResponse{
				Code:    500,
				Message: "查询库存状态失败",
			}, err
		}

		// 从所有车辆中排除不可用的车辆
		unavailableMap := make(map[uint]bool)
		for _, id := range unavailableVehicleIDs {
			unavailableMap[id] = true
		}

		for _, id := range vehicleIDs {
			if !unavailableMap[id] {
				filteredVehicleIDs = append(filteredVehicleIDs, id)
			}
		}
	}

	// 过滤出符合条件的车辆
	var allFilteredVehicles []*vehicle.VehicleInfo
	filteredMap := make(map[uint]bool)
	for _, id := range filteredVehicleIDs {
		filteredMap[id] = true
	}

	for _, v := range allVehicles {
		if filteredMap[v.ID] {
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
			allFilteredVehicles = append(allFilteredVehicles, vehicleInfo)
		}
	}

	// 处理分页
	total := int64(len(allFilteredVehicles))
	page := req.GetPage()
	pageSize := req.GetPageSize()

	// 设置默认分页参数
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 12
	}

	// 计算分页
	offset := (page - 1) * pageSize
	var pagedVehicles []*vehicle.VehicleInfo

	if offset < total {
		end := offset + pageSize
		if end > total {
			end = total
		}
		pagedVehicles = allFilteredVehicles[offset:end]
	}

	return &vehicle.GetAvailableVehiclesResponse{
		Code:     200,
		Message:  "获取成功",
		Vehicles: pagedVehicles,
		Total:    total,
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

// GetUserReservationList 获取用户预订列表
func GetUserReservationList(ctx context.Context, req *vehicle.GetUserReservationListRequest) (*vehicle.GetUserReservationListResponse, error) {
	// 参数验证
	if req.UserId <= 0 {
		return &vehicle.GetUserReservationListResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	// 设置默认分页参数
	page := req.Page
	pageSize := req.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 构建查询条件
	query := global.DB.Model(&model_mysql.VehicleInventory{}).
		Where("created_by = ?", req.UserId)

	// 如果指定了状态筛选
	if req.Status != "" {
		// 将前端状态映射到数据库状态
		var dbStatus int
		switch req.Status {
		case "processing":
			dbStatus = model_mysql.InventoryStatusReserved // 处理中 -> 已预订
		case "pending_payment":
			dbStatus = model_mysql.InventoryStatusReserved // 等待付款 -> 已预订
		case "confirmed":
			dbStatus = model_mysql.InventoryStatusReserved // 预订成功 -> 已预订
		case "in_use":
			dbStatus = model_mysql.InventoryStatusRented // 租赁中 -> 已租用
		case "completed":
			dbStatus = model_mysql.InventoryStatusAvailable // 已完成 -> 可用
		case "cancelled":
			dbStatus = model_mysql.InventoryStatusUnavailable // 已取消 -> 不可用
		default:
			// 不筛选状态
		}

		if dbStatus > 0 {
			query = query.Where("status = ?", dbStatus)
		}
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return &vehicle.GetUserReservationListResponse{
			Code:    500,
			Message: "查询预订总数失败",
		}, err
	}

	// 获取预订列表
	var inventories []model_mysql.VehicleInventory
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).
		Order("created_at DESC").Find(&inventories).Error; err != nil {
		return &vehicle.GetUserReservationListResponse{
			Code:    500,
			Message: "查询预订列表失败",
		}, err
	}

	// 转换为响应格式
	var reservations []*vehicle.ReservationInfo
	for _, inventory := range inventories {
		// 获取车辆信息
		var vehicleModel model_mysql.Vehicle
		if err := vehicleModel.GetByID(inventory.VehicleID); err != nil {
			continue // 跳过无法获取车辆信息的记录
		}

		// 映射状态 - 修复状态逻辑
		var status string
		switch inventory.Status {
		case model_mysql.InventoryStatusReserved:
			// 检查是否有关联订单且已支付
			if inventory.OrderID > 0 {
				// 有订单ID，检查订单支付状态
				var paymentStatus int
				if err := global.DB.Table("orders").Select("payment_status").Where("id = ?", inventory.OrderID).Scan(&paymentStatus).Error; err == nil && paymentStatus == 2 {
					status = "confirmed" // 已支付，预订成功 (payment_status = 2 表示已支付)
				} else {
					status = "pending_payment" // 有订单但未支付，等待付款
				}
			} else {
				// 没有订单ID，说明还未创建订单，等待付款
				status = "pending_payment"
			}
		case model_mysql.InventoryStatusRented:
			status = "in_use" // 租赁中
		case model_mysql.InventoryStatusAvailable:
			status = "completed" // 已完成
		case model_mysql.InventoryStatusUnavailable:
			status = "cancelled" // 已取消
		default:
			status = "processing" // 处理中
		}

		reservation := &vehicle.ReservationInfo{
			Id:             fmt.Sprintf("RES%d", inventory.ID),
			VehicleId:      int64(inventory.VehicleID),
			UserId:         int64(inventory.CreatedBy),
			StartDate:      inventory.StartDate.Format("2006-01-02"),
			EndDate:        inventory.EndDate.Format("2006-01-02"),
			PickupLocation: vehicleModel.Location,                                                                        // 使用车辆所在网点地址作为取车地点
			ReturnLocation: "",                                                                                           // 预订阶段不设置还车地点，支付时选择
			TotalAmount:    float64(vehicleModel.Price) * float64(inventory.EndDate.Sub(inventory.StartDate).Hours()/24), // 计算总金额
			Status:         status,
			CreatedAt:      inventory.CreatedAt.Format(time.RFC3339),
		}

		reservation.Vehicle = &vehicle.VehicleInfo{
			Id:          int64(vehicleModel.ID),
			MerchantId:  vehicleModel.MerchantID,
			TypeId:      vehicleModel.TypeID,
			BrandId:     vehicleModel.BrandID,
			Brand:       vehicleModel.Brand,
			Style:       vehicleModel.Style,
			Year:        vehicleModel.Year,
			Color:       vehicleModel.Color,
			Mileage:     vehicleModel.Mileage,
			Price:       vehicleModel.Price,
			Status:      vehicleModel.Status,
			Description: vehicleModel.Description,
			Images:      vehicleModel.Images,
			Location:    vehicleModel.Location,
			Contact:     vehicleModel.Contact,
			CreatedAt:   vehicleModel.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   vehicleModel.UpdatedAt.Format(time.RFC3339),
		}

		reservations = append(reservations, reservation)
	}

	return &vehicle.GetUserReservationListResponse{
		Code:         200,
		Message:      "获取成功",
		Reservations: reservations,
		Total:        total,
	}, nil
}
