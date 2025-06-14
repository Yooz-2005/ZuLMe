package internal

import (
	user "ZuLMe/ZuLMe/Srv/user_srv/proto_user"
	"ZuLMe/ZuLMe/Srv/user_srv/server"
	"google.golang.org/grpc"
)

func RegisterUserServer(ser *grpc.Server) {
	user.RegisterUserServer(ser, server.ServerUser{})
}
