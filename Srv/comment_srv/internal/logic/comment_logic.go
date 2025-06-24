package logic

import (
	"Common/global"
	"context"
	"fmt"
	"models/model_mongodb"
	"models/model_mysql"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// CommentLogic 评论业务逻辑
type CommentLogic struct{}

// NewCommentLogic 创建评论业务逻辑实例
func NewCommentLogic() *CommentLogic {
	return &CommentLogic{}
}

// CreateComment 创建评论
func (cl *CommentLogic) CreateComment(orderID, userID, vehicleID uint, rating int32, content string, images []string, serviceRating, vehicleRating, cleanRating int32, isAnonymous bool) (*model_mongodb.Comment, error) {
	fmt.Printf("开始创建评论: orderID=%d, userID=%d, vehicleID=%d\n", orderID, userID, vehicleID)

	// 1. 验证订单是否存在且属于该用户
	var order model_mysql.Orders
	if err := global.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		fmt.Printf("查询订单失败: orderID=%d, userID=%d, error=%v\n", orderID, userID, err)
		return nil, fmt.Errorf("订单不存在或不属于该用户: orderID=%d, userID=%d, error=%v", orderID, userID, err)
	}
	fmt.Printf("找到订单: ID=%d, VehicleId=%d, Status=%d\n", order.ID, order.VehicleId, order.Status)

	// 2. 从订单中获取车辆ID（如果传入的vehicleID为0）
	if vehicleID == 0 {
		vehicleID = order.VehicleId
		fmt.Printf("从订单中获取车辆ID: %d\n", vehicleID)
	}

	// 3. 验证订单状态（只有已完成或已还车的订单才能评论）
	// 为了测试，暂时允许已支付状态的订单也能评论
	if order.Status != model_mysql.OrderStatusCompleted && order.Status != model_mysql.OrderStatusReturned && order.Status != model_mysql.OrderStatusPaid {
		return nil, fmt.Errorf("订单状态不允许评论，当前状态: %d，只有已支付(2)、已完成(4)或已还车(6)的订单才能评论", order.Status)
	}
	fmt.Printf("订单状态验证通过: %d\n", order.Status)

	// 4. 检查是否已经评论过（暂时注释掉用于调试）
	fmt.Printf("跳过检查订单 %d 是否已评论过（调试模式）\n", orderID)
	// exists, err := model_mongodb.CheckOrderCommentExists(orderID)
	// if err != nil {
	// 	fmt.Printf("检查订单评论状态失败: orderID=%d, error=%v\n", orderID, err)
	// 	return nil, fmt.Errorf("检查订单评论状态失败: %v", err)
	// }
	// if exists {
	// 	return nil, fmt.Errorf("该订单已经评论过了")
	// }

	// 5. 验证评分范围
	if rating < 1 || rating > 5 || serviceRating < 1 || serviceRating > 5 || vehicleRating < 1 || vehicleRating > 5 || cleanRating < 1 || cleanRating > 5 {
		return nil, fmt.Errorf("评分必须在1-5之间")
	}

	// 6. 创建评论
	fmt.Printf("开始创建评论对象\n")
	comment := &model_mongodb.Comment{
		OrderID:       orderID,
		UserID:        userID,
		VehicleID:     vehicleID,
		Rating:        rating,
		Content:       content,
		Images:        images,
		ServiceRating: serviceRating,
		VehicleRating: vehicleRating,
		CleanRating:   cleanRating,
		IsAnonymous:   isAnonymous,
	}
	fmt.Printf("评论对象创建完成，准备保存到MongoDB\n")

	if err := comment.Create(); err != nil {
		fmt.Printf("保存评论到MongoDB失败: %v\n", err)
		return nil, fmt.Errorf("创建评论失败: %v", err)
	}
	fmt.Printf("评论创建成功，ID: %s\n", comment.ID.Hex())

	return comment, nil
}

// GetComment 获取评论详情
func (cl *CommentLogic) GetComment(commentID string) (*model_mongodb.Comment, error) {
	comment := &model_mongodb.Comment{}
	if err := comment.GetByID(commentID); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetOrderComment 获取订单评论
func (cl *CommentLogic) GetOrderComment(orderID uint) (*model_mongodb.Comment, error) {
	comment := &model_mongodb.Comment{}
	if err := comment.GetByOrderID(orderID); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetVehicleComments 获取车辆评论列表
func (cl *CommentLogic) GetVehicleComments(vehicleID uint, page, pageSize int32) ([]model_mongodb.Comment, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return model_mongodb.GetVehicleComments(vehicleID, page, pageSize)
}

// GetUserComments 获取用户评论列表
func (cl *CommentLogic) GetUserComments(userID uint, page, pageSize int32) ([]model_mongodb.Comment, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return model_mongodb.GetUserComments(userID, page, pageSize)
}

// GetVehicleStats 获取车辆评论统计
func (cl *CommentLogic) GetVehicleStats(vehicleID uint) (*model_mongodb.CommentStats, error) {
	return model_mongodb.GetVehicleStats(vehicleID)
}

// UpdateComment 更新评论
func (cl *CommentLogic) UpdateComment(commentID string, userID uint, rating int32, content string, images []string, serviceRating, vehicleRating, cleanRating int32, isAnonymous bool) (*model_mongodb.Comment, error) {
	// 1. 获取原评论
	comment := &model_mongodb.Comment{}
	if err := comment.GetByID(commentID); err != nil {
		return nil, err
	}

	// 2. 验证权限（只有评论作者可以修改）
	if comment.UserID != userID {
		return nil, fmt.Errorf("无权限修改该评论")
	}

	// 3. 验证评分范围
	if rating < 1 || rating > 5 || serviceRating < 1 || serviceRating > 5 || vehicleRating < 1 || vehicleRating > 5 || cleanRating < 1 || cleanRating > 5 {
		return nil, fmt.Errorf("评分必须在1-5之间")
	}

	// 4. 检查评论创建时间，只允许在创建后24小时内修改
	if time.Since(comment.CreatedAt) > 24*time.Hour {
		return nil, fmt.Errorf("评论创建超过24小时，不允许修改")
	}

	// 5. 更新评论内容
	comment.Rating = rating
	comment.Content = content
	comment.Images = images
	comment.ServiceRating = serviceRating
	comment.VehicleRating = vehicleRating
	comment.CleanRating = cleanRating
	comment.IsAnonymous = isAnonymous

	if err := comment.Update(); err != nil {
		return nil, fmt.Errorf("更新评论失败: %v", err)
	}

	return comment, nil
}

// DeleteComment 删除评论
func (cl *CommentLogic) DeleteComment(commentID string, userID uint) error {
	// 1. 获取评论
	comment := &model_mongodb.Comment{}
	if err := comment.GetByID(commentID); err != nil {
		return err
	}

	// 2. 验证权限（只有评论作者可以删除）
	if comment.UserID != userID {
		return fmt.Errorf("无权限删除该评论")
	}

	// 3. 检查评论创建时间，只允许在创建后24小时内删除
	if time.Since(comment.CreatedAt) > 24*time.Hour {
		return fmt.Errorf("评论创建超过24小时，不允许删除")
	}

	// 4. 软删除评论
	if err := comment.Delete(); err != nil {
		return fmt.Errorf("删除评论失败: %v", err)
	}

	return nil
}

// ReplyComment 商家回复评论
func (cl *CommentLogic) ReplyComment(commentID string, replyContent string, merchantID uint) (*model_mongodb.Comment, error) {
	// 1. 获取评论
	comment := &model_mongodb.Comment{}
	if err := comment.GetByID(commentID); err != nil {
		return nil, err
	}

	// 2. 验证商家权限（需要验证商家是否拥有该车辆）
	var vehicle model_mysql.Vehicle
	if err := global.DB.Where("id = ? AND merchant_id = ?", comment.VehicleID, merchantID).First(&vehicle).Error; err != nil {
		return nil, fmt.Errorf("无权限回复该评论")
	}

	// 3. 检查是否已经回复过
	if comment.ReplyContent != "" {
		return nil, fmt.Errorf("该评论已经回复过了")
	}

	// 4. 添加回复
	now := time.Now()
	comment.ReplyContent = replyContent
	comment.ReplyTime = &now

	if err := comment.Update(); err != nil {
		return nil, fmt.Errorf("回复评论失败: %v", err)
	}

	return comment, nil
}

// CheckOrderCommented 检查订单是否已评论
func (cl *CommentLogic) CheckOrderCommented(orderID uint) (bool, error) {
	return model_mongodb.CheckOrderCommentExists(orderID)
}

// ValidateCommentPermission 验证评论权限
func (cl *CommentLogic) ValidateCommentPermission(orderID, userID uint) error {
	// 验证订单是否存在且属于该用户
	var order model_mysql.Orders
	if err := global.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return fmt.Errorf("订单不存在或不属于该用户")
	}

	// 验证订单状态
	if order.Status != model_mysql.OrderStatusCompleted && order.Status != model_mysql.OrderStatusReturned {
		return fmt.Errorf("订单状态不允许评论")
	}

	return nil
}

// GetCommentsByDateRange 根据日期范围获取评论
func (cl *CommentLogic) GetCommentsByDateRange(vehicleID uint, startDate, endDate time.Time, page, pageSize int32) ([]model_mongodb.Comment, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	collection := global.MongoClient.Database("zulme").Collection("comments")

	filter := map[string]interface{}{
		"vehicle_id": vehicleID,
		"status":     model_mongodb.CommentStatusNormal,
		"created_at": map[string]interface{}{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	// 计算总数
	total, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, fmt.Errorf("统计评论数量失败: %v", err)
	}

	// 分页查询
	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)
	cursor, err := collection.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  map[string]interface{}{"created_at": -1},
	})
	if err != nil {
		return nil, 0, fmt.Errorf("查询评论失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var comments []model_mongodb.Comment
	if err = cursor.All(context.Background(), &comments); err != nil {
		return nil, 0, fmt.Errorf("解析评论数据失败: %v", err)
	}

	return comments, total, nil
}

// GetTopRatedVehicles 获取评分最高的车辆
func (cl *CommentLogic) GetTopRatedVehicles(limit int32) ([]map[string]interface{}, error) {
	collection := global.MongoClient.Database("zulme").Collection("comments")

	pipeline := []map[string]interface{}{
		{
			"$match": map[string]interface{}{
				"status": model_mongodb.CommentStatusNormal,
			},
		},
		{
			"$group": map[string]interface{}{
				"_id":            "$vehicle_id",
				"average_rating": map[string]interface{}{"$avg": "$rating"},
				"comment_count":  map[string]interface{}{"$sum": 1},
			},
		},
		{
			"$match": map[string]interface{}{
				"comment_count": map[string]interface{}{"$gte": 5}, // 至少5条评论
			},
		},
		{
			"$sort": map[string]interface{}{
				"average_rating": -1,
			},
		},
		{
			"$limit": limit,
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("查询高评分车辆失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var results []map[string]interface{}
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}

	return results, nil
}
