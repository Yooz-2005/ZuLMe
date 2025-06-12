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
