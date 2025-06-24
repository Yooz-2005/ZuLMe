package handler

import (
	"Api/client"
	"context"
	invoice "invoice_srv/proto_invoice"
)

// GenerateInvoice 商家直接开发票
func GenerateInvoice(ctx context.Context, req *invoice.GenerateInvoiceRequest) (*invoice.GenerateInvoiceResponse, error) {
	invoiceClient, err := client.InvoiceClient(ctx, func(ctx context.Context, in invoice.InvoiceClient) (interface{}, error) {
		generateInvoice, err := in.GenerateInvoice(ctx, req)
		if err != nil {
			return nil, err
		}
		return generateInvoice, nil
	})
	if err != nil {
		return nil, err
	}
	return invoiceClient.(*invoice.GenerateInvoiceResponse), nil
}

// ApplyInvoiceForUser 用户申请开发票
func ApplyInvoiceForUser(ctx context.Context, req *invoice.ApplyInvoiceForUserRequest) (*invoice.GenerateInvoiceResponse, error) {
	invoiceClient, err := client.InvoiceClient(ctx, func(ctx context.Context, in invoice.InvoiceClient) (interface{}, error) {
		generateInvoice, err := in.ApplyInvoiceForUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return generateInvoice, nil
	})
	if err != nil {
		return nil, err
	}
	return invoiceClient.(*invoice.GenerateInvoiceResponse), nil
}
