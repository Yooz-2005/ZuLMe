package router

import (
	"Api/trigger"
	jwt "Common/pkg"
	"github.com/gin-gonic/gin"
)

// RegisterMinioRoutes 注册 MinIO 相关的路由
func RegisterMinioRoutes(r *gin.Engine) {
	minio := r.Group("/minio")
	{
		// 獲取預簽名 URL
		minio.Use(jwt.JWTAuth("2209"))
		{
			minio.GET("/presigned-url", trigger.GetPresignedUrl)
		}

	}
}
