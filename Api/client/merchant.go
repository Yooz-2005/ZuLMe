package client

import (
	merchant "ZuLMe/ZuLMe/Srv/merchant_srv/proto_merchant"
	"context"
	"google.golang.org/grpc"
)

type HandlerMerchant func(ctx context.Context, in merchant.MerchantServiceClient) (interface{}, error)

func MerchantClient(ctx context.Context, handlerMerchant HandlerMerchant) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := merchant.NewMerchantServiceClient(dial)
	res, err := handlerMerchant(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
