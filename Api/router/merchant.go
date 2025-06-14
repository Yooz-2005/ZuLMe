package router

import (
	"Api/trigger"
	"github.com/gin-gonic/gin"
)

// RegisterMerchantRoutes 注册商户相关的路由
func RegisterMerchantRoutes(r *gin.Engine) {
	merchantGroup := r.Group("/merchant")
	{
		merchantGroup.POST("/register", trigger.MerchantRegisterHandler)
		merchantGroup.POST("/login", trigger.MerchantLoginHandler)
	}
}
