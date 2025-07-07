package manager

import (
	"Common/pkg"
	chat "chat_srv/proto_chat"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"models/model_mongodb"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// WebSocketManager WebSocket连接管理器
type WebSocketManager struct {
	connections map[uint]*websocket.Conn // 用户ID -> WebSocket连接
	mutex       *sync.RWMutex            // 读写锁
	upgrader    websocket.Upgrader       // WebSocket升级器
}

// NewWebSocketManager 创建WebSocket管理器
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[uint]*websocket.Conn),
		mutex:       &sync.RWMutex{},
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许跨域
			},
		},
	}
}

// WebSocket请求结构
type WSRequest struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

// WebSocket响应结构
type WSResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 上线请求数据
type OnlineData struct {
	Message string `json:"message"`
}

// 发送消息请求数据
type SendMessageData struct {
	ToUserID    uint   `json:"to_user_id" binding:"required"`
	MessageType int32  `json:"message_type"`
	Content     string `json:"content" binding:"required"`
}

// 群发消息请求数据
type BroadcastData struct {
	Message string `json:"message"`
}

// 群聊消息请求数据
type GroupMessageData struct {
	GroupID     string `json:"group_id"`     // 群聊ID
	MessageType int32  `json:"message_type"` // 消息类型
	Content     string `json:"content"`      // 消息内容
}

// HandleWebSocket 处理WebSocket连接
func (m *WebSocketManager) HandleWebSocket(c *gin.Context) {
	// 获取用户ID（从JWT中间件获取）
	userID, exists := c.Get("userId")
	if !exists {
		log.Printf("WebSocket连接失败：未找到用户ID")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	uid := userID.(uint)
	log.Printf("用户 %d 尝试建立WebSocket连接", uid)

	// 升级HTTP连接为WebSocket
	conn, err := m.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}
	defer conn.Close()

	// 处理用户上线
	m.handleUserOnline(uid, conn)

	// 消息处理循环
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		var req WSRequest
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("解析消息失败: %v", err)
			m.sendErrorResponse(conn, 1001, "请求参数错误")
			continue
		}

		// 根据命令类型处理消息
		switch req.Cmd {
		case "send_message": // 发送消息命令
			go m.handleSendCommand(conn, req.Data, uid)
		case "send_group_message": // 发送群聊消息命令
			go m.handleSendGroupMessageCommand(conn, req.Data, uid)
		case "broadcast": // 群发消息命令
			go m.handleBroadcastCommand(conn, req.Data, uid)
		case "ping": // 心跳命令
			go m.handlePingCommand(conn)
		default:
			m.sendErrorResponse(conn, 1000, "未知命令")
		}
	}

	// 用户下线处理
	m.handleUserOffline(uid)
}

// handleUserOnline 处理用户上线
func (m *WebSocketManager) handleUserOnline(userID uint, conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 如果用户已经在线，断开之前的连接
	if existingConn, exists := m.connections[userID]; exists {
		log.Printf("用户 %d 已在线，断开之前的连接", userID)
		existingConn.Close()
	}

	// 注册新连接
	m.connections[userID] = conn
	log.Printf("用户 %d 上线成功，当前在线人数: %d", userID, len(m.connections))

	// 发送欢迎消息（自动上线，无需客户端发送online命令）
	welcomeMsg := map[string]interface{}{
		"type":    "welcome",
		"message": "连接成功，自动上线！",
		"user_id": userID,
		"time":    time.Now().Format("2006-01-02 15:04:05"),
	}
	m.sendSuccessResponse(conn, welcomeMsg)
}

// handleUserOffline 处理用户下线
func (m *WebSocketManager) handleUserOffline(userID uint) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.connections[userID]; exists {
		delete(m.connections, userID)
		log.Printf("用户 %d 下线，当前在线人数: %d", userID, len(m.connections))
	}
}

// checkGroupChatExists 检查群聊是否存在并且用户是否是成员
func (m *WebSocketManager) checkGroupChatExists(groupID string, userID uint) bool {
	// 基本格式验证
	if len(groupID) < 5 {
		return false
	}

	// 创建Chat微服务连接
	conn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("连接Chat微服务失败: %v", err)
		return false
	}
	defer conn.Close()

	chatClient := chat.NewChatServiceClient(conn)

	// 调用Chat微服务获取群聊信息（同时验证用户是否是成员）
	req := &chat.GetGroupChatInfoRequest{
		GroupId: groupID,
		UserId:  uint32(userID),
	}

	resp, err := chatClient.GetGroupChatInfo(context.Background(), req)
	if err != nil {
		log.Printf("检查群聊存在性失败: %v", err)
		return false
	}

	return resp.Code == 200
}

// getGroupMembers 获取群聊成员列表
func (m *WebSocketManager) getGroupMembers(groupID string, userID uint) ([]uint32, error) {
	// 创建Chat微服务连接
	conn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("连接Chat微服务失败: %v", err)
		return nil, err
	}
	defer conn.Close()

	chatClient := chat.NewChatServiceClient(conn)

	// 调用Chat微服务获取群聊成员列表
	req := &chat.GetGroupMembersRequest{
		GroupId: groupID,
		UserId:  uint32(userID),
	}

	resp, err := chatClient.GetGroupMembers(context.Background(), req)
	if err != nil {
		log.Printf("获取群聊成员失败: %v", err)
		return nil, err
	}

	if resp.Code != 200 {
		log.Printf("获取群聊成员失败: %s", resp.Message)
		return nil, fmt.Errorf("获取群聊成员失败: %s", resp.Message)
	}

	// 提取成员ID列表
	memberIDs := make([]uint32, len(resp.Members))
	for i, member := range resp.Members {
		memberIDs[i] = member.UserId
	}

	return memberIDs, nil
}

// generatePrivateRoomID 生成一对一聊天房间ID
func (m *WebSocketManager) generatePrivateRoomID(userID1, userID2 uint) string {
	if userID1 < userID2 {
		return fmt.Sprintf("private_%d_%d", userID1, userID2)
	}
	return fmt.Sprintf("private_%d_%d", userID2, userID1)
}

// handleSendCommand 处理发送消息命令
func (m *WebSocketManager) handleSendCommand(conn *websocket.Conn, data interface{}, fromUserID uint) {
	// 解析发送消息数据
	dataBytes, _ := json.Marshal(data)
	var sendData SendMessageData
	if err := json.Unmarshal(dataBytes, &sendData); err != nil {
		m.sendErrorResponse(conn, 10001, "消息解析失败")
		return
	}

	if sendData.ToUserID == 0 || sendData.Content == "" {
		m.sendErrorResponse(conn, 10002, "目标用户ID和消息内容不能为空")
		return
	}

	// 生成一对一聊天房间ID
	roomID := m.generatePrivateRoomID(fromUserID, sendData.ToUserID)

	// 百度敏感词过滤
	isValid, reason, err := pkg.IsTextValid(sendData.Content)
	if err != nil {
		log.Printf("敏感词检测失败: %v", err)
		// API失败时允许消息通过
	} else if !isValid {
		m.sendErrorResponse(conn, 10007, fmt.Sprintf("消息包含敏感内容: %s", reason))
		return
	}

	// 存储聊天消息到MongoDB
	message := &model_mongodb.Message{
		RoomID:      roomID,
		FromUserID:  fromUserID,
		ToUserID:    sendData.ToUserID,
		MessageType: int(sendData.MessageType),
		Content:     sendData.Content,
	}

	if err := message.Create(); err != nil {
		log.Printf("保存消息失败: %v", err)
		m.sendErrorResponse(conn, 10008, "消息保存失败")
		return
	}

	// 查找目标用户连接
	m.mutex.RLock()
	targetConn, exists := m.connections[sendData.ToUserID]
	m.mutex.RUnlock()

	if !exists {
		m.sendErrorResponse(conn, 10003, "目标用户不在线")
		return
	}

	// 构造转发消息
	forwardMsg := map[string]interface{}{
		"type":         "new_message",
		"from_user_id": fromUserID,
		"to_user_id":   sendData.ToUserID,
		"content":      sendData.Content,
		"room_id":      roomID,
		"message_type": sendData.MessageType,
		"message_id":   message.ID.Hex(),
		"time":         time.Now().Format("2006-01-02 15:04:05"),
	}

	// 发送给目标用户
	if err := m.sendSuccessResponse(targetConn, forwardMsg); err != nil {
		log.Printf("转发消息给用户 %d 失败: %v", sendData.ToUserID, err)
		m.sendErrorResponse(conn, 10004, "消息发送失败")
		return
	}

	// 通知发送者发送成功
	m.sendSuccessResponse(conn, "消息发送成功")
}

// handleBroadcastCommand 处理广播消息命令
func (m *WebSocketManager) handleBroadcastCommand(conn *websocket.Conn, data interface{}, fromUserID uint) {
	// 解析广播数据
	dataBytes, _ := json.Marshal(data)
	var broadcastData BroadcastData
	if err := json.Unmarshal(dataBytes, &broadcastData); err != nil {
		m.sendErrorResponse(conn, 10001, "消息解析失败")
		return
	}

	if broadcastData.Message == "" {
		m.sendErrorResponse(conn, 10002, "广播消息不能为空")
		return
	}

	// 百度敏感词过滤
	isValid, reason, err := pkg.IsTextValid(broadcastData.Message)
	if err != nil {
		log.Printf("敏感词检测失败: %v", err)
		// API失败时允许消息通过
	} else if !isValid {
		m.sendErrorResponse(conn, 10007, fmt.Sprintf("消息包含敏感内容: %s", reason))
		return
	}

	// 构造广播消息
	broadcastMsg := map[string]interface{}{
		"type":         "broadcast",
		"from_user_id": fromUserID,
		"message":      broadcastData.Message,
		"time":         time.Now().Format("2006-01-02 15:04:05"),
	}

	// 广播给所有在线用户
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	successCount := 0
	for userID, userConn := range m.connections {
		if err := m.sendSuccessResponse(userConn, broadcastMsg); err != nil {
			log.Printf("广播消息给用户 %d 失败: %v", userID, err)
		} else {
			successCount++
		}
	}

	log.Printf("广播消息成功发送给 %d 个用户", successCount)
}

// handlePingCommand 处理心跳命令
func (m *WebSocketManager) handlePingCommand(conn *websocket.Conn) {
	pongMsg := map[string]interface{}{
		"type": "pong",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}
	m.sendSuccessResponse(conn, pongMsg)
}

// sendSuccessResponse 发送成功响应
func (m *WebSocketManager) sendSuccessResponse(conn *websocket.Conn, data interface{}) error {
	response := WSResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	return m.sendResponse(conn, response)
}

// sendErrorResponse 发送错误响应
func (m *WebSocketManager) sendErrorResponse(conn *websocket.Conn, code int, message string) error {
	response := WSResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	return m.sendResponse(conn, response)
}

// sendResponse 发送响应
func (m *WebSocketManager) sendResponse(conn *websocket.Conn, response WSResponse) error {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("序列化响应失败: %v", err)
	}

	return conn.WriteMessage(websocket.TextMessage, responseBytes)
}

// GetOnlineUsers 获取在线用户列表
func (m *WebSocketManager) GetOnlineUsers() []uint {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	users := make([]uint, 0, len(m.connections))
	for userID := range m.connections {
		users = append(users, userID)
	}
	return users
}

// GetOnlineCount 获取在线用户数量
func (m *WebSocketManager) GetOnlineCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.connections)
}

// SendMessageToUser 向指定用户发送消息
func (m *WebSocketManager) SendMessageToUser(userID uint, message interface{}) error {
	m.mutex.RLock()
	conn, exists := m.connections[userID]
	m.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("用户 %d 不在线", userID)
	}

	return m.sendSuccessResponse(conn, message)
}

// BroadcastMessage 广播消息给所有在线用户
func (m *WebSocketManager) BroadcastMessage(message interface{}) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for userID, conn := range m.connections {
		if err := m.sendSuccessResponse(conn, message); err != nil {
			log.Printf("广播消息给用户 %d 失败: %v", userID, err)
		}
	}
}

// handleSendGroupMessageCommand 处理发送群聊消息命令
func (m *WebSocketManager) handleSendGroupMessageCommand(conn *websocket.Conn, data interface{}, fromUserID uint) {
	// 解析群聊消息数据
	dataBytes, _ := json.Marshal(data)
	var groupMsgData GroupMessageData
	if err := json.Unmarshal(dataBytes, &groupMsgData); err != nil {
		m.sendErrorResponse(conn, 3004, "群聊消息解析失败")
		return
	}

	if groupMsgData.GroupID == "" || groupMsgData.Content == "" {
		m.sendErrorResponse(conn, 3005, "群聊ID和消息内容不能为空")
		return
	}

	// 验证群聊是否存在并且用户是否是成员
	if !m.checkGroupChatExists(groupMsgData.GroupID, fromUserID) {
		m.sendErrorResponse(conn, 3006, "群聊不存在或您不是该群聊成员")
		return
	}

	// 百度敏感词过滤
	isValid, reason, err := pkg.IsTextValid(groupMsgData.Content)
	if err != nil {
		log.Printf("敏感词检测失败: %v", err)
		// API失败时允许消息通过
	} else if !isValid {
		m.sendErrorResponse(conn, 3008, fmt.Sprintf("消息包含敏感内容: %s", reason))
		return
	}

	// 存储群聊消息到MongoDB
	message := &model_mongodb.Message{
		RoomID:      groupMsgData.GroupID, // 群聊使用GroupID作为RoomID
		FromUserID:  fromUserID,
		ToUserID:    0, // 群聊消息ToUserID为0
		MessageType: int(groupMsgData.MessageType),
		Content:     groupMsgData.Content,
	}

	if err := message.Create(); err != nil {
		log.Printf("保存群聊消息失败: %v", err)
		m.sendErrorResponse(conn, 3009, "群聊消息保存失败")
		return
	}

	log.Printf("用户 %d 在群聊 %s 发送消息，消息ID: %s", fromUserID, groupMsgData.GroupID, message.ID.Hex())

	// 构造群聊消息广播
	groupMessage := map[string]interface{}{
		"type":         "group_message",
		"group_id":     groupMsgData.GroupID,
		"from_user_id": fromUserID,
		"content":      groupMsgData.Content,
		"message_type": groupMsgData.MessageType,
		"message_id":   message.ID.Hex(),
		"time":         time.Now().Format("2006-01-02 15:04:05"),
	}

	// 获取群聊成员列表
	memberIDs, err := m.getGroupMembers(groupMsgData.GroupID, fromUserID)
	if err != nil {
		log.Printf("获取群聊成员失败: %v", err)
		return
	}

	// 只向群聊成员发送消息
	m.mutex.RLock()
	sentCount := 0
	for _, memberID := range memberIDs {
		// 不发给发送者自己
		if uint32(fromUserID) == memberID {
			continue
		}

		// 检查成员是否在线
		if userConn, exists := m.connections[uint(memberID)]; exists {
			if err := m.sendSuccessResponse(userConn, groupMessage); err != nil {
				log.Printf("发送群聊消息给用户 %d 失败: %v", memberID, err)
			} else {
				sentCount++
				log.Printf("成功发送群聊消息给群成员 %d", memberID)
			}
		} else {
			log.Printf("群成员 %d 不在线，跳过发送", memberID)
		}
	}
	m.mutex.RUnlock()

	// 发送成功响应给发送者
	successMsg := map[string]interface{}{
		"type":       "group_message_sent",
		"group_id":   groupMsgData.GroupID,
		"message_id": message.ID.Hex(),
		"sent_count": sentCount,
		"message":    "群聊消息已发送",
	}
	m.sendSuccessResponse(conn, successMsg)
}
