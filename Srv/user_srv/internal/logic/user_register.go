package logic

import user "user_srv/proto_user"

func UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	return &user.UserRegisterResponse{
		UserId: "okokok",
	}, nil
}
