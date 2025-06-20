package grpc_invoice

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func RegisterInvoiceGrpc(call func(grpc *grpc.Server)) {
	// 1.监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s", "127.0.0.1:8006"))
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	call(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("server failed to listening at %v", err)
	}

}
