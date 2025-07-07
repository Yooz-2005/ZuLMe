package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AmapConfig 高德地图配置
type AmapConfig struct {
	APIKey  string
	BaseURL string
}

// AmapService 高德地图服务
type AmapService struct {
	config AmapConfig
	client *http.Client
}

// GeocodeResponse 地理编码响应结构（简化版，只保留必要字段）
type GeocodeResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
	Count    string `json:"count"`
	Geocodes []struct {
		FormattedAddress string `json:"formatted_address"`
		Location         string `json:"location"`
		Level            string `json:"level"`
	} `json:"geocodes"`
}

// Coordinates 坐标结构
type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// NewAmapService 创建高德地图服务实例
func NewAmapService() *AmapService {
	return &AmapService{
		config: AmapConfig{
			APIKey:  "9a65114fc68c95e63f126827e26b026a",
			BaseURL: "https://restapi.amap.com/v3/geocode/geo",
		},
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetCoordinatesByAddress 根据地址获取经纬度
func (s *AmapService) GetCoordinatesByAddress(address string) (*Coordinates, error) {
	if address == "" {
		return nil, fmt.Errorf("地址不能为空")
	}

	// 构建请求URL
	params := url.Values{}
	params.Add("key", s.config.APIKey)
	params.Add("address", address)
	params.Add("output", "json")

	requestURL := fmt.Sprintf("%s?%s", s.config.BaseURL, params.Encode())

	// 发送HTTP请求
	resp, err := s.client.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("请求高德API失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析JSON响应
	var geocodeResp GeocodeResponse
	if err := json.Unmarshal(body, &geocodeResp); err != nil {
		return nil, fmt.Errorf("解析响应JSON失败: %v", err)
	}

	// 检查响应状态
	if geocodeResp.Status != "1" {
		return nil, fmt.Errorf("高德API返回错误: %s", geocodeResp.Info)
	}

	// 检查是否有结果
	if len(geocodeResp.Geocodes) == 0 {
		return nil, fmt.Errorf("未找到该地址的坐标信息")
	}

	// 解析坐标
	location := geocodeResp.Geocodes[0].Location
	if location == "" {
		return nil, fmt.Errorf("坐标信息为空")
	}

	// 分割经纬度 (格式: "longitude,latitude")
	coords, err := parseLocation(location)
	if err != nil {
		return nil, fmt.Errorf("解析坐标失败: %v", err)
	}

	return coords, nil
}

// parseLocation 解析位置字符串为坐标
func parseLocation(location string) (*Coordinates, error) {
	// 高德API返回格式: "longitude,latitude"
	parts := make([]string, 0, 2) // 分割字符串
	var current string
	for _, char := range location {
		if char == ',' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(char)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}

	if len(parts) != 2 {
		return nil, fmt.Errorf("坐标格式错误: %s", location)
	}

	longitude, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, fmt.Errorf("解析经度失败: %v", err)
	}

	latitude, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, fmt.Errorf("解析纬度失败: %v", err)
	}

	return &Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
	}, nil
}

// ValidateCoordinates 验证坐标是否有效
func (s *AmapService) ValidateCoordinates(longitude, latitude float64) bool {
	// 中国境内经纬度范围验证
	// 经度范围: 73°33′E 至 135°05′E
	// 纬度范围: 3°51′N 至 53°33′N
	if longitude < 73.33 || longitude > 135.05 {
		return false
	}
	if latitude < 3.51 || latitude > 53.33 {
		return false
	}
	return true
}
