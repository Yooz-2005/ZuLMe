package handler

import (
	"Api/client"
	"context"
	vehicle "vehicle_srv/proto_vehicle"
)

// CreateVehicle 创建车辆
func CreateVehicle(ctx context.Context, req *vehicle.CreateVehicleRequest) (*vehicle.CreateVehicleResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.CreateVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.CreateVehicleResponse), nil
}

// UpdateVehicle 更新车辆
func UpdateVehicle(ctx context.Context, req *vehicle.UpdateVehicleRequest) (*vehicle.UpdateVehicleResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.UpdateVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.UpdateVehicleResponse), nil
}

// DeleteVehicle 删除车辆
func DeleteVehicle(ctx context.Context, req *vehicle.DeleteVehicleRequest) (*vehicle.DeleteVehicleResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.DeleteVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.DeleteVehicleResponse), nil
}

// GetVehicle 获取车辆详情
func GetVehicle(ctx context.Context, req *vehicle.GetVehicleRequest) (*vehicle.GetVehicleResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetVehicleResponse), nil
}

// ListVehicles 获取车辆列表
func ListVehicles(ctx context.Context, req *vehicle.ListVehiclesRequest) (*vehicle.ListVehiclesResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.ListVehicles(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.ListVehiclesResponse), nil
}

// ==================== 车辆类型Handler ====================

// CreateVehicleType 创建车辆类型
func CreateVehicleType(ctx context.Context, req *vehicle.CreateVehicleTypeRequest) (*vehicle.CreateVehicleTypeResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.CreateVehicleType(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.CreateVehicleTypeResponse), nil
}

// UpdateVehicleType 更新车辆类型
func UpdateVehicleType(ctx context.Context, req *vehicle.UpdateVehicleTypeRequest) (*vehicle.UpdateVehicleTypeResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.UpdateVehicleType(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.UpdateVehicleTypeResponse), nil
}

// DeleteVehicleType 删除车辆类型
func DeleteVehicleType(ctx context.Context, req *vehicle.DeleteVehicleTypeRequest) (*vehicle.DeleteVehicleTypeResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.DeleteVehicleType(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.DeleteVehicleTypeResponse), nil
}

// GetVehicleType 获取车辆类型详情
func GetVehicleType(ctx context.Context, req *vehicle.GetVehicleTypeRequest) (*vehicle.GetVehicleTypeResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetVehicleType(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetVehicleTypeResponse), nil
}

// ListVehicleTypes 获取车辆类型列表
func ListVehicleTypes(ctx context.Context, req *vehicle.ListVehicleTypesRequest) (*vehicle.ListVehicleTypesResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.ListVehicleTypes(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.ListVehicleTypesResponse), nil
}
