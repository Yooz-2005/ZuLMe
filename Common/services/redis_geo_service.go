package services

import (
	"Common/global"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"

	"golang.org/x/net/context"
)

// RedisGeoService Redis地理位置服务
type RedisGeoService struct {
	client *redis.Client
	ctx    context.Context
}

// MerchantLocation 商家位置信息
type MerchantLocation struct {
	MerchantID int64   `json:"merchant_id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

// NewRedisGeoService 创建Redis地理位置服务实例
func NewRedisGeoService() *RedisGeoService {
	return &RedisGeoService{
		client: global.Rdb,
		ctx:    context.Background(),
	}
}

// AddMerchantLocation 添加商家位置到Redis Geo
func (s *RedisGeoService) AddMerchantLocation(merchant *MerchantLocation) error {
	key := "merchants:locations"
	memberName := fmt.Sprintf("merchant:%d", merchant.MerchantID)

	// 使用GEOADD命令添加商家位置
	err := s.client.GeoAdd(s.ctx, key, &redis.GeoLocation{
		Name:      memberName,
		Longitude: merchant.Longitude,
		Latitude:  merchant.Latitude,
	}).Err()

	if err != nil {
		return fmt.Errorf("添加商家位置失败: %v", err)
	}

	fmt.Printf("成功添加商家位置: ID=%d, 经纬度=(%f,%f)\n",
		merchant.MerchantID, merchant.Longitude, merchant.Latitude)
	return nil
}

// CalculateDistance 计算用户位置到指定商家的距离
func (s *RedisGeoService) CalculateDistance(userLongitude, userLatitude float64, merchantID int64) (float64, error) {
	key := "merchants:locations"
	memberName := fmt.Sprintf("merchant:%d", merchantID)

	// 首先检查商家是否存在
	exists, err := s.client.GeoPos(s.ctx, key, memberName).Result()
	if err != nil {
		return 0, fmt.Errorf("查询商家位置失败: %v", err)
	}

	if len(exists) == 0 || exists[0] == nil {
		return 0, fmt.Errorf("商家ID %d 不存在或位置信息未设置", merchantID)
	}

	// 临时添加用户位置
	userMemberName := fmt.Sprintf("temp_user_%d", merchantID)
	err = s.client.GeoAdd(s.ctx, key, &redis.GeoLocation{
		Name:      userMemberName,
		Longitude: userLongitude,
		Latitude:  userLatitude,
	}).Err()

	if err != nil {
		return 0, fmt.Errorf("添加临时用户位置失败: %v", err)
	}

	// 计算距离 (单位: 米)
	distances, err := s.client.GeoDist(s.ctx, key, userMemberName, memberName, "m").Result()
	if err != nil {
		// 清理临时用户位置
		s.client.ZRem(s.ctx, key, userMemberName)
		return 0, fmt.Errorf("计算距离失败: %v", err)
	}

	// 清理临时用户位置
	s.client.ZRem(s.ctx, key, userMemberName)

	fmt.Printf("计算距离成功: 用户(%f,%f) 到商家%d 距离=%.2f米\n",
		userLongitude, userLatitude, merchantID, distances)

	return distances, nil
}

// CalculateDistanceByAddress 根据地址计算到商家的距离
func (s *RedisGeoService) CalculateDistanceByAddress(address string, merchantID int64) (float64, error) {
	// 使用高德API获取地址的经纬度
	amapService := NewAmapService()
	coords, err := amapService.GetCoordinatesByAddress(address)
	if err != nil {
		return 0, fmt.Errorf("获取地址坐标失败: %v", err)
	}

	// 计算距离
	return s.CalculateDistance(coords.Longitude, coords.Latitude, merchantID)
}

// GetNearbyMerchants 获取用户附近的商家 (按距离排序)
func (s *RedisGeoService) GetNearbyMerchants(userLongitude, userLatitude float64, radiusKm float64, limit int) ([]MerchantDistance, error) {
	key := "merchants:locations"

	// 使用GEORADIUS命令查找附近的商家
	results, err := s.client.GeoRadius(s.ctx, key, userLongitude, userLatitude, &redis.GeoRadiusQuery{
		Radius:    radiusKm,
		Unit:      "km",
		WithCoord: true,
		WithDist:  true,
		Count:     limit,
		Sort:      "ASC", // 按距离升序排列
	}).Result()

	if err != nil {
		return nil, fmt.Errorf("查询附近商家失败: %v", err)
	}

	var merchants []MerchantDistance
	for _, result := range results {
		// 解析商家ID
		memberName := result.Name
		if !strings.HasPrefix(memberName, "merchant:") {
			continue
		}

		merchantIDStr := strings.TrimPrefix(memberName, "merchant:")
		merchantID, err := strconv.ParseInt(merchantIDStr, 10, 64)
		if err != nil {
			continue
		}

		merchants = append(merchants, MerchantDistance{
			MerchantID: merchantID,
			Distance:   result.Dist,
			Longitude:  result.Longitude,
			Latitude:   result.Latitude,
		})
	}

	return merchants, nil
}

// MerchantDistance 商家距离信息
type MerchantDistance struct {
	MerchantID int64   `json:"merchant_id"`
	Distance   float64 `json:"distance"` // 距离(公里)
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}

// InitializeMerchantLocations 初始化商家位置数据 (示例数据)
func (s *RedisGeoService) InitializeMerchantLocations() error {
	// 示例商家位置数据
	merchants := []MerchantLocation{
		{MerchantID: 1, Name: "北京朝阳店", Address: "北京市朝阳区三里屯", Longitude: 116.447, Latitude: 39.937},
		{MerchantID: 2, Name: "北京海淀店", Address: "北京市海淀区中关村", Longitude: 116.298, Latitude: 39.959},
		{MerchantID: 3, Name: "上海浦东店", Address: "上海市浦东新区陆家嘴", Longitude: 121.499, Latitude: 31.245},
		{MerchantID: 4, Name: "深圳南山店", Address: "深圳市南山区科技园", Longitude: 113.947, Latitude: 22.531},
		{MerchantID: 5, Name: "广州天河店", Address: "广州市天河区珠江新城", Longitude: 113.324, Latitude: 23.117},
	}

	for _, merchant := range merchants {
		if err := s.AddMerchantLocation(&merchant); err != nil {
			return fmt.Errorf("初始化商家%d位置失败: %v", merchant.MerchantID, err)
		}
	}

	fmt.Printf("成功初始化%d个商家位置\n", len(merchants))
	return nil
}

// RemoveMerchantLocation 删除商家位置
func (s *RedisGeoService) RemoveMerchantLocation(merchantID int64) error {
	key := "merchants:locations"
	memberName := fmt.Sprintf("merchant:%d", merchantID)

	err := s.client.ZRem(s.ctx, key, memberName).Err()
	if err != nil {
		return fmt.Errorf("删除商家位置失败: %v", err)
	}

	fmt.Printf("成功删除商家%d的位置信息\n", merchantID)
	return nil
}

// ClearAllMerchants 清理所有商家位置数据
func (s *RedisGeoService) ClearAllMerchants() error {
	key := "merchants:locations"

	err := s.client.Del(s.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("清理商家位置数据失败: %v", err)
	}

	fmt.Printf("成功清理所有商家位置数据\n")
	return nil
}

// FormatDistance 格式化距离显示
func FormatDistance(distanceMeters float64) string {
	if distanceMeters < 1000 {
		return fmt.Sprintf("%.0f米", distanceMeters)
	} else {
		return fmt.Sprintf("%.2f公里", distanceMeters/1000)
	}
}
