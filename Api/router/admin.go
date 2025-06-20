package router

import (
	"Api/trigger"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	{
		// 商户审核
		adminGroup.POST("/merchant/approve", trigger.AdminMerchantApproveHandler)
		// 商户编辑
		adminGroup.POST("/merchant/update", trigger.AdminMerchantUpdateHandler)
		// 商户删除
		adminGroup.POST("/merchant/delete", trigger.AdminMerchantDeleteHandler)
		// 商户列表
		adminGroup.GET("/merchant/list", trigger.AdminMerchantListHandler)
		// 商户详情
		adminGroup.GET("/merchant/:id", trigger.AdminMerchantDetailHandler)
	}
}
