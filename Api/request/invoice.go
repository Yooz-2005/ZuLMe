package request

type GenerateInvoiceRequest struct {
	OrderID      int    `json:"order_id" form:"order_id"`
	InvoiceTitle string `json:"invoice_title"  form:"invoice_title"`
	TaxNumber    string `json:"tax_number"  form:"tax_number"`
}
