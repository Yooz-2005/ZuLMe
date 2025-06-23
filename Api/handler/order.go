package handler

import (
	"Api/client"
	"context"
	order "order_srv/proto_order"
)

// CreateOrderFromReservation 基于预订创建订单
func CreateOrderFromReservation(ctx context.Context, req *order.CreateOrderFromReservationRequest) (*order.CreateOrderFromReservationResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.CreateOrderFromReservation(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CreateOrderFromReservationResponse), nil
}

// GetOrder 获取订单详情
func GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.GetOrder(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.GetOrderResponse), nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.UpdateOrderStatus(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.UpdateOrderStatusResponse), nil
}

// GetUserOrderList 获取用户订单列表
func GetUserOrderList(ctx context.Context, req *order.GetUserOrderListRequest) (*order.GetUserOrderListResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.GetUserOrderList(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.GetUserOrderListResponse), nil
}

// CancelOrder 取消订单
func CancelOrder(ctx context.Context, req *order.CancelOrderRequest) (*order.CancelOrderResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.CancelOrder(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CancelOrderResponse), nil
}

// CheckUserUnpaidOrder 检查用户未支付订单
func CheckUserUnpaidOrder(ctx context.Context, req *order.CheckUserUnpaidOrderRequest) (*order.CheckUserUnpaidOrderResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.CheckUserUnpaidOrder(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.CheckUserUnpaidOrderResponse), nil
}
