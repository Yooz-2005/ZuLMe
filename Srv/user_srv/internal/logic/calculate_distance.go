package logic

import (
	"Common/services"
	"fmt"
	user "user_srv/proto_user"
)

// CalculateDistance 计算用户到商家的距离
func CalculateDistance(in *user.CalculateDistanceRequest) (*user.CalculateDistanceResponse, error) {
	fmt.Printf("开始计算距离: 用户ID=%d, 地址=%s, 商家ID=%d\n", 
		in.Userid, in.Location, in.MerId)

	// 参数验证
	if in.Location == "" {
		return nil, fmt.Errorf("地址不能为空")
	}
	if in.MerId <= 0 {
		return nil, fmt.Errorf("商家ID无效")
	}

	// 创建Redis Geo服务实例
	geoService := services.NewRedisGeoService()

	// 根据地址计算到商家的距离
	distanceMeters, err := geoService.CalculateDistanceByAddress(in.Location, in.MerId)
	if err != nil {
		fmt.Printf("计算距离失败: %v\n", err)
		return nil, fmt.Errorf("计算距离失败: %v", err)
	}

	// 格式化距离显示
	distanceStr := services.FormatDistance(distanceMeters)
	
	fmt.Printf("距离计算成功: 用户ID=%d 到商家ID=%d 距离=%s\n", 
		in.Userid, in.MerId, distanceStr)

	return &user.CalculateDistanceResponse{
		Userid:   in.Userid,
		MerId:    in.MerId,
		Distance: distanceStr,
	}, nil
}

// GetNearbyMerchants 获取用户附近的商家 (可选功能)
func GetNearbyMerchants(userAddress string, radiusKm float64, limit int) ([]services.MerchantDistance, error) {
	fmt.Printf("查找附近商家: 地址=%s, 半径=%.1f公里, 限制=%d个\n", 
		userAddress, radiusKm, limit)

	// 获取用户地址的经纬度
	amapService := services.NewAmapService()
	coords, err := amapService.GetCoordinatesByAddress(userAddress)
	if err != nil {
		return nil, fmt.Errorf("获取用户地址坐标失败: %v", err)
	}

	// 创建Redis Geo服务实例
	geoService := services.NewRedisGeoService()

	// 查找附近商家
	merchants, err := geoService.GetNearbyMerchants(
		coords.Longitude, coords.Latitude, radiusKm, limit)
	if err != nil {
		return nil, fmt.Errorf("查找附近商家失败: %v", err)
	}

	fmt.Printf("找到%d个附近商家\n", len(merchants))
	return merchants, nil
}

// InitializeMerchantData 初始化商家位置数据 (管理员功能)
func InitializeMerchantData() error {
	fmt.Printf("开始初始化商家位置数据\n")
	
	geoService := services.NewRedisGeoService()
	err := geoService.InitializeMerchantLocations()
	if err != nil {
		fmt.Printf("初始化商家位置数据失败: %v\n", err)
		return err
	}
	
	fmt.Printf("商家位置数据初始化完成\n")
	return nil
}
