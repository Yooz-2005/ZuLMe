package grpc_minio

import (
	"google.golang.org/grpc"
	"minio_srv/internal"
)

func RegisterMinioServices(s *grpc.Server) {
	internal.RegisterMinioServer(s)
}
