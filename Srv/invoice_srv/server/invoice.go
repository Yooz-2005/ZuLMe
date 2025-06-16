package server

import (
	"ZuLMe/ZuLMe/Srv/invoice_srv/internal/logic"
	invoice "ZuLMe/ZuLMe/Srv/invoice_srv/proto_invoice"
	"context"
)

type ServerInvoice struct {
	invoice.UnimplementedInvoiceServer
}

func (s ServerInvoice) GenerateInvoice(ctx context.Context, in *invoice.GenerateInvoiceRequest) (*invoice.GenerateInvoiceResponse, error) {
	generateInvoice, err := logic.GenerateInvoice(in)
	if err != nil {
		return nil, err
	}
	return generateInvoice, nil
}
