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
		// 公共接口 - 不需要认证的图片上传
		minio.GET("/public/presigned-url", trigger.GetPresignedUrl)

		// 需要认证的接口
		minio.Use(jwt.JWTAuth("2209"))
		{
			minio.GET("/presigned-url", trigger.GetPresignedUrl)
		}

	}
}
