package router

import (
	"Api/trigger"
	"github.com/gin-gonic/gin"
)

func LoadUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", trigger.UserRegister)
		user.POST("/sendCode", trigger.SendCode)
	}
}
