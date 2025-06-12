package model_mysql

import "time"

type Orders struct {
	Id               int32     `gorm:"column:id;type:int UNSIGNED;comment:id;primaryKey;not null;" json:"id"`                                                     // id
	UserId           int32     `gorm:"column:user_id;type:int;comment:用户id;not null;" json:"user_id"`                                                           // 用户id
	VehicleId        int32     `gorm:"column:vehicle_id;type:int;comment:汽车id;not null;" json:"vehicle_id"`                                                     // 汽车id
	OrderSn          string    `gorm:"column:order_sn;type:varchar(50);comment:订单编号;not null;" json:"order_sn"`                                               // 订单编号
	PickupLocationId int32     `gorm:"column:pickup_location_id;type:int;comment:取车网点id;not null;" json:"pickup_location_id"`                                 // 取车网点id
	ReturnLocationId int32     `gorm:"column:return_location_id;type:int;comment:还车网点id;not null;" json:"return_location_id"`                                 // 还车网点id
	PickupTime       time.Time `gorm:"column:pickup_time;type:datetime;comment:取车时间;not null;" json:"pickup_time"`                                            // 取车时间
	ReturnTime       time.Time `gorm:"column:return_time;type:datetime;comment:还车时间;not null;" json:"return_time"`                                            // 还车时间
	RentalDays       int32     `gorm:"column:rental_days;type:int;comment:租赁天数;not null;" json:"rental_days"`                                                 // 租赁天数
	DailyRate        float64   `gorm:"column:daily_rate;type:decimal(10, 2);comment:日租金;not null;" json:"daily_rate"`                                          // 日租金
	TotalAmount      float64   `gorm:"column:total_amount;type:decimal(10, 2);comment:总金额;not null;" json:"total_amount"`                                      // 总金额
	Status           int32     `gorm:"column:status;type:int;comment:订单状态1:待支付2:已支付3:已取消4:已完成5:已取车6:已还车;not null;default:1;" json:"status"` // 订单状态1:待支付2:已支付3:已取消4:已完成5:已取车6:已还车
	Payment          int32     `gorm:"column:payment;type:int;comment:支付方式1:支付宝2:微信;not null;default:1;" json:"payment"`                                 // 支付方式1:支付宝2:微信
	PaymentStatus    int32     `gorm:"column:payment_status;type:int;comment:支付状态1:待支付2:已支付3:已取消4:已完成;not null;" json:"payment_status"`           // 支付状态1:待支付2:已支付3:已取消4:已完成
	Notes            string    `gorm:"column:notes;type:varchar(255);comment:备注;not null;" json:"notes"`                                                        // 备注
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`                                           // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                                           // 修改时间
	DeletedAt        time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                                       // 删除时间
}

func (O *Orders) TableName() string {
	return "orders"
}
