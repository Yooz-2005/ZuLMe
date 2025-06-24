package client

import (
	"context"
	"google.golang.org/grpc"
	minio "minio_srv/proto_minio"
)

type HandlerMinio func(ctx context.Context, in minio.MinioClient) (interface{}, error)

func MinioClient(ctx context.Context, handlerMinio HandlerMinio) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8007", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := minio.NewMinioClient(dial)
	res, err := handlerMinio(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil

}
