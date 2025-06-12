package handler

import (
	"Common/services"
	"context"
	"fmt"
)

// GeocodeResponse 地理编码响应
type GeocodeResponse struct {
	Code      int     `json:"code"`
	Message   string  `json:"message"`
	Address   string  `json:"address,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

// GetCoordinates 根据地址获取经纬度
func GetCoordinates(ctx context.Context, address string) (*GeocodeResponse, error) {
	if address == "" {
		return &GeocodeResponse{
			Code:    400,
			Message: "地址不能为空",
		}, nil
	}

	// 调用高德API服务
	amapService := services.NewAmapService()
	coords, err := amapService.GetCoordinatesByAddress(address)
	if err != nil {
		return &GeocodeResponse{
			Code:    500,
			Message: fmt.Sprintf("获取坐标失败: %v", err),
		}, nil
	}

	return &GeocodeResponse{
		Code:      200,
		Message:   "获取坐标成功",
		Address:   address,
		Longitude: coords.Longitude,
		Latitude:  coords.Latitude,
	}, nil
}
