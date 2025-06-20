package server

import (
	"context"

	"order_srv/internal/logic"
	order "order_srv/proto_order"
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

// GetUserOrderList 获取用户订单列表
func (s *OrderServer) GetUserOrderList(ctx context.Context, req *order.GetUserOrderListRequest) (*order.GetUserOrderListResponse, error) {
	return logic.GetUserOrderList(ctx, req)
}

// CancelOrder 取消订单
func (s *OrderServer) CancelOrder(ctx context.Context, req *order.CancelOrderRequest) (*order.CancelOrderResponse, error) {
	return logic.CancelOrder(ctx, req)
}
