package model_mongodb

import (
	"Common/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Comment 评论模型
type Comment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID     uint               `bson:"order_id" json:"order_id"`                         // 订单ID
	UserID      uint               `bson:"user_id" json:"user_id"`                           // 用户ID
	VehicleID   uint               `bson:"vehicle_id" json:"vehicle_id"`                     // 车辆ID
	Rating      int32              `bson:"rating" json:"rating"`                             // 评分 1-5星
	Content     string             `bson:"content" json:"content"`                           // 评论内容
	Images      []string           `bson:"images,omitempty" json:"images,omitempty"`         // 评论图片URLs
	ServiceRating int32            `bson:"service_rating" json:"service_rating"`             // 服务评分 1-5星
	VehicleRating int32            `bson:"vehicle_rating" json:"vehicle_rating"`             // 车辆评分 1-5星
	CleanRating   int32            `bson:"clean_rating" json:"clean_rating"`                 // 清洁度评分 1-5星
	IsAnonymous   bool             `bson:"is_anonymous" json:"is_anonymous"`                 // 是否匿名评论
	Status        int32            `bson:"status" json:"status"`                             // 评论状态 1:正常 2:隐藏 3:删除
	ReplyContent  string           `bson:"reply_content,omitempty" json:"reply_content,omitempty"` // 商家回复内容
	ReplyTime     *time.Time       `bson:"reply_time,omitempty" json:"reply_time,omitempty"`       // 商家回复时间
	CreatedAt     time.Time        `bson:"created_at" json:"created_at"`                     // 创建时间
	UpdatedAt     time.Time        `bson:"updated_at" json:"updated_at"`                     // 更新时间
}

// CommentStats 评论统计
type CommentStats struct {
	TotalComments   int64   `json:"total_comments"`   // 总评论数
	AverageRating   float64 `json:"average_rating"`   // 平均评分
	FiveStarCount   int64   `json:"five_star_count"`  // 5星评论数
	FourStarCount   int64   `json:"four_star_count"`  // 4星评论数
	ThreeStarCount  int64   `json:"three_star_count"` // 3星评论数
	TwoStarCount    int64   `json:"two_star_count"`   // 2星评论数
	OneStarCount    int64   `json:"one_star_count"`   // 1星评论数
	ServiceRating   float64 `json:"service_rating"`   // 平均服务评分
	VehicleRating   float64 `json:"vehicle_rating"`   // 平均车辆评分
	CleanRating     float64 `json:"clean_rating"`     // 平均清洁度评分
}

const (
	CommentStatusNormal = 1 // 正常
	CommentStatusHidden = 2 // 隐藏
	CommentStatusDeleted = 3 // 删除
)

// GetCollection 获取评论集合
func (c *Comment) GetCollection() *mongo.Collection {
	return global.MongoClient.Database("zulme").Collection("comments")
}

// Create 创建评论
func (c *Comment) Create() error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.Status = CommentStatusNormal
	
	collection := c.GetCollection()
	result, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		return fmt.Errorf("创建评论失败: %v", err)
	}
	
	c.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetByOrderID 根据订单ID获取评论
func (c *Comment) GetByOrderID(orderID uint) error {
	collection := c.GetCollection()
	filter := bson.M{
		"order_id": orderID,
		"status": bson.M{"$ne": CommentStatusDeleted},
	}
	
	err := collection.FindOne(context.Background(), filter).Decode(c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("订单评论不存在")
		}
		return fmt.Errorf("获取订单评论失败: %v", err)
	}
	
	return nil
}

// GetByID 根据ID获取评论
func (c *Comment) GetByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("无效的评论ID: %v", err)
	}
	
	collection := c.GetCollection()
	filter := bson.M{
		"_id": objectID,
		"status": bson.M{"$ne": CommentStatusDeleted},
	}
	
	err = collection.FindOne(context.Background(), filter).Decode(c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("评论不存在")
		}
		return fmt.Errorf("获取评论失败: %v", err)
	}
	
	return nil
}

// Update 更新评论
func (c *Comment) Update() error {
	c.UpdatedAt = time.Now()
	
	collection := c.GetCollection()
	filter := bson.M{"_id": c.ID}
	update := bson.M{"$set": c}
	
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("更新评论失败: %v", err)
	}
	
	return nil
}

// Delete 软删除评论
func (c *Comment) Delete() error {
	c.Status = CommentStatusDeleted
	c.UpdatedAt = time.Now()
	
	collection := c.GetCollection()
	filter := bson.M{"_id": c.ID}
	update := bson.M{"$set": bson.M{
		"status": CommentStatusDeleted,
		"updated_at": c.UpdatedAt,
	}}
	
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("删除评论失败: %v", err)
	}
	
	return nil
}

// GetVehicleComments 获取车辆的所有评论
func GetVehicleComments(vehicleID uint, page, pageSize int32) ([]Comment, int64, error) {
	collection := global.MongoClient.Database("zulme").Collection("comments")
	
	filter := bson.M{
		"vehicle_id": vehicleID,
		"status": CommentStatusNormal,
	}
	
	// 计算总数
	total, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, fmt.Errorf("统计评论数量失败: %v", err)
	}
	
	// 分页查询
	skip := (page - 1) * pageSize
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{"created_at", -1}}) // 按创建时间倒序
	
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("查询评论失败: %v", err)
	}
	defer cursor.Close(context.Background())
	
	var comments []Comment
	if err = cursor.All(context.Background(), &comments); err != nil {
		return nil, 0, fmt.Errorf("解析评论数据失败: %v", err)
	}
	
	return comments, total, nil
}

// GetUserComments 获取用户的所有评论
func GetUserComments(userID uint, page, pageSize int32) ([]Comment, int64, error) {
	collection := global.MongoClient.Database("zulme").Collection("comments")
	
	filter := bson.M{
		"user_id": userID,
		"status": bson.M{"$ne": CommentStatusDeleted},
	}
	
	// 计算总数
	total, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, fmt.Errorf("统计用户评论数量失败: %v", err)
	}
	
	// 分页查询
	skip := (page - 1) * pageSize
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{"created_at", -1}}) // 按创建时间倒序
	
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("查询用户评论失败: %v", err)
	}
	defer cursor.Close(context.Background())
	
	var comments []Comment
	if err = cursor.All(context.Background(), &comments); err != nil {
		return nil, 0, fmt.Errorf("解析用户评论数据失败: %v", err)
	}
	
	return comments, total, nil
}

// GetVehicleStats 获取车辆评论统计
func GetVehicleStats(vehicleID uint) (*CommentStats, error) {
	collection := global.MongoClient.Database("zulme").Collection("comments")
	
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"vehicle_id": vehicleID,
				"status": CommentStatusNormal,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"total_comments": bson.M{"$sum": 1},
				"average_rating": bson.M{"$avg": "$rating"},
				"service_rating": bson.M{"$avg": "$service_rating"},
				"vehicle_rating": bson.M{"$avg": "$vehicle_rating"},
				"clean_rating": bson.M{"$avg": "$clean_rating"},
				"five_star": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$rating", 5}}, 1, 0}}},
				"four_star": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$rating", 4}}, 1, 0}}},
				"three_star": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$rating", 3}}, 1, 0}}},
				"two_star": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$rating", 2}}, 1, 0}}},
				"one_star": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$rating", 1}}, 1, 0}}},
			},
		},
	}
	
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("统计车辆评论失败: %v", err)
	}
	defer cursor.Close(context.Background())
	
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("解析统计数据失败: %v", err)
	}
	
	if len(results) == 0 {
		return &CommentStats{}, nil
	}
	
	result := results[0]
	stats := &CommentStats{
		TotalComments:  getInt64FromBSON(result, "total_comments"),
		AverageRating:  getFloat64FromBSON(result, "average_rating"),
		ServiceRating:  getFloat64FromBSON(result, "service_rating"),
		VehicleRating:  getFloat64FromBSON(result, "vehicle_rating"),
		CleanRating:    getFloat64FromBSON(result, "clean_rating"),
		FiveStarCount:  getInt64FromBSON(result, "five_star"),
		FourStarCount:  getInt64FromBSON(result, "four_star"),
		ThreeStarCount: getInt64FromBSON(result, "three_star"),
		TwoStarCount:   getInt64FromBSON(result, "two_star"),
		OneStarCount:   getInt64FromBSON(result, "one_star"),
	}
	
	return stats, nil
}

// 辅助函数：从BSON结果中安全获取int64值
func getInt64FromBSON(result bson.M, key string) int64 {
	if val, ok := result[key]; ok {
		switch v := val.(type) {
		case int32:
			return int64(v)
		case int64:
			return v
		case int:
			return int64(v)
		}
	}
	return 0
}

// 辅助函数：从BSON结果中安全获取float64值
func getFloat64FromBSON(result bson.M, key string) float64 {
	if val, ok := result[key]; ok {
		switch v := val.(type) {
		case float64:
			return v
		case float32:
			return float64(v)
		case int32:
			return float64(v)
		case int64:
			return float64(v)
		case int:
			return float64(v)
		}
	}
	return 0.0
}

// CheckOrderCommentExists 检查订单是否已经评论过
func CheckOrderCommentExists(orderID uint) (bool, error) {
	collection := global.MongoClient.Database("zulme").Collection("comments")
	
	filter := bson.M{
		"order_id": orderID,
		"status": bson.M{"$ne": CommentStatusDeleted},
	}
	
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, fmt.Errorf("检查订单评论失败: %v", err)
	}
	
	return count > 0, nil
}
