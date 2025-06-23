package main

import (
	"Common/appconfig"
	"Common/global"
	"Common/initialize"
	"log"
	"models/model_mysql"
	"net"
	order "order_srv/proto_order"
	"order_srv/server"
	"time"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	initialize.InitES()

	// 启动订单超时检查定时任务
	startOrderTimeoutScheduler()

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

	log.Printf("Order gRPC Server started on %s (with timeout scheduler)\n", port)

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("订单服务启动失败: %v", err)
	}
}

// startOrderTimeoutScheduler 启动订单超时检查定时任务
func startOrderTimeoutScheduler() {
	log.Println("🚀 启动订单超时检查定时任务 (5分钟超时) - 使用Cron")

	// 创建cron调度器，支持秒级调度
	c := cron.New(cron.WithSeconds())

	// Cron表达式说明：
	// "0 */2 * * * *" = 每2分钟执行一次 (秒 分 时 日 月 周)
	// "0 */1 * * * *" = 每1分钟执行一次
	// "0 */5 * * * *" = 每5分钟执行一次
	// "0 0 */1 * * *" = 每1小时执行一次

	cronExpr := "0 */2 * * * *" // 每2分钟执行一次

	// 添加定时任务
	entryID, err := c.AddFunc(cronExpr, func() {
		checkAndCancelExpiredOrders()
	})

	if err != nil {
		log.Fatalf("❌ 添加cron任务失败: %v", err)
	}

	// 启动调度器
	c.Start()
	log.Printf("✅ Cron调度器已启动")
	log.Printf("   📅 Cron表达式: %s", cronExpr)
	log.Printf("   ⏰ 执行频率: 每2分钟检查一次超时订单")
	log.Printf("   🆔 任务ID: %d", entryID)

	// 立即执行一次检查
	go checkAndCancelExpiredOrders()
}

// checkAndCancelExpiredOrders 检查并取消超时订单
func checkAndCancelExpiredOrders() {
	log.Printf("🔍 [%s] 检查5分钟超时订单...", time.Now().Format("15:04:05"))

	// 查询5分钟前创建的未支付订单
	var expiredOrders []model_mysql.Orders
	err := global.DB.Where("payment_status = ? AND status = ? AND created_at < ?",
		model_mysql.PaymentStatusPending,
		model_mysql.OrderStatusPending,
		time.Now().Add(-5*time.Minute)).Find(&expiredOrders).Error

	if err != nil {
		log.Printf("❌ 查询超时订单失败: %v", err)
		return
	}

	if len(expiredOrders) == 0 {
		log.Println("✅ 无超时订单")
		return
	}

	log.Printf("⚠️ 发现 %d 个超时订单，开始取消...", len(expiredOrders))

	successCount := 0
	for _, expiredOrder := range expiredOrders {
		if cancelExpiredOrder(&expiredOrder) {
			successCount++
			log.Printf("✅ 已取消: %s", expiredOrder.OrderSn)
		}
	}

	log.Printf("📊 完成: %d/%d", successCount, len(expiredOrders))
}

// cancelExpiredOrder 取消超时订单
func cancelExpiredOrder(expiredOrder *model_mysql.Orders) bool {
	tx := global.DB.Begin()

	// 更新订单状态
	err := tx.Model(&model_mysql.Orders{}).Where("id = ?", expiredOrder.ID).Updates(map[string]interface{}{
		"status":         model_mysql.OrderStatusCancelled,
		"payment_status": model_mysql.PaymentStatusCancelled,
		"notes":          expiredOrder.Notes + " [系统自动取消：超过5分钟未支付]",
	}).Error

	if err != nil {
		tx.Rollback()
		log.Printf("❌ 更新订单失败: %v", err)
		return false
	}

	// 释放预订资源
	if expiredOrder.ReservationId > 0 {
		err = tx.Model(&model_mysql.VehicleInventory{}).Where("id = ?", expiredOrder.ReservationId).
			Update("status", model_mysql.InventoryStatusAvailable).Error
		if err != nil {
			tx.Rollback()
			log.Printf("❌ 释放预订失败: %v", err)
			return false
		}
	}

	tx.Commit()
	return true
}
