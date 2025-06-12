package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"github.com/gin-gonic/gin"
	user "user_srv/proto_user"
)

// todo用户注册登录
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

// todo发送验证码
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

// todo修改个人信息
func UpdateUserProfile(c *gin.Context) {
	var data request.UpdateUserProfileRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	userId := c.GetUint("userId")
	updateUserProfile, err := handler.UpdateUserProfile(c, &user.UpdateUserProfileRequest{
		UserId:         int64(userId),
		RealName:       data.RealName,
		IdType:         data.IdType,
		IdNumber:       data.IdNumber,
		IdExpireDate:   data.IdExpireDate,
		Email:          data.Email,
		Province:       data.Province,
		City:           data.City,
		District:       data.District,
		EmergencyName:  data.EmergencyName,
		EmergencyPhone: data.EmergencyPhone,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, updateUserProfile)
}

// todo修改手机号
func UpdateUserPhone(c *gin.Context) {
	var data request.UpdateUserPhoneRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	userId := c.GetUint("userId")
	updateUserPhone, err := handler.UpdateUserPhone(c, &user.UpdateUserPhoneRequest{
		UserId: int64(userId),
		Phone:  data.Phone,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, updateUserPhone)
}
