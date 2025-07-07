package request

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	RoomID      string `json:"room_id" form:"room_id" binding:"required"`
	ToUserID    uint   `json:"to_user_id" form:"to_user_id" binding:"required"`
	MessageType int32  `json:"message_type" form:"message_type"`
	Content     string `json:"content" form:"content" binding:"required"`
}

// GetRoomMessagesRequest 获取房间消息请求
type GetRoomMessagesRequest struct {
	RoomID   string `json:"room_id" form:"room_id" uri:"room_id" binding:"required"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// AddFriendRequest 添加好友请求
type AddFriendRequest struct {
	FriendID uint `json:"friend_id" form:"friend_id" binding:"required"`
}

// CreateGroupChatRequest 创建群聊请求
type CreateGroupChatRequest struct {
	GroupName   string `json:"group_name" form:"group_name" binding:"required"`
	Description string `json:"description" form:"description"`
	MemberIDs   []uint `json:"member_ids" form:"member_ids" binding:"required"`
}

// GetGroupChatInfoRequest 获取群聊信息请求
type GetGroupChatInfoRequest struct {
	GroupID string `json:"group_id" form:"group_id" uri:"group_id" binding:"required"`
}

// GetGroupMembersRequest 获取群聊成员请求
type GetGroupMembersRequest struct {
	GroupID string `json:"group_id" form:"group_id" uri:"group_id" binding:"required"`
}

// InviteToGroupURIRequest 邀请加入群聊URI请求
type InviteToGroupURIRequest struct {
	GroupID string `json:"group_id" form:"group_id" uri:"group_id" binding:"required"`
}

// InviteToGroupRequest 邀请加入群聊请求
type InviteToGroupRequest struct {
	UserIDs []uint `json:"user_ids" form:"user_ids" binding:"required"`
}

// LeaveGroupRequest 退出群聊请求
type LeaveGroupRequest struct {
	GroupID string `json:"group_id" form:"group_id" uri:"group_id" binding:"required"`
}

// IsFriendRequest 检查好友关系请求
type IsFriendRequest struct {
	FriendID uint `json:"friend_id" form:"friend_id" uri:"friend_id" binding:"required"`
}

// GetUserGroupsRequest 获取用户群聊列表请求（无需额外参数，用户ID从JWT获取）
type GetUserGroupsRequest struct {
}
