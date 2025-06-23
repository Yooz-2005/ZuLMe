package router

import (
	"Api/trigger"
	jwt "Common/pkg"

	"github.com/gin-gonic/gin"
)

func LoadUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", trigger.UserRegister)
		user.POST("/sendCode", trigger.SendCode)
		// 距离计算接口 - 公开接口，不需要登录
		user.POST("/calculateDistance", trigger.CalculateDistance)
		user.Use(jwt.JWTAuth("2209"))
		{
			user.POST("/profile", trigger.UpdateUserProfile)
			user.POST("/phone", trigger.UpdateUserPhone)
			user.POST("/realName", trigger.RealName)
			user.POST("/collect", trigger.CollectVehicle)
			user.GET("/collectList", trigger.CollectVehicleList)
		}
	}
}
