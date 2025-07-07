package logic

import (
	"Common/global"
	coupon "coupon_srv/proto_coupon"
	"fmt"
	"models/model_mysql"
	"time"

	"gorm.io/gorm"
)

// GrantCoupon 发放优惠券
func GrantCoupon(req *coupon.GrantCouponRequest) (*coupon.GrantCouponResponse, error) {
	if req.UserId == 0 {
		return &coupon.GrantCouponResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	if req.ActivityCode == "" {
		return &coupon.GrantCouponResponse{
			Code:    400,
			Message: "活动编码不能为空",
		}, nil
	}

	source := req.Source
	if source == "" {
		source = "AUTO_GRANT" // 默认自动发放
	}

	// 1. 获取活动配置
	var config model_mysql.PromotionConfig
	if err := config.GetByActivityCode(req.ActivityCode); err != nil {
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("活动不存在: %v", err),
		}, nil
	}

	// 2. 检查活动是否有效
	if !config.IsActive() {
		return &coupon.GrantCouponResponse{
			Code:    400,
			Message: "活动已结束或未开始",
		}, nil
	}

	// 3. 检查是否可以发放
	if !config.CanGrant() {
		return &coupon.GrantCouponResponse{
			Code:    400,
			Message: "活动已达到发放上限",
		}, nil
	}

	// 4. 检查用户是否已参与过该活动
	participated, err := model_mysql.HasParticipated(uint(req.UserId), req.ActivityCode)
	if err != nil {
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("检查用户参与状态失败: %v", err),
		}, nil
	}
	if participated {
		return &coupon.GrantCouponResponse{
			Code:    400,
			Message: "用户已参与过该活动",
		}, nil
	}

	// 5. 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 6. 创建优惠券
	newCoupon := &model_mysql.UserCoupon{
		UserID:            uint(req.UserId),
		CouponType:        config.ActivityType,
		CouponName:        config.ActivityName,
		CouponCode:        model_mysql.GenerateCouponCode(),
		DiscountType:      config.DiscountType,
		DiscountAmount:    config.DiscountAmount,
		DiscountRate:      config.DiscountRate,
		MinOrderAmount:    config.MinOrderAmount,
		MaxDiscountAmount: config.MaxDiscountAmount,
		Status:            model_mysql.CouponStatusUnused,
		Source:            source,
		ExpireTime:        time.Now().AddDate(0, 0, config.ValidDays),
	}

	if err := tx.Create(newCoupon).Error; err != nil {
		tx.Rollback()
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("创建优惠券失败: %v", err),
		}, nil
	}

	// 7. 增加活动发放计数
	if err := tx.Model(&config).UpdateColumn("current_grant_count", gorm.Expr("current_grant_count + 1")).Error; err != nil {
		tx.Rollback()
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("更新活动计数失败: %v", err),
		}, nil
	}

	// 8. 创建用户活动参与记录
	record := &model_mysql.UserActivityRecord{
		UserID:       uint(req.UserId),
		ActivityCode: req.ActivityCode,
		CouponCount:  1,
		Status:       1,
	}
	if err := tx.Create(record).Error; err != nil {
		tx.Rollback()
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("创建参与记录失败: %v", err),
		}, nil
	}

	// 9. 提交事务
	if err := tx.Commit().Error; err != nil {
		return &coupon.GrantCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("提交事务失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	couponInfo := convertToCouponInfo(newCoupon)

	return &coupon.GrantCouponResponse{
		Code:    200,
		Message: "优惠券发放成功",
		Coupon:  couponInfo,
	}, nil
}

// GetUserCoupons 获取用户优惠券列表
func GetUserCoupons(req *coupon.GetUserCouponsRequest) (*coupon.GetUserCouponsResponse, error) {
	if req.UserId == 0 {
		return &coupon.GetUserCouponsResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	var coupons []model_mysql.UserCoupon
	var total int64

	query := global.DB.Where("user_id = ?", req.UserId)
	if req.Status > 0 {
		query = query.Where("status = ?", req.Status)
	}

	// 获取总数
	if err := query.Model(&model_mysql.UserCoupon{}).Count(&total).Error; err != nil {
		return &coupon.GetUserCouponsResponse{
			Code:    500,
			Message: fmt.Sprintf("获取优惠券总数失败: %v", err),
		}, nil
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&coupons).Error; err != nil {
		return &coupon.GetUserCouponsResponse{
			Code:    500,
			Message: fmt.Sprintf("获取优惠券列表失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	couponInfos := make([]*coupon.CouponInfo, len(coupons))
	for i, c := range coupons {
		couponInfos[i] = convertToCouponInfo(&c)
	}

	return &coupon.GetUserCouponsResponse{
		Code:    200,
		Message: "获取优惠券列表成功",
		Coupons: couponInfos,
		Total:   total,
	}, nil
}

// GetAvailableCoupons 获取用户可用优惠券
func GetAvailableCoupons(req *coupon.GetAvailableCouponsRequest) (*coupon.GetAvailableCouponsResponse, error) {
	if req.UserId == 0 {
		return &coupon.GetAvailableCouponsResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	coupons, err := model_mysql.GetUserAvailableCoupons(uint(req.UserId), req.OrderAmount)
	if err != nil {
		return &coupon.GetAvailableCouponsResponse{
			Code:    500,
			Message: fmt.Sprintf("获取可用优惠券失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	couponInfos := make([]*coupon.CouponInfo, len(coupons))
	for i, c := range coupons {
		couponInfos[i] = convertToCouponInfo(&c)
	}

	return &coupon.GetAvailableCouponsResponse{
		Code:    200,
		Message: "获取可用优惠券成功",
		Coupons: couponInfos,
	}, nil
}

// ValidateCoupon 验证优惠券
func ValidateCoupon(req *coupon.ValidateCouponRequest) (*coupon.ValidateCouponResponse, error) {
	if req.CouponId == 0 || req.UserId == 0 {
		return &coupon.ValidateCouponResponse{
			Code:    400,
			Message: "优惠券ID和用户ID不能为空",
		}, nil
	}

	var userCoupon model_mysql.UserCoupon
	if err := userCoupon.GetByID(uint(req.CouponId)); err != nil {
		return &coupon.ValidateCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("优惠券不存在: %v", err),
			IsValid: false,
		}, nil
	}

	// 检查优惠券是否属于该用户
	if userCoupon.UserID != uint(req.UserId) {
		return &coupon.ValidateCouponResponse{
			Code:    400,
			Message: "优惠券不属于该用户",
			IsValid: false,
		}, nil
	}

	// 检查优惠券是否可用
	if !userCoupon.CanUseForOrder(req.OrderAmount) {
		return &coupon.ValidateCouponResponse{
			Code:    400,
			Message: "优惠券不可用",
			IsValid: false,
		}, nil
	}

	// 计算优惠金额
	discountAmount := userCoupon.CalculateDiscount(req.OrderAmount)

	return &coupon.ValidateCouponResponse{
		Code:           200,
		Message:        "验证成功",
		IsValid:        true,
		DiscountAmount: discountAmount,
	}, nil
}

// UseCoupon 使用优惠券
func UseCoupon(req *coupon.UseCouponRequest) (*coupon.UseCouponResponse, error) {
	if req.CouponId == 0 || req.UserId == 0 || req.OrderId == 0 {
		return &coupon.UseCouponResponse{
			Code:    400,
			Message: "优惠券ID、用户ID和订单ID不能为空",
		}, nil
	}

	// 1. 验证优惠券
	validateReq := &coupon.ValidateCouponRequest{
		CouponId:    req.CouponId,
		UserId:      req.UserId,
		OrderAmount: req.OriginalAmount,
	}
	validateResp, err := ValidateCoupon(validateReq)
	if err != nil {
		return &coupon.UseCouponResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}
	if !validateResp.IsValid {
		return &coupon.UseCouponResponse{
			Code:    400,
			Message: validateResp.Message,
		}, nil
	}

	// 2. 验证优惠金额
	if req.DiscountAmount != validateResp.DiscountAmount {
		return &coupon.UseCouponResponse{
			Code:    400,
			Message: "优惠金额不匹配",
		}, nil
	}

	// 3. 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 4. 更新优惠券状态
	var userCoupon model_mysql.UserCoupon
	if err := tx.Where("id = ?", req.CouponId).First(&userCoupon).Error; err != nil {
		tx.Rollback()
		return &coupon.UseCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("获取优惠券失败: %v", err),
		}, nil
	}

	if err := userCoupon.UseCoupon(uint(req.OrderId), req.OrderSn); err != nil {
		tx.Rollback()
		return &coupon.UseCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("更新优惠券状态失败: %v", err),
		}, nil
	}

	// 5. 创建使用记录
	usageLog := &model_mysql.CouponUsageLog{
		CouponID:       uint(req.CouponId),
		UserID:         uint(req.UserId),
		OrderID:        uint(req.OrderId),
		OrderSn:        req.OrderSn,
		OriginalAmount: req.OriginalAmount,
		DiscountAmount: req.DiscountAmount,
		FinalAmount:    req.OriginalAmount - req.DiscountAmount,
	}
	if err := tx.Create(usageLog).Error; err != nil {
		tx.Rollback()
		return &coupon.UseCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("创建使用记录失败: %v", err),
		}, nil
	}

	// 6. 提交事务
	if err := tx.Commit().Error; err != nil {
		return &coupon.UseCouponResponse{
			Code:    500,
			Message: fmt.Sprintf("提交事务失败: %v", err),
		}, nil
	}

	return &coupon.UseCouponResponse{
		Code:    200,
		Message: "优惠券使用成功",
	}, nil
}

// convertToCouponInfo 转换为优惠券信息
func convertToCouponInfo(userCoupon *model_mysql.UserCoupon) *coupon.CouponInfo {
	info := &coupon.CouponInfo{
		Id:                uint64(userCoupon.ID),
		UserId:            uint64(userCoupon.UserID),
		CouponType:        userCoupon.CouponType,
		CouponName:        userCoupon.CouponName,
		CouponCode:        userCoupon.CouponCode,
		DiscountType:      int32(userCoupon.DiscountType),
		DiscountAmount:    userCoupon.DiscountAmount,
		DiscountRate:      userCoupon.DiscountRate,
		MinOrderAmount:    userCoupon.MinOrderAmount,
		MaxDiscountAmount: userCoupon.MaxDiscountAmount,
		Status:            int32(userCoupon.Status),
		Source:            userCoupon.Source,
		ExpireTime:        userCoupon.ExpireTime.Format(time.RFC3339),
		CreatedAt:         userCoupon.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         userCoupon.UpdatedAt.Format(time.RFC3339),
	}

	if userCoupon.UsedTime != nil {
		info.UsedTime = userCoupon.UsedTime.Format(time.RFC3339)
	}

	if userCoupon.OrderID != nil {
		info.OrderId = uint64(*userCoupon.OrderID)
	}

	info.OrderSn = userCoupon.OrderSn

	return info
}

// ===== 活动管理相关函数 =====

// CreateActivity 创建活动
func CreateActivity(req *coupon.CreateActivityRequest) (*coupon.CreateActivityResponse, error) {
	if req.ActivityName == "" {
		return &coupon.CreateActivityResponse{
			Code:    400,
			Message: "活动名称不能为空",
		}, nil
	}

	if req.ActivityCode == "" {
		return &coupon.CreateActivityResponse{
			Code:    400,
			Message: "活动编码不能为空",
		}, nil
	}

	// 检查活动编码是否已存在
	var existingActivity model_mysql.PromotionConfig
	if err := existingActivity.GetByActivityCode(req.ActivityCode); err == nil {
		return &coupon.CreateActivityResponse{
			Code:    400,
			Message: "活动编码已存在",
		}, nil
	}

	// 创建活动
	activity := &model_mysql.PromotionConfig{
		ActivityName:      req.ActivityName,
		ActivityCode:      req.ActivityCode,
		ActivityType:      req.ActivityType,
		Description:       req.Description,
		DiscountType:      int(req.DiscountType),
		DiscountAmount:    req.DiscountAmount,
		DiscountRate:      req.DiscountRate,
		MinOrderAmount:    req.MinOrderAmount,
		MaxDiscountAmount: req.MaxDiscountAmount,
		ValidDays:         int(req.ValidDays),
		MaxGrantCount:     int(req.MaxGrantCount),
		TotalGrantLimit:   int(req.TotalGrantLimit),
		StartTime:         req.StartTime,
		EndTime:           req.EndTime,
		Status:            int(req.Status),
	}

	if err := activity.CreateActivity(); err != nil {
		return &coupon.CreateActivityResponse{
			Code:    500,
			Message: fmt.Sprintf("创建活动失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	activityInfo := convertToActivityInfo(activity)

	return &coupon.CreateActivityResponse{
		Code:     200,
		Message:  "活动创建成功",
		Activity: activityInfo,
	}, nil
}

// UpdateActivity 更新活动
func UpdateActivity(req *coupon.UpdateActivityRequest) (*coupon.UpdateActivityResponse, error) {
	if req.Id == 0 {
		return &coupon.UpdateActivityResponse{
			Code:    400,
			Message: "活动ID不能为空",
		}, nil
	}

	// 获取现有活动
	activity, err := model_mysql.GetActivityByID(uint(req.Id))
	if err != nil {
		return &coupon.UpdateActivityResponse{
			Code:    404,
			Message: "活动不存在",
		}, nil
	}

	// 更新活动信息
	activity.ActivityName = req.ActivityName
	activity.ActivityType = req.ActivityType
	activity.Description = req.Description
	activity.DiscountType = int(req.DiscountType)
	activity.DiscountAmount = req.DiscountAmount
	activity.DiscountRate = req.DiscountRate
	activity.MinOrderAmount = req.MinOrderAmount
	activity.MaxDiscountAmount = req.MaxDiscountAmount
	activity.ValidDays = int(req.ValidDays)
	activity.MaxGrantCount = int(req.MaxGrantCount)
	activity.TotalGrantLimit = int(req.TotalGrantLimit)
	activity.StartTime = req.StartTime
	activity.EndTime = req.EndTime
	activity.Status = int(req.Status)

	if err := activity.UpdateActivity(); err != nil {
		return &coupon.UpdateActivityResponse{
			Code:    500,
			Message: fmt.Sprintf("更新活动失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	activityInfo := convertToActivityInfo(activity)

	return &coupon.UpdateActivityResponse{
		Code:     200,
		Message:  "活动更新成功",
		Activity: activityInfo,
	}, nil
}

// DeleteActivity 删除活动
func DeleteActivity(req *coupon.DeleteActivityRequest) (*coupon.DeleteActivityResponse, error) {
	if req.Id == 0 {
		return &coupon.DeleteActivityResponse{
			Code:    400,
			Message: "活动ID不能为空",
		}, nil
	}

	// 获取活动
	activity, err := model_mysql.GetActivityByID(uint(req.Id))
	if err != nil {
		return &coupon.DeleteActivityResponse{
			Code:    404,
			Message: "活动不存在",
		}, nil
	}

	if err := activity.DeleteActivity(); err != nil {
		return &coupon.DeleteActivityResponse{
			Code:    500,
			Message: fmt.Sprintf("删除活动失败: %v", err),
		}, nil
	}

	return &coupon.DeleteActivityResponse{
		Code:    200,
		Message: "活动删除成功",
	}, nil
}

// GetActivity 获取活动详情
func GetActivity(req *coupon.GetActivityRequest) (*coupon.GetActivityResponse, error) {
	if req.Id == 0 {
		return &coupon.GetActivityResponse{
			Code:    400,
			Message: "活动ID不能为空",
		}, nil
	}

	activity, err := model_mysql.GetActivityByID(uint(req.Id))
	if err != nil {
		return &coupon.GetActivityResponse{
			Code:    404,
			Message: "活动不存在",
		}, nil
	}

	// 转换为响应格式
	activityInfo := convertToActivityInfo(activity)

	return &coupon.GetActivityResponse{
		Code:     200,
		Message:  "获取活动成功",
		Activity: activityInfo,
	}, nil
}

// GetActivityList 获取活动列表
func GetActivityList(req *coupon.GetActivityListRequest) (*coupon.GetActivityListResponse, error) {
	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	var status *int
	if req.Status > 0 {
		s := int(req.Status)
		status = &s
	}

	activities, total, err := model_mysql.GetActivityList(page, pageSize, status, req.Keyword)
	if err != nil {
		return &coupon.GetActivityListResponse{
			Code:    500,
			Message: fmt.Sprintf("获取活动列表失败: %v", err),
		}, nil
	}

	// 转换为响应格式
	activityInfos := make([]*coupon.ActivityInfo, len(activities))
	for i, activity := range activities {
		activityInfos[i] = convertToActivityInfo(&activity)
	}

	return &coupon.GetActivityListResponse{
		Code:       200,
		Message:    "获取活动列表成功",
		Activities: activityInfos,
		Total:      total,
	}, nil
}

// DistributeActivityCoupons 分发活动优惠券
func DistributeActivityCoupons(req *coupon.DistributeActivityCouponsRequest) (*coupon.DistributeActivityCouponsResponse, error) {
	if req.ActivityCode == "" {
		return &coupon.DistributeActivityCouponsResponse{
			Code:    400,
			Message: "活动编码不能为空",
		}, nil
	}

	if len(req.UserIds) == 0 {
		return &coupon.DistributeActivityCouponsResponse{
			Code:    400,
			Message: "用户ID列表不能为空",
		}, nil
	}

	source := req.Source
	if source == "" {
		source = "MANUAL_GRANT" // 默认手动发放
	}

	var successCount int32
	var failedCount int32
	var failedReasons []string

	// 逐个发放优惠券
	for _, userID := range req.UserIds {
		grantReq := &coupon.GrantCouponRequest{
			UserId:       userID,
			ActivityCode: req.ActivityCode,
			Source:       source,
		}

		grantResp, err := GrantCoupon(grantReq)
		if err != nil || grantResp.Code != 200 {
			failedCount++
			reason := fmt.Sprintf("用户ID %d: %s", userID, grantResp.Message)
			failedReasons = append(failedReasons, reason)
		} else {
			successCount++
		}
	}

	return &coupon.DistributeActivityCouponsResponse{
		Code:          200,
		Message:       fmt.Sprintf("分发完成，成功 %d 个，失败 %d 个", successCount, failedCount),
		SuccessCount:  successCount,
		FailedCount:   failedCount,
		FailedReasons: failedReasons,
	}, nil
}

// convertToActivityInfo 转换为活动信息
func convertToActivityInfo(activity *model_mysql.PromotionConfig) *coupon.ActivityInfo {
	return &coupon.ActivityInfo{
		Id:                uint64(activity.ID),
		ActivityName:      activity.ActivityName,
		ActivityCode:      activity.ActivityCode,
		ActivityType:      activity.ActivityType,
		Description:       activity.Description,
		DiscountType:      int32(activity.DiscountType),
		DiscountAmount:    activity.DiscountAmount,
		DiscountRate:      activity.DiscountRate,
		MinOrderAmount:    activity.MinOrderAmount,
		MaxDiscountAmount: activity.MaxDiscountAmount,
		ValidDays:         int32(activity.ValidDays),
		MaxGrantCount:     int32(activity.MaxGrantCount),
		TotalGrantLimit:   int32(activity.TotalGrantLimit),
		CurrentGrantCount: int32(activity.CurrentGrantCount),
		StartTime:         activity.StartTime,
		EndTime:           activity.EndTime,
		Status:            int32(activity.Status),
		CreatedAt:         activity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339),
	}
}
