package initialize

import (
	"Common/global"
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioInit minio連接初始化
func MinioInit() {
	cos := Nacos.Minio
	client, err := minio.New("14.103.149.192:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(cos.AccessKeyId, cos.AccessKeySecret, ""),
		Secure: false, // 根据需要设置为 true 或 false
	})
	if err != nil {
		log.Fatalf("无法创建 MinIO 客户端: %v", err)
	}

	// 保存客戶端實例到 global 包
	global.Minio = client

	// 检查连接是否成功
	ok, err := client.BucketExists(context.Background(), cos.Bucket)
	if err != nil {
		log.Fatalf("檢查 bucket 失敗: %v", err)
	}
	if !ok {
		log.Fatalf("bucket %s 不存在", cos.Bucket)
	}

	// 設置存儲桶為公開訪問
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
	}`, cos.Bucket)

	err = client.SetBucketPolicy(context.Background(), cos.Bucket, publicPolicy)
	if err != nil {
		log.Printf("設置存儲桶公開策略失敗: %v", err)
	} else {
		fmt.Printf("存儲桶 %s 已設置為公開訪問\n", cos.Bucket)
	}

	fmt.Printf("minio連接成功， 端點: 14.103.149.192:9000，默認 bucket: %s\n", cos.Bucket)
}
