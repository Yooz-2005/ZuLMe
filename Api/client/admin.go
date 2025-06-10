package client

import (
	admin "admin_srv/proto_admin"
	"context"
	"google.golang.org/grpc"
)

type HandlerAdmin func(ctx context.Context, in admin.AdminClient) (interface{}, error)

func AdminClient(ctx context.Context, handlerAdmin HandlerAdmin) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8003", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := admin.NewAdminClient(dial)
	res, err := handlerAdmin(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
