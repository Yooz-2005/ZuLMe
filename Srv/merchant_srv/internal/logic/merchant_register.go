package logic

import (
	"Common/global"
	"Common/services"
	"context"
	"fmt"
	merchant "merchant_srv/proto_merchant"
	"models/model_mysql"

	"golang.org/x/crypto/bcrypt"
)

func MerchantRegister(ctx context.Context, in *merchant.MerchantRegisterRequest) (*merchant.MerchantRegisterResponse, error) {
	// 1. 参数校验
	if in.Name == "" || in.Phone == "" || in.Email == "" || in.Password == "" || in.ConfirmPass == "" {
		return &merchant.MerchantRegisterResponse{Code: 400, Message: "所有字段都是必填项"}, nil
	}

	if in.Password != in.ConfirmPass {
		return &merchant.MerchantRegisterResponse{Code: 400, Message: "两次输入的密码不一致"}, nil
	}

	// TODO: 更多密码强度校验，例如长度、复杂性要求

	// 2. 检查手机号或邮箱是否已注册
	var existingMerchant model_mysql.Merchant
	result := global.DB.Where("phone = ? OR email = ?", in.Phone, in.Email).First(&existingMerchant)
	if result.Error == nil {
		return &merchant.MerchantRegisterResponse{Code: 400, Message: "手机号或邮箱已注册"}, nil
	}

	// 3. 密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &merchant.MerchantRegisterResponse{Code: 500, Message: "密码哈希失败"}, err
	}

	// 4. 获取经纬度坐标
	var longitude, latitude float64
	if in.Location != "" {
		// 如果前端没有提供经纬度，则调用高德API获取
		if in.Longitude == 0 && in.Latitude == 0 {
			amapService := services.NewAmapService()
			coords, err := amapService.GetCoordinatesByAddress(in.Location)
			if err != nil {
				fmt.Printf("获取坐标失败: %v, 地址: %s\n", err, in.Location)
				// 坐标获取失败不影响注册，使用默认值
				longitude = 0
				latitude = 0
			} else {
				longitude = coords.Longitude
				latitude = coords.Latitude
				fmt.Printf("获取坐标成功: 经度=%f, 纬度=%f, 地址=%s\n", longitude, latitude, in.Location)
			}
		} else {
			// 使用前端提供的经纬度
			longitude = in.Longitude
			latitude = in.Latitude
		}
	}

	// 5. 保存到数据库
	newMerchant := model_mysql.Merchant{
		Name:         in.Name,
		Phone:        in.Phone,
		Email:        in.Email,
		Password:     string(hashedPassword),
		Location:     in.Location,
		BusinessTime: in.BusinessTime,
		Longitude:    longitude,
		Latitude:     latitude,
	}

	result = global.DB.Create(&newMerchant)
	if result.Error != nil {
		return &merchant.MerchantRegisterResponse{Code: 500, Message: "商户注册失败"}, result.Error
	}

	// 6. 将商家位置信息存储到Redis Geo
	if longitude != 0 && latitude != 0 {
		geoService := services.NewRedisGeoService()
		merchantLocation := &services.MerchantLocation{
			MerchantID: int64(newMerchant.ID),
			Name:       in.Name,
			Address:    in.Location,
			Longitude:  longitude,
			Latitude:   latitude,
		}

		err := geoService.AddMerchantLocation(merchantLocation)
		if err != nil {
			fmt.Printf("警告: 商家位置信息存储到Redis失败: %v, 商家ID=%d\n", err, newMerchant.ID)
			// Redis存储失败不影响注册流程，只记录警告
		} else {
			fmt.Printf("商家位置信息已存储到Redis: 商家ID=%d, 经度=%f, 纬度=%f\n",
				newMerchant.ID, longitude, latitude)
		}
	} else {
		fmt.Printf("警告: 商家经纬度为空，跳过Redis存储: 商家ID=%d\n", newMerchant.ID)
	}

	fmt.Printf("商户注册成功: Name=%s, Phone=%s, Email=%s, ID=%d\n",
		in.Name, in.Phone, in.Email, newMerchant.ID)

	return &merchant.MerchantRegisterResponse{Code: 200, Message: "注册成功"}, nil
}
