package utils

import (
	"Common/global"
	"Common/services"
	"fmt"
	"models/model_mysql"
)

// SyncExistingMerchantsToRedis 同步已有商家数据到Redis
func SyncExistingMerchantsToRedis() error {
	fmt.Println("开始同步已有商家位置数据到Redis...")

	// 1. 从数据库获取所有商家
	var merchants []model_mysql.Merchant
	result := global.DB.Find(&merchants)
	if result.Error != nil {
		return fmt.Errorf("查询商家数据失败: %v", result.Error)
	}

	if len(merchants) == 0 {
		fmt.Println("数据库中没有商家数据")
		return nil
	}

	// 2. 创建Redis Geo服务
	geoService := services.NewRedisGeoService()

	// 3. 同步每个商家的位置信息
	successCount := 0
	failCount := 0

	for _, merchant := range merchants {
		// 跳过没有经纬度的商家
		if merchant.Longitude == 0 && merchant.Latitude == 0 {
			fmt.Printf("跳过商家 %s (ID: %d) - 没有经纬度信息\n",
				merchant.Name, merchant.ID)
			continue
		}

		merchantLocation := &services.MerchantLocation{
			MerchantID: int64(merchant.ID),
			Name:       merchant.Name,
			Address:    merchant.Location,
			Longitude:  merchant.Longitude,
			Latitude:   merchant.Latitude,
		}

		err := geoService.AddMerchantLocation(merchantLocation)
		if err != nil {
			fmt.Printf("❌ 同步失败: 商家 %s (ID: %d) - %v\n",
				merchant.Name, merchant.ID, err)
			failCount++
		} else {
			fmt.Printf("✓ 同步成功: 商家 %s (ID: %d) - %s\n",
				merchant.Name, merchant.ID, merchant.Location)
			successCount++
		}
	}

	fmt.Printf("\n✅ 同步完成: 成功 %d 个，失败 %d 个，总计 %d 个商家\n",
		successCount, failCount, len(merchants))

	return nil
}

// UpdateMerchantLocationInRedis 更新单个商家在Redis中的位置信息
func UpdateMerchantLocationInRedis(merchantID int64) error {
	fmt.Printf("更新商家 %d 的Redis位置信息...\n", merchantID)

	// 1. 从数据库获取商家信息
	var merchant model_mysql.Merchant
	result := global.DB.First(&merchant, merchantID)
	if result.Error != nil {
		return fmt.Errorf("查询商家信息失败: %v", result.Error)
	}

	// 2. 检查经纬度
	if merchant.Longitude == 0 && merchant.Latitude == 0 {
		return fmt.Errorf("商家 %d 没有经纬度信息", merchantID)
	}

	// 3. 更新Redis
	geoService := services.NewRedisGeoService()
	merchantLocation := &services.MerchantLocation{
		MerchantID: int64(merchant.ID),
		Name:       merchant.Name,
		Address:    merchant.Location,
		Longitude:  merchant.Longitude,
		Latitude:   merchant.Latitude,
	}

	err := geoService.AddMerchantLocation(merchantLocation)
	if err != nil {
		return fmt.Errorf("更新Redis失败: %v", err)
	}

	fmt.Printf("✓ 商家 %s (ID: %d) Redis位置信息更新成功\n",
		merchant.Name, merchant.ID)

	return nil
}

// RemoveMerchantFromRedis 从Redis中删除商家位置信息
func RemoveMerchantFromRedis(merchantID int64) error {
	fmt.Printf("从Redis中删除商家 %d 的位置信息...\n", merchantID)

	geoService := services.NewRedisGeoService()
	err := geoService.RemoveMerchantLocation(merchantID)
	if err != nil {
		return fmt.Errorf("从Redis删除商家位置失败: %v", err)
	}

	fmt.Printf("✓ 商家 %d 已从Redis中删除\n", merchantID)
	return nil
}

// GetMerchantLocationFromRedis 从Redis获取商家位置信息
func GetMerchantLocationFromRedis(merchantID int64) (*services.MerchantLocation, error) {
	//geoService := services.NewRedisGeoService()

	// 这里需要实现从Redis获取单个商家位置的方法
	// 由于当前的RedisGeoService没有这个方法，我们先从数据库获取
	var merchant model_mysql.Merchant
	result := global.DB.First(&merchant, merchantID)
	if result.Error != nil {
		return nil, fmt.Errorf("查询商家信息失败: %v", result.Error)
	}

	return &services.MerchantLocation{
		MerchantID: int64(merchant.ID),
		Name:       merchant.Name,
		Address:    merchant.Location,
		Longitude:  merchant.Longitude,
		Latitude:   merchant.Latitude,
	}, nil
}

// ValidateMerchantLocationData 验证商家位置数据的完整性
func ValidateMerchantLocationData() error {
	fmt.Println("开始验证商家位置数据完整性...")

	// 1. 获取数据库中的商家数据
	var merchants []model_mysql.Merchant
	result := global.DB.Find(&merchants)
	if result.Error != nil {
		return fmt.Errorf("查询商家数据失败: %v", result.Error)
	}

	// 2. 统计数据
	totalCount := len(merchants)
	withCoordsCount := 0
	withoutCoordsCount := 0

	fmt.Printf("数据库中共有 %d 个商家\n", totalCount)

	for _, merchant := range merchants {
		if merchant.Longitude != 0 && merchant.Latitude != 0 {
			withCoordsCount++
			fmt.Printf("✓ 商家 %s (ID: %d) 有坐标: (%.6f, %.6f)\n",
				merchant.Name, merchant.ID, merchant.Longitude, merchant.Latitude)
		} else {
			withoutCoordsCount++
			fmt.Printf("❌ 商家 %s (ID: %d) 缺少坐标信息\n",
				merchant.Name, merchant.ID)
		}
	}

	fmt.Printf("\n📊 统计结果:\n")
	fmt.Printf("- 有坐标的商家: %d 个\n", withCoordsCount)
	fmt.Printf("- 缺少坐标的商家: %d 个\n", withoutCoordsCount)
	fmt.Printf("- 坐标完整率: %.2f%%\n", float64(withCoordsCount)/float64(totalCount)*100)

	return nil
}

// FixMerchantCoordinates 修复缺少坐标的商家数据
func FixMerchantCoordinates() error {
	fmt.Println("开始修复缺少坐标的商家数据...")

	// 1. 查找缺少坐标的商家
	var merchants []model_mysql.Merchant
	result := global.DB.Where("(longitude = 0 OR latitude = 0) AND location != ''").Find(&merchants)
	if result.Error != nil {
		return fmt.Errorf("查询缺少坐标的商家失败: %v", result.Error)
	}

	if len(merchants) == 0 {
		fmt.Println("所有商家都已有坐标信息")
		return nil
	}

	fmt.Printf("找到 %d 个需要修复坐标的商家\n", len(merchants))

	// 2. 为每个商家获取坐标
	amapService := services.NewAmapService()
	geoService := services.NewRedisGeoService()

	successCount := 0
	failCount := 0

	for _, merchant := range merchants {
		fmt.Printf("正在处理商家: %s (ID: %d) - %s\n",
			merchant.Name, merchant.ID, merchant.Location)

		// 获取坐标
		coords, err := amapService.GetCoordinatesByAddress(merchant.Location)
		if err != nil {
			fmt.Printf("❌ 获取坐标失败: %v\n", err)
			failCount++
			continue
		}

		// 更新数据库
		result := global.DB.Model(&merchant).Updates(map[string]interface{}{
			"longitude": coords.Longitude,
			"latitude":  coords.Latitude,
		})
		if result.Error != nil {
			fmt.Printf("❌ 更新数据库失败: %v\n", result.Error)
			failCount++
			continue
		}

		// 更新Redis
		merchantLocation := &services.MerchantLocation{
			MerchantID: int64(merchant.ID),
			Name:       merchant.Name,
			Address:    merchant.Location,
			Longitude:  coords.Longitude,
			Latitude:   coords.Latitude,
		}

		err = geoService.AddMerchantLocation(merchantLocation)
		if err != nil {
			fmt.Printf("⚠️ Redis更新失败: %v (数据库已更新)\n", err)
		}

		fmt.Printf("✓ 修复成功: 经度=%.6f, 纬度=%.6f\n",
			coords.Longitude, coords.Latitude)
		successCount++
	}

	fmt.Printf("\n✅ 修复完成: 成功 %d 个，失败 %d 个\n", successCount, failCount)
	return nil
}
