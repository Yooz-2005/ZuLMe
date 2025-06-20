package logic

import (
	"ZuLMe/ZuLMe/Common/global"
	"ZuLMe/ZuLMe/Common/services"
	"ZuLMe/ZuLMe/models/model_mysql"
	admin "admin_srv/proto_admin"
	"context"
	"errors"
	"fmt"

	"time"

	"gorm.io/gorm"
)

// MerchantApprove 审核商户
func MerchantApprove(ctx context.Context, in *admin.MerchantApproveRequest) (*admin.MerchantApproveResponse, error) {
	var merchant model_mysql.Merchant
	result := global.DB.First(&merchant, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &admin.MerchantApproveResponse{Code: 404, Message: "商户不存在"}, nil
		}
		return &admin.MerchantApproveResponse{Code: 500, Message: "数据库查询失败"}, result.Error
	}

	merchant.Status = int(in.Status) // 0-未审核，1-审核通过，2-审核失败
	result = global.DB.Save(&merchant)
	if result.Error != nil {
		return &admin.MerchantApproveResponse{Code: 500, Message: "商户审核状态更新失败"}, result.Error
	}

	return &admin.MerchantApproveResponse{Code: 200, Message: "商户审核成功"}, nil
}

// MerchantUpdate 编辑商户
func MerchantUpdate(ctx context.Context, in *admin.MerchantUpdateRequest) (*admin.MerchantUpdateResponse, error) {
	var merchant model_mysql.Merchant
	result := global.DB.First(&merchant, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &admin.MerchantUpdateResponse{Code: 404, Message: "商户不存在"}, nil
		}
		return &admin.MerchantUpdateResponse{Code: 500, Message: "数据库查询失败"}, result.Error
	}

	if in.Name != "" {
		merchant.Name = in.Name
	}
	if in.Phone != "" {
		merchant.Phone = in.Phone
	}
	if in.Email != "" {
		merchant.Email = in.Email
	}
	if in.Location != "" {
		merchant.Location = in.Location

		// 如果地址发生变化，重新获取经纬度
		if in.Longitude == 0 && in.Latitude == 0 {
			// 调用高德API获取经纬度
			amapService := services.NewAmapService()
			coords, err := amapService.GetCoordinatesByAddress(in.Location)
			if err != nil {
				fmt.Printf("更新商家时获取坐标失败: %v, 地址: %s\n", err, in.Location)
				// 坐标获取失败不影响更新，保持原有坐标
			} else {
				merchant.Longitude = coords.Longitude
				merchant.Latitude = coords.Latitude
				fmt.Printf("更新商家坐标成功: 经度=%f, 纬度=%f, 地址=%s\n", coords.Longitude, coords.Latitude, in.Location)
			}
		} else {
			// 使用提供的经纬度
			merchant.Longitude = in.Longitude
			merchant.Latitude = in.Latitude
		}
	}
	if in.BusinessTime != "" {
		merchant.BusinessTime = in.BusinessTime
	}

	result = global.DB.Save(&merchant)
	if result.Error != nil {
		return &admin.MerchantUpdateResponse{Code: 500, Message: "商户信息更新失败"}, result.Error
	}

	return &admin.MerchantUpdateResponse{Code: 200, Message: "商户信息更新成功"}, nil
}

// MerchantDelete 删除商户
func MerchantDelete(ctx context.Context, in *admin.MerchantDeleteRequest) (*admin.MerchantDeleteResponse, error) {
	result := global.DB.Delete(&model_mysql.Merchant{}, in.Id)
	if result.Error != nil {
		return &admin.MerchantDeleteResponse{Code: 500, Message: "商户删除失败"}, result.Error
	}

	if result.RowsAffected == 0 {
		return &admin.MerchantDeleteResponse{Code: 404, Message: "商户不存在或已被删除"}, nil
	}

	return &admin.MerchantDeleteResponse{Code: 200, Message: "商户删除成功"}, nil
}

// MerchantList 获取商户列表
func MerchantList(ctx context.Context, in *admin.MerchantListRequest) (*admin.MerchantListResponse, error) {
	var merchants []model_mysql.Merchant
	var total int64

	query := global.DB.Model(&model_mysql.Merchant{})

	// 筛选条件
	if in.Keyword != "" {
		query = query.Where("name LIKE ? OR phone LIKE ? OR email LIKE ?", "%"+in.Keyword+"%", "%"+in.Keyword+"%", "%"+in.Keyword+"%")
	}
	if in.StatusFilter >= 0 { // 0, 1, 2 分别代表不同状态，-1表示不筛选状态
		query = query.Where("status = ?", in.StatusFilter)
	}

	// 获取总数
	query.Count(&total)

	// 分页
	page := in.Page
	pageSize := in.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	result := query.Offset(int(offset)).Limit(int(pageSize)).Find(&merchants)
	if result.Error != nil {
		return &admin.MerchantListResponse{Code: 500, Message: "获取商户列表失败"}, result.Error
	}

	// 转换为响应结构
	var merchantInfos []*admin.MerchantInfo
	for _, m := range merchants {
		merchantInfos = append(merchantInfos, &admin.MerchantInfo{
			Id:           int64(m.ID),
			Name:         m.Name,
			Phone:        m.Phone,
			Email:        m.Email,
			Status:       int64(m.Status), // 映射Status字段
			CreatedAt:    m.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    m.UpdatedAt.Format(time.RFC3339),
			Location:     m.Location,
			BusinessTime: m.BusinessTime,
			Longitude:    m.Longitude,
			Latitude:     m.Latitude,
		})
	}

	return &admin.MerchantListResponse{
		Code:      200,
		Message:   "获取商户列表成功",
		Merchants: merchantInfos,
		Total:     total,
	}, nil
}

// MerchantDetail 获取商户详情
func MerchantDetail(ctx context.Context, in *admin.MerchantDetailRequest) (*admin.MerchantDetailResponse, error) {
	var merchant model_mysql.Merchant
	result := global.DB.First(&merchant, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &admin.MerchantDetailResponse{Code: 404, Message: "商户不存在"}, nil
		}
		return &admin.MerchantDetailResponse{Code: 500, Message: "数据库查询失败"}, result.Error
	}

	// 将数据库模型转换为响应结构
	merchantInfo := &admin.MerchantInfo{
		Id:           int64(merchant.ID),
		Name:         merchant.Name,
		Phone:        merchant.Phone,
		Email:        merchant.Email,
		Status:       int64(merchant.Status), // 映射Status字段
		CreatedAt:    merchant.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    merchant.UpdatedAt.Format(time.RFC3339),
		Location:     merchant.Location,
		BusinessTime: merchant.BusinessTime,
		Longitude:    merchant.Longitude,
		Latitude:     merchant.Latitude,
	}

	return &admin.MerchantDetailResponse{
		Code:     200,
		Message:  "获取商户详情成功",
		Merchant: merchantInfo,
	}, nil
}
