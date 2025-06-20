package main

import (
	"Common/appconfig"
	"Common/initialize"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"vehicle_srv/grpc_vehicle"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.InitES()

	// 创建 gRPC 服务器
	gServer := grpc.NewServer()

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(gServer, health.NewServer())

	// 注册车辆服务
	grpc_vehicle.RegisterVehicleServices(gServer)

	// 监听端口
	lis, err := net.Listen("tcp", ":8004") // 车辆服务运行在8004端口
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	fmt.Println("Vehicle gRPC Server started on :8004")
	// 启动 gRPC 服务器
	if err := gServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
}
