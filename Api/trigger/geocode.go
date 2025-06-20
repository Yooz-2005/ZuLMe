package trigger

import (
	"Api/handler"
	"Api/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCoordinates 获取地址的经纬度坐标
func GetCoordinates(c *gin.Context) {
	var req request.GeocodeRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 调用handler处理
	response, err := handler.GetCoordinates(c, req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
