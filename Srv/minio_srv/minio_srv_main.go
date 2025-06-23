package main

import (
	"Common/appconfig"
	"Common/initialize"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"minio_srv/grpc_minio"
	"net"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.InitES()
	initialize.MinioInit()
	// 创建 gRPC 服务器
	gServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(gServer, health.NewServer())

	// 注册Minio服务
	grpc_minio.RegisterMinioServices(gServer)
	// 监听端口
	lis, err := net.Listen("tcp", ":8007") // minio服务运行在8007端口
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	fmt.Println("Vehicle gRPC Server started on :8007")
	// 启动 gRPC 服务器
	if err := gServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
}
