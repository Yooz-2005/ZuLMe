package internal

import (
	merchant "ZuLMe/ZuLMe/Srv/merchant_srv/proto_merchant"
	"ZuLMe/ZuLMe/Srv/merchant_srv/server"
	"google.golang.org/grpc"
)

func RegisterMerchantServer(ser *grpc.Server) {
	merchant.RegisterMerchantServiceServer(ser, server.ServerMerchant{})
}
