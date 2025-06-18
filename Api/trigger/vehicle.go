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
		BrandId:     req.BrandID,
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
		BrandId:     req.BrandID,
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
	// 从URL路径参数中获取ID
	idStr := c.Param("id")
	if idStr == "" {
		response.ResponseError400(c, "车辆ID不能为空")
		return
	}

	// 转换ID为int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "车辆ID格式错误")
		return
	}

	getRes, err := handler.GetVehicle(c, &vehicle.GetVehicleRequest{
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
		BrandId:    req.BrandID,
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

// ==================== 车辆品牌处理器 ====================

// CreateVehicleBrandHandler 创建车辆品牌处理器
func CreateVehicleBrandHandler(c *gin.Context) {
	var req request.CreateVehicleBrandRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	createRes, err := handler.CreateVehicleBrand(c, &vehicle.CreateVehicleBrandRequest{
		Name:        req.Name,
		EnglishName: req.EnglishName,
		Logo:        req.Logo,
		Country:     req.Country,
		Description: req.Description,
		Status:      req.Status,
		Sort:        req.Sort,
		IsHot:       req.IsHot,
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
		"message":       createRes.Message,
		"vehicle_brand": createRes.VehicleBrand,
	})
}

// UpdateVehicleBrandHandler 更新车辆品牌处理器
func UpdateVehicleBrandHandler(c *gin.Context) {
	var req request.UpdateVehicleBrandRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	updateRes, err := handler.UpdateVehicleBrand(c, &vehicle.UpdateVehicleBrandRequest{
		Id:          req.ID,
		Name:        req.Name,
		EnglishName: req.EnglishName,
		Logo:        req.Logo,
		Country:     req.Country,
		Description: req.Description,
		Status:      req.Status,
		Sort:        req.Sort,
		IsHot:       req.IsHot,
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
		"message":       updateRes.Message,
		"vehicle_brand": updateRes.VehicleBrand,
	})
}

// DeleteVehicleBrandHandler 删除车辆品牌处理器
func DeleteVehicleBrandHandler(c *gin.Context) {
	var req request.DeleteVehicleBrandRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	deleteRes, err := handler.DeleteVehicleBrand(c, &vehicle.DeleteVehicleBrandRequest{
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

// GetVehicleBrandHandler 获取车辆品牌详情处理器
func GetVehicleBrandHandler(c *gin.Context) {
	// 从URL路径参数中获取ID
	idStr := c.Param("id")
	if idStr == "" {
		response.ResponseError400(c, "车辆品牌ID不能为空")
		return
	}

	// 转换ID为int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "车辆品牌ID格式错误")
		return
	}

	getRes, err := handler.GetVehicleBrand(c, &vehicle.GetVehicleBrandRequest{
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
		"message":       getRes.Message,
		"vehicle_brand": getRes.VehicleBrand,
	})
}

// ListVehicleBrandsHandler 获取车辆品牌列表处理器
func ListVehicleBrandsHandler(c *gin.Context) {
	var req request.ListVehicleBrandsRequest
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

	// 处理is_hot参数：如果没有明确指定is_hot，设置为-1表示不筛选
	isHot := req.IsHot
	if c.Query("is_hot") == "" && c.PostForm("is_hot") == "" {
		isHot = -1 // 没有传递is_hot参数时，设置为-1表示不筛选
	}

	listRes, err := handler.ListVehicleBrands(c, &vehicle.ListVehicleBrandsRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Status:   status,
		IsHot:    isHot,
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
		"message":        listRes.Message,
		"vehicle_brands": listRes.VehicleBrands,
		"total":          listRes.Total,
		"page":           req.Page,
		"pageSize":       req.PageSize,
	})
}

// ==================== 车辆库存处理器 ====================

// CheckAvailabilityHandler 检查车辆可用性处理器
func CheckAvailabilityHandler(c *gin.Context) {
	var req request.CheckAvailabilityRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	checkRes, err := handler.CheckAvailability(c, &vehicle.CheckAvailabilityRequest{
		VehicleId: req.VehicleID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if checkRes.Code != 200 {
		response.ResponseError400(c, checkRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      checkRes.Message,
		"is_available": checkRes.IsAvailable,
	})
}

// CreateReservationHandler 创建预订处理器（新流程：先预订后订单）
func CreateReservationHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	var req request.CreateReservationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	createRes, err := handler.CreateReservation(c, &vehicle.CreateReservationRequest{
		VehicleId: req.VehicleID,
		UserId:    int64(userID), // 使用JWT中的用户ID
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Notes:     req.Notes,
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
		"message":        createRes.Message,
		"reservation_id": createRes.ReservationId,
	})
}

// CancelReservationHandler 取消预订处理器
func CancelReservationHandler(c *gin.Context) {
	// 从JWT中获取用户ID（用于权限验证）
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	var req request.CancelReservationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	cancelRes, err := handler.CancelReservation(c, &vehicle.CancelReservationRequest{
		ReservationId: req.ReservationID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if cancelRes.Code != 200 {
		response.ResponseError400(c, cancelRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": cancelRes.Message,
	})
}

// 订单相关的handler已移动到 Api/trigger/order.go 文件中

// GetUserReservationListHandler 获取用户预订列表处理器
func GetUserReservationListHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		response.ResponseError400(c, "用户ID不能为空")
		return
	}

	// 获取查询参数
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	status := c.Query("status") // 可选的状态筛选

	// 转换参数
	pageInt := 1
	pageSizeInt := 10
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageInt = p
	}
	if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
		pageSizeInt = ps
	}

	// 调用车辆微服务获取用户预订列表
	listRes, err := handler.GetUserReservationList(c, &vehicle.GetUserReservationListRequest{
		UserId:   int64(userID),
		Page:     int32(pageInt),
		PageSize: int32(pageSizeInt),
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
		"message":      listRes.Message,
		"reservations": listRes.Reservations,
		"total":        listRes.Total,
		"page":         pageInt,
		"page_size":    pageSizeInt,
	})
}

// UpdateReservationStatusHandler 更新预订状态处理器
func UpdateReservationStatusHandler(c *gin.Context) {
	var req request.UpdateReservationStatusRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	updateRes, err := handler.UpdateReservationStatus(c, &vehicle.UpdateReservationStatusRequest{
		OrderId: req.OrderID,
		Status:  req.Status,
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
	})
}

// GetAvailableVehiclesHandler 获取可用车辆处理器
func GetAvailableVehiclesHandler(c *gin.Context) {
	var req request.GetAvailableVehiclesRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 12
	}

	getRes, err := handler.GetAvailableVehicles(c, &vehicle.GetAvailableVehiclesRequest{
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		MerchantId: req.MerchantID,
		TypeId:     req.TypeID,
		BrandId:    req.BrandID,
		Status:     req.Status,
		PriceMin:   req.PriceMin,
		PriceMax:   req.PriceMax,
		Page:       req.Page,
		PageSize:   req.PageSize,
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
		"message":  getRes.Message,
		"vehicles": getRes.Vehicles,
		"total":    getRes.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// GetInventoryStatsHandler 获取库存统计处理器
func GetInventoryStatsHandler(c *gin.Context) {
	// 从JWT中获取商家ID
	mid := c.GetUint("userId")
	if mid == 0 {
		response.ResponseError400(c, "商家ID不能为空")
		return
	}

	statsRes, err := handler.GetInventoryStats(c, &vehicle.GetInventoryStatsRequest{
		MerchantId: int64(mid),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if statsRes.Code != 200 {
		response.ResponseError400(c, statsRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":     statsRes.Message,
		"total":       statsRes.Total,
		"available":   statsRes.Available,
		"reserved":    statsRes.Reserved,
		"rented":      statsRes.Rented,
		"maintenance": statsRes.Maintenance,
	})
}

// SetMaintenanceHandler 设置维护状态处理器
func SetMaintenanceHandler(c *gin.Context) {
	// 从JWT中获取商家ID作为创建人
	mid := c.GetUint("userId")

	var req request.SetMaintenanceRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 如果没有指定创建人，使用当前登录的商家ID
	if req.CreatedBy == 0 {
		req.CreatedBy = int64(mid)
	}

	setRes, err := handler.SetMaintenance(c, &vehicle.SetMaintenanceRequest{
		VehicleId: req.VehicleID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Notes:     req.Notes,
		CreatedBy: req.CreatedBy,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if setRes.Code != 200 {
		response.ResponseError400(c, setRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message": setRes.Message,
	})
}

// GetMaintenanceScheduleHandler 获取维护计划处理器
func GetMaintenanceScheduleHandler(c *gin.Context) {
	var req request.GetMaintenanceScheduleRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	scheduleRes, err := handler.GetMaintenanceSchedule(c, &vehicle.GetMaintenanceScheduleRequest{
		VehicleId: req.VehicleID,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if scheduleRes.Code != 200 {
		response.ResponseError400(c, scheduleRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":      scheduleRes.Message,
		"maintenances": scheduleRes.Maintenances,
		"total":        scheduleRes.Total,
	})
}

// GetInventoryCalendarHandler 获取库存日历处理器（支持公开访问）
func GetInventoryCalendarHandler(c *gin.Context) {
	var req request.GetInventoryCalendarRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	calendarRes, err := handler.GetInventoryCalendar(c, &vehicle.GetInventoryCalendarRequest{
		VehicleId: req.VehicleID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if calendarRes.Code != 200 {
		response.ResponseError400(c, calendarRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":  calendarRes.Message,
		"calendar": calendarRes.Calendar,
	})
}

// GetInventoryReportHandler 获取库存报表处理器
func GetInventoryReportHandler(c *gin.Context) {
	// 从JWT中获取商家ID
	mid := c.GetUint("userId")

	var req request.GetInventoryReportRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 如果没有指定商家ID，使用当前登录的商家ID
	if req.MerchantID == 0 {
		req.MerchantID = int64(mid)
	}

	reportRes, err := handler.GetInventoryReport(c, &vehicle.GetInventoryReportRequest{
		MerchantId: req.MerchantID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if reportRes.Code != 200 {
		response.ResponseError400(c, reportRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{
		"message":          reportRes.Message,
		"total_vehicles":   reportRes.TotalVehicles,
		"total_days":       reportRes.TotalDays,
		"total_capacity":   reportRes.TotalCapacity,
		"reservations":     reportRes.Reservations,
		"rentals":          reportRes.Rentals,
		"maintenances":     reportRes.Maintenances,
		"used_capacity":    reportRes.UsedCapacity,
		"utilization_rate": reportRes.UtilizationRate,
	})
}
