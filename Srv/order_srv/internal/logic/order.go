package logic

import (
	"Common/global"
	"Common/payment"
	"context"
	"fmt"
	"models/model_mysql"
	order "order_srv/proto_order"
	"strconv"
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

	// 3. 确定取车地点ID（如果前端没有传递，则使用车辆所在的商家ID）
	pickupLocationID := uint(req.PickupLocationId)
	if pickupLocationID == 0 {
		pickupLocationID = uint(vehicle.MerchantID) // 使用车辆所在的商家ID作为取车地点
	}

	// 4. 开启事务创建订单
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	var orderModel model_mysql.Orders
	if err := orderModel.CreateOrderFromReservationWithTx(tx, &reservation, &vehicle, pickupLocationID, uint(req.ReturnLocationId), req.Notes, req.PaymentMethod); err != nil {
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

	// 5. 更新预订的order_id
	if err := tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", reservation.ID).Update("order_id", orderModel.ID).Error; err != nil {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "更新预订失败: " + err.Error(),
		}, err
	}

	// 6. 创建支付链接
	pay := payment.NewAliPay()
	amount := strconv.FormatFloat(orderModel.TotalAmount, 'f', -1, 64)
	paymentURL := pay.Pay(vehicle.Brand, orderModel.OrderSn, amount)

	if paymentURL == "" {
		tx.Rollback()
		return &order.CreateOrderFromReservationResponse{
			Code:    500,
			Message: "创建支付链接失败",
		}, nil
	}

	// 7. 更新订单的支付链接
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

// GetUserOrderList 获取用户订单列表
func GetUserOrderList(ctx context.Context, req *order.GetUserOrderListRequest) (*order.GetUserOrderListResponse, error) {
	// 参数验证
	if req.UserId <= 0 {
		return &order.GetUserOrderListResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	// 设置默认分页参数
	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 调用模型方法获取订单列表
	orderModel := &model_mysql.Orders{}
	orders, total, err := orderModel.GetUserOrderList(uint(req.UserId), page, pageSize, req.Status, req.PaymentStatus)
	if err != nil {
		return &order.GetUserOrderListResponse{
			Code:    500,
			Message: "获取订单列表失败",
		}, err
	}

	// 转换为响应格式，只返回必要字段，并获取车辆信息
	var orderInfos []*order.OrderInfo
	for _, o := range orders {
		// 获取车辆信息
		var vehicleModel model_mysql.Vehicle
		var vehicleBrand, vehicleStyle, vehicleImages string
		if err := vehicleModel.GetByID(o.VehicleId); err == nil {
			vehicleBrand = vehicleModel.Brand
			vehicleStyle = vehicleModel.Style
			vehicleImages = vehicleModel.Images
			fmt.Printf("=== 订单 %s 的车辆信息 ===\n", o.OrderSn)
			fmt.Printf("车辆ID: %d\n", o.VehicleId)
			fmt.Printf("品牌: %s\n", vehicleBrand)
			fmt.Printf("型号: %s\n", vehicleStyle)
			fmt.Printf("图片: %s\n", vehicleImages)
		} else {
			fmt.Printf("获取车辆信息失败，车辆ID: %d, 错误: %v\n", o.VehicleId, err)
		}

		// 获取取车和还车地点信息
		var pickupLocation string
		var returnLocation string

		if o.PickupLocationId > 0 {
			var merchantModel model_mysql.Merchant
			if err := merchantModel.GetByID(o.PickupLocationId); err == nil {
				pickupLocation = merchantModel.Name + " - " + merchantModel.Location
			}
		}

		if o.ReturnLocationId > 0 {
			var merchantModel model_mysql.Merchant
			if err := merchantModel.GetByID(o.ReturnLocationId); err == nil {
				returnLocation = merchantModel.Name + " - " + merchantModel.Location
			}
		}

		orderInfo := &order.OrderInfo{
			Id:            int64(o.ID),
			OrderSn:       o.OrderSn,
			VehicleId:     int64(o.VehicleId),
			ReservationId: int64(o.ReservationId),
			PickupTime:    o.PickupTime.Format("2006-01-02"),
			ReturnTime:    o.ReturnTime.Format("2006-01-02"),
			TotalAmount:   o.TotalAmount,
			Status:        o.Status,
			PaymentStatus: o.PaymentStatus,
			PaymentUrl:    o.PaymentUrl,
			CreatedAt:     o.CreatedAt.Format(time.RFC3339),
			RentalDays:    o.RentalDays,
			Notes:         o.Notes,
		}

		// 添加额外信息到Notes字段中，方便前端显示
		extraInfo := ""
		if vehicleBrand != "" {
			extraInfo += "车辆: " + vehicleBrand + " " + vehicleStyle + "; "
		}
		if vehicleImages != "" {
			extraInfo += "图片: " + vehicleImages + "; "
		}
		if pickupLocation != "" {
			extraInfo += "取车: " + pickupLocation + "; "
		}
		if returnLocation != "" {
			extraInfo += "还车: " + returnLocation + "; "
		}
		if extraInfo != "" {
			if orderInfo.Notes != "" {
				orderInfo.Notes = extraInfo + "备注: " + orderInfo.Notes
			} else {
				orderInfo.Notes = extraInfo
			}
		}
		orderInfos = append(orderInfos, orderInfo)
	}

	return &order.GetUserOrderListResponse{
		Code:    200,
		Message: "获取成功",
		Orders:  orderInfos,
		Total:   total,
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

	fmt.Printf("=== 开始更新订单状态 ===\n")
	fmt.Printf("订单ID: %d, 订单号: %s, 预订ID: %d\n", orderModel.ID, orderModel.OrderSn, orderModel.ReservationId)
	fmt.Printf("当前订单状态: %d, 新状态: %d\n", orderModel.Status, req.Status)

	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态
	if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("status", req.Status).Error; err != nil {
		tx.Rollback()
		return &order.UpdateOrderStatusResponse{
			Code:    500,
			Message: "更新订单状态失败",
		}, err
	}

	// 如果是支付成功，同时更新支付状态和预订状态
	if req.Status == 2 { // 已支付状态
		// 更新支付状态
		if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Update("payment_status", model_mysql.PaymentStatusPaid).Error; err != nil {
			tx.Rollback()
			return &order.UpdateOrderStatusResponse{
				Code:    500,
				Message: "更新支付状态失败",
			}, err
		}

		// 更新预订状态为租用中
		if err := tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", orderModel.ReservationId).Update("status", model_mysql.InventoryStatusRented).Error; err != nil {
			tx.Rollback()
			return &order.UpdateOrderStatusResponse{
				Code:    500,
				Message: "更新预订状态失败",
			}, err
		}
		fmt.Printf("已更新预订状态为租用中\n")
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return &order.UpdateOrderStatusResponse{
			Code:    500,
			Message: "提交事务失败",
		}, err
	}

	fmt.Printf("订单状态更新成功\n")
	return &order.UpdateOrderStatusResponse{
		Code:    200,
		Message: "更新成功",
	}, nil
}

// CancelOrder 取消订单
func CancelOrder(ctx context.Context, req *order.CancelOrderRequest) (*order.CancelOrderResponse, error) {
	var orderModel model_mysql.Orders
	var err error

	// 根据订单ID或订单号查询
	if req.OrderId > 0 {
		err = orderModel.GetByID(uint(req.OrderId))
	} else if req.OrderSn != "" {
		err = orderModel.GetByOrderSn(req.OrderSn)
	} else {
		return &order.CancelOrderResponse{
			Code:    400,
			Message: "订单ID或订单号不能为空",
		}, nil
	}

	if err != nil {
		return &order.CancelOrderResponse{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	// 验证订单所有权
	if req.UserId > 0 && orderModel.UserId != uint(req.UserId) {
		return &order.CancelOrderResponse{
			Code:    403,
			Message: "无权限操作此订单",
		}, nil
	}

	// 检查订单状态是否可以取消
	if orderModel.Status != model_mysql.OrderStatusPending && orderModel.PaymentStatus != model_mysql.PaymentStatusPending {
		return &order.CancelOrderResponse{
			Code:    400,
			Message: "订单状态不允许取消",
		}, nil
	}

	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态为已取消
	cancelReason := " [订单已取消]"
	if req.Reason != "" {
		cancelReason = fmt.Sprintf(" [取消原因: %s]", req.Reason)
	}
	if err := tx.Model(&model_mysql.Orders{}).Where("id = ?", orderModel.ID).Updates(map[string]interface{}{
		"status":         model_mysql.OrderStatusCancelled,
		"payment_status": model_mysql.PaymentStatusCancelled,
		"notes":          orderModel.Notes + cancelReason,
	}).Error; err != nil {
		tx.Rollback()
		return &order.CancelOrderResponse{
			Code:    500,
			Message: "取消订单失败",
		}, err
	}

	// 释放预订资源 - 将预订状态改为可用
	if orderModel.ReservationId > 0 {
		if err := tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", orderModel.ReservationId).Update("status", model_mysql.InventoryStatusAvailable).Error; err != nil {
			tx.Rollback()
			return &order.CancelOrderResponse{
				Code:    500,
				Message: "释放预订资源失败",
			}, err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return &order.CancelOrderResponse{
			Code:    500,
			Message: "提交事务失败",
		}, err
	}

	return &order.CancelOrderResponse{
		Code:    200,
		Message: "订单取消成功",
	}, nil
}
