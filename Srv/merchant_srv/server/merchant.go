package server

import (
	"ZuLMe/ZuLMe/Srv/merchant_srv/internal/logic"
	merchant "ZuLMe/ZuLMe/Srv/merchant_srv/proto_merchant"
	"context"
)

type ServerMerchant struct {
	merchant.UnimplementedMerchantServiceServer
}

func (s ServerMerchant) MerchantRegister(ctx context.Context, in *merchant.MerchantRegisterRequest) (*merchant.MerchantRegisterResponse, error) {
	register, err := logic.MerchantRegister(ctx, in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

func (s ServerMerchant) MerchantLogin(ctx context.Context, in *merchant.MerchantLoginRequest) (*merchant.MerchantLoginResponse, error) {
	login, err := logic.MerchantLogin(ctx, in)
	if err != nil {
		return nil, err
	}
	return login, nil
}
