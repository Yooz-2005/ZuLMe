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

		// 商家位置管理接口 (管理员功能)
		locationGroup := merchantGroup.Group("/location")
		{
			locationGroup.POST("/sync", trigger.SyncMerchantLocationsHandler)           // 同步所有商家位置到Redis
			locationGroup.GET("/validate", trigger.ValidateMerchantLocationDataHandler) // 验证位置数据完整性
			locationGroup.POST("/fix", trigger.FixMerchantCoordinatesHandler)           // 修复缺少坐标的商家
			locationGroup.PUT("/:id", trigger.UpdateMerchantLocationHandler)            // 更新单个商家位置
		}
	}

	// 地理编码相关路由
	geocodeGroup := r.Group("/geocode")
	{
		geocodeGroup.POST("/coordinates", trigger.GetCoordinates)
	}
}
