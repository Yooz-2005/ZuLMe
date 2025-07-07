package request

import (
	"time"
)

// CreateActivityRequest 创建活动请求
type CreateActivityRequest struct {
	ActivityName      string  `json:"activity_name" form:"activity_name" binding:"required"` // 活动名称
	ActivityCode      string  `json:"activity_code" form:"activity_code" binding:"required"` // 活动编码
	ActivityType      string  `json:"activity_type" form:"activity_type" binding:"required"` // 活动类型
	Description       string  `json:"description" form:"description"`                        // 活动描述
	DiscountType      int     `json:"discount_type" form:"discount_type" binding:"required"` // 优惠类型:1-减免金额,2-折扣比例
	DiscountAmount    float64 `json:"discount_amount" form:"discount_amount"`                // 优惠金额
	DiscountRate      float64 `json:"discount_rate" form:"discount_rate"`                    // 折扣比例
	MinOrderAmount    float64 `json:"min_order_amount" form:"min_order_amount"`              // 最低订单金额
	MaxDiscountAmount float64 `json:"max_discount_amount" form:"max_discount_amount"`        // 最大优惠金额
	ValidDays         int     `json:"valid_days" form:"valid_days" binding:"required"`       // 有效期天数
	MaxGrantCount     int     `json:"max_grant_count" form:"max_grant_count"`                // 每人最多发放数量
	TotalGrantLimit   int     `json:"total_grant_limit" form:"total_grant_limit"`            // 总发放数量限制
	StartTime         string  `json:"start_time" form:"start_time" binding:"required"`       // 活动开始时间
	EndTime           string  `json:"end_time" form:"end_time" binding:"required"`           // 活动结束时间
	Status            int     `json:"status" form:"status"`                                  // 状态:1-启用,0-禁用
}

// UpdateActivityRequest 更新活动请求
type UpdateActivityRequest struct {
	ID                uint    `json:"id" form:"id" binding:"required"`                       // 活动ID
	ActivityName      string  `json:"activity_name" form:"activity_name" binding:"required"` // 活动名称
	ActivityType      string  `json:"activity_type" form:"activity_type" binding:"required"` // 活动类型
	Description       string  `json:"description" form:"description"`                        // 活动描述
	DiscountType      int     `json:"discount_type" form:"discount_type" binding:"required"` // 优惠类型
	DiscountAmount    float64 `json:"discount_amount" form:"discount_amount"`                // 优惠金额
	DiscountRate      float64 `json:"discount_rate" form:"discount_rate"`                    // 折扣比例
	MinOrderAmount    float64 `json:"min_order_amount" form:"min_order_amount"`              // 最低订单金额
	MaxDiscountAmount float64 `json:"max_discount_amount" form:"max_discount_amount"`        // 最大优惠金额
	ValidDays         int     `json:"valid_days" form:"valid_days" binding:"required"`       // 有效期天数
	MaxGrantCount     int     `json:"max_grant_count" form:"max_grant_count"`                // 每人最多发放数量
	TotalGrantLimit   int     `json:"total_grant_limit" form:"total_grant_limit"`            // 总发放数量限制
	StartTime         string  `json:"start_time" form:"start_time" binding:"required"`       // 活动开始时间
	EndTime           string  `json:"end_time" form:"end_time" binding:"required"`           // 活动结束时间
	Status            int     `json:"status" form:"status"`                                  // 状态
}

// GetActivityListRequest 获取活动列表请求
type GetActivityListRequest struct {
	Page     int    `json:"page" form:"page"`           // 页码
	PageSize int    `json:"page_size" form:"page_size"` // 每页数量
	Status   *int   `json:"status" form:"status"`       // 状态筛选
	Keyword  string `json:"keyword" form:"keyword"`     // 关键词搜索
}

// BatchGrantCouponRequest 批量发放优惠券请求
type BatchGrantCouponRequest struct {
	ActivityCode string   `json:"activity_code" form:"activity_code" binding:"required"` // 活动编码
	UserIDs      []uint64 `json:"user_ids" form:"user_ids" binding:"required"`           // 用户ID列表
	Source       string   `json:"source" form:"source"`                                  // 发放来源
}

// GrantCouponByConditionRequest 按条件发放优惠券请求
type GrantCouponByConditionRequest struct {
	ActivityCode   string    `json:"activity_code" form:"activity_code" binding:"required"` // 活动编码
	UserCondition  string    `json:"user_condition" form:"user_condition"`                  // 用户条件: all, new_users, vip_users
	RegisterAfter  time.Time `json:"register_after" form:"register_after"`                  // 注册时间之后
	RegisterBefore time.Time `json:"register_before" form:"register_before"`                // 注册时间之前
	MaxUsers       int       `json:"max_users" form:"max_users"`                            // 最大发放用户数
	Source         string    `json:"source" form:"source"`                                  // 发放来源
}

// DeleteActivityRequest 删除活动请求
type DeleteActivityRequest struct {
	ID uint `json:"id" form:"id" binding:"required"` // 活动ID
}
