package logic

import (
	"Common/global"
	"Common/pkg"
	"fmt"
	"models/model_mongodb"
	"models/model_mysql"
)

// ChatLogic 聊天业务逻辑
type ChatLogic struct{}

// NewChatLogic 创建聊天逻辑实例
func NewChatLogic() *ChatLogic {
	return &ChatLogic{}
}

// SendMessage 发送消息
func (cl *ChatLogic) SendMessage(roomID string, fromUserID, toUserID uint, messageType int32, content string) (*model_mongodb.Message, error) {
	fmt.Printf("开始发送消息: roomID=%s, fromUserID=%d, toUserID=%d\n", roomID, fromUserID, toUserID)

	// 1. 百度敏感词过滤审核
	fmt.Printf("开始进行消息内容审核: %s\n", content)
	isValid, reason, err := pkg.IsTextValid(content)
	if err != nil {
		fmt.Printf("内容审核API调用失败: %v\n", err)
		// 如果审核API失败，可以选择允许消息通过或者拒绝
		// 这里选择记录错误但允许消息通过，避免因为API问题影响用户体验
		fmt.Printf("警告：内容审核失败，但允许消息通过\n")
	} else if !isValid {
		fmt.Printf("消息内容审核不通过: %s\n", reason)
		return nil, fmt.Errorf("消息内容包含敏感信息，请修改后重新发送")
	} else {
		fmt.Printf("消息内容审核通过: %s\n", reason)
	}

	// 2. 检查好友关系（私聊时需要）
	if toUserID != 0 {
		isFriend, err := model_mysql.IsFriend(fromUserID, toUserID)
		if err != nil {
			fmt.Printf("检查好友关系失败: %v\n", err)
			return nil, fmt.Errorf("检查好友关系失败: %v", err)
		}
		if !isFriend {
			return nil, fmt.Errorf("对方不是您的好友，无法发送消息")
		}
	}

	// 3. 创建消息
	message := &model_mongodb.Message{
		RoomID:      roomID,
		FromUserID:  fromUserID,
		ToUserID:    toUserID,
		MessageType: int(messageType),
		Content:     content,
	}

	if err := message.Create(); err != nil {
		fmt.Printf("保存消息到MongoDB失败: %v\n", err)
		return nil, fmt.Errorf("发送消息失败: %v", err)
	}

	fmt.Printf("消息发送成功，ID: %s\n", message.ID.Hex())
	return message, nil
}

// GetRoomMessages 获取房间消息
func (cl *ChatLogic) GetRoomMessages(roomID string, page, pageSize int32) ([]model_mongodb.Message, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	return model_mongodb.GetRoomMessages(roomID, page, pageSize)
}

// AddFriend 添加好友
func (cl *ChatLogic) AddFriend(userID, friendID uint) error {
	if userID == friendID {
		return fmt.Errorf("不能添加自己为好友")
	}

	// 检查是否已经是好友
	isFriend, err := model_mysql.IsFriend(userID, friendID)
	if err != nil {
		return fmt.Errorf("检查好友关系失败: %v", err)
	}
	if isFriend {
		return fmt.Errorf("已经是好友关系")
	}

	// 检查目标用户是否存在
	var targetUser model_mysql.User
	if err := global.DB.Where("id = ?", friendID).First(&targetUser).Error; err != nil {
		return fmt.Errorf("目标用户不存在")
	}

	// 添加双向好友关系
	friend1 := &model_mysql.UserFriend{
		UserID:   userID,
		FriendID: friendID,
		Status:   model_mysql.FriendStatusNormal,
	}
	friend2 := &model_mysql.UserFriend{
		UserID:   friendID,
		FriendID: userID,
		Status:   model_mysql.FriendStatusNormal,
	}

	if err := friend1.AddFriend(); err != nil {
		return fmt.Errorf("添加好友失败: %v", err)
	}
	if err := friend2.AddFriend(); err != nil {
		return fmt.Errorf("添加好友失败: %v", err)
	}

	return nil
}

// GetFriendList 获取好友列表
func (cl *ChatLogic) GetFriendList(userID uint) ([]model_mysql.UserFriend, error) {
	return model_mysql.GetUserFriends(userID)
}

// IsFriend 检查好友关系
func (cl *ChatLogic) IsFriend(userID, friendID uint) (bool, error) {
	return model_mysql.IsFriend(userID, friendID)
}

// DeleteFriend 删除好友
func (cl *ChatLogic) DeleteFriend(userID, friendID uint) error {
	return model_mysql.DeleteFriend(userID, friendID)
}

// GetMessageStats 获取消息统计
func (cl *ChatLogic) GetMessageStats(roomID string) (map[string]interface{}, error) {
	return model_mongodb.GetMessageStats(roomID)
}

// CreateGroupChat 创建群聊
func (cl *ChatLogic) CreateGroupChat(creatorID uint, groupName, description string, memberIDs []uint) (*model_mysql.GroupChat, error) {
	// 验证群聊名称
	if groupName == "" {
		return nil, fmt.Errorf("群聊名称不能为空")
	}

	// 验证成员列表
	if len(memberIDs) == 0 {
		return nil, fmt.Errorf("群成员列表不能为空")
	}

	// 验证所有成员都是创建者的好友
	for _, memberID := range memberIDs {
		if memberID == creatorID {
			continue // 跳过创建者自己
		}

		isFriend, err := model_mysql.IsFriend(creatorID, memberID)
		if err != nil {
			return nil, fmt.Errorf("检查好友关系失败: %v", err)
		}
		if !isFriend {
			return nil, fmt.Errorf("用户 %d 不是您的好友，无法邀请加入群聊", memberID)
		}
	}

	// 创建群聊
	return model_mysql.CreateGroupChatWithMembers(creatorID, groupName, description, memberIDs)
}

// GetGroupChatInfo 获取群聊信息
func (cl *ChatLogic) GetGroupChatInfo(groupID string, userID uint) (*model_mysql.GroupChat, error) {
	// 检查用户是否是群聊成员
	isMember, err := model_mysql.IsGroupMember(groupID, userID)
	if err != nil {
		return nil, fmt.Errorf("检查群聊成员失败: %v", err)
	}
	if !isMember {
		return nil, fmt.Errorf("您不是该群聊的成员")
	}

	// 获取群聊信息
	var group model_mysql.GroupChat
	if err := group.GetGroupChatByID(groupID); err != nil {
		return nil, fmt.Errorf("获取群聊信息失败: %v", err)
	}

	return &group, nil
}

// GetGroupMembers 获取群聊成员列表
func (cl *ChatLogic) GetGroupMembers(groupID string, userID uint) ([]model_mysql.GroupMember, error) {
	// 检查用户是否是群聊成员
	isMember, err := model_mysql.IsGroupMember(groupID, userID)
	if err != nil {
		return nil, fmt.Errorf("检查群聊成员失败: %v", err)
	}
	if !isMember {
		return nil, fmt.Errorf("您不是该群聊的成员")
	}

	// 获取群聊成员列表
	return model_mysql.GetGroupMembers(groupID)
}

// InviteToGroup 邀请用户加入群聊
func (cl *ChatLogic) InviteToGroup(groupID string, inviterID uint, userIDs []uint) ([]uint, []uint, error) {
	// 检查邀请者是否是群聊成员
	isMember, err := model_mysql.IsGroupMember(groupID, inviterID)
	if err != nil {
		return nil, nil, fmt.Errorf("检查群聊成员失败: %v", err)
	}
	if !isMember {
		return nil, nil, fmt.Errorf("您不是该群聊的成员，无法邀请他人")
	}

	var successUserIDs []uint
	var failedUserIDs []uint

	for _, userID := range userIDs {
		// 检查是否已经是群聊成员
		isAlreadyMember, err := model_mysql.IsGroupMember(groupID, userID)
		if err != nil {
			failedUserIDs = append(failedUserIDs, userID)
			continue
		}
		if isAlreadyMember {
			failedUserIDs = append(failedUserIDs, userID)
			continue
		}

		// 检查是否是好友
		isFriend, err := model_mysql.IsFriend(inviterID, userID)
		if err != nil || !isFriend {
			failedUserIDs = append(failedUserIDs, userID)
			continue
		}

		// 添加到群聊
		member := &model_mysql.GroupMember{
			GroupID: groupID,
			UserID:  userID,
			Role:    model_mysql.GroupRoleMember,
		}
		if err := member.AddGroupMember(); err != nil {
			failedUserIDs = append(failedUserIDs, userID)
			continue
		}

		successUserIDs = append(successUserIDs, userID)
	}

	// 更新群聊成员数量
	var group model_mysql.GroupChat
	if err := group.GetGroupChatByID(groupID); err == nil {
		group.UpdateMemberCount()
	}

	return successUserIDs, failedUserIDs, nil
}

// LeaveGroup 退出群聊
func (cl *ChatLogic) LeaveGroup(groupID string, userID uint) error {
	// 检查用户是否是群聊成员
	isMember, err := model_mysql.IsGroupMember(groupID, userID)
	if err != nil {
		return fmt.Errorf("检查群聊成员失败: %v", err)
	}
	if !isMember {
		return fmt.Errorf("您不是该群聊的成员")
	}

	// 获取群聊信息
	var group model_mysql.GroupChat
	if err := group.GetGroupChatByID(groupID); err != nil {
		return fmt.Errorf("获取群聊信息失败: %v", err)
	}

	// 检查是否是群主
	if group.CreatorID == userID {
		return fmt.Errorf("群主不能退出群聊，请先转让群主身份或解散群聊")
	}

	// 移除群聊成员
	if err := model_mysql.RemoveGroupMember(groupID, userID); err != nil {
		return fmt.Errorf("退出群聊失败: %v", err)
	}

	// 更新群聊成员数量
	group.UpdateMemberCount()

	return nil
}

// GetUserGroups 获取用户参与的群聊列表
func (cl *ChatLogic) GetUserGroups(userID uint) ([]model_mysql.GroupChat, error) {
	return model_mysql.GetUserGroups(userID)
}
