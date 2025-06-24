package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	invoice "invoice_srv/proto_invoice"

	"github.com/gin-gonic/gin"
)

// GenerateInvoice 商家直接开发票
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

// ApplyInvoiceForUser 用户申请开发票
func ApplyInvoiceForUser(c *gin.Context) {
	var data request.ApplyInvoiceRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	// 获取当前登录用户ID
	userId := c.GetUint("userId")

	// 调用处理函数，自动查找订单对应的商家并开具发票
	generateInvoice, err := handler.ApplyInvoiceForUser(c, &invoice.ApplyInvoiceForUserRequest{
		OrderId: int32(data.OrderID),
		UserId:  int64(userId),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, generateInvoice)
}
