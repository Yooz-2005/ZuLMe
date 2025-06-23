package logic

import (
	"Common/global"
	"Common/pkg"
	"Common/services"
	"Common/utils"
	"errors"
	"fmt"
	"models/model_mysql"
	"models/model_redis"
	user "user_srv/proto_user"

	"gorm.io/gorm"

	"strconv"
)

// 生成Token响应的辅助函数
func generateTokenResponse(userID int64) (*user.UserRegisterResponse, error) {
	j := pkg.NewJWT("2209")
	token, err := j.CreateToken(pkg.CustomClaims{ID: uint(userID)})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &user.UserRegisterResponse{
		UserId: strconv.FormatUint(uint64(userID), 10),
		Token:  token,
	}, nil
}

func UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	// 验证验证码
	code, err := model_redis.GetVerificationCode("register", in.Phone)
	if err != nil {
		return nil, errors.New("验证码已过期")
	}
	if code != in.Code {
		return nil, errors.New("验证码错误")
	}

	// 檢查用戶是否已存在
	var existingUser model_mysql.User
	err = existingUser.Login(in.Phone)
	if err == nil {
		// 用戶已存在，生成 token
		return generateTokenResponse(existingUser.Id)
	}
	// 获取经纬度坐标
	var latAndLon string
	fmt.Printf("开始处理用户注册，手机号: %s, 地址: %s\n", in.Phone, in.Location)

	if in.Location != "" {
		// 调用高德API获取经纬度
		fmt.Printf("调用高德API获取坐标，地址: %s\n", in.Location)
		amapService := services.NewAmapService()
		coords, err := amapService.GetCoordinatesByAddress(in.Location)
		if err != nil {
			fmt.Printf("获取坐标失败: %v, 地址: %s\n", err, in.Location)
			// 坐标获取失败不影响注册，使用空值
			latAndLon = ""
		} else {
			latAndLon = fmt.Sprintf("%f,%f", coords.Longitude, coords.Latitude)
			fmt.Printf("获取坐标成功: 经纬度=%s, 地址=%s\n", latAndLon, in.Location)
		}
	} else {
		fmt.Printf("地址为空，跳过坐标获取\n")
	}

	// 创建新用户和用户资料
	newUser := &model_mysql.User{
		Phone:     in.Phone,
		Location:  in.Location,
		LatAndLon: latAndLon,
	}

	fmt.Printf("准备保存用户信息: Phone=%s, Location=%s, LatAndLon=%s\n",
		newUser.Phone, newUser.Location, newUser.LatAndLon)
	userProfile := &model_mysql.UserProfile{
		UserId: int32(newUser.Id),
		Phone:  utils.EncryptPhone(in.Phone),
	}
	//使用事务处理数据库操作
	fmt.Printf("开始数据库事务操作\n")
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		fmt.Printf("执行用户注册到数据库\n")
		if err := newUser.Register(); err != nil {
			fmt.Printf("用户注册失败: %v\n", err)
			return fmt.Errorf("用户注册失败: %v", err)
		}
		fmt.Printf("用户注册成功，用户ID: %d\n", newUser.Id)

		userProfile.UserId = int32(newUser.Id)
		fmt.Printf("创建用户资料\n")
		if err := userProfile.CreateUserProfile(); err != nil {
			fmt.Printf("用户资料创建失败: %v\n", err)
			return fmt.Errorf("用户资料创建失败: %v", err)
		}
		fmt.Printf("用户资料创建成功\n")
		return nil
	})

	if err != nil {
		fmt.Printf("数据库事务失败: %v\n", err)
		return nil, err
	}
	// 删除验证码
	err = model_redis.DeleteVerificationCode("register", in.Phone)
	if err != nil {
		return nil, errors.New("验证码删除失败")
	}
	return generateTokenResponse(newUser.Id)

}
