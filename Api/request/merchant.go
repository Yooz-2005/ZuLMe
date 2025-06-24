package request

type MerchantRegisterRequest struct {
	Name         string  `json:"name" form:"name"`
	Phone        string  `json:"phone" form:"phone"`
	Email        string  `json:"email" form:"email"`
	Password     string  `json:"password" form:"password"`
	ConfirmPass  string  `json:"confirm_pass" form:"confirm_pass"`
	Location     string  `json:"location" form:"location"`           // 网点地址
	BusinessTime string  `json:"business_time" form:"business_time"` // 营业时间
	Longitude    float64 `json:"longitude" form:"longitude"`         // 经度（可选，如果不提供会自动获取）
	Latitude     float64 `json:"latitude" form:"latitude"`           // 纬度（可选，如果不提供会自动获取）
}

type MerchantLoginRequest struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type GetLocationListRequest struct {
	Page         int32 `json:"page" form:"page"`
	PageSize     int32 `json:"page_size" form:"page_size"`
	StatusFilter int32 `json:"status_filter" form:"status_filter"` // 筛选审核状态：0-未审核，1-审核通过，2-审核失败
}
