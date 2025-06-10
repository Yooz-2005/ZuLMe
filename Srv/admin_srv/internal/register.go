package internal

import (
	admin "admin_srv/proto_admin"
	server2 "admin_srv/server"
	"google.golang.org/grpc"
)

func RegisterAdminServer(ser *grpc.Server) {
	admin.RegisterAdminServer(ser, server2.ServerAdmin{})
}
