package handler

import (
	"ZuLMe/ZuLMe/Api/client"
	user "ZuLMe/ZuLMe/Srv/user_srv/proto_user"
	"context"
)

// todo用户注册登录
func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		register, err := in.UserRegister(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserRegisterResponse), nil
}

// todo发送验证码
func SendCode(ctx context.Context, req *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		code, err := in.SendCode(ctx, req)
		if err != nil {
			return nil, err
		}
		return code, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.SendCodeResponse), nil
}

// todo修改个人信息
func UpdateUserProfile(ctx context.Context, req *user.UpdateUserProfileRequest) (*user.UpdateUserProfileResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		profile, err := in.UpdateUserProfile(ctx, req)
		if err != nil {
			return nil, err
		}
		return profile, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UpdateUserProfileResponse), nil
}

// todo修改用户手机号
func UpdateUserPhone(ctx context.Context, req *user.UpdateUserPhoneRequest) (*user.UpdateUserPhoneResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		phone, err := in.UpdateUserPhone(ctx, req)
		if err != nil {
			return nil, err
		}
		return phone, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UpdateUserPhoneResponse), nil
}
