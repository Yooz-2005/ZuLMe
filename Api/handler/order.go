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

// AlipayNotify 支付宝异步通知
func AlipayNotify(ctx context.Context, req *order.AlipayNotifyRequest) (*order.AlipayNotifyResponse, error) {
	orderClient, err := client.OrderClient(ctx, func(ctx context.Context, in order.OrderClient) (interface{}, error) {
		response, err := in.AlipayNotify(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	})
	if err != nil {
		return nil, err
	}
	return orderClient.(*order.AlipayNotifyResponse), nil
}
