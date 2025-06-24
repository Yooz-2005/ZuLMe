package server

import (
	"context"
	"minio_srv/internal/logic"
	minio "minio_srv/proto_minio"
)

type ServerMinio struct {
	minio.UnimplementedMinioServer
}

func (s *ServerMinio) GetPresignedUrl(ctx context.Context, in *minio.GetPresignedUrlRequest) (*minio.GetPresignedUrlResponse, error) {
	return logic.GetPresignedUrl(ctx, in)
}
