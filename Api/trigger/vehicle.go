package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"fmt"
	"strconv"
	vehicle "vehicle_srv/proto_vehicle"

	"github.com/gin-gonic/gin"
)

// CreateVehicleHandler 创建车辆处理器
func CreateVehicleHandler(c *gin.Context) {
	mid := c.GetUint("userId")
	fmt.Println("1111111111111111111111111")
	fmt.Println(mid)
	var req request.CreateVehicleRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	createRes, err := handler.CreateVehicle(c, &vehicle.CreateVehicleRequest{
		MerchantId:  int64(mid),
		TypeId:      req.TypeID,
		Brand:       req.Brand,
		Style:       req.Style,
		Year:        req.Year,
		Color:       req.Color,
		Mileage:     req.Mileage,
		Price:       req.Price,
		Status:      req.Status,
		Description: req.Description,
		Images:      req.Images,
		Contact:     req.Contact,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if createRes.Code != 200 {
		response.ResponseError400(c, createRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": createRes.Message,
		"vehicle": createRes.Vehicle,
	})
}

// UpdateVehicleHandler 更新车辆处理器
func UpdateVehicleHandler(c *gin.Context) {
	mid := c.GetUint64("userId")
	var req request.UpdateVehicleRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	updateRes, err := handler.UpdateVehicle(c, &vehicle.UpdateVehicleRequest{
		Id:          req.ID,
		MerchantId:  int64(mid),
		TypeId:      req.TypeID,
		Brand:       req.Brand,
		Style:       req.Style,
		Year:        req.Year,
		Color:       req.Color,
		Mileage:     req.Mileage,
		Price:       req.Price,
		Status:      req.Status,
		Description: req.Description,
		Images:      req.Images,
		Location:    req.Location,
		Contact:     req.Contact,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if updateRes.Code != 200 {
		response.ResponseError400(c, updateRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": updateRes.Message,
		"vehicle": updateRes.Vehicle,
	})
}

// DeleteVehicleHandler 删除车辆处理器
func DeleteVehicleHandler(c *gin.Context) {
	var req request.DeleteVehicleRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	deleteRes, err := handler.DeleteVehicle(c, &vehicle.DeleteVehicleRequest{
		Id: req.ID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if deleteRes.Code != 200 {
		response.ResponseError400(c, deleteRes.Message)
		return
	}

	response.ResponseSuccess(c, deleteRes.Message)
}

// GetVehicleHandler 获取车辆详情处理器
func GetVehicleHandler(c *gin.Context) {
	var req request.GetVehicleRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	getRes, err := handler.GetVehicle(c, &vehicle.GetVehicleRequest{
		Id: req.ID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if getRes.Code != 200 {
		response.ResponseError400(c, getRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": getRes.Message,
		"vehicle": getRes.Vehicle,
	})
}

// ListVehiclesHandler 获取车辆列表处理器
func ListVehiclesHandler(c *gin.Context) {
	var req request.ListVehiclesRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 处理status参数：如果没有明确指定status，设置为-1表示不筛选
	status := req.Status
	if c.Query("status") == "" && c.PostForm("status") == "" {
		status = -1 // 没有传递status参数时，设置为-1表示不筛选
	}

	listRes, err := handler.ListVehicles(c, &vehicle.ListVehiclesRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Keyword:    req.Keyword,
		MerchantId: req.MerchantID,
		TypeId:     req.TypeID,
		Status:     status,
		PriceMin:   req.PriceMin,
		PriceMax:   req.PriceMax,
		YearMin:    req.YearMin,
		YearMax:    req.YearMax,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if listRes.Code != 200 {
		response.ResponseError400(c, listRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":  listRes.Message,
		"vehicles": listRes.Vehicles,
		"total":    listRes.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// ==================== 车辆类型处理器 ====================

// CreateVehicleTypeHandler 创建车辆类型处理器
func CreateVehicleTypeHandler(c *gin.Context) {
	var req request.CreateVehicleTypeRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	createRes, err := handler.CreateVehicleType(c, &vehicle.CreateVehicleTypeRequest{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Sort:        req.Sort,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if createRes.Code != 200 {
		response.ResponseError400(c, createRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      createRes.Message,
		"vehicle_type": createRes.VehicleType,
	})
}

// UpdateVehicleTypeHandler 更新车辆类型处理器
func UpdateVehicleTypeHandler(c *gin.Context) {
	var req request.UpdateVehicleTypeRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	updateRes, err := handler.UpdateVehicleType(c, &vehicle.UpdateVehicleTypeRequest{
		Id:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Sort:        req.Sort,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if updateRes.Code != 200 {
		response.ResponseError400(c, updateRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      updateRes.Message,
		"vehicle_type": updateRes.VehicleType,
	})
}

// DeleteVehicleTypeHandler 删除车辆类型处理器
func DeleteVehicleTypeHandler(c *gin.Context) {
	var req request.DeleteVehicleTypeRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	deleteRes, err := handler.DeleteVehicleType(c, &vehicle.DeleteVehicleTypeRequest{
		Id: req.ID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if deleteRes.Code != 200 {
		response.ResponseError400(c, deleteRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": deleteRes.Message,
	})
}

// GetVehicleTypeHandler 获取车辆类型详情处理器
func GetVehicleTypeHandler(c *gin.Context) {
	// 从URL路径参数中获取ID
	idStr := c.Param("id")
	if idStr == "" {
		response.ResponseError400(c, "车辆类型ID不能为空")
		return
	}

	// 转换ID为int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "车辆类型ID格式错误")
		return
	}

	getRes, err := handler.GetVehicleType(c, &vehicle.GetVehicleTypeRequest{
		Id: id,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if getRes.Code != 200 {
		response.ResponseError400(c, getRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      getRes.Message,
		"vehicle_type": getRes.VehicleType,
	})
}

// ListVehicleTypesHandler 获取车辆类型列表处理器
func ListVehicleTypesHandler(c *gin.Context) {
	var req request.ListVehicleTypesRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 处理status参数：如果没有明确指定status，设置为-1表示不筛选
	status := req.Status
	if c.Query("status") == "" && c.PostForm("status") == "" {
		status = -1 // 没有传递status参数时，设置为-1表示不筛选
	}

	listRes, err := handler.ListVehicleTypes(c, &vehicle.ListVehicleTypesRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Status:   status,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if listRes.Code != 200 {
		response.ResponseError400(c, listRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":       listRes.Message,
		"vehicle_types": listRes.VehicleTypes,
		"total":         listRes.Total,
		"page":          req.Page,
		"pageSize":      req.PageSize,
	})
}
