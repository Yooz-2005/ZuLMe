package model_mysql

import (
	"Common/global"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserCoupon 用户优惠券模型
type UserCoupon struct {
	gorm.Model
	UserID            uint       `gorm:"column:user_id;not null;index;comment:'用户ID'" json:"user_id"`
	CouponType        string     `gorm:"column:coupon_type;type:varchar(50);not null;default:'NEWBIE_GIFT';comment:'优惠券类型'" json:"coupon_type"`
	CouponName        string     `gorm:"column:coupon_name;type:varchar(100);not null;comment:'优惠券名称'" json:"coupon_name"`
	CouponCode        string     `gorm:"column:coupon_code;type:varchar(50);not null;uniqueIndex;comment:'优惠券编码'" json:"coupon_code"`
	DiscountType      int        `gorm:"column:discount_type;type:tinyint;not null;default:1;comment:'优惠类型:1-减免金额,2-折扣比例'" json:"discount_type"`
	DiscountAmount    float64    `gorm:"column:discount_amount;type:decimal(10,2);not null;default:0.00;comment:'优惠金额(元)'" json:"discount_amount"`
	DiscountRate      float64    `gorm:"column:discount_rate;type:decimal(5,2);not null;default:0.00;comment:'折扣比例(0-100)'" json:"discount_rate"`
	MinOrderAmount    float64    `gorm:"column:min_order_amount;type:decimal(10,2);not null;default:0.00;comment:'最低订单金额要求'" json:"min_order_amount"`
	MaxDiscountAmount float64    `gorm:"column:max_discount_amount;type:decimal(10,2);not null;default:0.00;comment:'最大优惠金额(用于折扣券)'" json:"max_discount_amount"`
	Status            int        `gorm:"column:status;type:tinyint;not null;default:1;index;comment:'状态:1-未使用,2-已使用,3-已过期,4-已作废'" json:"status"`
	Source            string     `gorm:"column:source;type:varchar(50);not null;default:'AUTO_GRANT';comment:'来源:AUTO_GRANT-自动发放,MANUAL_GRANT-手动发放'" json:"source"`
	ExpireTime        time.Time  `gorm:"column:expire_time;not null;index;comment:'过期时间'" json:"expire_time"`
	UsedTime          *time.Time `gorm:"column:used_time;comment:'使用时间'" json:"used_time"`
	OrderID           *uint      `gorm:"column:order_id;index;comment:'使用的订单ID'" json:"order_id"`
	OrderSn           string     `gorm:"column:order_sn;type:varchar(50);comment:'使用的订单号'" json:"order_sn"`
}

// TableName 指定表名
func (uc *UserCoupon) TableName() string {
	return "user_coupons"
}

// 优惠券状态常量
const (
	CouponStatusUnused  = 1 // 未使用
	CouponStatusUsed    = 2 // 已使用
	CouponStatusExpired = 3 // 已过期
	CouponStatusVoid    = 4 // 已作废
)

// 优惠券类型常量
const (
	CouponTypeNewbieGift = "NEWBIE_GIFT" // 新人大礼包
)

// 优惠类型常量
const (
	DiscountTypeAmount = 1 // 减免金额
	DiscountTypeRate   = 2 // 折扣比例
)

// CreateCoupon 创建优惠券
func (uc *UserCoupon) CreateCoupon() error {
	return global.DB.Create(uc).Error
}

// GetByID 根据ID获取优惠券
func (uc *UserCoupon) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).First(uc).Error
}

// GetByCouponCode 根据优惠券编码获取优惠券
func (uc *UserCoupon) GetByCouponCode(code string) error {
	return global.DB.Where("coupon_code = ?", code).First(uc).Error
}

// GetUserCoupons 获取用户的优惠券列表
func GetUserCoupons(userID uint, status int) ([]UserCoupon, error) {
	var coupons []UserCoupon
	query := global.DB.Where("user_id = ?", userID)

	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Order("created_at DESC").Find(&coupons).Error
	return coupons, err
}

// GetUserAvailableCoupons 获取用户可用的优惠券
func GetUserAvailableCoupons(userID uint, orderAmount float64) ([]UserCoupon, error) {
	var coupons []UserCoupon
	err := global.DB.Where("user_id = ? AND status = ? AND expire_time > ? AND min_order_amount <= ?",
		userID, CouponStatusUnused, time.Now(), orderAmount).
		Order("discount_amount DESC").Find(&coupons).Error
	return coupons, err
}

// UseCoupon 使用优惠券
func (uc *UserCoupon) UseCoupon(orderID uint, orderSn string) error {
	now := time.Now()
	return global.DB.Model(uc).Updates(map[string]any{
		"status":     CouponStatusUsed,
		"used_time":  &now,
		"order_id":   orderID,
		"order_sn":   orderSn,
		"updated_at": now,
	}).Error
}

// IsValid 检查优惠券是否有效
func (uc *UserCoupon) IsValid() bool {
	return uc.Status == CouponStatusUnused && uc.ExpireTime.After(time.Now())
}

// CanUseForOrder 检查优惠券是否可用于指定订单
func (uc *UserCoupon) CanUseForOrder(orderAmount float64) bool {
	return uc.IsValid() && orderAmount >= uc.MinOrderAmount
}

// CalculateDiscount 计算优惠金额
func (uc *UserCoupon) CalculateDiscount(orderAmount float64) float64 {
	if !uc.CanUseForOrder(orderAmount) {
		return 0
	}

	var discount float64
	switch uc.DiscountType {
	case DiscountTypeAmount:
		discount = uc.DiscountAmount
	case DiscountTypeRate:
		discount = orderAmount * uc.DiscountRate / 100
		if uc.MaxDiscountAmount > 0 && discount > uc.MaxDiscountAmount {
			discount = uc.MaxDiscountAmount
		}
	}

	// 优惠金额不能超过订单金额
	if discount > orderAmount {
		discount = orderAmount
	}

	return discount
}

// UpdateExpiredCoupons 更新过期优惠券状态
func UpdateExpiredCoupons() error {
	return global.DB.Model(&UserCoupon{}).
		Where("status = ? AND expire_time < ?", CouponStatusUnused, time.Now()).
		Update("status", CouponStatusExpired).Error
}

// PromotionConfig 活动配置模型
type PromotionConfig struct {
	gorm.Model
	ActivityName      string  `gorm:"column:activity_name;type:varchar(100);not null;comment:'活动名称'" json:"activity_name"`
	ActivityCode      string  `gorm:"column:activity_code;type:varchar(50);not null;uniqueIndex;comment:'活动编码'" json:"activity_code"`
	ActivityType      string  `gorm:"column:activity_type;type:varchar(50);not null;comment:'活动类型'" json:"activity_type"`
	Description       string  `gorm:"column:description;type:text;comment:'活动描述'" json:"description"`
	DiscountType      int     `gorm:"column:discount_type;type:tinyint;not null;default:1;comment:'优惠类型:1-减免金额,2-折扣比例'" json:"discount_type"`
	DiscountAmount    float64 `gorm:"column:discount_amount;type:decimal(10,2);not null;default:0.00;comment:'优惠金额(元)'" json:"discount_amount"`
	DiscountRate      float64 `gorm:"column:discount_rate;type:decimal(5,2);not null;default:0.00;comment:'折扣比例(0-100)'" json:"discount_rate"`
	MinOrderAmount    float64 `gorm:"column:min_order_amount;type:decimal(10,2);not null;default:0.00;comment:'最低订单金额要求'" json:"min_order_amount"`
	MaxDiscountAmount float64 `gorm:"column:max_discount_amount;type:decimal(10,2);not null;default:0.00;comment:'最大优惠金额(用于折扣券)'" json:"max_discount_amount"`
	ValidDays         int     `gorm:"column:valid_days;not null;default:7;comment:'有效期天数'" json:"valid_days"`
	MaxGrantCount     int     `gorm:"column:max_grant_count;not null;default:1;comment:'每人最多可发放数量'" json:"max_grant_count"`
	TotalGrantLimit   int     `gorm:"column:total_grant_limit;not null;default:0;comment:'总发放数量限制,0表示无限制'" json:"total_grant_limit"`
	CurrentGrantCount int     `gorm:"column:current_grant_count;not null;default:0;comment:'当前已发放数量'" json:"current_grant_count"`
	Status            int     `gorm:"column:status;type:tinyint;not null;default:1;comment:'状态:1-启用,0-禁用'" json:"status"`
	StartTime         string  `gorm:"column:start_time;type:varchar(50);not null;comment:'活动开始时间'" json:"start_time"`
	EndTime           string  `gorm:"column:end_time;type:varchar(50);not null;comment:'活动结束时间'" json:"end_time"`
}

// TableName 指定表名
func (pc *PromotionConfig) TableName() string {
	return "promotion_configs"
}

// GetByActivityCode 根据活动编码获取配置
func (pc *PromotionConfig) GetByActivityCode(code string) error {
	return global.DB.Where("activity_code = ?", code).First(pc).Error
}

// IsActive 检查活动是否有效
func (pc *PromotionConfig) IsActive() bool {
	if pc.Status != 1 {
		return false
	}

	now := time.Now()

	// 解析开始时间
	startTime, err := parseTimeString(pc.StartTime)
	if err != nil {
		return false
	}

	// 解析结束时间
	endTime, err := parseTimeString(pc.EndTime)
	if err != nil {
		return false
	}

	return startTime.Before(now) && endTime.After(now)
}

// parseTimeString 解析时间字符串
func parseTimeString(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("时间字符串为空")
	}

	// 支持的时间格式
	timeFormats := []string{
		"2006-01-02T15:04:05Z07:00", // RFC3339
		"2006-01-02T15:04:05Z",      // RFC3339 UTC
		"2006-01-02T15:04:05",       // ISO8601 without timezone
		"2006-01-02 15:04:05",       // MySQL datetime
		"2006-01-02T15:04",          // HTML datetime-local
		"2006-01-02",                // Date only
	}

	// 尝试各种时间格式
	for _, format := range timeFormats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("无法解析时间格式: %s", timeStr)
}

// CanGrant 检查是否可以发放优惠券
func (pc *PromotionConfig) CanGrant() bool {
	if !pc.IsActive() {
		return false
	}

	if pc.TotalGrantLimit > 0 && pc.CurrentGrantCount >= pc.TotalGrantLimit {
		return false
	}

	return true
}

// IncrementGrantCount 增加发放数量
func (pc *PromotionConfig) IncrementGrantCount() error {
	return global.DB.Model(pc).UpdateColumn("current_grant_count", gorm.Expr("current_grant_count + 1")).Error
}

// GenerateCouponCode 生成优惠券编码
func GenerateCouponCode() string {
	return fmt.Sprintf("CPN%s%06d", time.Now().Format("20060102"), time.Now().UnixNano()%1000000)
}

// CreateActivity 创建活动
func (pc *PromotionConfig) CreateActivity() error {
	return global.DB.Create(pc).Error
}

// UpdateActivity 更新活动
func (pc *PromotionConfig) UpdateActivity() error {
	return global.DB.Save(pc).Error
}

// DeleteActivity 删除活动
func (pc *PromotionConfig) DeleteActivity() error {
	return global.DB.Delete(pc).Error
}

// GetActivityList 获取活动列表
func GetActivityList(page, pageSize int, status *int, keyword string) ([]PromotionConfig, int64, error) {
	var activities []PromotionConfig
	var total int64

	query := global.DB.Model(&PromotionConfig{})

	// 状态筛选
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 关键词搜索
	if keyword != "" {
		query = query.Where("activity_name LIKE ? OR activity_code LIKE ? OR description LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&activities).Error; err != nil {
		return nil, 0, err
	}

	return activities, total, nil
}

// GetActivityByID 根据ID获取活动
func GetActivityByID(id uint) (*PromotionConfig, error) {
	var activity PromotionConfig
	if err := global.DB.Where("id = ?", id).First(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// GetUsersByCondition 根据条件获取用户列表
func GetUsersByCondition(condition string, registerAfter, registerBefore time.Time, maxUsers int) ([]uint, error) {
	var userIDs []uint

	query := global.DB.Table("users").Select("id")

	switch condition {
	case "new_users":
		// 新用户：注册时间在指定范围内
		if !registerAfter.IsZero() {
			query = query.Where("created_at >= ?", registerAfter)
		}
		if !registerBefore.IsZero() {
			query = query.Where("created_at <= ?", registerBefore)
		}
	case "vip_users":
		// VIP用户：可以根据实际业务逻辑调整
		query = query.Where("user_level > 0")
	case "all":
		// 所有用户
		break
	default:
		return nil, fmt.Errorf("不支持的用户条件: %s", condition)
	}

	// 限制最大用户数
	if maxUsers > 0 {
		query = query.Limit(maxUsers)
	}

	if err := query.Pluck("id", &userIDs).Error; err != nil {
		return nil, err
	}

	return userIDs, nil
}

// CouponUsageLog 优惠券使用记录模型
type CouponUsageLog struct {
	gorm.Model
	CouponID       uint    `gorm:"column:coupon_id;not null;index;comment:'优惠券ID'" json:"coupon_id"`
	UserID         uint    `gorm:"column:user_id;not null;index;comment:'用户ID'" json:"user_id"`
	OrderID        uint    `gorm:"column:order_id;not null;index;comment:'订单ID'" json:"order_id"`
	OrderSn        string  `gorm:"column:order_sn;type:varchar(50);not null;comment:'订单号'" json:"order_sn"`
	OriginalAmount float64 `gorm:"column:original_amount;type:decimal(10,2);not null;comment:'原始金额'" json:"original_amount"`
	DiscountAmount float64 `gorm:"column:discount_amount;type:decimal(10,2);not null;comment:'优惠金额'" json:"discount_amount"`
	FinalAmount    float64 `gorm:"column:final_amount;type:decimal(10,2);not null;comment:'最终金额'" json:"final_amount"`
}

// TableName 指定表名
func (cul *CouponUsageLog) TableName() string {
	return "coupon_usage_logs"
}

// CreateUsageLog 创建使用记录
func (cul *CouponUsageLog) CreateUsageLog() error {
	return global.DB.Create(cul).Error
}

// UserActivityRecord 用户活动参与记录模型
type UserActivityRecord struct {
	gorm.Model
	UserID       uint   `gorm:"column:user_id;not null;index;comment:'用户ID'" json:"user_id"`
	ActivityCode string `gorm:"column:activity_code;type:varchar(50);not null;comment:'活动编码'" json:"activity_code"`
	CouponCount  int    `gorm:"column:coupon_count;not null;default:0;comment:'获得优惠券数量'" json:"coupon_count"`
	Status       int    `gorm:"column:status;type:tinyint;not null;default:1;comment:'状态:1-成功,0-失败'" json:"status"`
}

// TableName 指定表名
func (uar *UserActivityRecord) TableName() string {
	return "user_activity_records"
}

// CreateRecord 创建参与记录
func (uar *UserActivityRecord) CreateRecord() error {
	return global.DB.Create(uar).Error
}

// HasParticipated 检查用户是否已参与活动
func HasParticipated(userID uint, activityCode string) (bool, error) {
	var count int64
	err := global.DB.Model(&UserActivityRecord{}).
		Where("user_id = ? AND activity_code = ?", userID, activityCode).
		Count(&count).Error
	return count > 0, err
}
