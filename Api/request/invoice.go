package request

// GenerateInvoiceRequest 商家开发票请求
type GenerateInvoiceRequest struct {
	OrderID int `json:"order_id" form:"order_id"`
}

// ApplyInvoiceRequest 用户申请开发票请求
type ApplyInvoiceRequest struct {
	OrderID int `json:"order_id" form:"order_id" binding:"required"`
}
