package router

import (
	"Api/trigger"
	jwt "Common/pkg"

	"github.com/gin-gonic/gin"
)

// RegisterChatRoutes 注册聊天相关路由
func RegisterChatRoutes(r *gin.Engine) {
	// 聊天API路由组
	chatGroup := r.Group("/chat")
	{
		// 需要用户认证的聊天接口
		chatGroup.Use(jwt.JWTAuth("2209"))
		{
			// 消息相关接口
			chatGroup.POST("/send", trigger.SendMessageHandler)                      // 发送消息
			chatGroup.GET("/room/:room_id/messages", trigger.GetRoomMessagesHandler) // 获取房间消息

			// 好友相关接口
			chatGroup.POST("/friend/add", trigger.AddFriendHandler)      // 添加好友
			chatGroup.GET("/friends", trigger.GetFriendListHandler)      // 获取好友列表
			chatGroup.GET("/friend/:friend_id", trigger.IsFriendHandler) // 检查好友关系

			// 群聊相关接口
			chatGroup.POST("/group/create", trigger.CreateGroupChatHandler)           // 创建群聊
			chatGroup.GET("/groups", trigger.GetUserGroupsHandler)                    // 获取用户群聊列表
			chatGroup.GET("/group/:group_id", trigger.GetGroupChatInfoHandler)        // 获取群聊信息
			chatGroup.GET("/group/:group_id/members", trigger.GetGroupMembersHandler) // 获取群聊成员
			chatGroup.POST("/group/:group_id/invite", trigger.InviteToGroupHandler)   // 邀请加入群聊
			chatGroup.POST("/group/:group_id/leave", trigger.LeaveGroupHandler)       // 退出群聊
		}
	}

	// WebSocket连接路由（在websocket_srv中处理）
	// ws://localhost:8009/ws/chat
}
