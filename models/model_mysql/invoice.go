package model_mysql

import (
	"ZuLMe/ZuLMe/Common/global"
	"time"
)

type Invoice struct {
	Id          int32     `gorm:"column:id;type:int UNSIGNED;comment:id;primaryKey;not null;" json:"id"`                                     // id
	OrderId     int32     `gorm:"column:order_id;type:int;comment:订单id;not null;" json:"order_id"`                                         // 订单id
	MerchantId  int32     `gorm:"column:merchant_id;type:int;comment:开票人id;not null;" json:"merchant_id"`                                 // 开票人id
	OrderSn     string    `gorm:"column:order_sn;type:varchar(50);comment:订单编号;not null;" json:"order_sn"`                               // 订单编号
	InvoiceNo   string    `gorm:"column:invoice_no;type:varchar(50);comment:发票号码;not null;" json:"invoice_no"`                           // 发票号码
	InvoiceType int32     `gorm:"column:invoice_type;type:int;comment:发票类型1:电子发票2:纸质发票;not null;default:1;" json:"invoice_type"` // 发票类型1:电子发票2:纸质发票
	Title       string    `gorm:"column:title;type:varchar(50);comment:发票标题;not null;" json:"title"`                                     // 发票标题
	TaxNumber   string    `gorm:"column:tax_number;type:varchar(255);comment:税号;not null;" json:"tax_number"`                              // 税号
	Amount      float64   `gorm:"column:amount;type:decimal(10, 2);comment:发票金额;not null;" json:"amount"`                                // 发票金额
	Status      int32     `gorm:"column:status;type:int;comment:发票状态1:待开2:已开3:已作废;not null;default:1;" json:"status"`             // 发票状态1:待开2:已开3:已作废
	PdfUrl      string    `gorm:"column:pdf_url;type:varchar(255);comment:pdf文件url;not null;" json:"pdf_url"`                              // pdf文件url
	VehicleId   int32     `gorm:"column:vehicle_id;type:int;comment:车辆id;not null;" json:"vehicle_id"`                                     // 车辆id
	VehicleName string    `gorm:"column:vehicle_name;type:varchar(50);comment:车辆名称;not null;" json:"vehicle_name"`                       // 车辆名称
	RentalDays  int32     `gorm:"column:rental_days;type:int;comment:租赁天数;not null;" json:"rental_days"`                                 // 租赁天数
	DailyRate   float64   `gorm:"column:daily_rate;type:decimal(10, 2);comment:日租金;not null;" json:"daily_rate"`                          // 日租金
	PickupTime  time.Time `gorm:"column:pickup_time;type:datetime;comment:取车时间;not null;" json:"pickup_time"`                            // 取车时间
	ReturnTime  time.Time `gorm:"column:return_time;type:datetime;comment:还车时间;not null;" json:"return_time"`                            // 还车时间
	IssuedAt    time.Time `gorm:"column:issued_at;type:datetime;comment:开票时间;not null;" json:"issued_at"`                                // 开票时间
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`                           // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                           // 修改时间
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime;comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (i *Invoice) TableName() string {
	return "invoice"
}

// CreateInvoice 创建发票
func (i *Invoice) CreateInvoice() error {
	return global.DB.Create(i).Error
}

// GetInvoiceByID 根据ID获取发票
func (i *Invoice) GetInvoiceByID(id int32) error {
	return global.DB.First(i, id).Error
}

// GetInvoiceByOrderID 根据订单ID获取发票
func (i *Invoice) GetInvoiceByOrderID(orderID int32) error {
	return global.DB.Where("order_id = ?", orderID).First(i).Error
}

// UpdatePDFUrl 更新发票PDF URL
func (i *Invoice) UpdatePDFUrl(url string) error {
	return global.DB.Model(i).Update("pdf_url", url).Error
}

// UpdateStatus 更新发票状态
func (i *Invoice) UpdateStatus(status int32) error {
	return global.DB.Model(i).Update("status", status).Error
}
