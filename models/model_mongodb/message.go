package model_mongodb

import (
	"Common/global"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Message 消息模型
type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RoomID      string             `bson:"room_id" json:"room_id"`
	FromUserID  uint               `bson:"from_user_id" json:"from_user_id"`
	ToUserID    uint               `bson:"to_user_id" json:"to_user_id"`
	MessageType int                `bson:"message_type" json:"message_type"` // 消息类型 1: 文本消息 2: 图片消息 3: 文件消息
	Content     string             `bson:"content" json:"content"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

const (
	MessageTypeText  = 1 // 文本消息
	MessageTypeImage = 2 // 图片消息
	MessageTypeFile  = 3 // 文件消息
)

// GetCollection 获取消息集合
func (m *Message) GetCollection() *mongo.Collection {
	return global.MongoClient.Database("zulme").Collection("messages")
}

// Create 创建消息
func (m *Message) Create() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	collection := m.GetCollection()
	result, err := collection.InsertOne(context.Background(), m)
	if err != nil {
		return fmt.Errorf("创建消息失败: %v", err)
	}

	m.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetByID 根据ID获取消息
func (m *Message) GetByID(messageID string) error {
	objectID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return fmt.Errorf("无效的消息ID: %v", err)
	}

	collection := m.GetCollection()
	return collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(m)
}

// GetRoomMessages 获取房间消息列表
func GetRoomMessages(roomID string, page, pageSize int32) ([]Message, int64, error) {
	collection := global.MongoClient.Database("zulme").Collection("messages")

	// 计算跳过的文档数
	skip := (page - 1) * pageSize

	// 查询条件
	filter := bson.M{"room_id": roomID}

	// 获取总数
	total, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, fmt.Errorf("获取消息总数失败: %v", err)
	}

	// 查询选项
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.M{"created_at": -1}) // 按创建时间倒序

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("查询消息失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var messages []Message
	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, 0, fmt.Errorf("解析消息失败: %v", err)
	}

	return messages, total, nil
}

// GetUserChatRooms 获取用户参与的聊天室及最新消息
func GetUserChatRooms(userID uint) ([]map[string]interface{}, error) {
	collection := global.MongoClient.Database("zulme").Collection("messages")

	// 聚合管道
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$or": []bson.M{
					{"from_user_id": userID},
					{"to_user_id": userID},
				},
			},
		},
		{
			"$sort": bson.M{"created_at": -1},
		},
		{
			"$group": bson.M{
				"_id":            "$room_id",
				"latest_message": bson.M{"$first": "$$ROOT"},
			},
		},
		{
			"$sort": bson.M{"latest_message.created_at": -1},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("查询用户聊天室失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var results []map[string]interface{}
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("解析聊天室数据失败: %v", err)
	}

	return results, nil
}

// DeleteMessage 删除消息
func (m *Message) Delete() error {
	collection := m.GetCollection()

	filter := bson.M{"_id": m.ID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("删除消息失败: %v", err)
	}

	return nil
}

// GetMessageStats 获取消息统计
func GetMessageStats(roomID string) (map[string]interface{}, error) {
	collection := global.MongoClient.Database("zulme").Collection("messages")

	pipeline := []bson.M{
		{
			"$match": bson.M{"room_id": roomID},
		},
		{
			"$group": bson.M{
				"_id":                 nil,
				"total_messages":      bson.M{"$sum": 1},
				"latest_message_time": bson.M{"$max": "$created_at"},
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("获取消息统计失败: %v", err)
	}
	defer cursor.Close(context.Background())

	var result map[string]interface{}
	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("解析统计数据失败: %v", err)
		}
	}

	return result, nil
}
