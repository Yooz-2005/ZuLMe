package client

import (
	user "ZuLMe/ZuLMe/Srv/user_srv/proto_user"
	"context"
	"google.golang.org/grpc"
)

type HandlerUser func(ctx context.Context, in user.UserClient) (interface{}, error)

func UserClient(ctx context.Context, handlerUser HandlerUser) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := user.NewUserClient(dial)
	res, err := handlerUser(ctx, client)
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	return res, nil
}
