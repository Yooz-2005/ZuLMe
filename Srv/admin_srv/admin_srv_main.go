package main

import (
	"Common/appconfig"
	"Common/initialize"
	"admin_srv/grpc_admin"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"net"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()

	// 创建 gRPC 服务器
	gServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(gServer, health.NewServer())

	// 注册 Admin 服务
	grpc_admin.RegisterAdminServices(gServer)

	// 监听端口
	lis, err := net.Listen("tcp", ":8008") // 假设Admin服务运行在8003端口
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	fmt.Println("Admin gRPC Server started on :8008")
	// 启动 gRPC 服务器
	if err := gServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
}
