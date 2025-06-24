package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	minio "minio_srv/proto_minio"
)

// GetPresignedUrl获取预签名
func GetPresignedUrl(c *gin.Context) {
	var req request.GetPresignedUrlRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	uploadRes, err := handler.GetPresignedUrl(c, &minio.GetPresignedUrlRequest{
		Bucket:     req.Bucket,
		ObjectName: req.ObjectName,
		Expires:    req.Expires,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, uploadRes)
}
