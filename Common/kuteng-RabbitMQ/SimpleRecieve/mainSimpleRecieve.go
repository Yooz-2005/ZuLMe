package SimpleRecieve

import (
	"Common/global"
	"Common/kuteng-RabbitMQ/RabbitMQ"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	gorm.Model
	UserId           uint      `json:"user_id"`
	VehicleId        uint      `json:"vehicle_id"`
	ReservationId    uint      `json:"reservation_id"`
	OrderSn          string    `json:"order_sn"`
	PickupLocationId uint      `json:"pickup_location_id"`
	ReturnLocationId uint      `json:"return_location_id"`
	PickupTime       time.Time `json:"pickup_time"`
	ReturnTime       time.Time `json:"return_time"`
	RentalDays       int32     `json:"rental_days"`
	DailyRate        float64   `json:"daily_rate"`
	TotalAmount      float64   `json:"total_amount"`
	Status           int32     `json:"status"`
	Payment          int32     `json:"payment"`
	PaymentStatus    int32     `json:"payment_status"`
	PaymentUrl       string    `json:"payment_url"`
	AlipayTradeNo    string    `json:"alipay_trade_no"`
	Notes            string    `json:"notes"`
}

func SimpleReceive() error {
	fmt.Println("开始监听订单消息...")
	mq := RabbitMQ.NewRabbitMQSimple("Order")
	mq.ConsumeSimple(func(b []byte) {
		var orderMap map[string]interface{}
		err := json.Unmarshal(b, &orderMap)
		if err != nil {
			fmt.Println("消息解析失败:", err)
			return
		}
		// 加锁
		lockKey := "order_lock:" + fmt.Sprintf("%v", orderMap["order_sn"])
		lockValue := "locked"
		lockExpireTime := time.Second * 15
		result, err := global.Rdb.SetNX(context.Background(), lockKey, lockValue, lockExpireTime).Result()
		if err != nil || !result {
			fmt.Println("订单正在处理或加锁失败")
			return
		}
		defer global.Rdb.Del(context.Background(), lockKey)

		// 转 struct
		orderJson, _ := json.Marshal(orderMap)
		var order Orders
		if err := json.Unmarshal(orderJson, &order); err != nil {
			fmt.Println("map转struct失败:", err)
			return
		}

		// 入库
		if err := global.DB.Create(&order).Error; err != nil {
			zap.L().Error("订单入库失败", zap.Error(err))
			return
		}
		zap.L().Info("订单入库成功", zap.String("order_sn", order.OrderSn))
	})
	return nil
}
