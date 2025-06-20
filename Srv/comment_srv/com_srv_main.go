package main

import (
	"Common/appconfig"
	"Common/initialize"
	"comment_srv/internal/server"
	"comment_srv/proto_comment"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()

	// 初始化数据库连接
	_, err := initialize.MysqlInit()
	if err != nil {
		log.Fatalf("MySQL初始化失败: %v", err)
	}

	initialize.InitMongoDB()

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// 注册评论服务
	commentServer := server.NewCommentServer()
	proto_comment.RegisterCommentServiceServer(grpcServer, commentServer)

	// 监听端口
	port := ":8005" // 评论服务运行在8005端口
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("评论服务监听失败: %v", err)
	}

	fmt.Printf("Comment gRPC Server started on %s\n", port)

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("评论服务启动失败: %v", err)
	}
}
