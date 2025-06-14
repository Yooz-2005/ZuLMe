package logic

import (
	vehicle "ZuLMe/ZuLMe/Srv/vehicle_srv/proto_vehicle"
	"ZuLMe/ZuLMe/models/model_mysql"
	"context"
	"fmt"
	"time"
)

// CreateVehicleBrand 创建车辆品牌
func CreateVehicleBrand(ctx context.Context, req *vehicle.CreateVehicleBrandRequest) (*vehicle.CreateVehicleBrandResponse, error) {
	// 检查品牌名称是否已存在
	brandModel := &model_mysql.VehicleBrand{}
	exists, err := brandModel.CheckNameExists(req.Name, 0)
	if err != nil {
		return &vehicle.CreateVehicleBrandResponse{
			Code:    500,
			Message: "检查品牌名称失败",
		}, err
	}
	if exists {
		return &vehicle.CreateVehicleBrandResponse{
			Code:    400,
			Message: "品牌名称已存在",
		}, nil
	}

	// 创建品牌
	brand := &model_mysql.VehicleBrand{
		Name:        req.Name,
		EnglishName: req.EnglishName,
		Logo:        req.Logo,
		Country:     req.Country,
		Description: req.Description,
		Status:      int(req.Status),
		Sort:        int(req.Sort),
		IsHot:       int(req.IsHot),
	}

	if err := brand.Create(); err != nil {
		return &vehicle.CreateVehicleBrandResponse{
			Code:    500,
			Message: "创建品牌失败",
		}, err
	}

	// 返回创建的品牌信息
	brandInfo := &vehicle.VehicleBrandInfo{
		Id:          int64(brand.ID),
		Name:        brand.Name,
		EnglishName: brand.EnglishName,
		Logo:        brand.Logo,
		Country:     brand.Country,
		Description: brand.Description,
		Status:      int64(brand.Status),
		Sort:        int64(brand.Sort),
		IsHot:       int64(brand.IsHot),
		CreatedAt:   brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   brand.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.CreateVehicleBrandResponse{
		Code:         200,
		Message:      "创建成功",
		VehicleBrand: brandInfo,
	}, nil
}

// UpdateVehicleBrand 更新车辆品牌
func UpdateVehicleBrand(ctx context.Context, req *vehicle.UpdateVehicleBrandRequest) (*vehicle.UpdateVehicleBrandResponse, error) {
	// 检查品牌是否存在
	brand := &model_mysql.VehicleBrand{}
	if err := brand.GetByID(uint(req.Id)); err != nil {
		return &vehicle.UpdateVehicleBrandResponse{
			Code:    404,
			Message: "品牌不存在",
		}, nil
	}

	// 检查品牌名称是否已被其他品牌使用
	exists, err := brand.CheckNameExists(req.Name, uint(req.Id))
	if err != nil {
		return &vehicle.UpdateVehicleBrandResponse{
			Code:    500,
			Message: "检查品牌名称失败",
		}, err
	}
	if exists {
		return &vehicle.UpdateVehicleBrandResponse{
			Code:    400,
			Message: "品牌名称已存在",
		}, nil
	}

	// 更新品牌信息
	brand.Name = req.Name
	brand.EnglishName = req.EnglishName
	brand.Logo = req.Logo
	brand.Country = req.Country
	brand.Description = req.Description
	brand.Status = int(req.Status)
	brand.Sort = int(req.Sort)
	brand.IsHot = int(req.IsHot)

	if err := brand.Update(); err != nil {
		return &vehicle.UpdateVehicleBrandResponse{
			Code:    500,
			Message: "更新品牌失败",
		}, err
	}

	// 如果品牌名称发生变化，同步更新车辆表中的品牌名称
	vehicleModel := &model_mysql.Vehicle{}
	if err := vehicleModel.UpdateBrandInfo(int64(brand.ID), brand.Name); err != nil {
		// 记录日志，但不影响品牌更新的成功
		fmt.Printf("同步更新车辆品牌信息失败: %v\n", err)
	}

	// 返回更新后的品牌信息
	brandInfo := &vehicle.VehicleBrandInfo{
		Id:          int64(brand.ID),
		Name:        brand.Name,
		EnglishName: brand.EnglishName,
		Logo:        brand.Logo,
		Country:     brand.Country,
		Description: brand.Description,
		Status:      int64(brand.Status),
		Sort:        int64(brand.Sort),
		IsHot:       int64(brand.IsHot),
		CreatedAt:   brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   brand.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.UpdateVehicleBrandResponse{
		Code:         200,
		Message:      "更新成功",
		VehicleBrand: brandInfo,
	}, nil
}

// DeleteVehicleBrand 删除车辆品牌
func DeleteVehicleBrand(ctx context.Context, req *vehicle.DeleteVehicleBrandRequest) (*vehicle.DeleteVehicleBrandResponse, error) {
	// 检查品牌是否存在
	brand := &model_mysql.VehicleBrand{}
	if err := brand.GetByID(uint(req.Id)); err != nil {
		return &vehicle.DeleteVehicleBrandResponse{
			Code:    404,
			Message: "品牌不存在",
		}, nil
	}

	// 检查是否有车辆使用该品牌
	count, err := brand.GetVehicleCountByBrand(uint(req.Id))
	if err != nil {
		return &vehicle.DeleteVehicleBrandResponse{
			Code:    500,
			Message: "检查品牌使用情况失败",
		}, err
	}
	if count > 0 {
		return &vehicle.DeleteVehicleBrandResponse{
			Code:    400,
			Message: fmt.Sprintf("该品牌下还有 %d 辆车辆，无法删除", count),
		}, nil
	}

	// 删除品牌
	if err := brand.Delete(); err != nil {
		return &vehicle.DeleteVehicleBrandResponse{
			Code:    500,
			Message: "删除品牌失败",
		}, err
	}

	return &vehicle.DeleteVehicleBrandResponse{
		Code:    200,
		Message: "删除成功",
	}, nil
}

// GetVehicleBrand 获取车辆品牌详情
func GetVehicleBrand(ctx context.Context, req *vehicle.GetVehicleBrandRequest) (*vehicle.GetVehicleBrandResponse, error) {
	brand := &model_mysql.VehicleBrand{}
	if err := brand.GetByID(uint(req.Id)); err != nil {
		return &vehicle.GetVehicleBrandResponse{
			Code:    404,
			Message: "品牌不存在",
		}, nil
	}

	brandInfo := &vehicle.VehicleBrandInfo{
		Id:          int64(brand.ID),
		Name:        brand.Name,
		EnglishName: brand.EnglishName,
		Logo:        brand.Logo,
		Country:     brand.Country,
		Description: brand.Description,
		Status:      int64(brand.Status),
		Sort:        int64(brand.Sort),
		IsHot:       int64(brand.IsHot),
		CreatedAt:   brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   brand.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.GetVehicleBrandResponse{
		Code:         200,
		Message:      "获取成功",
		VehicleBrand: brandInfo,
	}, nil
}

// ListVehicleBrands 获取车辆品牌列表
func ListVehicleBrands(ctx context.Context, req *vehicle.ListVehicleBrandsRequest) (*vehicle.ListVehicleBrandsResponse, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	brand := &model_mysql.VehicleBrand{}

	// 处理热门筛选参数
	var isHot *int
	if req.IsHot >= 0 {
		hotValue := int(req.IsHot)
		isHot = &hotValue
	}

	brands, total, err := brand.GetList(page, pageSize, isHot)
	if err != nil {
		return &vehicle.ListVehicleBrandsResponse{
			Code:    500,
			Message: "获取品牌列表失败",
		}, err
	}

	// 转换为proto格式
	var brandInfos []*vehicle.VehicleBrandInfo
	for _, b := range brands {
		brandInfo := &vehicle.VehicleBrandInfo{
			Id:          int64(b.ID),
			Name:        b.Name,
			EnglishName: b.EnglishName,
			Logo:        b.Logo,
			Country:     b.Country,
			Description: b.Description,
			Status:      int64(b.Status),
			Sort:        int64(b.Sort),
			IsHot:       int64(b.IsHot),
			CreatedAt:   b.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   b.UpdatedAt.Format(time.RFC3339),
		}
		brandInfos = append(brandInfos, brandInfo)
	}

	return &vehicle.ListVehicleBrandsResponse{
		Code:          200,
		Message:       "获取成功",
		VehicleBrands: brandInfos,
		Total:         total,
	}, nil
}
