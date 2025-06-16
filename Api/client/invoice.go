package client

import (
	invoice "ZuLMe/ZuLMe/Srv/invoice_srv/proto_invoice"
	"context"
	"google.golang.org/grpc"
)

type HandlerInvoice func(ctx context.Context, in invoice.InvoiceClient) (interface{}, error)

func InvoiceClient(ctx context.Context, handlerInvoice HandlerInvoice) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8006", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := invoice.NewInvoiceClient(dial)
	res, err := handlerInvoice(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
