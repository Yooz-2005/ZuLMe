package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
)

// AdminMerchantApproveHandler 审核商户 HTTP 处理器
func AdminMerchantApproveHandler(c *gin.Context) {
	var req request.MerchantApproveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	approveRes, err := handler.MerchantApprove(c, &req)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if approveRes.Code != 200 {
		response.ResponseError400(c, approveRes.Message)
		return
	}

	response.ResponseSuccess(c, approveRes.Message)
}

// AdminMerchantUpdateHandler 编辑商户 HTTP 处理器
func AdminMerchantUpdateHandler(c *gin.Context) {
	var req request.MerchantUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	updateRes, err := handler.MerchantUpdate(c, &req)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if updateRes.Code != 200 {
		response.ResponseError400(c, updateRes.Message)
		return
	}

	response.ResponseSuccess(c, updateRes.Message)
}

// AdminMerchantDeleteHandler 删除商户 HTTP 处理器
func AdminMerchantDeleteHandler(c *gin.Context) {
	var req request.MerchantDeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	deleteRes, err := handler.MerchantDelete(c, &req)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if deleteRes.Code != 200 {
		response.ResponseError400(c, deleteRes.Message)
		return
	}

	response.ResponseSuccess(c, deleteRes.Message)
}

// AdminMerchantListHandler 获取商户列表 HTTP 处理器
func AdminMerchantListHandler(c *gin.Context) {
	var req request.MerchantListRequest
	if err := c.ShouldBindQuery(&req); err != nil { // 列表通常使用Query参数
		response.ResponseError400(c, err.Error())
		return
	}

	listRes, err := handler.MerchantList(c, &req)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if listRes.Code != 200 {
		response.ResponseError400(c, listRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":   listRes.Message,
		"merchants": listRes.Merchants,
		"total":     listRes.Total,
	})
}

// AdminMerchantDetailHandler 获取商户详情 HTTP 处理器
func AdminMerchantDetailHandler(c *gin.Context) {
	var req request.MerchantDetailRequest
	if err := c.ShouldBindUri(&req); err != nil { // 详情通常从URI获取ID
		response.ResponseError400(c, err.Error())
		return
	}

	detailRes, err := handler.MerchantDetail(c, &req)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if detailRes.Code != 200 {
		response.ResponseError400(c, detailRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":  detailRes.Message,
		"merchant": detailRes.Merchant,
	})
}
