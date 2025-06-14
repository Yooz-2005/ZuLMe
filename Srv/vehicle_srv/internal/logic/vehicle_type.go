package logic

import (
	"ZuLMe/ZuLMe/Common/global"
	vehicle "ZuLMe/ZuLMe/Srv/vehicle_srv/proto_vehicle"
	"ZuLMe/ZuLMe/models/model_mysql"
	"context"
	"strings"
	"time"
)

// CreateVehicleType 创建车辆类型
func CreateVehicleType(ctx context.Context, in *vehicle.CreateVehicleTypeRequest) (*vehicle.CreateVehicleTypeResponse, error) {
	// 参数验证
	if strings.TrimSpace(in.Name) == "" {
		return &vehicle.CreateVehicleTypeResponse{Code: 400, Message: "类型名称不能为空"}, nil
	}

	// 检查名称是否已存在
	var existingType model_mysql.VehicleType
	if err := global.DB.Where("name = ?", strings.TrimSpace(in.Name)).First(&existingType).Error; err == nil {
		return &vehicle.CreateVehicleTypeResponse{Code: 400, Message: "类型名称已存在"}, nil
	}

	// 创建车辆类型实例
	newVehicleType := model_mysql.VehicleType{
		Name:        strings.TrimSpace(in.Name),
		Description: strings.TrimSpace(in.Description),
		Status:      int(in.Status),
		Sort:        int(in.Sort),
	}

	// 设置默认状态
	if newVehicleType.Status == 0 {
		newVehicleType.Status = 1 // 默认启用
	}

	// 调用模型的Create方法
	if err := newVehicleType.Create(); err != nil {
		return &vehicle.CreateVehicleTypeResponse{Code: 500, Message: "车辆类型创建失败"}, err
	}

	// 转换为响应格式
	vehicleTypeInfo := &vehicle.VehicleTypeInfo{
		Id:          int64(newVehicleType.ID),
		Name:        newVehicleType.Name,
		Description: newVehicleType.Description,
		Status:      int64(newVehicleType.Status),
		Sort:        int64(newVehicleType.Sort),
		CreatedAt:   newVehicleType.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newVehicleType.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.CreateVehicleTypeResponse{
		Code:        200,
		Message:     "车辆类型创建成功",
		VehicleType: vehicleTypeInfo,
	}, nil
}

// UpdateVehicleType 更新车辆类型
func UpdateVehicleType(ctx context.Context, in *vehicle.UpdateVehicleTypeRequest) (*vehicle.UpdateVehicleTypeResponse, error) {
	// 参数验证
	if in.Id <= 0 {
		return &vehicle.UpdateVehicleTypeResponse{Code: 400, Message: "车辆类型ID不能为空"}, nil
	}

	if strings.TrimSpace(in.Name) == "" {
		return &vehicle.UpdateVehicleTypeResponse{Code: 400, Message: "类型名称不能为空"}, nil
	}

	// 查找车辆类型
	var vehicleTypeData model_mysql.VehicleType
	if err := vehicleTypeData.GetByID(uint(in.Id)); err != nil {
		return &vehicle.UpdateVehicleTypeResponse{Code: 404, Message: "车辆类型不存在"}, nil
	}

	// 检查名称是否已被其他记录使用
	var existingType model_mysql.VehicleType
	if err := global.DB.Where("name = ? AND id != ?", strings.TrimSpace(in.Name), in.Id).First(&existingType).Error; err == nil {
		return &vehicle.UpdateVehicleTypeResponse{Code: 400, Message: "类型名称已存在"}, nil
	}

	// 更新字段
	vehicleTypeData.Name = strings.TrimSpace(in.Name)
	vehicleTypeData.Description = strings.TrimSpace(in.Description)
	vehicleTypeData.Status = int(in.Status)
	vehicleTypeData.Sort = int(in.Sort)

	// 调用模型的Update方法
	if err := vehicleTypeData.Update(); err != nil {
		return &vehicle.UpdateVehicleTypeResponse{Code: 500, Message: "车辆类型更新失败"}, err
	}

	// 转换为响应格式
	vehicleTypeInfo := &vehicle.VehicleTypeInfo{
		Id:          int64(vehicleTypeData.ID),
		Name:        vehicleTypeData.Name,
		Description: vehicleTypeData.Description,
		Status:      int64(vehicleTypeData.Status),
		Sort:        int64(vehicleTypeData.Sort),
		CreatedAt:   vehicleTypeData.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   vehicleTypeData.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.UpdateVehicleTypeResponse{
		Code:        200,
		Message:     "车辆类型更新成功",
		VehicleType: vehicleTypeInfo,
	}, nil
}

// DeleteVehicleType 删除车辆类型
func DeleteVehicleType(ctx context.Context, in *vehicle.DeleteVehicleTypeRequest) (*vehicle.DeleteVehicleTypeResponse, error) {
	// 参数验证
	if in.Id <= 0 {
		return &vehicle.DeleteVehicleTypeResponse{Code: 400, Message: "车辆类型ID不能为空"}, nil
	}

	// 查找车辆类型
	var vehicleTypeData model_mysql.VehicleType
	if err := vehicleTypeData.GetByID(uint(in.Id)); err != nil {
		return &vehicle.DeleteVehicleTypeResponse{Code: 404, Message: "车辆类型不存在"}, nil
	}

	// 检查是否有车辆使用此类型
	var vehicleCount int64
	if err := global.DB.Model(&model_mysql.Vehicle{}).Where("type_id = ?", in.Id).Count(&vehicleCount).Error; err != nil {
		return &vehicle.DeleteVehicleTypeResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if vehicleCount > 0 {
		return &vehicle.DeleteVehicleTypeResponse{Code: 400, Message: "该车辆类型下还有车辆，无法删除"}, nil
	}

	// 调用模型的Delete方法
	if err := vehicleTypeData.Delete(); err != nil {
		return &vehicle.DeleteVehicleTypeResponse{Code: 500, Message: "车辆类型删除失败"}, err
	}

	return &vehicle.DeleteVehicleTypeResponse{Code: 200, Message: "车辆类型删除成功"}, nil
}

// GetVehicleType 获取车辆类型详情
func GetVehicleType(ctx context.Context, in *vehicle.GetVehicleTypeRequest) (*vehicle.GetVehicleTypeResponse, error) {
	// 参数验证
	if in.Id <= 0 {
		return &vehicle.GetVehicleTypeResponse{Code: 400, Message: "车辆类型ID不能为空"}, nil
	}

	// 查找车辆类型
	var vehicleTypeData model_mysql.VehicleType
	if err := vehicleTypeData.GetByID(uint(in.Id)); err != nil {
		return &vehicle.GetVehicleTypeResponse{Code: 404, Message: "车辆类型不存在"}, nil
	}

	// 转换为响应格式
	vehicleTypeInfo := &vehicle.VehicleTypeInfo{
		Id:          int64(vehicleTypeData.ID),
		Name:        vehicleTypeData.Name,
		Description: vehicleTypeData.Description,
		Status:      int64(vehicleTypeData.Status),
		Sort:        int64(vehicleTypeData.Sort),
		CreatedAt:   vehicleTypeData.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   vehicleTypeData.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.GetVehicleTypeResponse{
		Code:        200,
		Message:     "获取车辆类型详情成功",
		VehicleType: vehicleTypeInfo,
	}, nil
}

// ListVehicleTypes 获取车辆类型列表
func ListVehicleTypes(ctx context.Context, in *vehicle.ListVehicleTypesRequest) (*vehicle.ListVehicleTypesResponse, error) {
	var vehicleTypes []model_mysql.VehicleType
	var total int64

	query := global.DB.Model(&model_mysql.VehicleType{})

	// 筛选条件
	if strings.TrimSpace(in.Keyword) != "" {
		keyword := "%" + strings.TrimSpace(in.Keyword) + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", keyword, keyword)
	}

	// 状态筛选：只有当明确指定status且不为-1时才筛选
	if in.Status != -1 {
		query = query.Where("status = ?", in.Status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return &vehicle.ListVehicleTypesResponse{Code: 500, Message: "数据库查询失败"}, err
	}

	// 分页
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大页面大小
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Order("sort ASC, created_at DESC").Find(&vehicleTypes).Error; err != nil {
		return &vehicle.ListVehicleTypesResponse{Code: 500, Message: "数据库查询失败"}, err
	}

	// 转换为响应格式
	var vehicleTypeInfos []*vehicle.VehicleTypeInfo
	for _, vt := range vehicleTypes {
		vehicleTypeInfo := &vehicle.VehicleTypeInfo{
			Id:          int64(vt.ID),
			Name:        vt.Name,
			Description: vt.Description,
			Status:      int64(vt.Status),
			Sort:        int64(vt.Sort),
			CreatedAt:   vt.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   vt.UpdatedAt.Format(time.RFC3339),
		}
		vehicleTypeInfos = append(vehicleTypeInfos, vehicleTypeInfo)
	}

	return &vehicle.ListVehicleTypesResponse{
		Code:         200,
		Message:      "获取车辆类型列表成功",
		VehicleTypes: vehicleTypeInfos,
		Total:        total,
	}, nil
}
