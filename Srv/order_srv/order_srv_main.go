package main

import (
	"Common/appconfig"
	"Common/initialize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	order "order_srv/proto_order"
	"order_srv/server"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.InitES()

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()

	// 注册订单服务
	orderServer := server.NewOrderServer()
	order.RegisterOrderServer(grpcServer, orderServer)

	// 注册反射服务（用于调试）
	reflection.Register(grpcServer)

	// 自动迁移数据库
	//global.DB.AutoMigrate(&model_mysql.Orders{})

	// 监听端口
	port := ":9093"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("订单服务监听失败: %v", err)
	}

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("订单服务启动失败: %v", err)
	}
}
