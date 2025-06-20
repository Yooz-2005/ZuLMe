package handler

import (
	"Api/client"
	"context"
	user "user_srv/proto_user"
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

// todo用户实名认证
func RealName(ctx context.Context, req *user.RealNameRequest) (*user.RealNameResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		name, err := in.RealName(ctx, req)
		if err != nil {
			return nil, err
		}
		return name, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.RealNameResponse), nil
}

// todo收藏取消收藏车辆
func CollectVehicle(ctx context.Context, req *user.CollectVehicleRequest) (*user.CollectVehicleResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		collect, err := in.CollectVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return collect, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.CollectVehicleResponse), nil
}

// todo收藏车辆列表
func CollectVehicleList(ctx context.Context, req *user.CollectVehicleListRequest) (*user.CollectVehicleListResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		collect, err := in.CollectVehicleList(ctx, req)
		if err != nil {
			return nil, err
		}
		return collect, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.CollectVehicleListResponse), nil
}
