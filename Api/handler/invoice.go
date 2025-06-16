package handler

import (
	"ZuLMe/ZuLMe/Api/client"
	invoice "ZuLMe/ZuLMe/Srv/invoice_srv/proto_invoice"
	"context"
)

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
