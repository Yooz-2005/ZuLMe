package router

import (
	"Api/trigger"
	jwt "Common/pkg"

	"github.com/gin-gonic/gin"
)

// CommentRouter 评论相关路由
func CommentRouter(router *gin.Engine) {
	// 公开评论路由（无需认证）- 供游客和用户浏览
	publicCommentGroup := router.Group("/comment")
	{
		// 获取评论详情（公开）
		publicCommentGroup.GET("/:comment_id", trigger.GetCommentHandler)

		// 获取车辆评论列表（公开）- 用户选车时查看评价
		publicCommentGroup.GET("/vehicle/:vehicle_id", trigger.GetVehicleCommentsHandler)

		// 获取车辆评论统计（公开）- 显示车辆评分概况
		publicCommentGroup.GET("/stats/:vehicle_id", trigger.GetVehicleStatsHandler)
	}

	// 用户评论操作路由（需要用户认证）
	userCommentGroup := router.Group("/comment")
	{
		userCommentGroup.Use(jwt.JWTAuth("2209")) // 用户认证 - 暂时注释掉用于测试

		// 创建评论（需要用户登录）
		userCommentGroup.POST("/create", trigger.CreateCommentHandler)

		// 获取订单评论（需要用户登录）- 用户查看自己的评论
		userCommentGroup.GET("/order/:order_id", trigger.GetOrderCommentHandler)

		// 获取用户评论列表（需要用户登录）- 个人中心查看自己的评论
		userCommentGroup.GET("/user/:user_id", trigger.GetUserCommentsHandler)

		// 更新评论（需要用户登录）- 用户修改自己的评论
		userCommentGroup.PUT("/:comment_id", trigger.UpdateCommentHandler)

		// 删除评论（需要用户登录）- 用户删除自己的评论
		userCommentGroup.DELETE("/:comment_id", trigger.DeleteCommentHandler)

		// 检查订单是否已评论（需要用户登录）- 用户查看是否已评价
		userCommentGroup.GET("/check/:order_id", trigger.CheckOrderCommentedHandler)
	}

	// 商家评论管理路由（需要商家认证）
	merchantCommentGroup := router.Group("/merchant/comment")
	{
		merchantCommentGroup.Use(jwt.JWTAuth("merchant")) // 商家认证

		// 商家回复评论（需要商家登录）
		merchantCommentGroup.POST("/reply/:comment_id", trigger.ReplyCommentHandler)

		// 获取商家车辆的评论列表（需要商家登录）- 商家管理评论
		merchantCommentGroup.GET("/list", trigger.GetMerchantVehicleCommentsHandler)
	}
}
