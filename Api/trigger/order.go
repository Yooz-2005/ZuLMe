package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"fmt"
	"strconv"

	order "order_srv/proto_order"

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

	// 转换支付方式字符串为数字
	var paymentMethod int32
	switch req.PaymentMethod {
	case "alipay":
		paymentMethod = 1
	case "wechat":
		paymentMethod = 2
	default:
		response.ResponseError400(c, "不支持的支付方式")
		return
	}

	// 调用订单微服务创建订单
	createRes, err := handler.CreateOrderFromReservation(c, &order.CreateOrderFromReservationRequest{
		ReservationId:       req.ReservationID,
		UserId:              int64(userID),
		PickupLocationId:    req.PickupLocationID,
		ReturnLocationId:    req.ReturnLocationID,
		Notes:               req.Notes,
		PaymentMethod:       paymentMethod,
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

	// 转换状态字符串为数字
	var status int32
	switch req.Status {
	case "pending":
		status = 1 // 待支付
	case "paid":
		status = 2 // 已支付
	case "confirmed":
		status = 3 // 已确认
	case "in_progress":
		status = 4 // 进行中
	case "completed":
		status = 5 // 已完成
	case "cancelled":
		status = 6 // 已取消
	default:
		response.ResponseError400(c, "不支持的订单状态")
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
		Status:  status,
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

	// 解析表单数据
	if err := c.Request.ParseForm(); err != nil {
		fmt.Printf("解析表单数据失败: %v\n", err)
		c.String(400, "fail")
		return
	}

	// 获取支付宝通知参数
	params := make(map[string]string)
	for key, values := range c.Request.Form {
		if len(values) > 0 {
			params[key] = values[0]
			fmt.Printf("参数 %s: %s\n", key, values[0])
		}
	}

	// 验证必要参数
	outTradeNo := params["out_trade_no"]
	tradeNo := params["trade_no"]
	tradeStatus := params["trade_status"]
	totalAmount := params["total_amount"]

	if outTradeNo == "" || tradeStatus == "" {
		fmt.Printf("缺少必要参数: out_trade_no=%s, trade_status=%s\n", outTradeNo, tradeStatus)
		c.String(400, "fail")
		return
	}

	fmt.Printf("订单号: %s, 交易号: %s, 状态: %s, 金额: %s\n",
		outTradeNo, tradeNo, tradeStatus, totalAmount)

	// 直接调用logic层处理支付回调，不通过gRPC
	var orderStatus int32
	switch tradeStatus {
	case "TRADE_SUCCESS":
		orderStatus = 2 // 已支付
	case "TRADE_FINISHED":
		orderStatus = 5 // 已完成
	case "TRADE_CLOSED":
		orderStatus = 6 // 已取消
	default:
		fmt.Printf("未处理的交易状态: %s\n", tradeStatus)
		c.String(200, "success") // 即使状态未处理也返回成功，避免支付宝重复通知
		return
	}

	// 调用logic层更新订单状态
	updateRes, err := handler.UpdateOrderStatus(c, &order.UpdateOrderStatusRequest{
		OrderSn: outTradeNo,
		Status:  orderStatus,
	})
	if err != nil {
		fmt.Printf("更新订单状态失败: %v\n", err)
		c.String(500, "fail")
		return
	}

	if updateRes.Code != 200 {
		fmt.Printf("更新订单状态失败: %s\n", updateRes.Message)
		c.String(500, "fail")
		return
	}

	fmt.Printf("支付回调处理成功: %s\n", updateRes.Message)
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

// CancelOrderHandler 取消订单处理器
func CancelOrderHandler(c *gin.Context) {
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

	// 获取取消原因（可选）
	var req struct {
		Reason string `json:"reason"`
	}
	c.ShouldBind(&req)

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

	// 调用订单微服务取消订单
	cancelRes, err := handler.CancelOrder(c, &order.CancelOrderRequest{
		OrderId: orderID,
		UserId:  int64(userID),
		Reason:  req.Reason,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if cancelRes.Code != 200 {
		response.ResponseError400(c, cancelRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": cancelRes.Message,
	})
}

// GetUserOrderListHandler 获取用户订单列表处理器
func GetUserOrderListHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// 获取查询参数
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	status := c.Query("status")                // 可选的状态筛选
	paymentStatus := c.Query("payment_status") // 可选的支付状态筛选

	// 转换参数
	pageInt := 1
	pageSizeInt := 10
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageInt = p
	}
	if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
		pageSizeInt = ps
	}

	// 调用订单微服务获取用户订单列表
	listRes, err := handler.GetUserOrderList(c, &order.GetUserOrderListRequest{
		UserId:        int64(userID),
		Page:          int32(pageInt),
		PageSize:      int32(pageSizeInt),
		Status:        status,
		PaymentStatus: paymentStatus,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if listRes.Code != 200 {
		response.ResponseError400(c, listRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":   listRes.Message,
		"orders":    listRes.Orders,
		"total":     listRes.Total,
		"page":      pageInt,
		"page_size": pageSizeInt,
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
