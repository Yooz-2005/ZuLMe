package main

import (
	"Common/appconfig"
	"Common/global"
	"Common/initialize"
	"coupon_srv/grpc_coupon"
	"coupon_srv/internal"
	"google.golang.org/grpc"
	"models/model_mysql"
)

func main() {
	appconfig.GetViperConfData()
	initialize.NewNacos()
	initialize.MysqlInit()
	initialize.RedisInit()
	global.DB.AutoMigrate(&model_mysql.PromotionConfig{})
	grpc_coupon.RegisterCouponGrpc(func(server *grpc.Server) {
		internal.RegisterCouponServer(server)
	})
}
