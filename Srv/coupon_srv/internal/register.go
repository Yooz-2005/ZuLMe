package internal

import (
	"google.golang.org/grpc"
	coupon "coupon_srv/proto_coupon"
	"coupon_srv/server"
)

func RegisterCouponServer(ser *grpc.Server) {
	coupon.RegisterCouponServiceServer(ser, server.ServerCoupon{})
}
