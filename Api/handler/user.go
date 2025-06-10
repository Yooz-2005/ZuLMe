package handler

import (
	"Api/client"
	"context"
	user "user_srv/proto_user"
)

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
