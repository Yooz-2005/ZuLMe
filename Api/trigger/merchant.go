package trigger

import (
	"Api/handler"
	"Api/request"
	"Api/response"
	"Common/utils"
	merchant "merchant_srv/proto_merchant"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MerchantRegisterHandler(c *gin.Context) {
	var req request.MerchantRegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	registerRes, err := handler.MerchantRegister(c, &merchant.MerchantRegisterRequest{
		Name:         req.Name,
		Phone:        req.Phone,
		Email:        req.Email,
		Password:     req.Password,
		ConfirmPass:  req.ConfirmPass,
		Location:     req.Location,
		BusinessTime: req.BusinessTime,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if registerRes.Code != 200 {
		response.ResponseError400(c, registerRes.Message)
		return
	}

	response.ResponseSuccess(c, registerRes.Message)
}

func MerchantLoginHandler(c *gin.Context) {
	var req request.MerchantLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	loginRes, err := handler.MerchantLogin(c, &merchant.MerchantLoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if loginRes.Code != 200 {
		response.ResponseError400(c, loginRes.Message)
		return
	}

	response.ResponseSuccess(c, gin.H{"message": loginRes.Message, "token": loginRes.Token})
}

// SyncMerchantLocationsHandler 同步已有商家位置数据到Redis
func SyncMerchantLocationsHandler(c *gin.Context) {
	err := utils.SyncExistingMerchantsToRedis()
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, "商家位置数据同步成功")
}

// ValidateMerchantLocationDataHandler 验证商家位置数据完整性
func ValidateMerchantLocationDataHandler(c *gin.Context) {
	err := utils.ValidateMerchantLocationData()
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, "商家位置数据验证完成")
}

// FixMerchantCoordinatesHandler 修复缺少坐标的商家数据
func FixMerchantCoordinatesHandler(c *gin.Context) {
	err := utils.FixMerchantCoordinates()
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, "商家坐标数据修复完成")
}

// UpdateMerchantLocationHandler 更新单个商家在Redis中的位置信息
func UpdateMerchantLocationHandler(c *gin.Context) {
	merchantIDStr := c.Param("id")
	merchantID, err := strconv.ParseInt(merchantIDStr, 10, 64)
	if err != nil {
		response.ResponseError400(c, "无效的商家ID")
		return
	}

	err = utils.UpdateMerchantLocationInRedis(merchantID)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, "商家位置信息更新成功")
}

// GetMerchantLocationsHandler 获取所有审核通过的商家网点列表（公开接口）
func GetMerchantLocationsHandler(c *gin.Context) {
	var req request.GetLocationListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 100
	}
	if req.StatusFilter < 0 {
		req.StatusFilter = 1 // 默认只获取审核通过的商户
	}

	// 调用merchant服务
	locationListRes, err := handler.GetLocationList(c, &merchant.GetLocationListRequest{
		Page:         req.Page,
		PageSize:     req.PageSize,
		StatusFilter: req.StatusFilter,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	if locationListRes.Code != 200 {
		response.ResponseError400(c, locationListRes.Message)
		return
	}

	// 转换为前端需要的格式
	var merchantList []gin.H
	for _, location := range locationListRes.Locations {
		merchantList = append(merchantList, gin.H{
			"id":            location.Id,
			"name":          location.Name,
			"phone":         location.Phone,
			"email":         location.Email,
			"status":        location.Status,
			"location":      location.Location,
			"business_time": location.BusinessTime,
			"longitude":     location.Longitude,
			"latitude":      location.Latitude,
			"created_at":    location.CreatedAt,
			"updated_at":    location.UpdatedAt,
		})
	}

	response.ResponseSuccess(c, gin.H{
		"message":   locationListRes.Message,
		"merchants": merchantList,
		"total":     locationListRes.Total,
		"page":      req.Page,
		"page_size": req.PageSize,
	})
}
