package main

import (
	"Common/appconfig"
	"Common/initialize"
	"Common/utils"
	"fmt"
	"merchant_srv/grpc_merchant"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.InitES()

	// 创建 gRPC 服务器
	gServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(gServer, health.NewServer())

	// 注册商家服务
	grpc_merchant.RegisterMerchantServices(gServer)

	// 启动同步商家地理位置任务
	utils.StartupSyncMerchantLocations()

	// 监听端口
	lis, err := net.Listen("tcp", ":8002") // 假设商家服务运行在8002端口
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	fmt.Println("Merchant gRPC Server started on :8002")
	// 启动 gRPC 服务器
	if err := gServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
}
