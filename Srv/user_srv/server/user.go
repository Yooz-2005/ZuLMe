package server

import (
	"ZuLMe/ZuLMe/Srv/user_srv/internal/logic"
	user "ZuLMe/ZuLMe/Srv/user_srv/proto_user"
	"context"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

func (s ServerUser) SendCode(ctx context.Context, in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	code, err := logic.SendCode(in)
	if err != nil {
		return nil, err
	}
	return code, nil
}

func (s ServerUser) UserRegister(ctx context.Context, in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	register, err := logic.UserRegister(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

func (s ServerUser) UpdateUserProfile(ctx context.Context, in *user.UpdateUserProfileRequest) (*user.UpdateUserProfileResponse, error) {
	profile, err := logic.SetUserProfile(in)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s ServerUser) UpdateUserPhone(ctx context.Context, in *user.UpdateUserPhoneRequest) (*user.UpdateUserPhoneResponse, error) {
	phone, err := logic.UpdateUserPhone(in)
	if err != nil {
		return nil, err
	}
	return phone, nil
}
