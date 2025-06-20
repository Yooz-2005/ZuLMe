package request

// CreateOrderFromReservationRequest 从预订创建订单请求
type CreateOrderFromReservationRequest struct {
	ReservationID       int64   `json:"reservation_id" binding:"required"`       // 预订ID
	PickupLocationID    int64   `json:"pickup_location_id"`                      // 取车地点ID（可选，默认使用车辆所在门店）
	ReturnLocationID    int64   `json:"return_location_id"`                      // 还车地点ID（可选，默认使用取车地点）
	Notes               string  `json:"notes"`                                   // 订单备注
	PaymentMethod       string  `json:"payment_method"`                          // 支付方式（alipay/wechat）
	ExpectedTotalAmount float64 `json:"expected_total_amount"`                   // 预期总金额（用于验证）
}

// UpdateOrderStatusRequest 更新订单状态请求
type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"` // 订单状态
}

// GetOrderListRequest 获取订单列表请求
type GetOrderListRequest struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"page_size" json:"page_size"` // 每页数量
	Status   string `form:"status" json:"status"`     // 订单状态筛选
	OrderBy  string `form:"order_by" json:"order_by"` // 排序字段
}

// CancelOrderRequest 取消订单请求
type CancelOrderRequest struct {
	Reason string `json:"reason"` // 取消原因
}

// RefundOrderRequest 退款请求
type RefundOrderRequest struct {
	RefundAmount float64 `json:"refund_amount" binding:"required"` // 退款金额
	RefundReason string  `json:"refund_reason" binding:"required"` // 退款原因
}
