package initialize

import (
	"Common/global"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
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
	fmt.Printf("minio連接成功， 端點: 14.103.149.192:9000，默認 bucket: %s\n", cos.Bucket)
}
