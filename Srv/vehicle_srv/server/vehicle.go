package server

import (
	"context"
	"vehicle_srv/internal/logic"
	vehicle "vehicle_srv/proto_vehicle"
)

type ServerVehicle struct {
	vehicle.UnimplementedVehicleServer
}

// Ping 健康检查
func (s *ServerVehicle) Ping(ctx context.Context, in *vehicle.Request) (*vehicle.Response, error) {
	return &vehicle.Response{Pong: "pong"}, nil
}

// CreateVehicle 创建车辆
func (s *ServerVehicle) CreateVehicle(ctx context.Context, in *vehicle.CreateVehicleRequest) (*vehicle.CreateVehicleResponse, error) {
	return logic.CreateVehicle(ctx, in)
}

// UpdateVehicle 更新车辆
func (s *ServerVehicle) UpdateVehicle(ctx context.Context, in *vehicle.UpdateVehicleRequest) (*vehicle.UpdateVehicleResponse, error) {
	return logic.UpdateVehicle(ctx, in)
}

// DeleteVehicle 删除车辆
func (s *ServerVehicle) DeleteVehicle(ctx context.Context, in *vehicle.DeleteVehicleRequest) (*vehicle.DeleteVehicleResponse, error) {
	return logic.DeleteVehicle(ctx, in)
}

// GetVehicle 获取车辆详情
func (s *ServerVehicle) GetVehicle(ctx context.Context, in *vehicle.GetVehicleRequest) (*vehicle.GetVehicleResponse, error) {
	return logic.GetVehicle(ctx, in)
}

// ListVehicles 获取车辆列表
func (s *ServerVehicle) ListVehicles(ctx context.Context, in *vehicle.ListVehiclesRequest) (*vehicle.ListVehiclesResponse, error) {
	return logic.ListVehicles(ctx, in)
}

// ==================== 车辆类型gRPC方法 ====================

// CreateVehicleType 创建车辆类型
func (s *ServerVehicle) CreateVehicleType(ctx context.Context, in *vehicle.CreateVehicleTypeRequest) (*vehicle.CreateVehicleTypeResponse, error) {
	return logic.CreateVehicleType(ctx, in)
}

// UpdateVehicleType 更新车辆类型
func (s *ServerVehicle) UpdateVehicleType(ctx context.Context, in *vehicle.UpdateVehicleTypeRequest) (*vehicle.UpdateVehicleTypeResponse, error) {
	return logic.UpdateVehicleType(ctx, in)
}

// DeleteVehicleType 删除车辆类型
func (s *ServerVehicle) DeleteVehicleType(ctx context.Context, in *vehicle.DeleteVehicleTypeRequest) (*vehicle.DeleteVehicleTypeResponse, error) {
	return logic.DeleteVehicleType(ctx, in)
}

// GetVehicleType 获取车辆类型详情
func (s *ServerVehicle) GetVehicleType(ctx context.Context, in *vehicle.GetVehicleTypeRequest) (*vehicle.GetVehicleTypeResponse, error) {
	return logic.GetVehicleType(ctx, in)
}

// ListVehicleTypes 获取车辆类型列表
func (s *ServerVehicle) ListVehicleTypes(ctx context.Context, in *vehicle.ListVehicleTypesRequest) (*vehicle.ListVehicleTypesResponse, error) {
	return logic.ListVehicleTypes(ctx, in)
}

// ==================== 车辆品牌gRPC方法 ====================

// CreateVehicleBrand 创建车辆品牌
func (s *ServerVehicle) CreateVehicleBrand(ctx context.Context, in *vehicle.CreateVehicleBrandRequest) (*vehicle.CreateVehicleBrandResponse, error) {
	return logic.CreateVehicleBrand(ctx, in)
}

// UpdateVehicleBrand 更新车辆品牌
func (s *ServerVehicle) UpdateVehicleBrand(ctx context.Context, in *vehicle.UpdateVehicleBrandRequest) (*vehicle.UpdateVehicleBrandResponse, error) {
	return logic.UpdateVehicleBrand(ctx, in)
}

// DeleteVehicleBrand 删除车辆品牌
func (s *ServerVehicle) DeleteVehicleBrand(ctx context.Context, in *vehicle.DeleteVehicleBrandRequest) (*vehicle.DeleteVehicleBrandResponse, error) {
	return logic.DeleteVehicleBrand(ctx, in)
}

// GetVehicleBrand 获取车辆品牌详情
func (s *ServerVehicle) GetVehicleBrand(ctx context.Context, in *vehicle.GetVehicleBrandRequest) (*vehicle.GetVehicleBrandResponse, error) {
	return logic.GetVehicleBrand(ctx, in)
}

// ListVehicleBrands 获取车辆品牌列表
func (s *ServerVehicle) ListVehicleBrands(ctx context.Context, in *vehicle.ListVehicleBrandsRequest) (*vehicle.ListVehicleBrandsResponse, error) {
	return logic.ListVehicleBrands(ctx, in)
}

// ==================== 车辆库存管理gRPC方法 ====================

// CheckAvailability 检查车辆可用性
func (s *ServerVehicle) CheckAvailability(ctx context.Context, in *vehicle.CheckAvailabilityRequest) (*vehicle.CheckAvailabilityResponse, error) {
	return logic.CheckVehicleAvailability(ctx, in)
}

// CreateReservation 创建预订
func (s *ServerVehicle) CreateReservation(ctx context.Context, in *vehicle.CreateReservationRequest) (*vehicle.CreateReservationResponse, error) {
	return logic.CreateReservation(ctx, in)
}

// UpdateReservationStatus 更新预订状态
func (s *ServerVehicle) UpdateReservationStatus(ctx context.Context, in *vehicle.UpdateReservationStatusRequest) (*vehicle.UpdateReservationStatusResponse, error) {
	return logic.UpdateReservationStatus(ctx, in)
}

// CancelReservation 取消预订
func (s *ServerVehicle) CancelReservation(ctx context.Context, in *vehicle.CancelReservationRequest) (*vehicle.CancelReservationResponse, error) {
	return logic.CancelReservation(ctx, in)
}

// GetAvailableVehicles 获取可用车辆列表
func (s *ServerVehicle) GetAvailableVehicles(ctx context.Context, in *vehicle.GetAvailableVehiclesRequest) (*vehicle.GetAvailableVehiclesResponse, error) {
	return logic.GetAvailableVehicles(ctx, in)
}

// GetUserReservationList 获取用户预订列表
func (s *ServerVehicle) GetUserReservationList(ctx context.Context, in *vehicle.GetUserReservationListRequest) (*vehicle.GetUserReservationListResponse, error) {
	return logic.GetUserReservationList(ctx, in)
}

// GetInventoryStats 获取库存统计
func (s *ServerVehicle) GetInventoryStats(ctx context.Context, in *vehicle.GetInventoryStatsRequest) (*vehicle.GetInventoryStatsResponse, error) {
	return logic.GetInventoryStatistics(ctx, in)
}

// SetMaintenance 设置维护状态
func (s *ServerVehicle) SetMaintenance(ctx context.Context, in *vehicle.SetMaintenanceRequest) (*vehicle.SetMaintenanceResponse, error) {
	return logic.SetVehicleMaintenance(ctx, in)
}

// GetMaintenanceSchedule 获取维护计划
func (s *ServerVehicle) GetMaintenanceSchedule(ctx context.Context, in *vehicle.GetMaintenanceScheduleRequest) (*vehicle.GetMaintenanceScheduleResponse, error) {
	return logic.GetMaintenanceSchedule(ctx, in)
}

// GetInventoryCalendar 获取库存日历
func (s *ServerVehicle) GetInventoryCalendar(ctx context.Context, in *vehicle.GetInventoryCalendarRequest) (*vehicle.GetInventoryCalendarResponse, error) {
	return logic.GetInventoryCalendar(ctx, in)
}

// GetInventoryReport 获取库存报表
func (s *ServerVehicle) GetInventoryReport(ctx context.Context, in *vehicle.GetInventoryReportRequest) (*vehicle.GetInventoryReportResponse, error) {
	return logic.GetInventoryReport(ctx, in)
}

// 批量操作方法暂时注释，等proto文件正确生成后再启用
// BatchCreateReservations 批量创建预订
// func (s *ServerVehicle) BatchCreateReservations(ctx context.Context, in *vehicle.BatchCreateReservationsRequest) (*vehicle.BatchCreateReservationsResponse, error) {
// 	return logic.BatchCreateReservations(ctx, in)
// }

// BatchCancelReservations 批量取消预订
// func (s *ServerVehicle) BatchCancelReservations(ctx context.Context, in *vehicle.BatchCancelReservationsRequest) (*vehicle.BatchCancelReservationsResponse, error) {
// 	return logic.BatchCancelReservations(ctx, in)
// }
