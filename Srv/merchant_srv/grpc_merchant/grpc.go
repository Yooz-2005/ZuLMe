package grpc_merchant

import (
	"ZuLMe/ZuLMe/Srv/merchant_srv/internal"
	"google.golang.org/grpc"
)

// RegisterMerchantServices 负责将Merchant服务注册到gRPC服务器
func RegisterMerchantServices(s *grpc.Server) {
	internal.RegisterMerchantServer(s)
}
