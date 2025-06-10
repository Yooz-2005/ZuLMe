package logic

import (
	"Common/global"
	"Common/pkg"

	"errors"
	"gorm.io/gorm"
	"models/model_mysql"
	"models/model_redis"
	"strconv"
	user "user_srv/proto_user"
)

// 生成Token响应的辅助函数
func generateTokenResponse(userID uint64) (*user.UserRegisterResponse, error) {
	j := pkg.NewJWT("2209")
	token, err := j.CreateToken(pkg.CustomClaims{ID: uint(userID)})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &user.UserRegisterResponse{
		UserId: strconv.FormatUint(userID, 10),
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
	// 创建新用户和用户资料
	newUser := &model_mysql.User{
		Phone: in.Phone,
	}
	userProfile := &model_mysql.UserProfile{
		UserId: int32(newUser.Id),
		Phone:  utils.EncryptPhone(in.Phone),
	}
	//使用事务处理数据库操作
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := newUser.Register(); err != nil {
			return errors.New("用户注册失败")
		}
		userProfile.UserId = int32(newUser.Id)
		return userProfile.CreateUserProfile()
	})
	// 删除验证码
	err = model_redis.DeleteVerificationCode("register", in.Phone)
	if err != nil {
		return nil, errors.New("验证码删除失败")
	}
	return generateTokenResponse(newUser.Id)

}
