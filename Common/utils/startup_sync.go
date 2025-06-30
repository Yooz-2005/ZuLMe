package utils

import (
	"Common/services"
	"fmt"
	"time"
)

// StartupSyncMerchantLocations 服务启动时同步商家位置数据
func StartupSyncMerchantLocations() {
	fmt.Println("🚀 服务启动中，开始同步商家位置数据...")

	// 延迟一下确保数据库连接已建立
	time.Sleep(2 * time.Second)

	// 1. 验证数据完整性
	fmt.Println("📊 步骤1: 验证商家位置数据完整性...")
	err := ValidateMerchantLocationData()
	if err != nil {
		fmt.Printf("⚠️ 数据验证失败: %v\n", err)
	}

	// 2. 修复缺少坐标的商家数据
	fmt.Println("\n🔧 步骤2: 修复缺少坐标的商家数据...")
	err = FixMerchantCoordinates()
	if err != nil {
		fmt.Printf("⚠️ 坐标修复失败: %v\n", err)
	}

	// 3. 同步所有商家位置到Redis
	fmt.Println("\n🔄 步骤3: 同步商家位置数据到Redis...")
	err = SyncExistingMerchantsToRedis()
	if err != nil {
		fmt.Printf("⚠️ Redis同步失败: %v\n", err)
	} else {
		fmt.Println("✅ 商家位置数据同步完成")
	}

	fmt.Println("🎉 商家位置数据初始化完成!")
}

// PeriodicSyncMerchantLocations 定期同步商家位置数据 (可选)
//定期同步防止数据不一致
//商家位置发生变更 及时更新 定期同步双重保障
func PeriodicSyncMerchantLocations(intervalHours int) {
	ticker := time.NewTicker(time.Duration(intervalHours) * time.Hour) // 每小时执行一次
	defer ticker.Stop()

	fmt.Printf("⏰ 启动定期同步任务，间隔: %d 小时\n", intervalHours)

	for {
		select {//使用select语句监听定时器通道
		case <-ticker.C: //监听定时器信号
			fmt.Println("🔄 开始定期同步商家位置数据...")//执行同步逻辑

			err := SyncExistingMerchantsToRedis()
			if err != nil {
				fmt.Printf("❌ 定期同步失败: %v\n", err)
			} else {
				fmt.Println("✅ 定期同步完成")
			}
		}
	}
}


// CheckRedisConnection 检查Redis连接状态
func CheckRedisConnection() error {
	fmt.Println("🔍 检查Redis连接状态...")

	geoService := services.NewRedisGeoService()

	// 尝试添加一个测试位置
	testLocation := &services.MerchantLocation{
		MerchantID: -1, // 使用负数ID作为测试
		Name:       "测试位置",
		Address:    "测试地址",
		Longitude:  116.397,
		Latitude:   39.916,
	}

	err := geoService.AddMerchantLocation(testLocation)
	if err != nil {
		return fmt.Errorf("Redis连接失败: %v", err)
	}

	// 删除测试位置
	err = geoService.RemoveMerchantLocation(-1)
	if err != nil {
		fmt.Printf("⚠️ 清理测试数据失败: %v\n", err)
	}

	fmt.Println("✅ Redis连接正常")
	return nil
}

// CheckAmapService 检查高德地图服务状态
func CheckAmapService() error {
	fmt.Println("🗺️ 检查高德地图服务状态...")

	amapService := services.NewAmapService()

	// 测试地址解析
	testAddress := "北京市朝阳区三里屯"
	coords, err := amapService.GetCoordinatesByAddress(testAddress)
	if err != nil {
		return fmt.Errorf("高德地图服务异常: %v", err)
	}

	fmt.Printf("✅ 高德地图服务正常，测试坐标: (%.6f, %.6f)\n",
		coords.Longitude, coords.Latitude)
	return nil
}

// HealthCheck 健康检查
func HealthCheck() error {
	fmt.Println("🏥 开始健康检查...")

	// 检查Redis连接
	if err := CheckRedisConnection(); err != nil {
		return err
	}

	// 检查高德地图服务
	if err := CheckAmapService(); err != nil {
		return err
	}

	fmt.Println("✅ 所有服务健康检查通过")
	return nil
}

// GetSystemStatus 获取系统状态信息
func GetSystemStatus() map[string]interface{} {
	status := make(map[string]interface{})

	// Redis状态
	redisErr := CheckRedisConnection()
	status["redis"] = map[string]interface{}{
		"status": redisErr == nil,
		"error":  getErrorString(redisErr),
	}

	// 高德地图服务状态
	amapErr := CheckAmapService()
	status["amap"] = map[string]interface{}{
		"status": amapErr == nil,
		"error":  getErrorString(amapErr),
	}

	// 商家数据统计
	merchantStats := getMerchantStats()
	status["merchants"] = merchantStats

	return status
}

// 获取错误字符串
func getErrorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

// 获取商家数据统计
func getMerchantStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// 这里可以添加更详细的统计信息
	stats["total"] = 0
	stats["with_coordinates"] = 0
	stats["without_coordinates"] = 0

	// 实际实现需要查询数据库
	// 为了简化，这里返回默认值

	return stats
}
