package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	user "user_srv/proto_user"
)

func UserRegister(c *gin.Context) {
	var data request.UserRegisterRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
	}
	register, err := handler.UserRegister(c, &user.UserRegisterRequest{
		Phone: data.Phone,
		Code:  data.Code,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

func SendCode(c *gin.Context) {
	var data request.SendCodeRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
	}
	code, err := handler.SendCode(c, &user.SendCodeRequest{
		Phone:  data.Phone,
		Source: data.Source,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, code)
}
