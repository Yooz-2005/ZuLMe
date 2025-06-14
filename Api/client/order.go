package client

import (
	order "ZuLMe/ZuLMe/Srv/order_srv/proto_order"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OrderClient 订单服务客户端
func OrderClient(ctx context.Context, call func(context.Context, order.OrderClient) (interface{}, error)) (interface{}, error) {
	// 连接订单服务
	conn, err := grpc.Dial("localhost:9093", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("连接订单服务失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := order.NewOrderClient(conn)

	// 调用服务
	return call(ctx, client)
}
