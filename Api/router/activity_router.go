package router

import (
	"Api/trigger"
	"Common/pkg"
	"github.com/gin-gonic/gin"
)

// ActivityRouter 活动管理路由
func ActivityRouter(r *gin.Engine) {
	// 活动管理路由组 - 需要管理员权限
	activityGroup := r.Group("/admin/activity")
	{
		// 管理员认证中间件
		//activityGroup.Use(pkg.JWTAuth("admin"))
		{
			// 活动CRUD操作
			activityGroup.POST("/create", trigger.CreateActivityHandler)   // 创建活动
			activityGroup.PUT("/update", trigger.UpdateActivityHandler)    // 更新活动
			activityGroup.GET("/list", trigger.GetActivityListHandler)     // 获取活动列表
			activityGroup.GET("/:id", trigger.GetActivityDetailHandler)    // 获取活动详情
			activityGroup.DELETE("/delete", trigger.DeleteActivityHandler) // 删除活动

			// 优惠券发放操作
			activityGroup.POST("/grant/batch", trigger.BatchGrantCouponHandler)           // 批量发放优惠券
			activityGroup.POST("/grant/condition", trigger.GrantCouponByConditionHandler) // 按条件发放优惠券
		}
	}

	// 公开的活动查询接口（用户可查看）
	publicActivityGroup := r.Group("/activity")
	{
		// 需要用户认证
		publicActivityGroup.Use(pkg.JWTAuth("2209"))
		{
			publicActivityGroup.GET("/list", trigger.GetPublicActivityListHandler)  // 获取公开活动列表
			publicActivityGroup.GET("/:id", trigger.GetPublicActivityDetailHandler) // 获取公开活动详情
		}
	}
}
