package handler

import (
	"Api/client"
	"context"
	minio "minio_srv/proto_minio"
)

// GetPresignedUrl 獲取預簽名 URL
func GetPresignedUrl(ctx context.Context, req *minio.GetPresignedUrlRequest) (res *minio.GetPresignedUrlResponse, err error) {
	minioClient, err := client.MinioClient(ctx, func(ctx context.Context, in minio.MinioClient) (interface{}, error) {
		url, err := in.GetPresignedUrl(ctx, req)
		if err != nil {
			return nil, err
		}
		return url, nil
	})
	if err != nil {
		return nil, err
	}
	return minioClient.(*minio.GetPresignedUrlResponse), nil
}
