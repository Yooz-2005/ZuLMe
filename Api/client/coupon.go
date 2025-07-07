package client

import (
	proto_coupon "Srv/coupon_srv/proto_coupon"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CouponClient 优惠券服务客户端
func CouponClient(ctx context.Context, fn func(ctx context.Context, in proto_coupon.CouponServiceClient) (interface{}, error)) (interface{}, error) {
	// 连接优惠券微服务
	conn, err := grpc.Dial("localhost:8006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("连接优惠券服务失败: %v", err)
		return nil, fmt.Errorf("连接优惠券服务失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := proto_coupon.NewCouponServiceClient(conn)

	// 调用服务方法
	return fn(ctx, client)
}
