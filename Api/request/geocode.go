package request

// GeocodeRequest 地理编码请求
type GeocodeRequest struct {
	Address string `json:"address" form:"address" binding:"required"` // 地址
}
