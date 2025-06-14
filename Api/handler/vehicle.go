package handler

import (
	"ZuLMe/ZuLMe/Api/client"
	vehicle "ZuLMe/ZuLMe/Srv/vehicle_srv/proto_vehicle"
	"context"
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

// ==================== 车辆品牌Handler ====================

// CreateVehicleBrand 创建车辆品牌
func CreateVehicleBrand(ctx context.Context, req *vehicle.CreateVehicleBrandRequest) (*vehicle.CreateVehicleBrandResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.CreateVehicleBrand(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.CreateVehicleBrandResponse), nil
}

// UpdateVehicleBrand 更新车辆品牌
func UpdateVehicleBrand(ctx context.Context, req *vehicle.UpdateVehicleBrandRequest) (*vehicle.UpdateVehicleBrandResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.UpdateVehicleBrand(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.UpdateVehicleBrandResponse), nil
}

// DeleteVehicleBrand 删除车辆品牌
func DeleteVehicleBrand(ctx context.Context, req *vehicle.DeleteVehicleBrandRequest) (*vehicle.DeleteVehicleBrandResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.DeleteVehicleBrand(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.DeleteVehicleBrandResponse), nil
}

// GetVehicleBrand 获取车辆品牌详情
func GetVehicleBrand(ctx context.Context, req *vehicle.GetVehicleBrandRequest) (*vehicle.GetVehicleBrandResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetVehicleBrand(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetVehicleBrandResponse), nil
}

// ListVehicleBrands 获取车辆品牌列表
func ListVehicleBrands(ctx context.Context, req *vehicle.ListVehicleBrandsRequest) (*vehicle.ListVehicleBrandsResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.ListVehicleBrands(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.ListVehicleBrandsResponse), nil
}

// ==================== 车辆库存Handler ====================

// CheckAvailability 检查车辆可用性
func CheckAvailability(ctx context.Context, req *vehicle.CheckAvailabilityRequest) (*vehicle.CheckAvailabilityResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.CheckAvailability(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.CheckAvailabilityResponse), nil
}

// CreateReservation 创建预订
func CreateReservation(ctx context.Context, req *vehicle.CreateReservationRequest) (*vehicle.CreateReservationResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.CreateReservation(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.CreateReservationResponse), nil
}

// UpdateReservationStatus 更新预订状态
func UpdateReservationStatus(ctx context.Context, req *vehicle.UpdateReservationStatusRequest) (*vehicle.UpdateReservationStatusResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.UpdateReservationStatus(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.UpdateReservationStatusResponse), nil
}

// GetAvailableVehicles 获取可用车辆
func GetAvailableVehicles(ctx context.Context, req *vehicle.GetAvailableVehiclesRequest) (*vehicle.GetAvailableVehiclesResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetAvailableVehicles(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetAvailableVehiclesResponse), nil
}

// GetInventoryStats 获取库存统计
func GetInventoryStats(ctx context.Context, req *vehicle.GetInventoryStatsRequest) (*vehicle.GetInventoryStatsResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetInventoryStats(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetInventoryStatsResponse), nil
}

// SetMaintenance 设置维护状态
func SetMaintenance(ctx context.Context, req *vehicle.SetMaintenanceRequest) (*vehicle.SetMaintenanceResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.SetMaintenance(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.SetMaintenanceResponse), nil
}

// GetMaintenanceSchedule 获取维护计划
func GetMaintenanceSchedule(ctx context.Context, req *vehicle.GetMaintenanceScheduleRequest) (*vehicle.GetMaintenanceScheduleResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetMaintenanceSchedule(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetMaintenanceScheduleResponse), nil
}

// GetInventoryCalendar 获取库存日历
func GetInventoryCalendar(ctx context.Context, req *vehicle.GetInventoryCalendarRequest) (*vehicle.GetInventoryCalendarResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetInventoryCalendar(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetInventoryCalendarResponse), nil
}

// GetInventoryReport 获取库存报表
func GetInventoryReport(ctx context.Context, req *vehicle.GetInventoryReportRequest) (*vehicle.GetInventoryReportResponse, error) {
	vehicleClient, err := client.VehicleClient(ctx, func(ctx context.Context, in vehicle.VehicleClient) (interface{}, error) {
		response, err := in.GetInventoryReport(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return vehicleClient.(*vehicle.GetInventoryReportResponse), nil
}
