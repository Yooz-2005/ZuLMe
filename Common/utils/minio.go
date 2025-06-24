package utils

import (
	"Common/global"
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
)

// ValidateMinioClient 驗證 MinIO 客戶端
func ValidateMinioClient() error {
	if global.Minio == nil {
		return fmt.Errorf("MinIO 客戶端未初始化")
	}
	return nil
}

// EnsureBucketExists 確保 bucket 存在
func EnsureBucketExists(ctx context.Context, bucketName string) error {
	client := global.Minio
	if client == nil {
		return fmt.Errorf("MinIO 客戶端未初始化")
	}

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("檢查 bucket 是否存在失敗: %v", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("創建 bucket 失敗: %v", err)
		}
	}

	return nil
}

// GeneratePresignedUrl 生成預簽名 URL
func GeneratePresignedUrl(ctx context.Context, bucket, objectName string, expires time.Duration) (string, error) {
	client := global.Minio
	if client == nil {
		return "", fmt.Errorf("MinIO 客戶端未初始化")
	}

	// 確保 bucket 存在
	err := EnsureBucketExists(ctx, bucket)
	if err != nil {
		return "", err
	}

	// 生成預簽名 URL
	url, err := client.PresignedPutObject(ctx, bucket, objectName, expires)
	if err != nil {
		return "", fmt.Errorf("生成預簽名 URL 失敗: %v", err)
	}

	return url.String(), nil
}

// GetFileUrl 獲取文件訪問 URL
func GetFileUrl(bucket, objectName string) string {
	return fmt.Sprintf("http://14.103.149.192:9000/%s/%s", bucket, objectName)
}

// SetBucketPublic 設置存儲桶為公開訪問
func SetBucketPublic(ctx context.Context, bucketName string) error {
	client := global.Minio
	if client == nil {
		return fmt.Errorf("MinIO 客戶端未初始化")
	}

	// 確保 bucket 存在
	err := EnsureBucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	// 設置公開讀取策略
	publicPolicy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": "*",
				"Action": "s3:GetObject",
				"Resource": "arn:aws:s3:::%s/*"
			}
		]
	}`, bucketName)

	err = client.SetBucketPolicy(ctx, bucketName, publicPolicy)
	if err != nil {
		return fmt.Errorf("設置存儲桶公開策略失敗: %v", err)
	}

	fmt.Printf("存儲桶 %s 已設置為公開訪問\n", bucketName)
	return nil
}
