package router

import (
	"Api/trigger"
	"Common/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterInvoiceRoutes(r *gin.Engine) {

	// 商家发票管理路由（需要商家认证）
	merchantInvoiceGroup := r.Group("/merchant/invoice")
	{
		merchantInvoiceGroup.Use(pkg.JWTAuth("merchant"))
		{
			merchantInvoiceGroup.POST("/generate", trigger.GenerateInvoice) // 商家直接开发票
		}
	}

	// 用户发票申请路由（需要用户认证）
	userInvoiceGroup := r.Group("/invoice")
	{
		userInvoiceGroup.Use(pkg.JWTAuth("2209")) // 用户认证
		{
			userInvoiceGroup.POST("/apply", trigger.ApplyInvoiceForUser) // 用户申请开发票
		}
	}

}
