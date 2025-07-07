package trigger

import (
	"Api/client"
	"Api/request"
	"context"
	"net/http"
	"strconv"

	proto_coupon "Srv/coupon_srv/proto_coupon"

	"github.com/gin-gonic/gin"
)

// CreateActivityHandler 创建活动
func CreateActivityHandler(c *gin.Context) {
	var req request.CreateActivityRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 通过gRPC调用coupon_srv创建活动
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.CreateActivity(ctx, &proto_coupon.CreateActivityRequest{
			ActivityName:      req.ActivityName,
			ActivityCode:      req.ActivityCode,
			ActivityType:      req.ActivityType,
			Description:       req.Description,
			DiscountType:      int32(req.DiscountType),
			DiscountAmount:    req.DiscountAmount,
			DiscountRate:      req.DiscountRate,
			MinOrderAmount:    req.MinOrderAmount,
			MaxDiscountAmount: req.MaxDiscountAmount,
			ValidDays:         int32(req.ValidDays),
			MaxGrantCount:     int32(req.MaxGrantCount),
			TotalGrantLimit:   int32(req.TotalGrantLimit),
			StartTime:         req.StartTime,
			EndTime:           req.EndTime,
			Status:            int32(req.Status),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.CreateActivityResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "活动创建成功",
		"data":    response.Activity,
	})
}

// UpdateActivityHandler 更新活动
func UpdateActivityHandler(c *gin.Context) {
	var req request.UpdateActivityRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 通过gRPC调用coupon_srv更新活动
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.UpdateActivity(ctx, &proto_coupon.UpdateActivityRequest{
			Id:                uint64(req.ID),
			ActivityName:      req.ActivityName,
			ActivityType:      req.ActivityType,
			Description:       req.Description,
			DiscountType:      int32(req.DiscountType),
			DiscountAmount:    req.DiscountAmount,
			DiscountRate:      req.DiscountRate,
			MinOrderAmount:    req.MinOrderAmount,
			MaxDiscountAmount: req.MaxDiscountAmount,
			ValidDays:         int32(req.ValidDays),
			MaxGrantCount:     int32(req.MaxGrantCount),
			TotalGrantLimit:   int32(req.TotalGrantLimit),
			StartTime:         req.StartTime,
			EndTime:           req.EndTime,
			Status:            int32(req.Status),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.UpdateActivityResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "活动更新成功",
		"data":    response.Activity,
	})
}

// GetActivityListHandler 获取活动列表
func GetActivityListHandler(c *gin.Context) {
	var req request.GetActivityListRequest
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

	// 通过gRPC调用coupon_srv获取活动列表
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		status := int32(0)
		if req.Status != nil {
			status = int32(*req.Status)
		}
		return client.GetActivityList(ctx, &proto_coupon.GetActivityListRequest{
			Page:     int32(req.Page),
			PageSize: int32(req.PageSize),
			Status:   status,
			Keyword:  req.Keyword,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.GetActivityListResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"activities": response.Activities,
			"total":      response.Total,
			"page":       req.Page,
			"page_size":  req.PageSize,
		},
	})
}

// GetActivityDetailHandler 获取活动详情
func GetActivityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的活动ID",
		})
		return
	}

	// 通过gRPC调用coupon_srv获取活动详情
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.GetActivity(ctx, &proto_coupon.GetActivityRequest{
			Id: id,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.GetActivityResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response.Activity,
	})
}

// DeleteActivityHandler 删除活动
func DeleteActivityHandler(c *gin.Context) {
	var req request.DeleteActivityRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 通过gRPC调用coupon_srv删除活动
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.DeleteActivity(ctx, &proto_coupon.DeleteActivityRequest{
			Id: uint64(req.ID),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.DeleteActivityResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "活动删除成功",
	})
}

// BatchGrantCouponHandler 批量发放优惠券
func BatchGrantCouponHandler(c *gin.Context) {
	var req request.BatchGrantCouponRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认来源
	if req.Source == "" {
		req.Source = "BATCH_GRANT"
	}

	// 转换用户ID列表
	userIds := make([]uint64, len(req.UserIDs))
	for i, id := range req.UserIDs {
		userIds[i] = uint64(id)
	}

	// 通过gRPC调用coupon_srv分发活动优惠券
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.DistributeActivityCoupons(ctx, &proto_coupon.DistributeActivityCouponsRequest{
			ActivityCode: req.ActivityCode,
			UserIds:      userIds,
			Source:       req.Source,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.DistributeActivityCouponsResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": response.Message,
		"data": gin.H{
			"activity_code":  req.ActivityCode,
			"success_count":  response.SuccessCount,
			"failed_count":   response.FailedCount,
			"failed_reasons": response.FailedReasons,
		},
	})
}

// GrantCouponByConditionHandler 按条件发放优惠券
func GrantCouponByConditionHandler(c *gin.Context) {
	var req request.GrantCouponByConditionRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认来源
	if req.Source == "" {
		req.Source = "CONDITION_GRANT"
	}

	// TODO: 这里需要调用user_srv获取符合条件的用户列表
	// 然后调用coupon_srv的DistributeActivityCoupons接口
	// 暂时返回功能未实现的响应
	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    501,
		"message": "按条件发放功能需要集成user_srv，暂未实现",
		"data": gin.H{
			"activity_code":  req.ActivityCode,
			"user_condition": req.UserCondition,
			"source":         req.Source,
		},
	})
}

// GetPublicActivityListHandler 获取公开活动列表（用户端）
func GetPublicActivityListHandler(c *gin.Context) {
	var req request.GetActivityListRequest
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
		req.PageSize = 10
	}

	// 通过gRPC调用coupon_srv获取活动列表（只显示启用的活动）
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.GetActivityList(ctx, &proto_coupon.GetActivityListRequest{
			Page:     int32(req.Page),
			PageSize: int32(req.PageSize),
			Status:   1, // 只显示启用的活动
			Keyword:  req.Keyword,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.GetActivityListResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	// 过滤敏感信息，只返回用户需要的字段
	var publicActivities []gin.H
	for _, activity := range response.Activities {
		publicActivities = append(publicActivities, gin.H{
			"id":                  activity.Id,
			"activity_name":       activity.ActivityName,
			"activity_code":       activity.ActivityCode,
			"activity_type":       activity.ActivityType,
			"description":         activity.Description,
			"discount_type":       activity.DiscountType,
			"discount_amount":     activity.DiscountAmount,
			"discount_rate":       activity.DiscountRate,
			"min_order_amount":    activity.MinOrderAmount,
			"max_discount_amount": activity.MaxDiscountAmount,
			"valid_days":          activity.ValidDays,
			"start_time":          activity.StartTime,
			"end_time":            activity.EndTime,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"activities": publicActivities,
			"total":      len(publicActivities),
			"page":       req.Page,
			"page_size":  req.PageSize,
		},
	})
}

// GetPublicActivityDetailHandler 获取公开活动详情（用户端）
func GetPublicActivityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的活动ID",
		})
		return
	}

	// 通过gRPC调用coupon_srv获取活动详情
	ctx := context.Background()
	result, err := client.CouponClient(ctx, func(ctx context.Context, client proto_coupon.CouponServiceClient) (interface{}, error) {
		return client.GetActivity(ctx, &proto_coupon.GetActivityRequest{
			Id: id,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务调用失败: " + err.Error(),
		})
		return
	}

	response := result.(*proto_coupon.GetActivityResponse)
	if response.Code != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    response.Code,
			"message": response.Message,
		})
		return
	}

	activity := response.Activity
	// 只有启用的活动才能查看详情
	if activity.Status != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "活动不存在或已结束",
		})
		return
	}

	// 返回用户需要的活动信息
	publicActivity := gin.H{
		"id":                  activity.Id,
		"activity_name":       activity.ActivityName,
		"activity_code":       activity.ActivityCode,
		"activity_type":       activity.ActivityType,
		"description":         activity.Description,
		"discount_type":       activity.DiscountType,
		"discount_amount":     activity.DiscountAmount,
		"discount_rate":       activity.DiscountRate,
		"min_order_amount":    activity.MinOrderAmount,
		"max_discount_amount": activity.MaxDiscountAmount,
		"valid_days":          activity.ValidDays,
		"start_time":          activity.StartTime,
		"end_time":            activity.EndTime,
		"remaining_count":     activity.TotalGrantLimit - activity.CurrentGrantCount, // 剩余数量
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    publicActivity,
	})
}
