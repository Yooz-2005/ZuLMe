package router

import (
	"Api/trigger"
	jwt "Common/pkg"

	"github.com/gin-gonic/gin"
)

// LoadOrder 加载订单路由（与主文件中的调用保持一致）
func LoadOrder(r *gin.Engine) {
	InitOrderRouter(r)
}

// InitOrderRouter 初始化订单路由
func InitOrderRouter(r *gin.Engine) {
	// 用户订单操作路由（需要用户认证）
	userOrderGroup := r.Group("/order")
	{
		userOrderGroup.Use(jwt.JWTAuth("2209"))                                                    // 用户认证
		userOrderGroup.POST("/create-from-reservation", trigger.CreateOrderFromReservationHandler) // 基于预订创建订单
		userOrderGroup.GET("/detail/:order_id", trigger.GetOrderDetailHandler)                     // 获取订单详情
		userOrderGroup.GET("/detail-by-sn/:order_sn", trigger.GetOrderDetailBySnHandler)           // 根据订单号获取详情
		userOrderGroup.PUT("/status/:order_id", trigger.UpdateOrderStatusHandler)                  // 更新订单状态
		userOrderGroup.POST("/cancel/:order_id", trigger.CancelOrderHandler)                       // 取消订单
		userOrderGroup.GET("/list", trigger.GetUserOrderListHandler)                               // 获取用户订单列表
		userOrderGroup.GET("/check-unpaid", trigger.CheckUserUnpaidOrderHandler)                   // 检查用户未支付订单
	}

	// 商家订单管理路由（需要商家认证）
	merchantOrderGroup := r.Group("/merchant/order")
	{
		merchantOrderGroup.Use(jwt.JWTAuth("merchant"))                                 // 商家认证
		merchantOrderGroup.GET("/list", trigger.GetMerchantOrderListHandler)            // 获取商家订单列表
		merchantOrderGroup.PUT("/status/:order_id", trigger.MerchantUpdateOrderHandler) // 商家更新订单状态
		merchantOrderGroup.GET("/statistics", trigger.GetOrderStatisticsHandler)        // 获取订单统计
	}

	// 支付相关路由（公开，供支付宝回调）
	paymentGroup := r.Group("/payment")
	{
		paymentGroup.POST("/alipay/notify", trigger.AlipayNotifyHandler) // 支付宝异步通知
		paymentGroup.GET("/alipay/return", trigger.AlipayReturnHandler)  // 支付宝同步返回
		paymentGroup.GET("/test", trigger.TestCallbackHandler)           // 测试回调接口
		paymentGroup.POST("/test", trigger.TestCallbackHandler)          // 测试回调接口
	}

	// 测试相关路由（仅用于开发测试）
	testGroup := r.Group("/test")
	{
		testGroup.PUT("/payment/status/:order_sn", trigger.ManualUpdatePaymentStatusHandler) // 手动更新支付状态
	}

	// 订单管理路由（内部调用，可选择是否需要认证）
	adminOrderGroup := r.Group("/admin/order")
	{
		// adminOrderGroup.Use(jwt.JWTAuth("admin")) // 管理员认证（可选）
		adminOrderGroup.GET("/list", trigger.GetAllOrderListHandler)     // 获取所有订单列表
		adminOrderGroup.GET("/export", trigger.ExportOrderDataHandler)   // 导出订单数据
		adminOrderGroup.DELETE("/:order_id", trigger.DeleteOrderHandler) // 删除订单（软删除）
	}
}
