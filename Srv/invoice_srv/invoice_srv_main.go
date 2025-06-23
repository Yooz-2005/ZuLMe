package main

import (
	"Common/appconfig"
	"Common/initialize"
	"google.golang.org/grpc"
	"invoice_srv/grpc_invoice"
	"invoice_srv/internal"
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
