package client

import (
	"context"
	"fmt"
	"log"
	"time"

	chat "chat_srv/proto_chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ChatClient Chat微服务客户端
type ChatClient struct {
	conn   *grpc.ClientConn
	client chat.ChatServiceClient
}

// NewChatClient 创建Chat客户端
func NewChatClient() (*ChatClient, error) {
	// 连接Chat微服务 (端口8003)
	conn, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("连接Chat服务失败: %v", err)
	}

	client := chat.NewChatServiceClient(conn)
	
	return &ChatClient{
		conn:   conn,
		client: client,
	}, nil
}

// Close 关闭连接
func (c *ChatClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetGroupChatInfo 获取群聊信息
func (c *ChatClient) GetGroupChatInfo(groupID string, userID uint) (*chat.GetGroupChatInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &chat.GetGroupChatInfoRequest{
		GroupId: groupID,
		UserId:  uint32(userID),
	}

	resp, err := c.client.GetGroupChatInfo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取群聊信息失败: %v", err)
	}

	return resp, nil
}

// GetGroupMembers 获取群聊成员列表
func (c *ChatClient) GetGroupMembers(groupID string, userID uint) (*chat.GetGroupMembersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &chat.GetGroupMembersRequest{
		GroupId: groupID,
		UserId:  uint32(userID),
	}

	resp, err := c.client.GetGroupMembers(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取群聊成员失败: %v", err)
	}

	return resp, nil
}

// IsFriend 检查好友关系
func (c *ChatClient) IsFriend(userID, friendID uint) (*chat.IsFriendResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &chat.IsFriendRequest{
		UserId:   uint32(userID),
		FriendId: uint32(friendID),
	}

	resp, err := c.client.IsFriend(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("检查好友关系失败: %v", err)
	}

	return resp, nil
}

// ValidateGroupAccess 验证用户是否有权限访问群聊
func (c *ChatClient) ValidateGroupAccess(groupID string, userID uint) (bool, error) {
	// 获取群聊成员列表
	resp, err := c.GetGroupMembers(groupID, userID)
	if err != nil {
		log.Printf("获取群聊成员失败: %v", err)
		return false, err
	}

	if resp.Code != 200 {
		log.Printf("获取群聊成员失败: %s", resp.Message)
		return false, fmt.Errorf("获取群聊成员失败: %s", resp.Message)
	}

	// 检查用户是否在成员列表中
	for _, member := range resp.Members {
		if member.UserId == uint32(userID) {
			return true, nil
		}
	}

	return false, nil
}

// GetGroupMemberIDs 获取群聊成员ID列表
func (c *ChatClient) GetGroupMemberIDs(groupID string, userID uint) ([]uint, error) {
	resp, err := c.GetGroupMembers(groupID, userID)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("获取群聊成员失败: %s", resp.Message)
	}

	var memberIDs []uint
	for _, member := range resp.Members {
		memberIDs = append(memberIDs, uint(member.UserId))
	}

	return memberIDs, nil
}

// 全局Chat客户端实例
var globalChatClient *ChatClient

// InitChatClient 初始化全局Chat客户端
func InitChatClient() error {
	client, err := NewChatClient()
	if err != nil {
		return fmt.Errorf("初始化Chat客户端失败: %v", err)
	}
	
	globalChatClient = client
	log.Println("Chat客户端初始化成功")
	return nil
}

// GetChatClient 获取全局Chat客户端
func GetChatClient() *ChatClient {
	return globalChatClient
}

// CloseChatClient 关闭全局Chat客户端
func CloseChatClient() {
	if globalChatClient != nil {
		globalChatClient.Close()
		globalChatClient = nil
		log.Println("Chat客户端已关闭")
	}
}
