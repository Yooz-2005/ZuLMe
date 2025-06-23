package internal

import (
	"google.golang.org/grpc"
	minio "minio_srv/proto_minio"
	"minio_srv/server"
)

func RegisterMinioServer(ser *grpc.Server) {
	minio.RegisterMinioServer(ser, &server.ServerMinio{})
}
