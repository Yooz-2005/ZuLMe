package internal

import (
	"google.golang.org/grpc"
	merchant "merchant_srv/proto_merchant"
	"merchant_srv/server"
)

func RegisterMerchantServer(srv *grpc.Server) {
	merchant.RegisterMerchantServer(srv, server.ServerMerchant{})
}
