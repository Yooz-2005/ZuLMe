package grpc_merchant

import (
	"google.golang.org/grpc"
	"merchant_srv/internal"
)

// RegisterMerchantServices 负责将Merchant服务注册到gRPC服务器
func RegisterMerchantServices(s *grpc.Server) {
	internal.RegisterMerchantServer(s)
}
