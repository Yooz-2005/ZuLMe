package router

import (
	"Api/trigger"
	"Common/pkg"
	"github.com/gin-gonic/gin"
)

// CouponRouter 优惠券路由
func CouponRouter(r *gin.Engine) {
	// 优惠券相关路由组
	couponGroup := r.Group("/coupon")
	{
		// 需要认证的路由
		authGroup := couponGroup.Group("")
		// 发放优惠券 (管理员功能)
		authGroup.POST("/grant", trigger.GrantCouponHandler)
		{

			authGroup.Use(pkg.JWTAuth("2209"))
			// 获取用户优惠券列表
			authGroup.GET("/list", trigger.GetUserCouponsHandler)

			// 获取用户可用优惠券
			authGroup.GET("/available", trigger.GetAvailableCouponsHandler)

			// 验证优惠券
			authGroup.POST("/validate", trigger.ValidateCouponHandler)

			// 使用优惠券
			authGroup.POST("/use", trigger.UseCouponHandler)

			// 根据ID获取优惠券详情
			authGroup.GET("/:coupon_id", trigger.GetCouponByIdHandler)
		}
	}
}
