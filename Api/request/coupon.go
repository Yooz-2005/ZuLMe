package request

// GrantCouponRequest 发放优惠券请求
type GrantCouponRequest struct {
	UserID       uint64 `json:"user_id" form:"user_id" binding:"required"`             // 用户ID
	ActivityCode string `json:"activity_code" form:"activity_code" binding:"required"` // 活动编码
	Source       string `json:"source" form:"source"`                                  // 发放来源，可选
}

// GetUserCouponsRequest 获取用户优惠券列表请求
type GetUserCouponsRequest struct {
	Status   int `json:"status" form:"status"`       // 状态筛选 0-全部 1-未使用 2-已使用 3-已过期
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 每页数量
}

// GetAvailableCouponsRequest 获取可用优惠券请求
type GetAvailableCouponsRequest struct {
	OrderAmount float64 `json:"order_amount" form:"order_amount" binding:"required"` // 订单金额
}

// ValidateCouponRequest 验证优惠券请求
type ValidateCouponRequest struct {
	CouponID    uint64  `json:"coupon_id" form:"coupon_id" binding:"required"`       // 优惠券ID
	OrderAmount float64 `json:"order_amount" form:"order_amount" binding:"required"` // 订单金额
}

// UseCouponRequest 使用优惠券请求
type UseCouponRequest struct {
	CouponID       uint64  `json:"coupon_id" form:"coupon_id" binding:"required"`             // 优惠券ID
	OrderID        uint64  `json:"order_id" form:"order_id" binding:"required"`               // 订单ID
	OrderSn        string  `json:"order_sn" form:"order_sn" binding:"required"`               // 订单号
	OriginalAmount float64 `json:"original_amount" form:"original_amount" binding:"required"` // 原始金额
	DiscountAmount float64 `json:"discount_amount" form:"discount_amount" binding:"required"` // 优惠金额
}
