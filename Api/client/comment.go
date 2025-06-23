package client

import (
	"comment_srv/proto_comment"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HandlerComment func(ctx context.Context, in proto_comment.CommentServiceClient) (interface{}, error)

func CommentClient(ctx context.Context, handlerComment HandlerComment) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer dial.Close()

	client := proto_comment.NewCommentServiceClient(dial)
	res, err := handlerComment(ctx, client)
	if err != nil {
		return nil, err
	}

	return res, nil
}
