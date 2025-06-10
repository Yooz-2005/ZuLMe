package grpc_admin

import (
	"admin_srv/internal"
	"google.golang.org/grpc"
)

// RegisterAdminServices 负责将Admin服务注册到gRPC服务器
func RegisterAdminServices(s *grpc.Server) {
	internal.RegisterAdminServer(s)
}
