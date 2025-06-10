package logic

import (
	"Common/global"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	merchant "merchant_srv/proto_merchant"
	"models/model_mysql"
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

	// 4. 保存到数据库
	newMerchant := model_mysql.Merchant{
		Name:     in.Name,
		Phone:    in.Phone,
		Email:    in.Email,
		Password: string(hashedPassword),
	}

	result = global.DB.Create(&newMerchant)
	if result.Error != nil {
		return &merchant.MerchantRegisterResponse{Code: 500, Message: "商户注册失败"}, result.Error
	}

	fmt.Printf("商户注册成功: Name=%s, Phone=%s, Email=%s\n", in.Name, in.Phone, in.Email)

	return &merchant.MerchantRegisterResponse{Code: 200, Message: "注册成功"}, nil
}
