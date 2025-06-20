package main

import (
	"Common/appconfig"
	"Common/initialize"
	"invoice_srv/grpc_invoice"
	"invoice_srv/internal"

	"google.golang.org/grpc"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	grpc_invoice.RegisterInvoiceGrpc(func(grpc *grpc.Server) {
		internal.RegisterInvoiceServer(grpc)
	})
}
