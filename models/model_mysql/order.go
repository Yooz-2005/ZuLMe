package model_mysql

import (
	"Common/global"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Orders 订单模型（使用gorm.Model）
type Orders struct {
	gorm.Model                 // 包含ID, CreatedAt, UpdatedAt, DeletedAt
	UserId           uint      `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                                             // 用户id
	VehicleId        uint      `gorm:"column:vehicle_id;type:int;comment:汽车id;not null;" json:"vehicle_id"`                                       // 汽车id
	ReservationId    uint      `gorm:"column:reservation_id;type:int;comment:预订id;not null;" json:"reservation_id"`                               // 预订id（新增）
	OrderSn          string    `gorm:"column:order_sn;type:varchar(50);comment:订单编号;not null;uniqueIndex;" json:"order_sn"`                       // 订单编号
	PickupLocationId uint      `gorm:"column:pickup_location_id;type:int;comment:取车网点id;not null;" json:"pickup_location_id"`                     // 取车网点id
	ReturnLocationId uint      `gorm:"column:return_location_id;type:int;comment:还车网点id;not null;" json:"return_location_id"`                     // 还车网点id
	PickupTime       time.Time `gorm:"column:pickup_time;type:datetime;comment:取车时间;not null;" json:"pickup_time"`                                // 取车时间
	ReturnTime       time.Time `gorm:"column:return_time;type:datetime;comment:还车时间;not null;" json:"return_time"`                                // 还车时间
	RentalDays       int32     `gorm:"column:rental_days;type:int;comment:租赁天数;not null;" json:"rental_days"`                                     // 租赁天数
	DailyRate        float64   `gorm:"column:daily_rate;type:decimal(10, 2);comment:日租金;not null;" json:"daily_rate"`                             // 日租金
	TotalAmount      float64   `gorm:"column:total_amount;type:decimal(10, 2);comment:总金额;not null;" json:"total_amount"`                         // 总金额
	Status           int32     `gorm:"column:status;type:int;comment:订单状态1:待支付2:已支付3:已取消4:已完成5:已取车6:已还车;not null;default:1;" json:"status"`       // 订单状态1:待支付2:已支付3:已取消4:已完成5:已取车6:已还车
	Payment          int32     `gorm:"column:payment;type:int;comment:支付方式1:支付宝2:微信;not null;default:1;" json:"payment"`                          // 支付方式1:支付宝2:微信
	PaymentStatus    int32     `gorm:"column:payment_status;type:int;comment:支付状态1:待支付2:已支付3:已取消4:已完成;not null;default:1;" json:"payment_status"` // 支付状态1:待支付2:已支付3:已取消4:已完成
	PaymentUrl       string    `gorm:"column:payment_url;type:varchar(500);comment:支付链接;" json:"payment_url"`                                     // 支付链接（新增）
	AlipayTradeNo    string    `gorm:"column:alipay_trade_no;type:varchar(100);comment:支付宝交易号;" json:"alipay_trade_no"`                           // 支付宝交易号（新增）
	Notes            string    `gorm:"column:notes;type:varchar(500);comment:备注;" json:"notes"`                                                   // 备注
}

// TableName 指定表名
func (o *Orders) TableName() string {
	return "orders"
}

// Create 创建订单
func (o *Orders) Create() error {
	return global.DB.Create(o).Error
}

// GetByID 根据ID获取订单
func (o *Orders) GetByID(id uint) error {
	return global.DB.Where("id = ?", id).First(o).Error
}

// GetByOrderSn 根据订单号获取订单
func (o *Orders) GetByOrderSn(orderSn string) error {
	return global.DB.Where("order_sn = ?", orderSn).First(o).Error
}

// GetByReservationID 根据预订ID获取订单
func (o *Orders) GetByReservationID(reservationId uint) error {
	return global.DB.Where("reservation_id = ?", reservationId).First(o).Error
}

// UpdateStatus 更新订单状态
func (o *Orders) UpdateStatus(id uint, status int32) error {
	return global.DB.Model(&Orders{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateStatus 根据订单号更新订单状态
func (o *Orders) UpdateStatusByOrderSn(orderSn string, status int32) error {
	return global.DB.Model(&Orders{}).Where("order_sn =?", orderSn).Update("status", status).Error
}

// UpdatePaymentStatus 更新支付状态
func (o *Orders) UpdatePaymentStatus(id uint, paymentStatus int32) error {
	return global.DB.Model(&Orders{}).Where("id = ?", id).Update("payment_status", paymentStatus).Error
}

// UpdatePaymentInfo 更新支付信息
func (o *Orders) UpdatePaymentInfo(id uint, paymentUrl, alipayTradeNo string) error {
	updates := map[string]interface{}{
		"payment_url": paymentUrl,
	}
	if alipayTradeNo != "" {
		updates["alipay_trade_no"] = alipayTradeNo
	}
	return global.DB.Model(&Orders{}).Where("id = ?", id).Updates(updates).Error
}

// GenerateOrderSn 生成订单号
func GenerateOrderSn() string {
	return fmt.Sprintf("ORD%d", time.Now().Unix())
}

// GetUserOrderList 获取用户订单列表
func (o *Orders) GetUserOrderList(userID uint, page, pageSize int, status, paymentStatus string) ([]Orders, int64, error) {
	var orders []Orders
	var total int64

	query := global.DB.Model(&Orders{}).Where("user_id = ?", userID)

	// 状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 支付状态筛选
	if paymentStatus != "" {
		query = query.Where("payment_status = ?", paymentStatus)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// CreateOrderFromReservation 基于预订创建订单
func (o *Orders) CreateOrderFromReservation(reservation *VehicleInventory, vehicle *Vehicle, pickupLocationID, returnLocationID uint, notes string, paymentMethod int32) error {
	// 计算租赁天数
	rentalDays := int32(reservation.EndDate.Sub(reservation.StartDate).Hours()/24) + 1

	// 计算总金额
	totalAmount := vehicle.Price * float64(rentalDays)

	// 生成订单号
	orderSn := GenerateOrderSn()

	// 填充订单信息
	o.UserId = reservation.CreatedBy
	o.VehicleId = reservation.VehicleID
	o.ReservationId = reservation.ID
	o.OrderSn = orderSn
	o.PickupLocationId = pickupLocationID
	o.ReturnLocationId = returnLocationID
	o.PickupTime = reservation.StartDate
	o.ReturnTime = reservation.EndDate
	o.RentalDays = rentalDays
	o.DailyRate = vehicle.Price
	o.TotalAmount = totalAmount
	o.Status = 1 // 待支付
	o.Payment = paymentMethod
	o.PaymentStatus = 1 // 待支付
	o.Notes = notes

	return o.Create()
}

// CreateOrderFromReservationWithTx 基于预订创建订单（带事务）
func (o *Orders) CreateOrderFromReservationWithTx(tx *gorm.DB, reservation *VehicleInventory, vehicle *Vehicle, pickupLocationID, returnLocationID uint, notes string, paymentMethod int32) error {
	// 计算租赁天数
	rentalDays := int32(reservation.EndDate.Sub(reservation.StartDate).Hours()/24) + 1

	// 计算总金额
	totalAmount := vehicle.Price * float64(rentalDays)

	// 生成订单号
	orderSn := GenerateOrderSn()

	// 填充订单信息
	o.UserId = reservation.CreatedBy
	o.VehicleId = reservation.VehicleID
	o.ReservationId = reservation.ID
	o.OrderSn = orderSn
	o.PickupLocationId = pickupLocationID
	o.ReturnLocationId = returnLocationID
	o.PickupTime = reservation.StartDate
	o.ReturnTime = reservation.EndDate
	o.RentalDays = rentalDays
	o.DailyRate = vehicle.Price
	o.TotalAmount = totalAmount
	o.Status = 1 // 待支付
	o.Payment = paymentMethod
	o.PaymentStatus = 1 // 待支付
	o.Notes = notes

	// 使用事务创建订单
	return tx.Create(o).Error
}

// 订单状态常量
const (
	OrderStatusPending   = 1 // 待支付
	OrderStatusPaid      = 2 // 已支付
	OrderStatusCancelled = 3 // 已取消
	OrderStatusCompleted = 4 // 已完成
	OrderStatusPickedUp  = 5 // 已取车
	OrderStatusReturned  = 6 // 已还车
)

// 支付状态常量
const (
	PaymentStatusPending   = 1 // 待支付
	PaymentStatusPaid      = 2 // 已支付
	PaymentStatusCancelled = 3 // 已取消
	PaymentStatusCompleted = 4 // 已完成
)

// 支付方式常量
const (
	PaymentMethodAlipay = 1 // 支付宝
	PaymentMethodWechat = 2 // 微信
)
