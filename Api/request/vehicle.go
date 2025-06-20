package request

// CreateVehicleRequest 创建车辆请求
type CreateVehicleRequest struct {
	MerchantID  int64   `json:"merchant_id" form:"merchant_id"`
	TypeID      int64   `json:"type_id" form:"type_id" binding:"required"`
	BrandID     int64   `json:"brand_id" form:"brand_id" binding:"required"`
	Brand       string  `json:"brand" form:"brand"` // 保留用于兼容性，但不再必需
	Style       string  `json:"style" form:"style" binding:"required"`
	Year        int64   `json:"year" form:"year" binding:"required"`
	Color       string  `json:"color" form:"color"`
	Mileage     int64   `json:"mileage" form:"mileage"`
	Price       float64 `json:"price" form:"price" binding:"required"`
	Status      int64   `json:"status" form:"status"`
	Description string  `json:"description" form:"description"`
	Images      string  `json:"images" form:"images"`
	Contact     string  `json:"contact" form:"contact"`
}

// UpdateVehicleRequest 更新车辆请求
type UpdateVehicleRequest struct {
	ID          int64   `json:"id" form:"id" binding:"required"`
	MerchantID  int64   `json:"merchant_id"`
	TypeID      int64   `json:"type_id" form:"type_id"`
	BrandID     int64   `json:"brand_id" form:"brand_id"`
	Brand       string  `json:"brand" form:"brand"` // 保留用于兼容性
	Style       string  `json:"style" form:"style"`
	Year        int64   `json:"year" form:"year"`
	Color       string  `json:"color" form:"color"`
	Mileage     int64   `json:"mileage" form:"mileage"`
	Price       float64 `json:"price" form:"price"`
	Status      int64   `json:"status" form:"status"`
	Description string  `json:"description" form:"description"`
	Images      string  `json:"images" form:"images"`
	Location    string  `json:"location" form:"location"`
	Contact     string  `json:"contact" form:"contact"`
}

// DeleteVehicleRequest 删除车辆请求
type DeleteVehicleRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// GetVehicleRequest 获取车辆详情请求
type GetVehicleRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// ListVehiclesRequest 获取车辆列表请求
type ListVehiclesRequest struct {
	Page       int64   `json:"page" form:"page"`
	PageSize   int64   `json:"page_size" form:"page_size"`
	Keyword    string  `json:"keyword" form:"keyword"`
	MerchantID int64   `json:"merchant_id" form:"merchant_id"`
	TypeID     int64   `json:"type_id" form:"type_id"`
	BrandID    int64   `json:"brand_id" form:"brand_id"`
	Status     int64   `json:"status" form:"status"`
	PriceMin   float64 `json:"price_min" form:"price_min"`
	PriceMax   float64 `json:"price_max" form:"price_max"`
	YearMin    int64   `json:"year_min" form:"year_min"`
	YearMax    int64   `json:"year_max" form:"year_max"`
}

// ==================== 车辆类型请求结构体 ====================

// CreateVehicleTypeRequest 创建车辆类型请求
type CreateVehicleTypeRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Status      int64  `json:"status" form:"status"`
	Sort        int64  `json:"sort" form:"sort"`
}

// UpdateVehicleTypeRequest 更新车辆类型请求
type UpdateVehicleTypeRequest struct {
	ID          int64  `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Status      int64  `json:"status" form:"status"`
	Sort        int64  `json:"sort" form:"sort"`
}

// DeleteVehicleTypeRequest 删除车辆类型请求
type DeleteVehicleTypeRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// GetVehicleTypeRequest 获取车辆类型详情请求
type GetVehicleTypeRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// ListVehicleTypesRequest 获取车辆类型列表请求
type ListVehicleTypesRequest struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   int64  `json:"status" form:"status"`
}

// ==================== 车辆品牌请求结构体 ====================

// CreateVehicleBrandRequest 创建车辆品牌请求
type CreateVehicleBrandRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	EnglishName string `json:"english_name" form:"english_name"`
	Logo        string `json:"logo" form:"logo"`
	Country     string `json:"country" form:"country"`
	Description string `json:"description" form:"description"`
	Status      int64  `json:"status" form:"status"`
	Sort        int64  `json:"sort" form:"sort"`
	IsHot       int64  `json:"is_hot" form:"is_hot"`
}

// UpdateVehicleBrandRequest 更新车辆品牌请求
type UpdateVehicleBrandRequest struct {
	ID          int64  `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	EnglishName string `json:"english_name" form:"english_name"`
	Logo        string `json:"logo" form:"logo"`
	Country     string `json:"country" form:"country"`
	Description string `json:"description" form:"description"`
	Status      int64  `json:"status" form:"status"`
	Sort        int64  `json:"sort" form:"sort"`
	IsHot       int64  `json:"is_hot" form:"is_hot"`
}

// DeleteVehicleBrandRequest 删除车辆品牌请求
type DeleteVehicleBrandRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// GetVehicleBrandRequest 获取车辆品牌详情请求
type GetVehicleBrandRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// ListVehicleBrandsRequest 获取车辆品牌列表请求
type ListVehicleBrandsRequest struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   int64  `json:"status" form:"status"`
	IsHot    int64  `json:"is_hot" form:"is_hot"`
}

// ==================== 车辆库存请求结构体 ====================

// CheckAvailabilityRequest 检查车辆可用性请求
type CheckAvailabilityRequest struct {
	VehicleID int64  `json:"vehicle_id" form:"vehicle_id" binding:"required"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate   string `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
}

// CreateReservationRequest 创建预订请求（新流程：先预订后订单）
type CreateReservationRequest struct {
	VehicleID int64  `json:"vehicle_id" form:"vehicle_id" binding:"required"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate   string `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
	Notes     string `json:"notes" form:"notes"`                              // 预订备注
	// UserID 从JWT token中获取，不需要在请求中传递
	// OrderID 将在创建订单时关联，预订时不需要
}

// UpdateReservationStatusRequest 更新预订状态请求
type UpdateReservationStatusRequest struct {
	OrderID int64  `json:"order_id" form:"order_id" binding:"required"`
	Status  string `json:"status" form:"status" binding:"required"` // rented, completed, cancelled
}

// CancelReservationRequest 取消预订请求
type CancelReservationRequest struct {
	ReservationID string `json:"reservation_id" form:"reservation_id" binding:"required"` // 预订ID，格式如 RES123
}

// GetAvailableVehiclesRequest 获取可用车辆请求
type GetAvailableVehiclesRequest struct {
	StartDate  string  `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate    string  `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
	MerchantID int64   `json:"merchant_id" form:"merchant_id"`                  // 可选，按商家筛选
	TypeID     int64   `json:"type_id" form:"type_id"`                          // 可选，按类型筛选
	BrandID    int64   `json:"brand_id" form:"brand_id"`                        // 可选，按品牌筛选
	Status     int64   `json:"status" form:"status"`                            // 可选，按库存状态筛选 -1:全部 0:默认可用 1:可租用 2:已预订 3:租用中 4:维护中 5:不可用
	PriceMin   float64 `json:"price_min" form:"price_min"`                      // 可选，最低价格
	PriceMax   float64 `json:"price_max" form:"price_max"`                      // 可选，最高价格
	Page       int64   `json:"page" form:"page"`                                // 页码
	PageSize   int64   `json:"page_size" form:"page_size"`                      // 每页数量
}

// GetInventoryStatsRequest 获取库存统计请求
type GetInventoryStatsRequest struct {
	MerchantID int64 `json:"merchant_id" form:"merchant_id" binding:"required"`
}

// SetMaintenanceRequest 设置维护状态请求
type SetMaintenanceRequest struct {
	VehicleID int64  `json:"vehicle_id" form:"vehicle_id" binding:"required"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate   string `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
	Notes     string `json:"notes" form:"notes"`                              // 维护备注
	CreatedBy int64  `json:"created_by" form:"created_by"`                    // 创建人ID
}

// GetMaintenanceScheduleRequest 获取维护计划请求
type GetMaintenanceScheduleRequest struct {
	VehicleID int64 `json:"vehicle_id" form:"vehicle_id" binding:"required"`
}

// GetInventoryCalendarRequest 获取库存日历请求
type GetInventoryCalendarRequest struct {
	VehicleID int64  `json:"vehicle_id" form:"vehicle_id" binding:"required"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate   string `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
}

// GetInventoryReportRequest 获取库存报表请求
type GetInventoryReportRequest struct {
	MerchantID int64  `json:"merchant_id" form:"merchant_id"`
	StartDate  string `json:"start_date" form:"start_date" binding:"required"` // 格式: YYYY-MM-DD
	EndDate    string `json:"end_date" form:"end_date" binding:"required"`     // 格式: YYYY-MM-DD
}
