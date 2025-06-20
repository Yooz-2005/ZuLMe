package logic

import (
	"Common/global"
	"context"
	"models/model_mysql"
	"strings"
	"time"
	vehicle "vehicle_srv/proto_vehicle"
)

// CreateVehicle 创建车辆
func CreateVehicle(ctx context.Context, in *vehicle.CreateVehicleRequest) (*vehicle.CreateVehicleResponse, error) {
	// 参数验证
	if in.TypeId <= 0 {
		return &vehicle.CreateVehicleResponse{Code: 400, Message: "车辆类型ID不能为空"}, nil
	}
	if in.BrandId <= 0 {
		return &vehicle.CreateVehicleResponse{Code: 400, Message: "品牌ID不能为空"}, nil
	}
	if strings.TrimSpace(in.Style) == "" {
		return &vehicle.CreateVehicleResponse{Code: 400, Message: "型号不能为空"}, nil
	}
	if in.Year <= 0 {
		return &vehicle.CreateVehicleResponse{Code: 400, Message: "年份不能为空"}, nil
	}
	if in.Price <= 0 {
		return &vehicle.CreateVehicleResponse{Code: 400, Message: "价格必须大于0"}, nil
	}

	// 验证车辆类型是否存在
	var vehicleType model_mysql.VehicleType
	if err := global.DB.Where("id = ?", in.TypeId).Limit(1).Find(&vehicleType).Error; err != nil {
		return &vehicle.CreateVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if vehicleType.ID == 0 {
		return &vehicle.CreateVehicleResponse{Code: 404, Message: "车辆类型不存在"}, nil
	}

	// 验证品牌是否存在
	var vehicleBrand model_mysql.VehicleBrand
	if err := vehicleBrand.GetByID(uint(in.BrandId)); err != nil {
		return &vehicle.CreateVehicleResponse{Code: 404, Message: "品牌不存在"}, nil
	}

	// 验证商家是否存在
	var merchant model_mysql.Merchant
	if err := global.DB.Where("id =?", in.MerchantId).Limit(1).Find(&merchant).Error; err != nil {
		return &vehicle.CreateVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if merchant.ID == 0 {
		return &vehicle.CreateVehicleResponse{Code: 404, Message: "商家不存在"}, nil
	}

	// 创建车辆实例
	newVehicle := model_mysql.Vehicle{
		MerchantID:  in.MerchantId,
		TypeID:      in.TypeId,
		BrandID:     in.BrandId,
		Brand:       vehicleBrand.Name, // 从品牌表获取品牌名称
		Style:       strings.TrimSpace(in.Style),
		Year:        in.Year,
		Color:       strings.TrimSpace(in.Color),
		Mileage:     in.Mileage,
		Price:       in.Price,
		Status:      in.Status,
		Description: strings.TrimSpace(in.Description),
		Images:      strings.TrimSpace(in.Images),
		Location:    strings.TrimSpace(merchant.Location),
		Contact:     strings.TrimSpace(in.Contact),
	}

	// 调用模型的Create方法
	if err := newVehicle.Create(); err != nil {
		return &vehicle.CreateVehicleResponse{Code: 500, Message: "车辆创建失败"}, err
	}

	// 转换为响应格式
	vehicleInfo := &vehicle.VehicleInfo{
		Id:          int64(newVehicle.ID),
		MerchantId:  newVehicle.MerchantID,
		TypeId:      newVehicle.TypeID,
		BrandId:     newVehicle.BrandID,
		Brand:       newVehicle.Brand,
		Style:       newVehicle.Style,
		Year:        newVehicle.Year,
		Color:       newVehicle.Color,
		Mileage:     newVehicle.Mileage,
		Price:       newVehicle.Price,
		Status:      newVehicle.Status,
		Description: newVehicle.Description,
		Images:      newVehicle.Images,
		Location:    newVehicle.Location,
		Contact:     newVehicle.Contact,
		CreatedAt:   newVehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newVehicle.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.CreateVehicleResponse{
		Code:    200,
		Message: "车辆创建成功",
		Vehicle: vehicleInfo,
	}, nil
}

// UpdateVehicle 更新车辆
func UpdateVehicle(ctx context.Context, in *vehicle.UpdateVehicleRequest) (*vehicle.UpdateVehicleResponse, error) {
	if in.Id <= 0 {
		return &vehicle.UpdateVehicleResponse{Code: 400, Message: "车辆ID不能为空"}, nil
	}

	// 查找车辆
	var existingVehicle model_mysql.Vehicle
	if err := existingVehicle.GetByID(uint(in.Id)); err != nil {
		return &vehicle.UpdateVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if existingVehicle.ID == 0 {
		return &vehicle.UpdateVehicleResponse{Code: 404, Message: "车辆不存在"}, nil
	}

	// 验证车辆类型
	if in.TypeId > 0 {
		var vehicleType model_mysql.VehicleType
		if err := global.DB.Where("id = ?", in.TypeId).Limit(1).Find(&vehicleType).Error; err != nil {
			return &vehicle.UpdateVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
		}
		if vehicleType.ID == 0 {
			return &vehicle.UpdateVehicleResponse{Code: 404, Message: "车辆类型不存在"}, nil
		}
		existingVehicle.TypeID = in.TypeId
	}

	// 更新商家ID（中间件已验证商家身份）
	if in.MerchantId > 0 {
		existingVehicle.MerchantID = in.MerchantId
	}

	// 更新其他字段（只更新非空字段）
	// 暂时注释品牌更新，等proto重新生成后启用
	// if strings.TrimSpace(in.Brand) != "" {
	//	existingVehicle.Brand = strings.TrimSpace(in.Brand)
	// }
	if strings.TrimSpace(in.Style) != "" {
		existingVehicle.Style = strings.TrimSpace(in.Style)
	}
	if in.Year > 0 {
		existingVehicle.Year = in.Year
	}
	if strings.TrimSpace(in.Color) != "" {
		existingVehicle.Color = strings.TrimSpace(in.Color)
	}
	if in.Mileage >= 0 {
		existingVehicle.Mileage = in.Mileage
	}
	if in.Price > 0 {
		existingVehicle.Price = in.Price
	}
	if in.Status >= 0 {
		existingVehicle.Status = in.Status
	}
	if strings.TrimSpace(in.Description) != "" {
		existingVehicle.Description = strings.TrimSpace(in.Description)
	}
	if strings.TrimSpace(in.Images) != "" {
		existingVehicle.Images = strings.TrimSpace(in.Images)
	}
	if strings.TrimSpace(in.Location) != "" {
		existingVehicle.Location = strings.TrimSpace(in.Location)
	}
	if strings.TrimSpace(in.Contact) != "" {
		existingVehicle.Contact = strings.TrimSpace(in.Contact)
	}

	// 调用模型的Update方法
	if err := existingVehicle.Update(); err != nil {
		return &vehicle.UpdateVehicleResponse{Code: 500, Message: "车辆更新失败"}, err
	}

	// 转换为响应格式
	vehicleInfo := &vehicle.VehicleInfo{
		Id:          int64(existingVehicle.ID),
		MerchantId:  existingVehicle.MerchantID,
		TypeId:      existingVehicle.TypeID,
		Brand:       existingVehicle.Brand,
		Style:       existingVehicle.Style,
		Year:        existingVehicle.Year,
		Color:       existingVehicle.Color,
		Mileage:     existingVehicle.Mileage,
		Price:       existingVehicle.Price,
		Status:      existingVehicle.Status,
		Description: existingVehicle.Description,
		Images:      existingVehicle.Images,
		Location:    existingVehicle.Location,
		Contact:     existingVehicle.Contact,
		CreatedAt:   existingVehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   existingVehicle.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.UpdateVehicleResponse{
		Code:    200,
		Message: "车辆更新成功",
		Vehicle: vehicleInfo,
	}, nil
}

// DeleteVehicle 删除车辆
func DeleteVehicle(ctx context.Context, in *vehicle.DeleteVehicleRequest) (*vehicle.DeleteVehicleResponse, error) {
	if in.Id <= 0 {
		return &vehicle.DeleteVehicleResponse{Code: 400, Message: "车辆ID不能为空"}, nil
	}

	// 先查找车辆是否存在
	var vehicleData model_mysql.Vehicle
	if err := vehicleData.GetByID(uint(in.Id)); err != nil {
		return &vehicle.DeleteVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if vehicleData.ID == 0 {
		return &vehicle.DeleteVehicleResponse{Code: 404, Message: "车辆不存在"}, nil
	}

	// 调用模型的Delete方法
	if err := vehicleData.Delete(); err != nil {
		return &vehicle.DeleteVehicleResponse{Code: 500, Message: "车辆删除失败"}, err
	}

	return &vehicle.DeleteVehicleResponse{Code: 200, Message: "车辆删除成功"}, nil
}

// GetVehicle 获取车辆详情
func GetVehicle(ctx context.Context, in *vehicle.GetVehicleRequest) (*vehicle.GetVehicleResponse, error) {
	if in.Id <= 0 {
		return &vehicle.GetVehicleResponse{Code: 400, Message: "车辆ID不能为空"}, nil
	}

	var vehicleData model_mysql.Vehicle
	if err := vehicleData.GetByID(uint(in.Id)); err != nil {
		return &vehicle.GetVehicleResponse{Code: 500, Message: "数据库查询失败"}, err
	}
	if vehicleData.ID == 0 {
		return &vehicle.GetVehicleResponse{Code: 404, Message: "车辆不存在"}, nil
	}

	// 转换为响应格式
	vehicleInfo := &vehicle.VehicleInfo{
		Id:          int64(vehicleData.ID),
		MerchantId:  vehicleData.MerchantID,
		TypeId:      vehicleData.TypeID,
		Brand:       vehicleData.Brand,
		Style:       vehicleData.Style,
		Year:        vehicleData.Year,
		Color:       vehicleData.Color,
		Mileage:     vehicleData.Mileage,
		Price:       vehicleData.Price,
		Status:      vehicleData.Status,
		Description: vehicleData.Description,
		Images:      vehicleData.Images,
		Location:    vehicleData.Location,
		Contact:     vehicleData.Contact,
		CreatedAt:   vehicleData.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   vehicleData.UpdatedAt.Format(time.RFC3339),
	}

	return &vehicle.GetVehicleResponse{
		Code:    200,
		Message: "获取车辆详情成功",
		Vehicle: vehicleInfo,
	}, nil
}

// ListVehicles 获取车辆列表
func ListVehicles(ctx context.Context, in *vehicle.ListVehiclesRequest) (*vehicle.ListVehiclesResponse, error) {
	var vehicles []model_mysql.Vehicle
	var total int64

	query := global.DB.Model(&model_mysql.Vehicle{})

	// 筛选条件
	if strings.TrimSpace(in.Keyword) != "" {
		keyword := "%" + strings.TrimSpace(in.Keyword) + "%"
		query = query.Where("brand LIKE ? OR style LIKE ? OR description LIKE ? OR location LIKE ?",
			keyword, keyword, keyword, keyword)
	}

	if in.MerchantId > 0 {
		query = query.Where("merchant_id = ?", in.MerchantId)
	}

	if in.TypeId > 0 {
		query = query.Where("type_id = ?", in.TypeId)
	}

	// 只有当明确指定status且不为-1时才筛选
	// -1表示不筛选状态，0和1是有效的状态值
	if in.Status != -1 {
		query = query.Where("status = ?", in.Status)
	}

	// 价格范围筛选
	if in.PriceMin > 0 {
		query = query.Where("price >= ?", in.PriceMin)
	}
	if in.PriceMax > 0 {
		query = query.Where("price <= ?", in.PriceMax)
	}

	// 年份范围筛选
	if in.YearMin > 0 {
		query = query.Where("year >= ?", in.YearMin)
	}
	if in.YearMax > 0 {
		query = query.Where("year <= ?", in.YearMax)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return &vehicle.ListVehiclesResponse{Code: 500, Message: "数据库查询失败"}, err
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
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Order("created_at DESC").Find(&vehicles).Error; err != nil {
		return &vehicle.ListVehiclesResponse{Code: 500, Message: "数据库查询失败"}, err
	}

	// 转换为响应格式
	var vehicleInfos []*vehicle.VehicleInfo
	for _, v := range vehicles {
		vehicleInfo := &vehicle.VehicleInfo{
			Id:          int64(v.ID),
			MerchantId:  v.MerchantID,
			TypeId:      v.TypeID,
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
		vehicleInfos = append(vehicleInfos, vehicleInfo)
	}

	return &vehicle.ListVehiclesResponse{
		Code:     200,
		Message:  "获取车辆列表成功",
		Vehicles: vehicleInfos,
		Total:    total,
	}, nil
}
