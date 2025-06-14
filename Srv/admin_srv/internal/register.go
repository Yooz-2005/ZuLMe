package internal

import (
	admin "ZuLMe/ZuLMe/Srv/admin_srv/proto_admin"
	"ZuLMe/ZuLMe/Srv/admin_srv/server"
	"google.golang.org/grpc"
)

func RegisterAdminServer(ser *grpc.Server) {
	admin.RegisterAdminServer(ser, server.ServerAdmin{})
}
