package logic

import (
	"Common/global"
	"context"
	"time"
	merchant "merchant_srv/proto_merchant"
	"models/model_mysql"
)

// GetLocationList 获取网点列表
func GetLocationList(ctx context.Context, in *merchant.GetLocationListRequest) (*merchant.GetLocationListResponse, error) {
	var merchants []model_mysql.Merchant
	var total int64

	query := global.DB.Model(&model_mysql.Merchant{})

	// 筛选审核状态
	if in.StatusFilter >= 0 {
		query = query.Where("status = ?", in.StatusFilter)
	}

	// 获取总数
	query.Count(&total)

	// 分页参数处理
	page := int(in.Page)
	pageSize := int(in.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	// 分页查询
	offset := (page - 1) * pageSize
	result := query.Offset(offset).Limit(pageSize).Find(&merchants)
	if result.Error != nil {
		return &merchant.GetLocationListResponse{
			Code:    500,
			Message: "获取商家网点列表失败",
		}, result.Error
	}

	// 转换为响应格式
	var locationList []*merchant.LocationInfo
	for _, m := range merchants {
		locationList = append(locationList, &merchant.LocationInfo{
			Id:           int64(m.ID),
			Name:         m.Name,
			Phone:        m.Phone,
			Email:        m.Email,
			Status:       int32(m.Status),
			Location:     m.Location,
			BusinessTime: m.BusinessTime,
			Longitude:    m.Longitude,
			Latitude:     m.Latitude,
			CreatedAt:    m.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    m.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &merchant.GetLocationListResponse{
		Code:      200,
		Message:   "获取商家网点列表成功",
		Locations: locationList,
		Total:     total,
	}, nil
}
