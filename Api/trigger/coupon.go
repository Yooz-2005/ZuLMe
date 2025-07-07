package trigger

import (
	"Api/handler"
	"Api/request"
	proto_coupon "Srv/coupon_srv/proto_coupon"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GrantCouponHandler 发放优惠券
func GrantCouponHandler(c *gin.Context) {
	var req request.GrantCouponRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认来源
	if req.Source == "" {
		req.Source = "MANUAL_GRANT"
	}

	// 调用优惠券服务的gRPC接口
	grpcReq := &proto_coupon.GrantCouponRequest{
		UserId:       req.UserID,       // 用户ID
		ActivityCode: req.ActivityCode, // 活动代码
		Source:       req.Source,       // 发放来源
	}

	response, err := handler.GrantCoupon(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	if response.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "优惠券发放成功",
		"data":    response.Coupon,
	})
}

// GetUserCouponsHandler 获取用户优惠券列表
func GetUserCouponsHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	var req request.GetUserCouponsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// 调用优惠券服务的gRPC接口
	grpcReq := &proto_coupon.GetUserCouponsRequest{
		UserId:   uint64(userID),
		Status:   int32(req.Status),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}

	response, err := handler.GetUserCoupons(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取优惠券列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取优惠券列表成功",
		"data":    response,
	})
}

// GetAvailableCouponsHandler 获取用户可用优惠券
func GetAvailableCouponsHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	var req request.GetAvailableCouponsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 调用优惠券服务的gRPC接口
	grpcReq := &proto_coupon.GetAvailableCouponsRequest{
		UserId:      uint64(userID),
		OrderAmount: req.OrderAmount,
	}

	response, err := handler.GetAvailableCoupons(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取可用优惠券失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取可用优惠券成功",
		"data":    response,
	})
}

// ValidateCouponHandler 验证优惠券
func ValidateCouponHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	var req request.ValidateCouponRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 调用优惠券服务的gRPC接口
	grpcReq := &proto_coupon.ValidateCouponRequest{
		CouponId:    req.CouponID,
		UserId:      uint64(userID),
		OrderAmount: req.OrderAmount,
	}

	response, err := handler.ValidateCoupon(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "验证优惠券失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证优惠券成功",
		"data":    response,
	})
}

// UseCouponHandler 使用优惠券
func UseCouponHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	var req request.UseCouponRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 调用优惠券服务的gRPC接口
	grpcReq := &proto_coupon.UseCouponRequest{
		CouponId:       req.CouponID,
		UserId:         uint64(userID),
		OrderId:        req.OrderID,
		OrderSn:        req.OrderSn,
		OriginalAmount: req.OriginalAmount,
		DiscountAmount: req.DiscountAmount,
	}

	response, err := handler.UseCoupon(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "使用优惠券失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "使用优惠券成功",
		"data":    response,
	})
}

// GetCouponByIdHandler 根据ID获取优惠券详情
func GetCouponByIdHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 获取优惠券ID
	couponIDStr := c.Param("coupon_id")
	couponID, err := strconv.ParseUint(couponIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "优惠券ID格式错误",
		})
		return
	}

	// 通过获取用户优惠券列表来查找指定优惠券
	grpcReq := &proto_coupon.GetUserCouponsRequest{
		UserId:   uint64(userID),
		Status:   0, // 获取所有状态的优惠券
		Page:     1,
		PageSize: 1000, // 设置较大的页面大小以确保能找到目标优惠券
	}

	response, err := handler.GetUserCoupons(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取优惠券详情失败: " + err.Error(),
		})
		return
	}

	// 查找指定ID的优惠券
	var targetCoupon *proto_coupon.CouponInfo
	for _, coupon := range response.Coupons {
		if coupon.Id == couponID {
			targetCoupon = coupon
			break
		}
	}

	if targetCoupon == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "优惠券不存在或不属于当前用户",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取优惠券详情成功",
		"data":    targetCoupon,
	})
}
