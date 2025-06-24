package logic

import (
	"Common/utils"
	"context"
	"time"

	minio "minio_srv/proto_minio"
)

// GetPresignedUrl 获取预签名URL
func GetPresignedUrl(ctx context.Context, in *minio.GetPresignedUrlRequest) (*minio.GetPresignedUrlResponse, error) {
	// 参数验证
	if in.Bucket == "" {
		return &minio.GetPresignedUrlResponse{
			Success: false,
			Message: "存储桶名称不能为空",
		}, nil
	}
	if in.ObjectName == "" {
		return &minio.GetPresignedUrlResponse{
			Success: false,
			Message: "对象名称不能为空",
		}, nil
	}

	// 限制最大有效期
	expires := time.Duration(in.Expires) * time.Second
	if expires > 7*24*time.Hour {
		expires = 7 * 24 * time.Hour
	}

	// 驗證 MinIO 客戶端
	if err := utils.ValidateMinioClient(); err != nil {
		return &minio.GetPresignedUrlResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// 生成預簽名 URL
	url, err := utils.GeneratePresignedUrl(ctx, in.Bucket, in.ObjectName, expires)
	if err != nil {
		return &minio.GetPresignedUrlResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// 計算過期時間
	expiresAt := time.Now().Add(expires).Format(time.RFC3339)

	return &minio.GetPresignedUrlResponse{
		Success:   true,
		Url:       url,
		Method:    "PUT",
		ExpiresAt: expiresAt,
		Message:   "获取预签名URL成功",
	}, nil
}
