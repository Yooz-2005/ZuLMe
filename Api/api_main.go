package main

import (
	"Api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadUser(r)
	r.Run(":8888")
}
