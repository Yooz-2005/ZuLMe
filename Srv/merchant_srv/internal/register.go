package internal

import (
	"google.golang.org/grpc"
	merchant "merchant_srv/proto_merchant"
	"merchant_srv/server"
)

func RegisterMerchantServer(ser *grpc.Server) {
	merchant.RegisterMerchantServiceServer(ser, server.ServerMerchant{})
}
