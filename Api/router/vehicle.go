package router

import (
	"Api/trigger"
	jwt "Common/pkg"

	"github.com/gin-gonic/gin"
)

// RegisterVehicleRoutes 注册车辆相关的路由
func RegisterVehicleRoutes(r *gin.Engine) {
	vehicleGroup := r.Group("/vehicle")
	{
		// 需要认证的路由
		vehicleGroup.Use(jwt.JWTAuth("merchant"))                  // 应用JWT认证中间件
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
}
