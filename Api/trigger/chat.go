package trigger

import (
	"Api/handler"
	"Api/request"
	chat "chat_srv/proto_chat"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendMessageHandler 发送消息
func SendMessageHandler(c *gin.Context) {
	var req request.SendMessageRequest

	// 支持JSON和Form表单绑定
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 默认消息类型为文本
	if req.MessageType == 0 {
		req.MessageType = 1
	}

	// 调用聊天服务的gRPC接口
	grpcReq := &chat.SendMessageRequest{
		RoomId:      req.RoomID,
		FromUserId:  uint32(userID),
		ToUserId:    uint32(req.ToUserID),
		MessageType: req.MessageType,
		Content:     req.Content,
	}

	response, err := handler.SendMessage(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "发送消息失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "消息发送成功",
		"data":    response,
	})
}

// GetRoomMessagesHandler 获取房间消息
func GetRoomMessagesHandler(c *gin.Context) {
	var req request.GetRoomMessagesRequest

	// 绑定URI参数和查询参数
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "房间ID参数错误: " + err.Error(),
		})
		return
	}

	// 绑定查询参数，支持form表单
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口获取消息
	grpcReq := &chat.GetRoomMessagesRequest{
		RoomId:   req.RoomID,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}

	response, err := handler.GetRoomMessages(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取消息失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取消息成功",
		"data":    response,
	})
}

// AddFriendHandler 添加好友
func AddFriendHandler(c *gin.Context) {
	var req request.AddFriendRequest

	// 支持JSON和Form表单绑定
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口添加好友
	grpcReq := &chat.AddFriendRequest{
		UserId:   uint32(userID),
		FriendId: uint32(req.FriendID),
	}

	response, err := handler.AddFriend(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "添加好友失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加好友成功",
		"data":    response,
	})
}

// GetFriendListHandler 获取好友列表
func GetFriendListHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口获取好友列表
	grpcReq := &chat.GetFriendListRequest{
		UserId: uint32(userID),
	}

	response, err := handler.GetFriendList(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取好友列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取好友列表成功",
		"data":    response,
	})
}

// IsFriendHandler 检查好友关系
func IsFriendHandler(c *gin.Context) {
	var req request.IsFriendRequest

	// 绑定URI参数
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "好友ID参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口检查好友关系
	grpcReq := &chat.IsFriendRequest{
		UserId:   uint32(userID),
		FriendId: uint32(req.FriendID),
	}

	response, err := handler.IsFriend(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "检查好友关系失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "检查好友关系成功",
		"data":    response,
	})
}

// CreateGroupChatHandler 创建群聊
func CreateGroupChatHandler(c *gin.Context) {
	var req request.CreateGroupChatRequest

	// 支持JSON和Form表单绑定
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 转换成员ID列表
	memberIDs := make([]uint32, len(req.MemberIDs))
	for i, id := range req.MemberIDs {
		memberIDs[i] = uint32(id)
	}

	// 调用聊天服务的gRPC接口创建群聊
	grpcReq := &chat.CreateGroupChatRequest{
		CreatorId:   uint32(userID),
		GroupName:   req.GroupName,
		Description: req.Description,
		MemberIds:   memberIDs,
	}

	response, err := handler.CreateGroupChat(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建群聊失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建群聊成功",
		"data":    response,
	})
}

// GetGroupChatInfoHandler 获取群聊信息
func GetGroupChatInfoHandler(c *gin.Context) {
	var req request.GetGroupChatInfoRequest

	// 绑定URI参数
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "群聊ID参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口获取群聊信息
	grpcReq := &chat.GetGroupChatInfoRequest{
		GroupId: req.GroupID,
		UserId:  uint32(userID),
	}

	response, err := handler.GetGroupChatInfo(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取群聊信息失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取群聊信息成功",
		"data":    response,
	})
}

// GetGroupMembersHandler 获取群聊成员列表
func GetGroupMembersHandler(c *gin.Context) {
	var req request.GetGroupMembersRequest

	// 绑定URI参数
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "群聊ID参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口获取群聊成员
	grpcReq := &chat.GetGroupMembersRequest{
		GroupId: req.GroupID,
		UserId:  uint32(userID),
	}

	response, err := handler.GetGroupMembers(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取群聊成员失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取群聊成员成功",
		"data":    response,
	})
}

// InviteToGroupHandler 邀请用户加入群聊
func InviteToGroupHandler(c *gin.Context) {
	var uriReq request.InviteToGroupURIRequest
	var bodyReq request.InviteToGroupRequest

	// 绑定URI参数
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "群聊ID参数错误: " + err.Error(),
		})
		return
	}

	// 绑定请求体参数
	if err := c.ShouldBind(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 转换用户ID列表
	userIDs := make([]uint32, len(bodyReq.UserIDs))
	for i, id := range bodyReq.UserIDs {
		userIDs[i] = uint32(id)
	}

	// 调用聊天服务的gRPC接口邀请用户
	grpcReq := &chat.InviteToGroupRequest{
		GroupId:   uriReq.GroupID,
		InviterId: uint32(userID),
		UserIds:   userIDs,
	}

	response, err := handler.InviteToGroup(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "邀请用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "邀请用户成功",
		"data":    response,
	})
}

// LeaveGroupHandler 退出群聊
func LeaveGroupHandler(c *gin.Context) {
	var req request.LeaveGroupRequest

	// 绑定URI参数
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "群聊ID参数错误: " + err.Error(),
		})
		return
	}

	// 从JWT中获取用户ID
	userID := c.GetUint("userId")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口退出群聊
	grpcReq := &chat.LeaveGroupRequest{
		GroupId: req.GroupID,
		UserId:  uint32(userID),
	}

	response, err := handler.LeaveGroup(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "退出群聊失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出群聊成功",
		"data":    response,
	})
}

// GetUserGroupsHandler 获取用户参与的群聊列表
func GetUserGroupsHandler(c *gin.Context) {
	// 从JWT中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登录",
		})
		return
	}

	// 调用聊天服务的gRPC接口获取用户群聊列表
	grpcReq := &chat.GetUserGroupsRequest{
		UserId: uint32(userID.(uint)),
	}

	response, err := handler.GetUserGroups(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取群聊列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取群聊列表成功",
		"data":    response,
	})
}
