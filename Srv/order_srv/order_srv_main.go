package main

import (
	"Common/appconfig"
	"Common/initialize"
	"google.golang.org/grpc"
	"order_srv/grpc_order"
	"order_srv/internal"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	grpc_order.RegisterOrderGrpc(func(server *grpc.Server) {
		internal.RegisterOrderServer(server)
	})
}
