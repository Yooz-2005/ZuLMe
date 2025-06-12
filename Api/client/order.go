package client

import (
	"context"
	"google.golang.org/grpc"
	order "order_srv/proto_order"
)

type HandlerOrder func(ctx context.Context, in order.OrderClient) (interface{}, error)

func OrderClient(ctx context.Context, handlerOrder HandlerOrder) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8004", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := order.NewOrderClient(dial)
	res, err := handlerOrder(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
