package internal

import (
	"google.golang.org/grpc"
	user "user_srv/proto_user"
	"user_srv/server"
)

func RegisterUserServer(ser *grpc.Server) {
	user.RegisterUserServer(ser, server.ServerUser{})
}
