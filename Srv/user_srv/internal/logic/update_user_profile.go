package logic

import (
	"ZuLMe/ZuLMe/models/model_mysql"
	"errors"
	"time"
	user "user_srv/proto_user"
)

func SetUserProfile(in *user.UpdateUserProfileRequest) (*user.UpdateUserProfileResponse, error) {
	// 定义日期格式 假设日期格式为 "YYYY-MM-DD"
	const dateFormat = "2006-01-02"

	println("SetUserProfile: 接收到的 IdExpireDate:", in.IdExpireDate) // 添加日誌

	var parsedDate *time.Time //用于存储解析后的时间
	//为什么使用指针？
	//如果 IdExpireDate 为空字符串，那么 parsedDate 应该为 nil，而不是指向一个无效的时间。
	//如果 IdExpireDate 不为空字符串，那么 parsedDate 应该指向一个有效的时间。
	if in.IdExpireDate != "" {
		d, err := time.Parse(dateFormat, in.IdExpireDate)
		if err != nil {
			return nil, errors.New("无效的证件有效期格式")
		}
		parsedDate = &d
	} else {
		parsedDate = nil // 如果 IdExpireDate 为空字符串，则设置为 nil
	}

	updates := map[string]interface{}{
		"real_name":       in.RealName,
		"id_type":         in.IdType,
		"id_number":       in.IdNumber,
		"email":           in.Email,
		"province":        in.Province,
		"city":            in.City,
		"district":        in.District,
		"emergency_name":  in.EmergencyName,
		"emergency_phone": in.EmergencyPhone,
	}

	// 只有当 parsedDate 不为 nil 时才添加 IdExpireDate，避免 GORM 尝试更新一个不存在的空值
	if parsedDate != nil {
		updates["id_expire_date"] = parsedDate
	} else {
		// 如果 parsedDate 为 nil，明确设置 id_expire_date 为 NULL
		updates["id_expire_date"] = nil
	}

	// 调用 UserProfile 的方法来更新数据库
	userProfile := &model_mysql.UserProfile{}
	err := userProfile.UpdateUserProfileByMap(in.UserId, updates)
	if err != nil {
		println("SetUserProfile: 更新用户资料失败:", err.Error()) // 增加日誌
		return nil, errors.New("更新用户资料失败")
	}

	return &user.UpdateUserProfileResponse{
		UserId:  in.UserId,
		Message: "保存成功",
	}, nil
}
