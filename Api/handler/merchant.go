package handler

import (
	"Api/client"
	"context"
	merchant "merchant_srv/proto_merchant"
)

func MerchantRegister(ctx context.Context, req *merchant.MerchantRegisterRequest) (*merchant.MerchantRegisterResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantServiceClient) (interface{}, error) {
		register, err := in.MerchantRegister(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.MerchantRegisterResponse), nil
}

func MerchantLogin(ctx context.Context, req *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	merchantClient, err := client.MerchantClient(ctx, func(ctx context.Context, in merchant.MerchantServiceClient) (interface{}, error) {
		register, err := in.MerchantLogin(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return merchantClient.(*merchant.MerchantLoginResponse), nil
}
