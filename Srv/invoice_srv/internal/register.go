package internal

import (
	"google.golang.org/grpc"
	invoice "invoice_srv/proto_invoice"
	"invoice_srv/server"
)

func RegisterInvoiceServer(ser *grpc.Server) {
	invoice.RegisterInvoiceServer(ser, server.ServerInvoice{})
}
