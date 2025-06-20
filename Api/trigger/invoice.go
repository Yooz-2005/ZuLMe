package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	invoice "invoice_srv/proto_invoice"
)

func GenerateInvoice(c *gin.Context) {
	var data request.GenerateInvoiceRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	merchantId := c.GetUint("userId")
	generateInvoice, err := handler.GenerateInvoice(c, &invoice.GenerateInvoiceRequest{
		OrderId:    int32(data.OrderID),
		MerchantId: int64(merchantId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, generateInvoice)
}
