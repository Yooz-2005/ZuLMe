package router

import (
	"ZuLMe/ZuLMe/Api/trigger"
	"ZuLMe/ZuLMe/Common/pkg"
	"github.com/gin-gonic/gin"
)

// RegisterVehicleRoutes 注册车辆相关的路由
func RegisterVehicleRoutes(r *gin.Engine) {
	vehicleGroup := r.Group("/vehicle")
	{
		// 需要认证的路由
		vehicleGroup.Use(pkg.JWTAuth("merchant"))                  // 应用JWT认证中间件
		vehicleGroup.POST("/create", trigger.CreateVehicleHandler) // 创建车辆
		vehicleGroup.PUT("/update", trigger.UpdateVehicleHandler)  // 更新车辆
		vehicleGroup.POST("/delete", trigger.DeleteVehicleHandler) // 删除车辆
	}

	// 公开路由（不需要认证）
	publicVehicleGroup := r.Group("/vehicle")
	{
		publicVehicleGroup.GET("/:id", trigger.GetVehicleHandler)    // 获取车辆详情
		publicVehicleGroup.GET("/list", trigger.ListVehiclesHandler) // 获取车辆列表
	}

	// 车辆类型管理路由（不需要身份验证）
	vehicleTypeGroup := r.Group("/vehicle-type")
	{
		vehicleTypeGroup.POST("/create", trigger.CreateVehicleTypeHandler) // 创建车辆类型
		vehicleTypeGroup.PUT("/update", trigger.UpdateVehicleTypeHandler)  // 更新车辆类型
		vehicleTypeGroup.POST("/delete", trigger.DeleteVehicleTypeHandler) // 删除车辆类型
		vehicleTypeGroup.GET("/:id", trigger.GetVehicleTypeHandler)        // 获取车辆类型详情
		vehicleTypeGroup.GET("/list", trigger.ListVehicleTypesHandler)     // 获取车辆类型列表
	}

	// 车辆品牌管理路由
	vehicleBrandGroup := r.Group("/vehicle-brand")
	{
		// 管理员路由（需要认证）
		vehicleBrandGroup.Use(pkg.JWTAuth("admin"))                          // 应用JWT认证中间件
		vehicleBrandGroup.POST("/create", trigger.CreateVehicleBrandHandler) // 创建车辆品牌
		vehicleBrandGroup.PUT("/update", trigger.UpdateVehicleBrandHandler)  // 更新车辆品牌
		vehicleBrandGroup.POST("/delete", trigger.DeleteVehicleBrandHandler) // 删除车辆品牌
	}

	// 车辆品牌公开路由（不需要认证）
	publicBrandGroup := r.Group("/vehicle-brand")
	{
		publicBrandGroup.GET("/:id", trigger.GetVehicleBrandHandler)    // 获取车辆品牌详情
		publicBrandGroup.GET("/list", trigger.ListVehicleBrandsHandler) // 获取车辆品牌列表
	}

	// 车辆库存公开路由（不需要认证）
	publicInventoryGroup := r.Group("/vehicle-inventory")
	{
		publicInventoryGroup.POST("/check-availability", trigger.CheckAvailabilityHandler)    // 检查车辆可用性
		publicInventoryGroup.POST("/available-vehicles", trigger.GetAvailableVehiclesHandler) // 获取可用车辆
		publicInventoryGroup.GET("/calendar", trigger.GetInventoryCalendarHandler)            // 获取库存日历（用户选择日期）
	}

	// 用户库存操作路由（需要用户认证）
	userInventoryGroup := r.Group("/vehicle-inventory")
	{
		userInventoryGroup.Use(pkg.JWTAuth("2209"))                                      // 用户认证
		userInventoryGroup.POST("/reservation/create", trigger.CreateReservationHandler) // 用户创建预订
	}

	// 订单相关路由已移动到 router/order.go 文件中

	// 系统内部调用路由（无需认证，供其他微服务调用）
	systemInventoryGroup := r.Group("/vehicle-inventory")
	{
		systemInventoryGroup.PUT("/reservation/status", trigger.UpdateReservationStatusHandler) // 系统更新预订状态
	}

	// 商家库存管理路由（需要商家认证）
	merchantInventoryGroup := r.Group("/vehicle-inventory")
	{
		merchantInventoryGroup.Use(pkg.JWTAuth("merchant"))                                        // 商家认证
		merchantInventoryGroup.GET("/stats", trigger.GetInventoryStatsHandler)                     // 获取库存统计
		merchantInventoryGroup.POST("/maintenance/set", trigger.SetMaintenanceHandler)             // 设置维护状态
		merchantInventoryGroup.GET("/maintenance/schedule", trigger.GetMaintenanceScheduleHandler) // 获取维护计划
		merchantInventoryGroup.GET("/report", trigger.GetInventoryReportHandler)                   // 获取库存报表
	}
}
