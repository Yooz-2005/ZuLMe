package internal

import (
	admin "admin_srv/proto_admin"
	"admin_srv/server"
	"google.golang.org/grpc"
)

func RegisterAdminServer(ser *grpc.Server) {
	admin.RegisterAdminServer(ser, server.ServerAdmin{})
}
