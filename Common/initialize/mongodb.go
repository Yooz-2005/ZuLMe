package initialize

import (
	"Common/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitMongoDB() {
	var err error
	// 设置 MongoDB 连接 URI
	//需要动态配置
	data := Nacos.Mongodb
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", data.Username, data.Password, data.Host, data.Port)
	// 设置上下文超时，20 秒后若操作未完成则会取消操作
	// 这样可以避免程序因长时间等待连接而阻塞
	global.Ctx, global.Cancel = context.WithTimeout(context.Background(), 20*time.Second)
	// 确保在函数结束时取消上下文，防止资源泄漏
	defer global.Cancel()

	// 尝试连接到 MongoDB 服务器
	// 使用 options.Client().ApplyURI 方法设置连接的 URI
	// 这里的 URI 是示例，你需要替换为实际的 MongoDB 连接信息
	global.MongoClient, err = mongo.Connect(global.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// 若连接失败，输出错误信息并终止程序
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}

	// 检查是否成功连接到 MongoDB
	// 向 MongoDB 服务器发送一个 ping 请求，若成功则表示连接正常
	err = global.MongoClient.Ping(global.Ctx, nil)
	if err != nil {
		// 若 ping 失败，输出错误信息
		fmt.Println("Failed to ping MongoDB:", err)
	} else {
		// 若 ping 成功，输出连接成功信息
		fmt.Println("Successfully connected to MongoDB")
		fmt.Println("MongoDB 链接成功")
	}

	// 当程序结束或不再需要连接时，断开与 MongoDB 的连接
	//err = global.Client.Disconnect(global.Ctx)
	//if err != nil {
	//	// 若断开连接失败，输出错误信息
	//	fmt.Println("Failed to disconnect from MongoDB:", err)
	//}
}
