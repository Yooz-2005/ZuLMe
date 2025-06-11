package main

import (
	"Api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadUser(r)
	router.RegisterMerchantRoutes(r)
	router.RegisterAdminRoutes(r)
	router.RegisterVehicleRoutes(r)
	r.Run(":8888")
}
