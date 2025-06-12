package request

// CreateVehicleRequest 创建车辆请求
type CreateVehicleRequest struct {
	MerchantID  int64   `json:"merchant_id" form:"merchant_id"`
	TypeID      int64   `json:"type_id" form:"type_id" binding:"required"`
	Brand       string  `json:"brand" form:"brand" binding:"required"`
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
	Brand       string  `json:"brand" form:"brand"`
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
