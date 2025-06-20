package main

import (
	"ZuLMe/ZuLMe/Common/appconfig"
	"ZuLMe/ZuLMe/Common/initialize"
	"google.golang.org/grpc"
	"user_srv/grpc_user"
	"user_srv/internal"
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
