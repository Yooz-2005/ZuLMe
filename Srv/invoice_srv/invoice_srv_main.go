package main

import (
	"ZuLMe/ZuLMe/Common/appconfig"
	"ZuLMe/ZuLMe/Common/initialize"
	"ZuLMe/ZuLMe/Srv/invoice_srv/grpc_invoice"
	"ZuLMe/ZuLMe/Srv/invoice_srv/internal"
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
