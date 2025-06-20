package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"Common/payment"
	"fmt"
	order "order_srv/proto_order"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrderFromReservationHandler 基于预订创建订单处理器
func CreateOrderFromReservationHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	var req request.CreateOrderFromReservationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 调用订单微服务创建订单
	createRes, err := handler.CreateOrderFromReservation(c, &order.CreateOrderFromReservationRequest{
		ReservationId:       req.ReservationID,
		UserId:              int64(userID),
		PickupLocationId:    req.PickupLocationID,
		ReturnLocationId:    req.ReturnLocationID,
		Notes:               req.Notes,
		PaymentMethod:       req.PaymentMethod,
		ExpectedTotalAmount: req.ExpectedTotalAmount,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if createRes.Code != 200 {
		response.ResponseError400(c, createRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      createRes.Message,
		"order_id":     createRes.OrderId,
		"order_sn":     createRes.OrderSn,
		"total_amount": createRes.TotalAmount,
		"payment_url":  createRes.PaymentUrl,
		"status":       "待支付",
	})
}

// GetOrderDetailHandler 获取订单详情处理器
func GetOrderDetailHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// 获取订单ID
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "订单ID格式错误")
		return
	}

	// 调用订单微服务获取订单详情
	getRes, err := handler.GetOrder(c, &order.GetOrderRequest{
		OrderId: orderID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if getRes.Code != 200 {
		response.ResponseError400(c, getRes.Message)
		return
	}

	// 验证订单归属（用户只能查看自己的订单）
	if getRes.Order.UserId != int64(userID) {
		response.ResponseError400(c, "无权查看此订单")
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": getRes.Message,
		"order":   getRes.Order,
	})
}

// GetOrderDetailBySnHandler 根据订单号获取订单详情
func GetOrderDetailBySnHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// 获取订单号
	orderSn := c.Param("order_sn")
	if orderSn == "" {
		response.ResponseError400(c, "订单号不能为空")
		return
	}

	// 调用订单微服务获取订单详情
	getRes, err := handler.GetOrder(c, &order.GetOrderRequest{
		OrderSn: orderSn,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if getRes.Code != 200 {
		response.ResponseError400(c, getRes.Message)
		return
	}

	// 验证订单归属
	if getRes.Order.UserId != int64(userID) {
		response.ResponseError400(c, "无权查看此订单")
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": getRes.Message,
		"order":   getRes.Order,
	})
}

// UpdateOrderStatusHandler 更新订单状态处理器
func UpdateOrderStatusHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// 获取订单ID
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseInt(orderIDStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "订单ID格式错误")
		return
	}

	var req request.UpdateOrderStatusRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 先验证订单归属
	getRes, err := handler.GetOrder(c, &order.GetOrderRequest{
		OrderId: orderID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if getRes.Code != 200 {
		response.ResponseError400(c, getRes.Message)
		return
	}

	if getRes.Order.UserId != int64(userID) {
		response.ResponseError400(c, "无权操作此订单")
		return
	}

	// 调用订单微服务更新状态
	updateRes, err := handler.UpdateOrderStatus(c, &order.UpdateOrderStatusRequest{
		OrderId: orderID,
		Status:  req.Status,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if updateRes.Code != 200 {
		response.ResponseError400(c, updateRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": updateRes.Message,
	})
}

// AlipayNotifyHandler 支付宝异步通知处理器
func AlipayNotifyHandler(c *gin.Context) {
	// 记录收到回调的日志
	fmt.Printf("=== 收到支付宝回调通知 ===\n")
	fmt.Printf("请求方法: %s\n", c.Request.Method)
	fmt.Printf("请求URL: %s\n", c.Request.URL.String())
	fmt.Printf("Content-Type: %s\n", c.GetHeader("Content-Type"))

	// 解析表单数据
	if err := c.Request.ParseForm(); err != nil {
		fmt.Printf("解析表单数据失败: %v\n", err)
		c.String(400, "解析表单数据失败")
		return
	}

	// 获取支付宝通知参数
	params := make(map[string]string)
	for key, values := range c.Request.Form {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	// 打印所有参数
	fmt.Printf("回调参数:\n")
	for key, value := range params {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// 验证通知
	alipayService := payment.NewAlipayService()
	if !alipayService.VerifyNotify(params) {
		fmt.Printf("支付宝通知验证失败\n")
		c.String(400, "验证失败")
		return
	}
	fmt.Printf("支付宝通知验证成功\n")

	// 解析通知数据
	result := alipayService.ParseNotify(params)
	fmt.Printf("解析结果: 订单号=%s, 交易号=%s, 状态=%s, 金额=%.2f\n",
		result.OrderSn, result.TradeNo, result.TradeStatus, result.TotalAmount)

	// 调用订单微服务处理支付通知
	fmt.Printf("调用订单微服务处理支付通知...\n")
	notifyRes, err := handler.AlipayNotify(c, &order.AlipayNotifyRequest{
		OutTradeNo:  result.OrderSn,
		TradeNo:     result.TradeNo,
		TradeStatus: result.TradeStatus,
		TotalAmount: fmt.Sprintf("%.2f", result.TotalAmount),
		GmtPayment:  result.PayTime,
	})
	if err != nil {
		fmt.Printf("调用订单微服务失败: %v\n", err)
		c.String(500, "处理支付通知失败")
		return
	}

	fmt.Printf("订单微服务响应: Code=%d, Message=%s\n", notifyRes.Code, notifyRes.Message)

	if notifyRes.Code != 200 {
		fmt.Printf("订单微服务处理失败: %s\n", notifyRes.Message)
		c.String(500, notifyRes.Message)
		return
	}

	fmt.Printf("支付回调处理成功\n")
	// 返回成功响应给支付宝
	c.String(200, "success")
}

// AlipayReturnHandler 支付宝同步返回处理器
func AlipayReturnHandler(c *gin.Context) {
	// 获取返回参数
	orderSn := c.Query("out_trade_no")
	tradeNo := c.Query("trade_no")

	if orderSn == "" {
		response.ResponseError400(c, "订单号不能为空")
		return
	}

	// 重定向到前端支付成功页面
	redirectURL := fmt.Sprintf("http://localhost:3000/payment/success?order_sn=%s&trade_no=%s", orderSn, tradeNo)
	c.Redirect(302, redirectURL)
}

// GetUserOrderListHandler 获取用户订单列表处理器
func GetUserOrderListHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// TODO: 实现获取用户订单列表的逻辑
	// 这里需要在订单微服务中添加相应的接口

	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
		"orders":  []interface{}{},
	})
}

// GetMerchantOrderListHandler 获取商家订单列表处理器
func GetMerchantOrderListHandler(c *gin.Context) {
	// 从JWT中获取商家ID
	merchantID := c.GetUint("userId")
	if merchantID == 0 {
		response.ResponseError400(c, "商家ID不能为空")
		return
	}

	// TODO: 实现获取商家订单列表的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
		"orders":  []interface{}{},
	})
}

// MerchantUpdateOrderHandler 商家更新订单状态处理器
func MerchantUpdateOrderHandler(c *gin.Context) {
	// TODO: 实现商家更新订单状态的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
	})
}

// GetOrderStatisticsHandler 获取订单统计处理器
func GetOrderStatisticsHandler(c *gin.Context) {
	// TODO: 实现订单统计的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
		"statistics": gin.H{
			"total_orders":     0,
			"pending_orders":   0,
			"paid_orders":      0,
			"completed_orders": 0,
		},
	})
}

// WechatNotifyHandler 微信支付异步通知处理器（预留）
func WechatNotifyHandler(c *gin.Context) {
	// TODO: 实现微信支付通知处理
	c.String(200, "success")
}

// GetAllOrderListHandler 获取所有订单列表处理器
func GetAllOrderListHandler(c *gin.Context) {
	// TODO: 实现获取所有订单列表的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
		"orders":  []interface{}{},
	})
}

// ExportOrderDataHandler 导出订单数据处理器
func ExportOrderDataHandler(c *gin.Context) {
	// TODO: 实现订单数据导出的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
	})
}

// DeleteOrderHandler 删除订单处理器
func DeleteOrderHandler(c *gin.Context) {
	// TODO: 实现订单软删除的逻辑
	response.ResponseSuccess(c, gin.H{
		"message": "功能开发中",
	})
}
