package test

import (
	"context"
	"testing"
	"time"
	"vehicle_srv/internal/logic"
	vehicle "vehicle_srv/proto_vehicle"

	"github.com/stretchr/testify/assert"
)

// TestCheckVehicleAvailability 测试检查车辆可用性
func TestCheckVehicleAvailability(t *testing.T) {
	ctx := context.Background()

	// 测试正常情况
	req := &vehicle.CheckAvailabilityRequest{
		VehicleId: 1,
		StartDate: "2024-01-15",
		EndDate:   "2024-01-20",
	}

	resp, err := logic.CheckVehicleAvailability(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)

	// 测试无效的车辆ID
	req.VehicleId = 0
	resp, err = logic.CheckVehicleAvailability(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int64(400), resp.Code)
	assert.Contains(t, resp.Message, "车辆ID不能为空")

	// 测试无效的日期格式
	req.VehicleId = 1
	req.StartDate = "invalid-date"
	resp, err = logic.CheckVehicleAvailability(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int64(400), resp.Code)
	assert.Contains(t, resp.Message, "开始日期格式错误")
}

// TestCreateReservation 测试创建预订
func TestCreateReservation(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.CreateReservationRequest{
		VehicleId: 1,
		StartDate: "2024-01-25",
		EndDate:   "2024-01-30",
		OrderId:   1001,
		UserId:    1,
	}

	resp, err := logic.CreateReservation(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// 注意：实际测试中可能需要模拟数据库
}

// TestGetAvailableVehicles 测试获取可用车辆列表
func TestGetAvailableVehicles(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.GetAvailableVehiclesRequest{
		StartDate:  "2024-02-01",
		EndDate:    "2024-02-05",
		MerchantId: 1,
	}

	resp, err := logic.GetAvailableVehicles(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)
}

// TestSetVehicleMaintenance 测试设置车辆维护状态
func TestSetVehicleMaintenance(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.SetMaintenanceRequest{
		VehicleId: 1,
		StartDate: "2024-03-01",
		EndDate:   "2024-03-05",
		Notes:     "定期保养",
		CreatedBy: 1,
	}

	resp, err := logic.SetVehicleMaintenance(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// 实际测试中需要检查具体的响应码
}

// TestGetMaintenanceSchedule 测试获取维护计划
func TestGetMaintenanceSchedule(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.GetMaintenanceScheduleRequest{
		VehicleId: 1,
	}

	resp, err := logic.GetMaintenanceSchedule(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)
}

// TestGetInventoryCalendar 测试获取库存日历
func TestGetInventoryCalendar(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.GetInventoryCalendarRequest{
		VehicleId: 1,
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}

	resp, err := logic.GetInventoryCalendar(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)
	assert.NotNil(t, resp.Calendar)
}

// TestGetInventoryStatistics 测试获取库存统计
func TestGetInventoryStatistics(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.GetInventoryStatsRequest{
		MerchantId: 1,
	}

	resp, err := logic.GetInventoryStatistics(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)
}

// TestGetInventoryReport 测试获取库存报表
func TestGetInventoryReport(t *testing.T) {
	ctx := context.Background()

	req := &vehicle.GetInventoryReportRequest{
		MerchantId: 1,
		StartDate:  "2024-01-01",
		EndDate:    "2024-01-31",
	}

	resp, err := logic.GetInventoryReport(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(200), resp.Code)
}

// TestUpdateReservationStatus 测试更新预订状态
func TestUpdateReservationStatus(t *testing.T) {
	ctx := context.Background()

	// 测试更新为租用中
	req := &vehicle.UpdateReservationStatusRequest{
		OrderId: 1001,
		Status:  "rented",
	}

	resp, err := logic.UpdateReservationStatus(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// 测试完成租用
	req.Status = "completed"
	resp, err = logic.UpdateReservationStatus(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// 测试取消预订
	req.Status = "cancelled"
	resp, err = logic.UpdateReservationStatus(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// 测试无效状态
	req.Status = "invalid"
	resp, err = logic.UpdateReservationStatus(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int64(400), resp.Code)
	assert.Contains(t, resp.Message, "无效的状态")
}
