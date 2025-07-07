package server

import (
	"chat_srv/internal/logic"
	"chat_srv/proto_chat"
	"context"
	"models/model_mongodb"
	"models/model_mysql"
	"time"
)

// ChatServer 聊天服务器
type ChatServer struct {
	proto_chat.UnimplementedChatServiceServer
	chatLogic *logic.ChatLogic
}

// NewChatServer 创建聊天服务器
func NewChatServer() *ChatServer {
	return &ChatServer{
		chatLogic: logic.NewChatLogic(),
	}
}

// SendMessage 发送消息
func (s *ChatServer) SendMessage(ctx context.Context, req *proto_chat.SendMessageRequest) (*proto_chat.SendMessageResponse, error) {
	// 参数验证
	if req.RoomId == "" || req.FromUserId == 0 || req.Content == "" {
		return &proto_chat.SendMessageResponse{
			Code:    400,
			Message: "房间ID、发送者ID和消息内容不能为空",
		}, nil
	}

	// 发送消息
	message, err := s.chatLogic.SendMessage(
		req.RoomId,
		uint(req.FromUserId),
		uint(req.ToUserId),
		req.MessageType,
		req.Content,
	)

	if err != nil {
		return &proto_chat.SendMessageResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	messageInfo := s.convertToMessageInfo(message)

	return &proto_chat.SendMessageResponse{
		Code:        200,
		Message:     "消息发送成功",
		MessageInfo: messageInfo,
	}, nil
}

// GetRoomMessages 获取房间消息
func (s *ChatServer) GetRoomMessages(ctx context.Context, req *proto_chat.GetRoomMessagesRequest) (*proto_chat.GetRoomMessagesResponse, error) {
	if req.RoomId == "" {
		return &proto_chat.GetRoomMessagesResponse{
			Code:    400,
			Message: "房间ID不能为空",
		}, nil
	}

	messages, total, err := s.chatLogic.GetRoomMessages(req.RoomId, req.Page, req.PageSize)
	if err != nil {
		return &proto_chat.GetRoomMessagesResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	var messageInfos []*proto_chat.MessageInfo
	for _, message := range messages {
		messageInfos = append(messageInfos, s.convertToMessageInfo(&message))
	}

	return &proto_chat.GetRoomMessagesResponse{
		Code:     200,
		Message:  "获取消息成功",
		Messages: messageInfos,
		Total:    total,
	}, nil
}

// AddFriend 添加好友
func (s *ChatServer) AddFriend(ctx context.Context, req *proto_chat.AddFriendRequest) (*proto_chat.AddFriendResponse, error) {
	if req.UserId == 0 || req.FriendId == 0 {
		return &proto_chat.AddFriendResponse{
			Code:    400,
			Message: "用户ID和好友ID不能为空",
		}, nil
	}

	err := s.chatLogic.AddFriend(uint(req.UserId), uint(req.FriendId))
	if err != nil {
		return &proto_chat.AddFriendResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_chat.AddFriendResponse{
		Code:    200,
		Message: "添加好友成功",
	}, nil
}

// GetFriendList 获取好友列表
func (s *ChatServer) GetFriendList(ctx context.Context, req *proto_chat.GetFriendListRequest) (*proto_chat.GetFriendListResponse, error) {
	if req.UserId == 0 {
		return &proto_chat.GetFriendListResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	friends, err := s.chatLogic.GetFriendList(uint(req.UserId))
	if err != nil {
		return &proto_chat.GetFriendListResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	var friendInfos []*proto_chat.FriendInfo
	for _, friend := range friends {
		friendInfos = append(friendInfos, s.convertToFriendInfo(&friend))
	}

	return &proto_chat.GetFriendListResponse{
		Code:    200,
		Message: "获取好友列表成功",
		Friends: friendInfos,
	}, nil
}

// IsFriend 检查好友关系
func (s *ChatServer) IsFriend(ctx context.Context, req *proto_chat.IsFriendRequest) (*proto_chat.IsFriendResponse, error) {
	if req.UserId == 0 || req.FriendId == 0 {
		return &proto_chat.IsFriendResponse{
			Code:    400,
			Message: "用户ID和好友ID不能为空",
		}, nil
	}

	isFriend, err := s.chatLogic.IsFriend(uint(req.UserId), uint(req.FriendId))
	if err != nil {
		return &proto_chat.IsFriendResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_chat.IsFriendResponse{
		Code:     200,
		Message:  "检查好友关系成功",
		IsFriend: isFriend,
	}, nil
}

// 辅助方法：转换消息信息
func (s *ChatServer) convertToMessageInfo(message *model_mongodb.Message) *proto_chat.MessageInfo {
	return &proto_chat.MessageInfo{
		Id:          message.ID.Hex(),
		RoomId:      message.RoomID,
		FromUserId:  uint32(message.FromUserID),
		ToUserId:    uint32(message.ToUserID),
		MessageType: int32(message.MessageType),
		Content:     message.Content,
		CreatedAt:   message.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   message.UpdatedAt.Format(time.RFC3339),
	}
}

// 辅助方法：转换好友信息
func (s *ChatServer) convertToFriendInfo(friend *model_mysql.UserFriend) *proto_chat.FriendInfo {
	return &proto_chat.FriendInfo{
		UserId:    uint32(friend.UserID),
		FriendId:  uint32(friend.FriendID),
		Status:    int32(friend.Status),
		CreatedAt: friend.CreatedAt.Format(time.RFC3339),
	}
}

// CreateGroupChat 创建群聊
func (s *ChatServer) CreateGroupChat(ctx context.Context, req *proto_chat.CreateGroupChatRequest) (*proto_chat.CreateGroupChatResponse, error) {
	if req.CreatorId == 0 {
		return &proto_chat.CreateGroupChatResponse{
			Code:    400,
			Message: "创建者ID不能为空",
		}, nil
	}

	if req.GroupName == "" {
		return &proto_chat.CreateGroupChatResponse{
			Code:    400,
			Message: "群聊名称不能为空",
		}, nil
	}

	if len(req.MemberIds) == 0 {
		return &proto_chat.CreateGroupChatResponse{
			Code:    400,
			Message: "群成员列表不能为空",
		}, nil
	}

	// 转换成员ID列表
	memberIDs := make([]uint, len(req.MemberIds))
	for i, id := range req.MemberIds {
		memberIDs[i] = uint(id)
	}

	// 创建群聊
	group, err := s.chatLogic.CreateGroupChat(uint(req.CreatorId), req.GroupName, req.Description, memberIDs)
	if err != nil {
		return &proto_chat.CreateGroupChatResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	groupInfo := s.convertToGroupChatInfo(group)

	return &proto_chat.CreateGroupChatResponse{
		Code:      200,
		Message:   "创建群聊成功",
		GroupInfo: groupInfo,
	}, nil
}

// GetGroupChatInfo 获取群聊信息
func (s *ChatServer) GetGroupChatInfo(ctx context.Context, req *proto_chat.GetGroupChatInfoRequest) (*proto_chat.GetGroupChatInfoResponse, error) {
	if req.GroupId == "" {
		return &proto_chat.GetGroupChatInfoResponse{
			Code:    400,
			Message: "群聊ID不能为空",
		}, nil
	}

	if req.UserId == 0 {
		return &proto_chat.GetGroupChatInfoResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	group, err := s.chatLogic.GetGroupChatInfo(req.GroupId, uint(req.UserId))
	if err != nil {
		return &proto_chat.GetGroupChatInfoResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	groupInfo := s.convertToGroupChatInfo(group)

	return &proto_chat.GetGroupChatInfoResponse{
		Code:      200,
		Message:   "获取群聊信息成功",
		GroupInfo: groupInfo,
	}, nil
}

// GetGroupMembers 获取群聊成员列表
func (s *ChatServer) GetGroupMembers(ctx context.Context, req *proto_chat.GetGroupMembersRequest) (*proto_chat.GetGroupMembersResponse, error) {
	if req.GroupId == "" {
		return &proto_chat.GetGroupMembersResponse{
			Code:    400,
			Message: "群聊ID不能为空",
		}, nil
	}

	if req.UserId == 0 {
		return &proto_chat.GetGroupMembersResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	members, err := s.chatLogic.GetGroupMembers(req.GroupId, uint(req.UserId))
	if err != nil {
		return &proto_chat.GetGroupMembersResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	memberInfos := make([]*proto_chat.GroupMemberInfo, len(members))
	for i, member := range members {
		memberInfos[i] = s.convertToGroupMemberInfo(&member)
	}

	return &proto_chat.GetGroupMembersResponse{
		Code:    200,
		Message: "获取群聊成员成功",
		Members: memberInfos,
	}, nil
}

// InviteToGroup 邀请用户加入群聊
func (s *ChatServer) InviteToGroup(ctx context.Context, req *proto_chat.InviteToGroupRequest) (*proto_chat.InviteToGroupResponse, error) {
	if req.GroupId == "" {
		return &proto_chat.InviteToGroupResponse{
			Code:    400,
			Message: "群聊ID不能为空",
		}, nil
	}

	if req.InviterId == 0 {
		return &proto_chat.InviteToGroupResponse{
			Code:    400,
			Message: "邀请者ID不能为空",
		}, nil
	}

	if len(req.UserIds) == 0 {
		return &proto_chat.InviteToGroupResponse{
			Code:    400,
			Message: "被邀请用户列表不能为空",
		}, nil
	}

	// 转换用户ID列表
	userIDs := make([]uint, len(req.UserIds))
	for i, id := range req.UserIds {
		userIDs[i] = uint(id)
	}

	successUserIDs, failedUserIDs, err := s.chatLogic.InviteToGroup(req.GroupId, uint(req.InviterId), userIDs)
	if err != nil {
		return &proto_chat.InviteToGroupResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	successIDs := make([]uint32, len(successUserIDs))
	for i, id := range successUserIDs {
		successIDs[i] = uint32(id)
	}

	failedIDs := make([]uint32, len(failedUserIDs))
	for i, id := range failedUserIDs {
		failedIDs[i] = uint32(id)
	}

	return &proto_chat.InviteToGroupResponse{
		Code:           200,
		Message:        "邀请处理完成",
		SuccessUserIds: successIDs,
		FailedUserIds:  failedIDs,
	}, nil
}

// LeaveGroup 退出群聊
func (s *ChatServer) LeaveGroup(ctx context.Context, req *proto_chat.LeaveGroupRequest) (*proto_chat.LeaveGroupResponse, error) {
	if req.GroupId == "" {
		return &proto_chat.LeaveGroupResponse{
			Code:    400,
			Message: "群聊ID不能为空",
		}, nil
	}

	if req.UserId == 0 {
		return &proto_chat.LeaveGroupResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	err := s.chatLogic.LeaveGroup(req.GroupId, uint(req.UserId))
	if err != nil {
		return &proto_chat.LeaveGroupResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &proto_chat.LeaveGroupResponse{
		Code:    200,
		Message: "退出群聊成功",
	}, nil
}

// GetUserGroups 获取用户参与的群聊列表
func (s *ChatServer) GetUserGroups(ctx context.Context, req *proto_chat.GetUserGroupsRequest) (*proto_chat.GetUserGroupsResponse, error) {
	if req.UserId == 0 {
		return &proto_chat.GetUserGroupsResponse{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	groups, err := s.chatLogic.GetUserGroups(uint(req.UserId))
	if err != nil {
		return &proto_chat.GetUserGroupsResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	// 转换为响应格式
	groupInfos := make([]*proto_chat.GroupChatInfo, len(groups))
	for i, group := range groups {
		groupInfos[i] = &proto_chat.GroupChatInfo{
			GroupId:     group.GroupID,
			GroupName:   group.GroupName,
			Description: group.Description,
			CreatorId:   uint32(group.CreatorID),
			MemberCount: int32(group.MemberCount),
			CreatedAt:   group.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   group.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &proto_chat.GetUserGroupsResponse{
		Code:    200,
		Message: "获取群聊列表成功",
		Groups:  groupInfos,
	}, nil
}

// 辅助方法：转换群聊信息
func (s *ChatServer) convertToGroupChatInfo(group *model_mysql.GroupChat) *proto_chat.GroupChatInfo {
	return &proto_chat.GroupChatInfo{
		GroupId:     group.GroupID,
		GroupName:   group.GroupName,
		CreatorId:   uint32(group.CreatorID),
		Description: group.Description,
		MemberCount: int32(group.MemberCount),
		CreatedAt:   group.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   group.UpdatedAt.Format(time.RFC3339),
	}
}

// 辅助方法：转换群聊成员信息
func (s *ChatServer) convertToGroupMemberInfo(member *model_mysql.GroupMember) *proto_chat.GroupMemberInfo {
	return &proto_chat.GroupMemberInfo{
		GroupId:  member.GroupID,
		UserId:   uint32(member.UserID),
		Role:     int32(member.Role),
		JoinedAt: member.JoinedAt.Format(time.RFC3339),
	}
}
