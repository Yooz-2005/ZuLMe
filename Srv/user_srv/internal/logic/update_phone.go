package logic

import (
	"Common/utils"
	"errors"
	"models/model_mysql"
	user "user_srv/proto_user"
)

func UpdateUserPhone(in *user.UpdateUserPhoneRequest) (*user.UpdateUserPhoneResponse, error) {
	// 调用 UserProfile 的方法来更新数据库
	userProfile := &model_mysql.UserProfile{}
	err := userProfile.UpdateUserPhoneByUserId(in.UserId, utils.EncryptPhone(in.Phone))
	if err != nil {
		return nil, errors.New("修改用户手机号失败")
	}
	return &user.UpdateUserPhoneResponse{
		UserId:  in.UserId,
		Message: "修改用户手机号成功",
	}, nil
}
