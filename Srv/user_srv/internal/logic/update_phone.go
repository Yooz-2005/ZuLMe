package logic

import (
	"Common/global"
	"Common/utils"
	"errors"
	"models/model_mysql"
	user "user_srv/proto_user"
)

func UpdateUserPhone(in *user.UpdateUserPhoneRequest) (*user.UpdateUserPhoneResponse, error) {
	// 检查手机号是否已被其他用户注册
	users := &model_mysql.User{}
	phoneExist, err := users.CheckPhoneExistExcludingUser(in.Phone)
	if err != nil {
		return nil, errors.New("查询手机号失败")
	}
	if phoneExist {
		return nil, errors.New("该手机号已被其他用户注册")
	}

	userProfile := &model_mysql.UserProfile{}

	tx := global.DB.Begin()
	err = userProfile.UpdateUserPhoneByUserId(in.UserId, utils.EncryptPhone(in.Phone))
	if err != nil {
		tx.Rollback()
		return nil, errors.New("修改个人手机号失败")
	}
	err = users.UpdateUserByPhone(in.UserId, in.Phone)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("修改用户手机号失败")
	}
	tx.Commit()
	return &user.UpdateUserPhoneResponse{
		UserId:  in.UserId,
		Message: "修改用户手机号成功",
	}, nil
}
