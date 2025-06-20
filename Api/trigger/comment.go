package trigger

import (
	"Api/client"
	"Api/request"
	"comment_srv/proto_comment"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCommentHandler 创建评论
func CreateCommentHandler(c *gin.Context) {
	var req request.CreateCommentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID（暂时硬编码用于测试）
	userID := c.GetUint("userId")

	// 添加调试信息
	fmt.Printf("API层调试信息: OrderID=%d, UserID=%d\n", req.OrderID, userID)

	// 调用评论服务
	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.CreateComment(ctx, &proto_comment.CreateCommentRequest{
			OrderId:       uint32(req.OrderID),
			UserId:        uint32(userID),
			Rating:        req.Rating,
			Content:       req.Content,
			Images:        req.Images,
			ServiceRating: req.ServiceRating,
			VehicleRating: req.VehicleRating,
			CleanRating:   req.CleanRating,
			IsAnonymous:   req.IsAnonymous,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.CreateCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Comment,
	})
}

// GetCommentHandler 获取评论详情
func GetCommentHandler(c *gin.Context) {
	commentID := c.Param("comment_id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论ID不能为空",
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.GetComment(ctx, &proto_comment.GetCommentRequest{
			CommentId: commentID,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.GetCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Comment,
	})
}

// GetOrderCommentHandler 获取订单评论
func GetOrderCommentHandler(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "订单ID格式错误",
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.GetOrderComment(ctx, &proto_comment.GetOrderCommentRequest{
			OrderId: uint32(orderID),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.GetCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Comment,
	})
}

// GetVehicleCommentsHandler 获取车辆评论列表
func GetVehicleCommentsHandler(c *gin.Context) {
	vehicleIDStr := c.Param("vehicle_id")
	vehicleID, err := strconv.ParseUint(vehicleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "车辆ID格式错误",
		})
		return
	}

	// 获取分页参数
	var req request.GetVehicleCommentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		// 使用默认值
		req.Page = 1
		req.PageSize = 10
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.GetVehicleComments(ctx, &proto_comment.GetVehicleCommentsRequest{
			VehicleId: uint32(vehicleID),
			Page:      req.Page,
			PageSize:  req.PageSize,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.GetVehicleCommentsResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"message":   resp.Message,
		"data":      resp.Comments,
		"total":     resp.Total,
		"page":      resp.Page,
		"page_size": resp.PageSize,
	})
}

// GetUserCommentsHandler 获取用户评论列表
func GetUserCommentsHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "用户ID格式错误",
		})
		return
	}

	// 获取分页参数
	var req request.GetUserCommentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		// 使用默认值
		req.Page = 1
		req.PageSize = 10
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.GetUserComments(ctx, &proto_comment.GetUserCommentsRequest{
			UserId:   uint32(userID),
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.GetUserCommentsResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"message":   resp.Message,
		"data":      resp.Comments,
		"total":     resp.Total,
		"page":      resp.Page,
		"page_size": resp.PageSize,
	})
}

// GetVehicleStatsHandler 获取车辆评论统计
func GetVehicleStatsHandler(c *gin.Context) {
	vehicleIDStr := c.Param("vehicle_id")
	vehicleID, err := strconv.ParseUint(vehicleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "车辆ID格式错误",
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.GetVehicleStats(ctx, &proto_comment.GetVehicleStatsRequest{
			VehicleId: uint32(vehicleID),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.GetVehicleStatsResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Stats,
	})
}

// UpdateCommentHandler 更新评论
func UpdateCommentHandler(c *gin.Context) {
	commentID := c.Param("comment_id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论ID不能为空",
		})
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	var req request.UpdateCommentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.UpdateComment(ctx, &proto_comment.UpdateCommentRequest{
			CommentId:     commentID,
			Rating:        req.Rating,
			Content:       req.Content,
			Images:        req.Images,
			ServiceRating: req.ServiceRating,
			VehicleRating: req.VehicleRating,
			CleanRating:   req.CleanRating,
			IsAnonymous:   req.IsAnonymous,
			UserId:        uint32(userID.(uint)),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.UpdateCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Comment,
	})
}

// DeleteCommentHandler 删除评论
func DeleteCommentHandler(c *gin.Context) {
	commentID := c.Param("comment_id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论ID不能为空",
		})
		return
	}

	// 从JWT中获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.DeleteComment(ctx, &proto_comment.DeleteCommentRequest{
			CommentId: commentID,
			UserId:    uint32(userID.(uint)),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.DeleteCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
	})
}

// ReplyCommentHandler 商家回复评论
func ReplyCommentHandler(c *gin.Context) {
	commentID := c.Param("comment_id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "评论ID不能为空",
		})
		return
	}

	// 从JWT中获取商家ID
	merchantID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "商家未登录",
		})
		return
	}

	var req request.ReplyCommentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.ReplyComment(ctx, &proto_comment.ReplyCommentRequest{
			CommentId:    commentID,
			ReplyContent: req.ReplyContent,
			MerchantId:   uint32(merchantID.(uint)),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.ReplyCommentResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Comment,
	})
}

// CheckOrderCommentedHandler 检查订单是否已评论
func CheckOrderCommentedHandler(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "订单ID格式错误",
		})
		return
	}

	commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
		return commentClient.CheckOrderCommented(ctx, &proto_comment.CheckOrderCommentedRequest{
			OrderId: uint32(orderID),
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "服务器内部错误",
		})
		return
	}

	resp := commentResp.(*proto_comment.CheckOrderCommentedResponse)
	c.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"message":   resp.Message,
		"commented": resp.Commented,
	})
}

// GetMerchantVehicleCommentsHandler 获取商家车辆的评论列表
func GetMerchantVehicleCommentsHandler(c *gin.Context) {
	// 从JWT中获取商家ID
	merchantID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "商家未登录",
		})
		return
	}

	// 获取分页参数
	var req request.GetVehicleCommentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		req.Page = 1
		req.PageSize = 10
	}

	vehicleIDStr := c.Query("vehicle_id")
	if vehicleIDStr != "" {
		// 获取指定车辆的评论
		vehicleID, err := strconv.ParseUint(vehicleIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "车辆ID格式错误",
			})
			return
		}

		commentResp, err := client.CommentClient(context.Background(), func(ctx context.Context, commentClient proto_comment.CommentServiceClient) (interface{}, error) {
			return commentClient.GetVehicleComments(ctx, &proto_comment.GetVehicleCommentsRequest{
				VehicleId: uint32(vehicleID),
				Page:      req.Page,
				PageSize:  req.PageSize,
			})
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "服务器内部错误",
			})
			return
		}

		resp := commentResp.(*proto_comment.GetVehicleCommentsResponse)
		c.JSON(http.StatusOK, gin.H{
			"code":      resp.Code,
			"message":   resp.Message,
			"data":      resp.Comments,
			"total":     resp.Total,
			"page":      resp.Page,
			"page_size": resp.PageSize,
		})
	} else {
		// 获取商家所有车辆的评论（需要先获取商家的车辆列表）
		// 这里需要调用车辆服务获取商家的车辆列表，然后获取评论
		// 暂时返回空数据，实际实现需要根据具体需求调整
		_ = merchantID // 避免未使用变量警告
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"message":   "获取成功",
			"data":      []interface{}{},
			"total":     0,
			"page":      req.Page,
			"page_size": req.PageSize,
		})
	}
}
