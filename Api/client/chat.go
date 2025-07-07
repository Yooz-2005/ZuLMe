package client

import (
	chat "chat_srv/proto_chat"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ChatClient 聊天服务客户端
func ChatClient(ctx context.Context, call func(context.Context, chat.ChatServiceClient) (interface{}, error)) (interface{}, error) {
	// 连接聊天服务
	conn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("连接聊天服务失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := chat.NewChatServiceClient(conn)

	// 调用服务
	return call(ctx, client)
}
