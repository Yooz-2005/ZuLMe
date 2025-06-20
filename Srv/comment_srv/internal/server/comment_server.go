package server

import (
	"comment_srv/internal/logic"
	"comment_srv/proto_comment"
	"context"
	"models/model_mongodb"
	"time"
)

// CommentServer 评论服务实现
type CommentServer struct {
	proto_comment.UnimplementedCommentServiceServer
	commentLogic *logic.CommentLogic
}

// NewCommentServer 创建评论服务实例
func NewCommentServer() *CommentServer {
	return &CommentServer{
		commentLogic: logic.NewCommentLogic(),
	}
}

// CreateComment 创建评论
func (s *CommentServer) CreateComment(ctx context.Context, req *proto_comment.CreateCommentRequest) (*proto_comment.CreateCommentResponse, error) {
	// 参数验证
	if req.OrderId == 0 || req.UserId == 0 {
		return &proto_comment.CreateCommentResponse{
			Code:    400,
			Message: "订单ID和用户ID不能为空",
		}, nil
	}

	if req.Rating < 1 || req.Rating > 5 {
		return &proto_comment.CreateCommentResponse{
			Code:    400,
			Message: "评分必须在1-5之间",
		}, nil
	}

	// 创建评论
	comment, err := s.commentLogic.CreateComment(
		uint(req.OrderId),
		uint(req.UserId),
		uint(req.VehicleId),
		req.Rating,
		req.Content,
		req.Images,
		req.ServiceRating,
		req.VehicleRating,
		req.CleanRating,
		req.IsAnonymous,
	)

	if err != nil {
		return &proto_comment.CreateCommentResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	commentInfo := s.convertToCommentInfo(comment)

	return &proto_comment.CreateCommentResponse{
		Code:    200,
		Message: "评论创建成功",
		Comment: commentInfo,
	}, nil
}

// GetComment 获取评论详情
func (s *CommentServer) GetComment(ctx context.Context, req *proto_comment.GetCommentRequest) (*proto_comment.GetCommentResponse, error) {
	if req.CommentId == "" {
		return &proto_comment.GetCommentResponse{
			Code:    400,
			Message: "评论ID不能为空",
		}, nil
	}

	comment, err := s.commentLogic.GetComment(req.CommentId)
	if err != nil {
		return &proto_comment.GetCommentResponse{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	commentInfo := s.convertToCommentInfo(comment)

	return &proto_comment.GetCommentResponse{
		Code:    200,
		Message: "获取评论成功",
		Comment: commentInfo,
	}, nil
}

// GetOrderComment 获取订单评论
func (s *CommentServer) GetOrderComment(ctx context.Context, req *proto_comment.GetOrderCommentRequest) (*proto_comment.GetCommentResponse, error) {
	if req.OrderId == 0 {
		return &proto_comment.GetCommentResponse{
			Code:    400,
			Message: "订单ID不能为空",
		}, nil
	}

	comment, err := s.commentLogic.GetOrderComment(uint(req.OrderId))
	if err != nil {
		return &proto_comment.GetCommentResponse{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	commentInfo := s.convertToCommentInfo(comment)

	return &proto_comment.GetCommentResponse{
		Code:    200,
		Message: "获取订单评论成功",
		Comment: commentInfo,
	}, nil
}

// GetVehicleComments 获取车辆评论列表
func (s *CommentServer) GetVehicleComments(ctx context.Context, req *proto_comment.GetVehicleCommentsRequest) (*proto_comment.GetVehicleCommentsResponse, error) {
	if req.VehicleId == 0 {
		return &proto_comment.GetVehicleCommentsResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	comments, total, err := s.commentLogic.GetVehicleComments(uint(req.VehicleId), page, pageSize)
	if err != nil {
		return &proto_comment.GetVehicleCommentsResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	var commentInfos []*proto_comment.CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, s.convertToCommentInfo(&comment))
	}

	return &proto_comment.GetVehicleCommentsResponse{
		Code:     200,
		Message:  "获取车辆评论成功",
		Comments: commentInfos,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// GetUserComments 获取用户评论列表
func (s *CommentServer) GetUserComments(ctx context.Context, req *proto_comment.GetUserCommentsRequest) (*proto_comment.GetUserCommentsResponse, error) {
	if req.UserId == 0 {
		return &proto_comment.GetUserCommentsResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	comments, total, err := s.commentLogic.GetUserComments(uint(req.UserId), page, pageSize)
	if err != nil {
		return &proto_comment.GetUserCommentsResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	var commentInfos []*proto_comment.CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, s.convertToCommentInfo(&comment))
	}

	return &proto_comment.GetUserCommentsResponse{
		Code:     200,
		Message:  "获取用户评论成功",
		Comments: commentInfos,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// GetVehicleStats 获取车辆评论统计
func (s *CommentServer) GetVehicleStats(ctx context.Context, req *proto_comment.GetVehicleStatsRequest) (*proto_comment.GetVehicleStatsResponse, error) {
	if req.VehicleId == 0 {
		return &proto_comment.GetVehicleStatsResponse{
			Code:    400,
			Message: "车辆ID不能为空",
		}, nil
	}

	stats, err := s.commentLogic.GetVehicleStats(uint(req.VehicleId))
	if err != nil {
		return &proto_comment.GetVehicleStatsResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_comment.GetVehicleStatsResponse{
		Code:    200,
		Message: "获取车辆评论统计成功",
		Stats: &proto_comment.CommentStats{
			TotalComments:  stats.TotalComments,
			AverageRating:  stats.AverageRating,
			FiveStarCount:  stats.FiveStarCount,
			FourStarCount:  stats.FourStarCount,
			ThreeStarCount: stats.ThreeStarCount,
			TwoStarCount:   stats.TwoStarCount,
			OneStarCount:   stats.OneStarCount,
			ServiceRating:  stats.ServiceRating,
			VehicleRating:  stats.VehicleRating,
			CleanRating:    stats.CleanRating,
		},
	}, nil
}

// UpdateComment 更新评论
func (s *CommentServer) UpdateComment(ctx context.Context, req *proto_comment.UpdateCommentRequest) (*proto_comment.UpdateCommentResponse, error) {
	if req.CommentId == "" {
		return &proto_comment.UpdateCommentResponse{
			Code:    400,
			Message: "评论ID不能为空",
		}, nil
	}

	if req.UserId == 0 {
		return &proto_comment.UpdateCommentResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	comment, err := s.commentLogic.UpdateComment(
		req.CommentId,
		uint(req.UserId),
		req.Rating,
		req.Content,
		req.Images,
		req.ServiceRating,
		req.VehicleRating,
		req.CleanRating,
		req.IsAnonymous,
	)

	if err != nil {
		return &proto_comment.UpdateCommentResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	commentInfo := s.convertToCommentInfo(comment)

	return &proto_comment.UpdateCommentResponse{
		Code:    200,
		Message: "评论更新成功",
		Comment: commentInfo,
	}, nil
}

// DeleteComment 删除评论
func (s *CommentServer) DeleteComment(ctx context.Context, req *proto_comment.DeleteCommentRequest) (*proto_comment.DeleteCommentResponse, error) {
	if req.CommentId == "" {
		return &proto_comment.DeleteCommentResponse{
			Code:    400,
			Message: "评论ID不能为空",
		}, nil
	}

	if req.UserId == 0 {
		return &proto_comment.DeleteCommentResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	err := s.commentLogic.DeleteComment(req.CommentId, uint(req.UserId))
	if err != nil {
		return &proto_comment.DeleteCommentResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_comment.DeleteCommentResponse{
		Code:    200,
		Message: "评论删除成功",
	}, nil
}

// ReplyComment 商家回复评论
func (s *CommentServer) ReplyComment(ctx context.Context, req *proto_comment.ReplyCommentRequest) (*proto_comment.ReplyCommentResponse, error) {
	if req.CommentId == "" || req.ReplyContent == "" {
		return &proto_comment.ReplyCommentResponse{
			Code:    400,
			Message: "评论ID和回复内容不能为空",
		}, nil
	}

	if req.MerchantId == 0 {
		return &proto_comment.ReplyCommentResponse{
			Code:    400,
			Message: "商家ID不能为空",
		}, nil
	}

	comment, err := s.commentLogic.ReplyComment(req.CommentId, req.ReplyContent, uint(req.MerchantId))
	if err != nil {
		return &proto_comment.ReplyCommentResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	commentInfo := s.convertToCommentInfo(comment)

	return &proto_comment.ReplyCommentResponse{
		Code:    200,
		Message: "回复评论成功",
		Comment: commentInfo,
	}, nil
}

// CheckOrderCommented 检查订单是否已评论
func (s *CommentServer) CheckOrderCommented(ctx context.Context, req *proto_comment.CheckOrderCommentedRequest) (*proto_comment.CheckOrderCommentedResponse, error) {
	if req.OrderId == 0 {
		return &proto_comment.CheckOrderCommentedResponse{
			Code:    400,
			Message: "订单ID不能为空",
		}, nil
	}

	commented, err := s.commentLogic.CheckOrderCommented(uint(req.OrderId))
	if err != nil {
		return &proto_comment.CheckOrderCommentedResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_comment.CheckOrderCommentedResponse{
		Code:      200,
		Message:   "检查成功",
		Commented: commented,
	}, nil
}

// convertToCommentInfo 转换评论模型为gRPC响应格式
func (s *CommentServer) convertToCommentInfo(comment *model_mongodb.Comment) *proto_comment.CommentInfo {
	commentInfo := &proto_comment.CommentInfo{
		Id:            comment.ID.Hex(),
		OrderId:       uint32(comment.OrderID),
		UserId:        uint32(comment.UserID),
		VehicleId:     uint32(comment.VehicleID),
		Rating:        comment.Rating,
		Content:       comment.Content,
		Images:        comment.Images,
		ServiceRating: comment.ServiceRating,
		VehicleRating: comment.VehicleRating,
		CleanRating:   comment.CleanRating,
		IsAnonymous:   comment.IsAnonymous,
		Status:        comment.Status,
		ReplyContent:  comment.ReplyContent,
		CreatedAt:     comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     comment.UpdatedAt.Format(time.RFC3339),
	}

	if comment.ReplyTime != nil {
		commentInfo.ReplyTime = comment.ReplyTime.Format(time.RFC3339)
	}

	return commentInfo
}
