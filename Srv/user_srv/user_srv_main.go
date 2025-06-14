package main

import (
	"ZuLMe/ZuLMe/Common/appconfig"
	"ZuLMe/ZuLMe/Common/initialize"
	"ZuLMe/ZuLMe/Srv/user_srv/grpc_user"
	"ZuLMe/ZuLMe/Srv/user_srv/internal"
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
