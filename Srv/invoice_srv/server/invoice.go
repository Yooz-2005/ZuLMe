package server

import (
	"context"
	"invoice_srv/internal/logic"
	invoice "invoice_srv/proto_invoice"
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

func (s ServerInvoice) ApplyInvoiceForUser(ctx context.Context, in *invoice.ApplyInvoiceForUserRequest) (*invoice.GenerateInvoiceResponse, error) {
	generateInvoice, err := logic.ApplyInvoiceForUser(in)
	if err != nil {
		return nil, err
	}
	return generateInvoice, nil
}
