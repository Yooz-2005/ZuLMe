package router

import (
	"Api/trigger"
	"ZuLMe/ZuLMe/Common/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterInvoiceRoutes(r *gin.Engine) {

	invoiceGroup := r.Group("/invoice")
	{
		invoiceGroup.Use(pkg.JWTAuth("merchant"))
		{
			invoiceGroup.POST("/generate", trigger.GenerateInvoice)
		}

	}

}
