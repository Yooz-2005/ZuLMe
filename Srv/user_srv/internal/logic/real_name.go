package logic

import (
	"ZuLMe/ZuLMe/Common/pkg"
	user "user_srv/proto_user"
)

func RealName(in *user.RealNameRequest) (*user.RealNameResponse, error) {
	ok := pkg.RealName(in.IdNumber, in.RealName)
	var msg string
	if ok {
		msg = "实名认证成功"
	} else {
		msg = "实名认证失败"
	}
	return &user.RealNameResponse{
		UserId:  in.UserId,
		Message: msg,
	}, nil
}
