package handler

import (
	"Api/client"
	chat "chat_srv/proto_chat"
	"context"
)

// SendMessage 发送消息
func SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		message, err := in.SendMessage(ctx, req)
		if err != nil {
			return nil, err
		}
		return message, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.SendMessageResponse), nil
}

// GetRoomMessages 获取房间消息
func GetRoomMessages(ctx context.Context, req *chat.GetRoomMessagesRequest) (*chat.GetRoomMessagesResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		messages, err := in.GetRoomMessages(ctx, req)
		if err != nil {
			return nil, err
		}
		return messages, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.GetRoomMessagesResponse), nil
}

// AddFriend 添加好友
func AddFriend(ctx context.Context, req *chat.AddFriendRequest) (*chat.AddFriendResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		friend, err := in.AddFriend(ctx, req)
		if err != nil {
			return nil, err
		}
		return friend, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.AddFriendResponse), nil
}

// GetFriendList 获取好友列表
func GetFriendList(ctx context.Context, req *chat.GetFriendListRequest) (*chat.GetFriendListResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		friends, err := in.GetFriendList(ctx, req)
		if err != nil {
			return nil, err
		}
		return friends, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.GetFriendListResponse), nil
}

// CreateGroupChat 创建群聊
func CreateGroupChat(ctx context.Context, req *chat.CreateGroupChatRequest) (*chat.CreateGroupChatResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		group, err := in.CreateGroupChat(ctx, req)
		if err != nil {
			return nil, err
		}
		return group, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.CreateGroupChatResponse), nil
}

// GetGroupChatInfo 获取群聊信息
func GetGroupChatInfo(ctx context.Context, req *chat.GetGroupChatInfoRequest) (*chat.GetGroupChatInfoResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		group, err := in.GetGroupChatInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		return group, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.GetGroupChatInfoResponse), nil
}

// GetGroupMembers 获取群聊成员
func GetGroupMembers(ctx context.Context, req *chat.GetGroupMembersRequest) (*chat.GetGroupMembersResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		members, err := in.GetGroupMembers(ctx, req)
		if err != nil {
			return nil, err
		}
		return members, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.GetGroupMembersResponse), nil
}

// InviteToGroup 邀请用户加入群聊
func InviteToGroup(ctx context.Context, req *chat.InviteToGroupRequest) (*chat.InviteToGroupResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		result, err := in.InviteToGroup(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.InviteToGroupResponse), nil
}

// LeaveGroup 退出群聊
func LeaveGroup(ctx context.Context, req *chat.LeaveGroupRequest) (*chat.LeaveGroupResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		result, err := in.LeaveGroup(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.LeaveGroupResponse), nil
}

// IsFriend 检查好友关系
func IsFriend(ctx context.Context, req *chat.IsFriendRequest) (*chat.IsFriendResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		result, err := in.IsFriend(ctx, req)
		if err != nil {
			return nil, err
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.IsFriendResponse), nil
}

// GetUserGroups 获取用户参与的群聊列表
func GetUserGroups(ctx context.Context, req *chat.GetUserGroupsRequest) (*chat.GetUserGroupsResponse, error) {
	chatClient, err := client.ChatClient(ctx, func(ctx context.Context, in chat.ChatServiceClient) (interface{}, error) {
		groups, err := in.GetUserGroups(ctx, req)
		if err != nil {
			return nil, err
		}
		return groups, nil
	})
	if err != nil {
		return nil, err
	}
	return chatClient.(*chat.GetUserGroupsResponse), nil
}
