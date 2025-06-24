package utils

import (
	"Common/services"
	"fmt"
)

// InitMerchantLocations 初始化商家位置数据到Redis
func InitMerchantLocations() error {
	fmt.Println("开始初始化商家位置数据...")
	
	geoService := services.NewRedisGeoService()
	
	// 示例商家位置数据
	merchants := []services.MerchantLocation{
		{
			MerchantID: 1,
			Name:       "北京朝阳店",
			Address:    "北京市朝阳区三里屯太古里",
			Longitude:  116.447,
			Latitude:   39.937,
		},
		{
			MerchantID: 2,
			Name:       "北京海淀店", 
			Address:    "北京市海淀区中关村大街",
			Longitude:  116.298,
			Latitude:   39.959,
		},
		{
			MerchantID: 3,
			Name:       "上海浦东店",
			Address:    "上海市浦东新区陆家嘴金融中心",
			Longitude:  121.499,
			Latitude:   31.245,
		},
		{
			MerchantID: 4,
			Name:       "深圳南山店",
			Address:    "深圳市南山区科技园",
			Longitude:  113.947,
			Latitude:   22.531,
		},
		{
			MerchantID: 5,
			Name:       "广州天河店",
			Address:    "广州市天河区珠江新城",
			Longitude:  113.324,
			Latitude:   23.117,
		},
		{
			MerchantID: 6,
			Name:       "杭州西湖店",
			Address:    "杭州市西湖区文三路",
			Longitude:  120.131,
			Latitude:   30.279,
		},
		{
			MerchantID: 7,
			Name:       "成都锦江店",
			Address:    "成都市锦江区春熙路",
			Longitude:  104.081,
			Latitude:   30.660,
		},
		{
			MerchantID: 8,
			Name:       "武汉江汉店",
			Address:    "武汉市江汉区江汉路",
			Longitude:  114.273,
			Latitude:   30.584,
		},
		{
			MerchantID: 9,
			Name:       "西安雁塔店",
			Address:    "西安市雁塔区小寨",
			Longitude:  108.953,
			Latitude:   34.218,
		},
		{
			MerchantID: 10,
			Name:       "南京鼓楼店",
			Address:    "南京市鼓楼区新街口",
			Longitude:  118.778,
			Latitude:   32.041,
		},
	}
	
	// 添加所有商家位置
	for _, merchant := range merchants {
		err := geoService.AddMerchantLocation(&merchant)
		if err != nil {
			fmt.Printf("添加商家 %s (ID: %d) 失败: %v\n", merchant.Name, merchant.MerchantID, err)
			return err
		}
		fmt.Printf("✓ 成功添加商家: %s (ID: %d) - %s\n", 
			merchant.Name, merchant.MerchantID, merchant.Address)
	}
	
	fmt.Printf("✅ 成功初始化 %d 个商家位置数据\n", len(merchants))
	return nil
}

// TestDistanceCalculation 测试距离计算功能
func TestDistanceCalculation() error {
	fmt.Println("开始测试距离计算功能...")
	
	geoService := services.NewRedisGeoService()
	
	// 测试地址
	testCases := []struct {
		address    string
		merchantID int64
		expected   string
	}{
		{"北京市朝阳区国贸", 1, "北京朝阳店"},
		{"上海市黄浦区外滩", 3, "上海浦东店"},
		{"深圳市福田区华强北", 4, "深圳南山店"},
		{"广州市越秀区北京路", 5, "广州天河店"},
	}
	
	for i, testCase := range testCases {
		fmt.Printf("\n测试案例 %d: %s -> %s\n", i+1, testCase.address, testCase.expected)
		
		distance, err := geoService.CalculateDistanceByAddress(testCase.address, testCase.merchantID)
		if err != nil {
			fmt.Printf("❌ 计算失败: %v\n", err)
			continue
		}
		
		distanceStr := services.FormatDistance(distance)
		fmt.Printf("✓ 距离: %s\n", distanceStr)
	}
	
	fmt.Println("\n✅ 距离计算测试完成")
	return nil
}

// GetNearbyMerchantsExample 获取附近商家示例
func GetNearbyMerchantsExample() error {
	fmt.Println("开始测试附近商家查询...")
	
	geoService := services.NewRedisGeoService()
	
	// 测试地址: 北京市中心
	testAddress := "北京市东城区天安门"
	radiusKm := 50.0 // 50公里范围内
	limit := 5       // 最多5个商家
	
	fmt.Printf("查询地址: %s\n", testAddress)
	fmt.Printf("查询范围: %.1f公里\n", radiusKm)
	fmt.Printf("结果限制: %d个\n", limit)
	
	merchants, err := geoService.GetNearbyMerchants(116.397, 39.916, radiusKm, limit)
	if err != nil {
		fmt.Printf("❌ 查询失败: %v\n", err)
		return err
	}
	
	fmt.Printf("\n找到 %d 个附近商家:\n", len(merchants))
	for i, merchant := range merchants {
		fmt.Printf("%d. 商家ID: %d, 距离: %.2f公里\n", 
			i+1, merchant.MerchantID, merchant.Distance)
	}
	
	fmt.Println("\n✅ 附近商家查询测试完成")
	return nil
}
