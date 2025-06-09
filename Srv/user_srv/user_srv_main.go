package main

import (
	"Common/appconfig"
	"Common/global"
	"Common/initialize"
	"google.golang.org/grpc"
	"models/model_mysql"
	"user_srv/grpc_user"
	"user_srv/internal"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	global.DB.AutoMigrate(&model_mysql.User{})
	grpc_user.RegisterUserGrpc(func(server *grpc.Server) {
		internal.RegisterUserServer(server)
	})
}
