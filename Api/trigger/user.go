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
		Username: data.UserName,
		Password: data.Password,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}
