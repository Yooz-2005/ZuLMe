package internal

import (
	invoice "ZuLMe/ZuLMe/Srv/invoice_srv/proto_invoice"
	"ZuLMe/ZuLMe/Srv/invoice_srv/server"
	"google.golang.org/grpc"
)

func RegisterInvoiceServer(ser *grpc.Server) {
	invoice.RegisterInvoiceServer(ser, server.ServerInvoice{})
}
