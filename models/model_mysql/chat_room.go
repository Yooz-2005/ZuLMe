package model_mysql

import (
	"Common/global"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// GroupChat 群聊模型
type GroupChat struct {
	gorm.Model
	GroupID     string `gorm:"type:varchar(50);uniqueIndex;not null" json:"group_id"` //
	GroupName   string `gorm:"type:varchar(100);not null" json:"group_name"`          // 群聊名称
	CreatorID   uint   `gorm:"not null;index" json:"creator_id"`                      // 创建者ID
	Description string `gorm:"type:text" json:"description"`                          // 群聊描述
	MemberCount int    `gorm:"default:0" json:"member_count"`                         // 成员数量
}

// GroupMember 群聊成员模型
type GroupMember struct {
	gorm.Model
	GroupID  string    `gorm:"type:varchar(50);not null;index:idx_group_user" json:"group_id"`
	UserID   uint      `gorm:"not null;index:idx_group_user" json:"user_id"`
	Role     int       `gorm:"type:tinyint;default:3;comment:'1:群主 2:管理员 3:普通成员'" json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

// UserFriend 好友关系模型
type UserFriend struct {
	gorm.Model
	UserID   uint `gorm:"not null;uniqueIndex:uk_user_friend" json:"user_id"`
	FriendID uint `gorm:"not null;uniqueIndex:uk_user_friend" json:"friend_id"`
	Status   int  `gorm:"type:tinyint;default:1;comment:'1:正常 2:已删除'" json:"status"`
}

const (
	FriendStatusNormal  = 1 // 正常
	FriendStatusDeleted = 2 // 已删除
)

const (
	GroupRoleOwner  = 1 // 群主
	GroupRoleAdmin  = 2 // 管理员
	GroupRoleMember = 3 // 普通成员
)

// AddFriend 添加好友
func (uf *UserFriend) AddFriend() error {
	return global.DB.Create(uf).Error
}

// IsFriend 检查是否是好友关系
func IsFriend(userID, friendID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&UserFriend{}).Where(
		"((user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)) AND status = ?",
		userID, friendID, friendID, userID, FriendStatusNormal,
	).Count(&count).Error

	return count > 0, err
}

// GetUserFriends 获取用户好友列表
func GetUserFriends(userID uint) ([]UserFriend, error) {
	var friends []UserFriend
	err := global.DB.Where("user_id = ? AND status = ?", userID, FriendStatusNormal).Find(&friends).Error
	return friends, err
}

// DeleteFriend 删除好友关系
func DeleteFriend(userID, friendID uint) error {
	return global.DB.Model(&UserFriend{}).Where(
		"((user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)) AND status = ?",
		userID, friendID, friendID, userID, FriendStatusNormal,
	).Update("status", FriendStatusDeleted).Error
}

// CreateGroupChat 创建群聊
func (gc *GroupChat) CreateGroupChat() error {
	return global.DB.Create(gc).Error
}

// GetGroupChatByID 根据群聊ID获取群聊信息
func (gc *GroupChat) GetGroupChatByID(groupID string) error {
	return global.DB.Where("group_id = ?", groupID).First(gc).Error
}

// UpdateMemberCount 更新群聊成员数量
func (gc *GroupChat) UpdateMemberCount() error {
	var count int64
	err := global.DB.Model(&GroupMember{}).Where("group_id = ?", gc.GroupID).Count(&count).Error
	if err != nil {
		return err
	}

	gc.MemberCount = int(count)
	return global.DB.Model(gc).Update("member_count", gc.MemberCount).Error
}

// AddGroupMember 添加群聊成员
func (gm *GroupMember) AddGroupMember() error {
	gm.JoinedAt = time.Now()
	return global.DB.Create(gm).Error
}

// GetGroupMembers 获取群聊成员列表
func GetGroupMembers(groupID string) ([]GroupMember, error) {
	var members []GroupMember
	err := global.DB.Where("group_id = ?", groupID).Find(&members).Error
	return members, err
}

// IsGroupMember 检查用户是否是群聊成员
func IsGroupMember(groupID string, userID uint) (bool, error) {
	var count int64
	err := global.DB.Model(&GroupMember{}).Where("group_id = ? AND user_id = ?", groupID, userID).Count(&count).Error
	return count > 0, err
}

// GetUserGroups 获取用户参与的群聊列表
func GetUserGroups(userID uint) ([]GroupChat, error) {
	var groups []GroupChat
	err := global.DB.Table("group_chats").
		Joins("JOIN group_members ON group_chats.group_id = group_members.group_id").
		Where("group_members.user_id = ?", userID).
		Find(&groups).Error
	return groups, err
}

// RemoveGroupMember 移除群聊成员
func RemoveGroupMember(groupID string, userID uint) error {
	return global.DB.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&GroupMember{}).Error
}

// GenerateGroupID 生成群聊ID
func GenerateGroupID(creatorID uint) string {
	return fmt.Sprintf("group_%d_%d", time.Now().UnixNano(), creatorID)
}

// CreateGroupChatWithMembers 创建群聊并添加成员
func CreateGroupChatWithMembers(creatorID uint, groupName, description string, memberIDs []uint) (*GroupChat, error) {
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 生成群聊ID
	groupID := GenerateGroupID(creatorID)

	// 创建群聊
	group := &GroupChat{
		GroupID:     groupID,
		GroupName:   groupName,
		CreatorID:   creatorID,
		Description: description,
		MemberCount: 0,
	}

	if err := tx.Create(group).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建群聊失败: %v", err)
	}

	// 确保创建者在成员列表中
	memberSet := make(map[uint]bool)
	memberSet[creatorID] = true
	for _, memberID := range memberIDs {
		memberSet[memberID] = true
	}

	// 添加群主
	ownerMember := &GroupMember{
		GroupID:  groupID,
		UserID:   creatorID,
		Role:     GroupRoleOwner,
		JoinedAt: time.Now(),
	}
	if err := tx.Create(ownerMember).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("添加群主失败: %v", err)
	}

	// 添加其他成员
	for memberID := range memberSet {
		if memberID != creatorID {
			member := &GroupMember{
				GroupID:  groupID,
				UserID:   memberID,
				Role:     GroupRoleMember,
				JoinedAt: time.Now(),
			}
			if err := tx.Create(member).Error; err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("添加成员 %d 失败: %v", memberID, err)
			}
		}
	}

	// 更新成员数量
	group.MemberCount = len(memberSet)
	if err := tx.Model(group).Update("member_count", group.MemberCount).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新成员数量失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败: %v", err)
	}

	return group, nil
}
