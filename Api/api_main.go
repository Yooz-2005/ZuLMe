package main

import (
	"Api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadUser(r)
	router.RegisterMerchantRoutes(r)
	r.Run(":8888")
}
