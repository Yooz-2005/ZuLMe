package trigger

import (
	"ZuLMe/ZuLMe/Api/handler"
	"ZuLMe/ZuLMe/Api/request"
	"ZuLMe/ZuLMe/Api/response"
	invoice "ZuLMe/ZuLMe/Srv/invoice_srv/proto_invoice"
	"github.com/gin-gonic/gin"
)

func GenerateInvoice(c *gin.Context) {
	var data request.GenerateInvoiceRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	merchantId := c.GetUint("userId")
	generateInvoice, err := handler.GenerateInvoice(c, &invoice.GenerateInvoiceRequest{
		OrderId:      int32(data.OrderID),
		InvoiceTitle: data.InvoiceTitle,
		TaxNumber:    data.TaxNumber,
		MerchantId:   int64(merchantId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, generateInvoice)
}
