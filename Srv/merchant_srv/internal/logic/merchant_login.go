package logic

import (
	"Common/global"
	jwt "Common/pkg"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	merchant "merchant_srv/proto_merchant"
	"models/model_mysql"
)

func MerchantLogin(ctx context.Context, in *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	// 1. 参数校验
	if in.Phone == "" || in.Password == "" {
		return &merchant.MerchantLoginResponse{Code: 400, Message: "手机号和密码是必填项"}, nil
	}

	// 2. 从数据库查询用户
	var existingMerchant model_mysql.Merchant
	result := global.DB.Where("phone = ?", in.Phone).First(&existingMerchant)
	if result.Error != nil {
		// 如果找不到用户或者有其他数据库错误
		return &merchant.MerchantLoginResponse{Code: 404, Message: "用户不存在或数据库错误"}, result.Error
	}

	// 3. 密码比对
	err := bcrypt.CompareHashAndPassword([]byte(existingMerchant.Password), []byte(in.Password))
	if err != nil {
		// 密码不匹配
		return &merchant.MerchantLoginResponse{Code: 401, Message: "密码不正确"}, nil
	}

	// 4. 生成Token (使用JWT库)
	claims := jwt.CustomClaims{ID: existingMerchant.ID}
	token, err := jwt.NewJWT("merchant").CreateToken(claims)
	if err != nil {
		return &merchant.MerchantLoginResponse{Code: 500, Message: "生成Token失败"}, err
	}

	fmt.Printf("商户登录成功: Phone=%s\n", in.Phone)

	return &merchant.MerchantLoginResponse{Code: 200, Message: "登录成功", Token: token}, nil
}
