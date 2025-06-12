package internal

import (
	"google.golang.org/grpc"
	order "order_srv/proto_order"
	"order_srv/server"
)

func RegisterOrderServer(ser *grpc.Server) {
	order.RegisterOrderServer(ser, server.ServerOrder{})
}
