package main

import (
	"Common/appconfig"
	"Common/global"
	"Common/initialize"
	"chat_srv/internal/server"
	"chat_srv/proto_chat"
	"fmt"
	"log"
	"models/model_mysql"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
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

	// 初始化MongoDB
	initialize.InitMongoDB()

	// 初始化Redis
	initialize.RedisInit()

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// 注册聊天服务
	chatServer := server.NewChatServer()
	proto_chat.RegisterChatServiceServer(grpcServer, chatServer)

	// 注册反射服务（用于调试）
	reflection.Register(grpcServer)

	// 迁移
	global.DB.AutoMigrate(&model_mysql.GroupChat{}, &model_mysql.GroupMember{})

	// 监听端口
	port := ":8003"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("聊天服务监听失败: %v", err)
	}

	fmt.Printf("Chat gRPC Server started on %s\n", port)

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("聊天服务启动失败: %v", err)
	}
}
