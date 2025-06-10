package router

import (
	"Api/trigger"
	"Common/pkg"
	"github.com/gin-gonic/gin"
)

func LoadUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		// 不需要驗證的路由
		user.POST("/register", trigger.UserRegister)
		user.POST("/sendCode", trigger.SendCode)

		// 需要驗證的路由
		auth := user.Group("")
		auth.Use(pkg.JWTAuth("your-secret-key")) // 請使用配置中的密鑰
		{
			// 在這裡添加需要驗證的路由
			// 例如：獲取用戶信息、更新用戶資料等
		}
	}
}
