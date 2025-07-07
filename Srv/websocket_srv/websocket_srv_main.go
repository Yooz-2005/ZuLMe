package main

import (
	"Common/appconfig"
	"Common/initialize"
	"Common/pkg"
	"ZuLMe/Srv/websocket_srv/internal/manager"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	appconfig.GetViperConfData()
	initialize.NewNacos()

	// 初始化数据库连接
	_, err := initialize.MysqlInit()
	if err != nil {
		log.Fatalf("MySQL初始化失败: %v", err)
	}

	// 初始化MongoDB
	initialize.InitMongoDB()

	// 初始化Redis
	initialize.RedisInit()

	// 创建WebSocket管理器
	wsManager := manager.NewWebSocketManager()

	// 创建Gin路由
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// WebSocket路由（需要JWT认证）
	wsGroup := r.Group("/ws")
	{
		// WebSocket连接需要JWT认证
		wsGroup.Use(pkg.JWTAuth("2209"))
		wsGroup.GET("/chat", wsManager.HandleWebSocket)
	}

	// API路由
	apiGroup := r.Group("/api/ws")
	{
		// 获取在线用户数量（公开接口）
		apiGroup.GET("/online-count", func(c *gin.Context) {
			count := wsManager.GetOnlineCount()
			c.JSON(200, gin.H{
				"code":    200,
				"message": "success",
				"data": gin.H{
					"online_count": count,
				},
			})
		})

		// 需要认证的API
		authAPI := apiGroup.Group("")
		authAPI.Use(pkg.JWTAuth("2209"))
		{
			// 获取在线用户列表
			authAPI.GET("/online-users", func(c *gin.Context) {
				users := wsManager.GetOnlineUsers()
				c.JSON(200, gin.H{
					"code":    200,
					"message": "success",
					"data": gin.H{
						"online_users": users,
					},
				})
			})

			// 向指定用户发送消息
			authAPI.POST("/send-message", func(c *gin.Context) {
				var req struct {
					UserID  uint        `json:"user_id" binding:"required"`
					Message interface{} `json:"message" binding:"required"`
				}

				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(400, gin.H{
						"code":    400,
						"message": "参数错误",
						"error":   err.Error(),
					})
					return
				}

				err := wsManager.SendMessageToUser(req.UserID, req.Message)
				if err != nil {
					c.JSON(500, gin.H{
						"code":    500,
						"message": "发送失败",
						"error":   err.Error(),
					})
					return
				}

				c.JSON(200, gin.H{
					"code":    200,
					"message": "发送成功",
				})
			})

			// 广播消息
			authAPI.POST("/broadcast", func(c *gin.Context) {
				var req struct {
					Message interface{} `json:"message" binding:"required"`
				}

				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(400, gin.H{
						"code":    400,
						"message": "参数错误",
						"error":   err.Error(),
					})
					return
				}

				wsManager.BroadcastMessage(req.Message)

				c.JSON(200, gin.H{
					"code":    200,
					"message": "广播成功",
				})
			})
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "websocket-service",
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	// 启动服务
	port := ":8009"
	fmt.Printf("WebSocket Server started on %s\n", port)
	log.Printf("WebSocket连接地址: ws://localhost%s/ws/chat", port)
	log.Printf("API接口地址: http://localhost%s/api/ws", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("WebSocket服务启动失败: %v", err)
	}
}
