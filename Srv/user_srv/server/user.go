package server

import (
	"context"
	"user_srv/internal/logic"
	user "user_srv/proto_user"
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

func (s ServerUser) RealName(ctx context.Context, in *user.RealNameRequest) (*user.RealNameResponse, error) {
	name, err := logic.RealName(in)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (s ServerUser) CollectVehicle(ctx context.Context, in *user.CollectVehicleRequest) (*user.CollectVehicleResponse, error) {
	collect, err := logic.CollectVehicle(in)
	if err != nil {
		return nil, err
	}
	return collect, nil
}

func (s ServerUser) CollectVehicleList(ctx context.Context, in *user.CollectVehicleListRequest) (*user.CollectVehicleListResponse, error) {
	collect, err := logic.CollectVehicleList(in)
	if err != nil {
		return nil, err
	}
	return collect, nil
}

// CalculateDistance 计算用户到商家的距离
func (s ServerUser) CalculateDistance(ctx context.Context, in *user.CalculateDistanceRequest) (*user.CalculateDistanceResponse, error) {
	distance, err := logic.CalculateDistance(in)
	if err != nil {
		return nil, err
	}
	return distance, nil
}
