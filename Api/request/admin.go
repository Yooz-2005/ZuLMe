package request

// MerchantApproveRequest 审核商户请求
type MerchantApproveRequest struct {
	ID     int64 `json:"id" form:"id" binding:"required"`         // 商户ID
	Status int64 `json:"status" form:"status" binding:"required"` // 审核状态：0-未审核，1-审核通过，2-审核失败
}

// MerchantUpdateRequest 编辑商户请求
type MerchantUpdateRequest struct {
	ID    int64  `json:"id" form:"id" binding:"required"`
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Email string `json:"email" form:"email" binding:"omitempty,email"`
}

// MerchantDeleteRequest 删除商户请求
type MerchantDeleteRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

// MerchantListRequest 获取商户列表请求
type MerchantListRequest struct {
	Page        int64  `json:"page" form:"page"`
	PageSize    int64  `json:"page_size" form:"page_size"`
	Keyword     string `json:"keyword" form:"keyword"`             // 搜索关键词
	StatusFilter int64  `json:"status_filter" form:"status_filter"` // 筛选审核状态：0-未审核，1-审核通过，2-审核失败
}

// MerchantDetailRequest 获取商户详情请求
type MerchantDetailRequest struct {
	ID int64 `json:"id" form:"id" uri:"id" binding:"required"`
}
