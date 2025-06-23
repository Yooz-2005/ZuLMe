package request

type GenerateInvoiceRequest struct {
	OrderID int `json:"order_id" form:"order_id"`
}
