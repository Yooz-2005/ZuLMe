package router

import (
	"ZuLMe/ZuLMe/Api/trigger"
	"ZuLMe/ZuLMe/Common/pkg"
	"github.com/gin-gonic/gin"
)

func LoadUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", trigger.UserRegister)
		user.POST("/sendCode", trigger.SendCode)
		user.Use(pkg.JWTAuth("2209"))
		{
			user.POST("/profile", trigger.UpdateUserProfile)
			user.POST("/phone", trigger.UpdateUserPhone)
		}
	}
}
