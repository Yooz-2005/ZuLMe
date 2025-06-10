package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	merchant "merchant_srv/proto_merchant"
)

func MerchantRegisterHandler(c *gin.Context) {
	var req request.MerchantRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	registerRes, err := handler.MerchantRegister(c, &merchant.MerchantRegisterRequest{
		Name:        req.Name,
		Phone:       req.Phone,
		Email:       req.Email,
		Password:    req.Password,
		ConfirmPass: req.ConfirmPass,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if registerRes.Code != 200 {
		response.ResponseError400(c, registerRes.Message)
		return
	}

	response.ResponseSuccess(c, registerRes.Message)
}

func MerchantLoginHandler(c *gin.Context) {
	var req request.MerchantLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	loginRes, err := handler.MerchantLogin(c, &merchant.MerchantLoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if loginRes.Code != 200 {
		response.ResponseError400(c, loginRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{"message": loginRes.Message, "token": loginRes.Token})
}
