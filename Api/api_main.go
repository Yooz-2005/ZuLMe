package main

import (
	"Api/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 允许前端域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 註冊路由
	router.LoadUser(r)
	router.RegisterMerchantRoutes(r)
	router.RegisterAdminRoutes(r)
	router.RegisterVehicleRoutes(r)
	router.RegisterInvoiceRoutes(r)
	router.LoadOrder(r)
	router.CommentRouter(r)
	router.RegisterMinioRoutes(r)
	r.Run(":8888")
}
