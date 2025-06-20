package logic

import (
	"ZuLMe/ZuLMe/Common/global"
	"ZuLMe/ZuLMe/Common/payment"
	"ZuLMe/ZuLMe/models/model_mysql"
	"context"
	"fmt"
	order "order_srv/proto_order"
	"time"
)

// Ping 健康检查
func Ping(ctx context.Context, req *order.Request) (*order.Response, error) {
	return &order.Response{
		Pong: "order service pong",
	}, nil
}

// CreateOrderFromReservation 基于预订创建订单
func CreateOrderFromReservation(ctx context.Context, req *order.CreateOrderFromReservationRequest) (*order.CreateOrderFromReservationResponse, error) {
	// 参数验证
	if req.ReservationId <= 0 {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "预订ID不能为空",
		}, nil
	}

	if req.UserId <= 0 {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	// 1. 验证预订是否存在且属于当前用户
	var reservation model_mysql.VehicleInventory
	if err := reservation.GetByID(uint(req.ReservationId)); err != nil {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "预订不存在",
		}, nil
	}

	if reservation.CreatedBy != uint(req.UserId) {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "无权操作此预订",
		}, nil
	}

	if reservation.OrderID != 0 {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "该预订已创建订单",
		}, nil
	}

	// 2. 获取车辆信息
	var vehicle model_mysql.Vehicle
	if err := vehicle.GetByID(reservation.VehicleID); err != nil {
		return &order.CreateOrderFromReservationResponse{
			Code:    400,
			Message: "车辆不存在",
		}, nil
	}

	// 3. 开启事务创建订单
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	var orderModel model_mysql.Orders
	if err := orderModel.CreateOrderFromReservationWithTx(tx, &reservation, &vehicle, uint(req.PickupLocationId), uint(req.ReturnLocationId), req.Notes, req.PaymentMethod); err != nil {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "创建订单失败: " + err.Error(),
		}, err
	}

	// 验证订单是否真的创建成功
	if orderModel.ID == 0 {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "订单创建失败，未获取到订单ID",
		}, fmt.Errorf("order ID is 0 after creation")
	}

	// 4. 更新预订的order_id
	if err := tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", reservation.ID).Update("order_id", orderModel.ID).Error; err != nil {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "更新预订失败: " + err.Error(),
		}, err
	}

	// 5. 创建支付链接
	alipayService := payment.NewAlipayService()
	paymentURL, err := alipayService.CreatePaymentURL(orderModel.OrderSn, orderModel.TotalAmount, fmt.Sprintf("租车订单-%s", vehicle.Style))
	if err != nil {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "创建支付链接失败: " + err.Error(),
		}, err
	}

	// 6. 更新订单的支付链接
	if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("payment_url", paymentURL).Error; err != nil {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "更新支付信息失败: " + err.Error(),
		}, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "提交事务失败: " + err.Error(),
		}, err
	}

	return &order.CreateOrderFromReservationResponse{
		Code:        200,
		Message:     "订单创建成功",
		OrderId:     int64(orderModel.ID),
		OrderSn:     orderModel.OrderSn,
		TotalAmount: orderModel.TotalAmount,
		PaymentUrl:  paymentURL,
	}, nil
}

// GetOrder 获取订单详情
func GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	var orderModel model_mysql.Orders
	var err error

	// 根据订单ID或订单号查询
	if req.OrderId > 0 {
		err = orderModel.GetByID(uint(req.OrderId))
	} else if req.OrderSn != "" {
		err = orderModel.GetByOrderSn(req.OrderSn)
	} else {
		return &order.GetOrderResponse{
			Code:    400,
			Message: "订单ID或订单号不能为空",
		}, nil
	}

	if err != nil {
		return &order.GetOrderResponse{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	// 转换为响应格式
	orderInfo := &order.OrderInfo{
		Id:               int64(orderModel.ID),
		UserId:           int64(orderModel.UserId),
		VehicleId:        int64(orderModel.VehicleId),
		ReservationId:    int64(orderModel.ReservationId),
		OrderSn:          orderModel.OrderSn,
		PickupLocationId: int64(orderModel.PickupLocationId),
		ReturnLocationId: int64(orderModel.ReturnLocationId),
		PickupTime:       orderModel.PickupTime.Format(time.RFC3339),
		ReturnTime:       orderModel.ReturnTime.Format(time.RFC3339),
		RentalDays:       orderModel.RentalDays,
		DailyRate:        orderModel.DailyRate,
		TotalAmount:      orderModel.TotalAmount,
		Status:           orderModel.Status,
		Payment:          orderModel.Payment,
		PaymentStatus:    orderModel.PaymentStatus,
		PaymentUrl:       orderModel.PaymentUrl,
		AlipayTradeNo:    orderModel.AlipayTradeNo,
		Notes:            orderModel.Notes,
		CreatedAt:        orderModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        orderModel.UpdatedAt.Format(time.RFC3339),
	}

	return &order.GetOrderResponse{
		Code:    200,
		Message: "获取成功",
		Order:   orderInfo,
	}, nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	var orderModel model_mysql.Orders
	var err error

	// 根据订单ID或订单号查询
	if req.OrderId > 0 {
		err = orderModel.GetByID(uint(req.OrderId))
	} else if req.OrderSn != "" {
		err = orderModel.GetByOrderSn(req.OrderSn)
	} else {
		return &order.UpdateOrderStatusResponse{
			Code:    400,
			Message: "订单ID或订单号不能为空",
		}, nil
	}

	if err != nil {
		return &order.UpdateOrderStatusResponse{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	// 更新订单状态
	if err := orderModel.UpdateStatus(orderModel.ID, req.Status); err != nil {
		return &order.UpdateOrderStatusResponse{
			Code:    500,
			Message: "更新订单状态失败",
		}, err
	}

	return &order.UpdateOrderStatusResponse{
		Code:    200,
		Message: "更新成功",
	}, nil
}

// AlipayNotify 支付宝异步通知处理
func AlipayNotify(ctx context.Context, req *order.AlipayNotifyRequest) (*order.AlipayNotifyResponse, error) {
	fmt.Printf("=== 订单微服务收到支付宝通知 ===\n")
	fmt.Printf("订单号: %s\n", req.OutTradeNo)
	fmt.Printf("交易号: %s\n", req.TradeNo)
	fmt.Printf("交易状态: %s\n", req.TradeStatus)
	fmt.Printf("金额: %s\n", req.TotalAmount)

	// 查找订单
	var orderModel model_mysql.Orders
	if err := orderModel.GetByOrderSn(req.OutTradeNo); err != nil {
		fmt.Printf("订单查找失败: %v\n", err)
		return &order.AlipayNotifyResponse{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	fmt.Printf("找到订单: ID=%d, 当前状态=%d, 支付状态=%d\n",
		orderModel.ID, orderModel.Status, orderModel.PaymentStatus)

	// 处理支付成功
	if req.TradeStatus == "TRADE_SUCCESS" {
		fmt.Printf("处理支付成功状态...\n")
		// 开启事务
		tx := global.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// 更新支付状态
		if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("payment_status", model_mysql.PaymentStatusPaid).Error; err != nil {
			tx.Rollback()
			return &order.AlipayNotifyResponse{
				Code:    500,
				Message: "更新支付状态失败",
			}, err
		}

		// 更新订单状态
		if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("status", model_mysql.OrderStatusPaid).Error; err != nil {
			tx.Rollback()
			return &order.AlipayNotifyResponse{
				Code:    500,
				Message: "更新订单状态失败",
			}, err
		}

		// 更新支付宝交易号
		if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("alipay_trade_no", req.TradeNo).Error; err != nil {
			tx.Rollback()
			return &order.AlipayNotifyResponse{
				Code:    500,
				Message: "更新交易号失败",
			}, err
		}

		// 更新预订状态为租用中
		if err := tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", orderModel.ReservationId).Update("status", model_mysql.InventoryStatusRented).Error; err != nil {
			tx.Rollback()
			return &order.AlipayNotifyResponse{
				Code:    500,
				Message: "更新预订状态失败",
			}, err
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			fmt.Printf("提交事务失败: %v\n", err)
			return &order.AlipayNotifyResponse{
				Code:    500,
				Message: "提交事务失败",
			}, err
		}
		fmt.Printf("支付状态更新成功，事务已提交\n")
	} else {
		fmt.Printf("交易状态不是TRADE_SUCCESS，跳过处理: %s\n", req.TradeStatus)
	}

	return &order.AlipayNotifyResponse{
		Code:    200,
		Message: "处理成功",
	}, nil
}
