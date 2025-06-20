package main

import (
	"Common/appconfig"
	"Common/initialize"
	"user_srv/grpc_user"
	"user_srv/internal"

	"google.golang.org/grpc"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	grpc_user.RegisterUserGrpc(func(server *grpc.Server) {
		internal.RegisterUserServer(server)
	})
}
