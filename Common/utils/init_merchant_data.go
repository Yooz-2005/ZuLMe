package utils

import (
	"Common/global"
	"Common/services"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"models/model_mysql"
)

// MerchantTestData 测试商家数据结构
type MerchantTestData struct {
	Name         string
	Phone        string
	Email        string
	Password     string
	Location     string
	BusinessTime string
	Longitude    float64
	Latitude     float64
}

// InitMerchantTestData 初始化商家测试数据到数据库
func InitMerchantTestData() error {
	fmt.Println("开始初始化商家测试数据到数据库...")

	// 测试商家数据
	merchantsData := []MerchantTestData{
		{
			Name:         "北京朝阳店",
			Phone:        "13800001001",
			Email:        "beijing_chaoyang@zulme.com",
			Password:     "123456",
			Location:     "北京市朝阳区三里屯太古里",
			BusinessTime: "08:00-22:00",
			Longitude:    116.447,
			Latitude:     39.937,
		},
		{
			Name:         "北京海淀店",
			Phone:        "13800001002",
			Email:        "beijing_haidian@zulme.com",
			Password:     "123456",
			Location:     "北京市海淀区中关村大街",
			BusinessTime: "08:00-22:00",
			Longitude:    116.298,
			Latitude:     39.959,
		},
		{
			Name:         "上海浦东店",
			Phone:        "13800001003",
			Email:        "shanghai_pudong@zulme.com",
			Password:     "123456",
			Location:     "上海市浦东新区陆家嘴金融中心",
			BusinessTime: "08:00-22:00",
			Longitude:    121.499,
			Latitude:     31.245,
		},
		{
			Name:         "深圳南山店",
			Phone:        "13800001004",
			Email:        "shenzhen_nanshan@zulme.com",
			Password:     "123456",
			Location:     "深圳市南山区科技园",
			BusinessTime: "08:00-22:00",
			Longitude:    113.947,
			Latitude:     22.531,
		},
		{
			Name:         "广州天河店",
			Phone:        "13800001005",
			Email:        "guangzhou_tianhe@zulme.com",
			Password:     "123456",
			Location:     "广州市天河区珠江新城",
			BusinessTime: "08:00-22:00",
			Longitude:    113.324,
			Latitude:     23.117,
		},
		{
			Name:         "杭州西湖店",
			Phone:        "13800001006",
			Email:        "hangzhou_xihu@zulme.com",
			Password:     "123456",
			Location:     "杭州市西湖区文三路",
			BusinessTime: "08:00-22:00",
			Longitude:    120.131,
			Latitude:     30.279,
		},
		{
			Name:         "成都锦江店",
			Phone:        "13800001007",
			Email:        "chengdu_jinjiang@zulme.com",
			Password:     "123456",
			Location:     "成都市锦江区春熙路",
			BusinessTime: "08:00-22:00",
			Longitude:    104.081,
			Latitude:     30.660,
		},
		{
			Name:         "武汉江汉店",
			Phone:        "13800001008",
			Email:        "wuhan_jianghan@zulme.com",
			Password:     "123456",
			Location:     "武汉市江汉区江汉路",
			BusinessTime: "08:00-22:00",
			Longitude:    114.273,
			Latitude:     30.584,
		},
		{
			Name:         "西安雁塔店",
			Phone:        "13800001009",
			Email:        "xian_yanta@zulme.com",
			Password:     "123456",
			Location:     "西安市雁塔区小寨",
			BusinessTime: "08:00-22:00",
			Longitude:    108.953,
			Latitude:     34.218,
		},
		{
			Name:         "南京鼓楼店",
			Phone:        "13800001010",
			Email:        "nanjing_gulou@zulme.com",
			Password:     "123456",
			Location:     "南京市鼓楼区新街口",
			BusinessTime: "08:00-22:00",
			Longitude:    118.778,
			Latitude:     32.041,
		},
	}

	// 创建Redis Geo服务实例
	geoService := services.NewRedisGeoService()

	// 逐个创建商家记录
	for _, data := range merchantsData {
		// 检查是否已存在
		var existingMerchant model_mysql.Merchant
		result := global.DB.Where("phone = ? OR email = ?", data.Phone, data.Email).First(&existingMerchant)
		if result.Error == nil {
			fmt.Printf("商家 %s 已存在，跳过创建\n", data.Name)
			continue
		}

		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("❌ 加密密码失败: %v\n", err)
			continue
		}

		// 创建商家记录
		merchant := model_mysql.Merchant{
			Name:         data.Name,
			Phone:        data.Phone,
			Email:        data.Email,
			Password:     string(hashedPassword),
			Location:     data.Location,
			BusinessTime: data.BusinessTime,
			Longitude:    data.Longitude,
			Latitude:     data.Latitude,
			Status:       1, // 设置为审核通过状态
		}

		result = global.DB.Create(&merchant)
		if result.Error != nil {
			fmt.Printf("❌ 创建商家 %s 失败: %v\n", data.Name, result.Error)
			continue
		}

		fmt.Printf("✓ 成功创建商家: %s (ID: %d)\n", data.Name, merchant.ID)

		// 同时添加到Redis Geo
		merchantLocation := &services.MerchantLocation{
			MerchantID: int64(merchant.ID),
			Name:       data.Name,
			Address:    data.Location,
			Longitude:  data.Longitude,
			Latitude:   data.Latitude,
		}

		err = geoService.AddMerchantLocation(merchantLocation)
		if err != nil {
			fmt.Printf("⚠️ 添加商家 %s 到Redis失败: %v\n", data.Name, err)
		} else {
			fmt.Printf("✓ 成功添加商家 %s 到Redis\n", data.Name)
		}
	}

	fmt.Printf("✅ 商家测试数据初始化完成\n")
	return nil
}

// CleanMerchantTestData 清理测试数据
func CleanMerchantTestData() error {
	fmt.Println("开始清理商家测试数据...")

	// 删除测试邮箱的商家
	result := global.DB.Where("email LIKE ?", "%@zulme.com").Delete(&model_mysql.Merchant{})
	if result.Error != nil {
		fmt.Printf("❌ 清理数据库数据失败: %v\n", result.Error)
		return result.Error
	}

	fmt.Printf("✓ 成功删除 %d 条数据库记录\n", result.RowsAffected)

	// 清理Redis数据
	geoService := services.NewRedisGeoService()
	err := geoService.ClearAllMerchants()
	if err != nil {
		fmt.Printf("⚠️ 清理Redis数据失败: %v\n", err)
	} else {
		fmt.Printf("✓ 成功清理Redis数据\n")
	}

	fmt.Printf("✅ 测试数据清理完成\n")
	return nil
}
