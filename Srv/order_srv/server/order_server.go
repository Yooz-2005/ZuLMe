package server

import (
	"ZuLMe/ZuLMe/Srv/order_srv/internal/logic"
	order "ZuLMe/ZuLMe/Srv/order_srv/proto_order"
	"context"
)

// OrderServer 订单服务器
type OrderServer struct {
	order.UnimplementedOrderServer
}

// NewOrderServer 创建订单服务器
func NewOrderServer() *OrderServer {
	return &OrderServer{}
}

// Ping 健康检查
func (s *OrderServer) Ping(ctx context.Context, req *order.Request) (*order.Response, error) {
	return logic.Ping(ctx, req)
}

// CreateOrderFromReservation 基于预订创建订单
func (s *OrderServer) CreateOrderFromReservation(ctx context.Context, req *order.CreateOrderFromReservationRequest) (*order.CreateOrderFromReservationResponse, error) {
	return logic.CreateOrderFromReservation(ctx, req)
}

// GetOrder 获取订单详情
func (s *OrderServer) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	return logic.GetOrder(ctx, req)
}

// UpdateOrderStatus 更新订单状态
func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	return logic.UpdateOrderStatus(ctx, req)
}

// AlipayNotify 支付宝异步通知
func (s *OrderServer) AlipayNotify(ctx context.Context, req *order.AlipayNotifyRequest) (*order.AlipayNotifyResponse, error) {
	return logic.AlipayNotify(ctx, req)
}
