package request

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	OrderID       uint     `json:"order_id" form:"order_id" binding:"required"`
	Rating        int32    `json:"rating" form:"rating" binding:"required,min=1,max=5"`                 // 1-5分
	Content       string   `json:"content" form:"content" binding:"required"`                           // 评论内容
	Images        []string `json:"images" form:"images"`                                                // 图片链接
	ServiceRating int32    `json:"service_rating" form:"service_rating" binding:"required,min=1,max=5"` // 服务评分
	VehicleRating int32    `json:"vehicle_rating" form:"vehicle_rating" binding:"required,min=1,max=5"` // 车辆评分
	CleanRating   int32    `json:"clean_rating" form:"clean_rating" binding:"required,min=1,max=5"`     // 清洁评分
	IsAnonymous   bool     `json:"is_anonymous" form:"is_anonymous"`                                    // 是否匿名
}

// UpdateCommentRequest 更新评论请求
type UpdateCommentRequest struct {
	Rating        int32    `json:"rating" form:"rating" binding:"required,min=1,max=5"`
	Content       string   `json:"content" form:"content" binding:"required"`
	Images        []string `json:"images" form:"images"`
	ServiceRating int32    `json:"service_rating" form:"service_rating" binding:"required,min=1,max=5"`
	VehicleRating int32    `json:"vehicle_rating" form:"vehicle_rating" binding:"required,min=1,max=5"`
	CleanRating   int32    `json:"clean_rating" form:"clean_rating" binding:"required,min=1,max=5"`
	IsAnonymous   bool     `json:"is_anonymous" form:"is_anonymous"`
}

// ReplyCommentRequest 商家回复评论请求
type ReplyCommentRequest struct {
	ReplyContent string `json:"reply_content" form:"reply_content" binding:"required"`
}

// GetVehicleCommentsRequest 获取车辆评论列表请求
type GetVehicleCommentsRequest struct {
	Page     int32 `json:"page" form:"page"`
	PageSize int32 `json:"page_size" form:"page_size"`
}

// GetUserCommentsRequest 获取用户评论列表请求
type GetUserCommentsRequest struct {
	Page     int32 `json:"page" form:"page"`
	PageSize int32 `json:"page_size" form:"page_size"`
}
