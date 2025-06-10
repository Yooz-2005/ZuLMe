package main

import (
	"Common/appconfig"
	"Common/global"
	"Common/initialize"
	"google.golang.org/grpc"
	"merchant_srv/grpc_merchant"
	"merchant_srv/internal"
	"models/model_mysql"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	// 自动迁移 Merchant 模型到数据库
	global.DB.AutoMigrate(&model_mysql.Merchant{})
	// 注册商家服务
	grpc_merchant.RegisterMerchant(func(server *grpc.Server) {
		internal.RegisterMerchantServer(server)
	})

}
