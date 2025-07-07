# WebSocket 聊天服务

## 功能特性

- 实时聊天消息
- 用户在线状态管理
- 房间管理（加入/离开房间）
- 群发消息
- 心跳检测
- JWT Token 自动认证（无需手动传入用户ID）
- 百度敏感词过滤

## API 接口

### WebSocket 连接
- **地址**: `ws://localhost:8009/ws/chat`
- **认证**: 需要在请求头中携带 JWT Token
- **自动上线**: 连接成功后自动上线，无需发送online命令

### 消息格式

#### 客户端发送消息格式
```json
{
    "cmd": "命令类型",
    "data": {
        // 具体数据
    }
}
```

#### 服务端响应格式
```json
{
    "code": 0,
    "message": "success",
    "data": {
        // 响应数据
    }
}
```

## 支持的命令

### 1. 自动上线（连接时自动触发）
连接成功后会自动收到欢迎消息：
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "welcome",
        "message": "连接成功，自动上线！",
        "user_id": 1,
        "time": "2025-07-01 11:03:23"
    }
}
```

### 2. 发送一对一消息 (send_message)
```json
{
    "cmd": "send_message",
    "data": {
        "to_user_id": 2,
        "message_type": 1,
        "content": "你好！"
    }
}
```
**说明**:
- `room_id` 会根据发送者和接收者的用户ID自动生成 (格式: `private_{小ID}_{大ID}`)
- 消息会自动存储到MongoDB数据库
- 支持百度敏感词过滤

### 3. 发送群聊消息 (send_group_message)
```json
{
    "cmd": "send_group_message",
    "data": {
        "group_id": "group_1719816465123456789",
        "message_type": 1,
        "content": "群聊消息内容"
    }
}
```
**说明**:
- 支持群聊ID格式验证
- 消息会自动存储到MongoDB数据库
- 支持百度敏感词过滤
- 广播给所有在线用户（简化处理）

### 4. 群发消息 (broadcast)
```json
{
    "cmd": "broadcast",
    "data": {
        "message": "群发消息内容"
    }
}
```

### 5. 心跳检测 (ping)
```json
{
    "cmd": "ping",
    "data": {}
}
```

## 服务端推送消息类型

### 新消息通知
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "new_message",
        "from_user_id": 1,
        "to_user_id": 2,
        "content": "你好！",
        "room_id": "private_1_2",
        "message_type": 1,
        "message_id": "60f7b3b3b3b3b3b3b3b3b3b3",
        "time": "2025-07-01 11:03:23"
    }
}
```

### 群聊消息通知
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "group_message",
        "group_id": "group_1719816465123456789",
        "group_name": "我的群聊",
        "from_user_id": 1,
        "content": "群聊消息内容",
        "message_type": 1,
        "time": "2025-07-01 11:03:23"
    }
}
```

### 群聊消息发送成功通知
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "group_message_sent",
        "group_id": "group_1719816465123456789",
        "message_id": "60f7b3b3b3b3b3b3b3b3b3b3",
        "sent_count": 5,
        "message": "群聊消息已发送"
    }
}
```

### 群聊消息通知
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "group_message",
        "group_id": "group_1719816465123456789",
        "from_user_id": 1,
        "content": "群聊消息内容",
        "message_type": 1,
        "message_id": "60f7b3b3b3b3b3b3b3b3b3b3",
        "time": "2025-07-01 11:03:23"
    }
}
```

### 房间操作成功通知
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "type": "join_room_success",
        "room_id": "room_123",
        "user_id": 1,
        "time": "2025-07-01 11:03:23"
    }
}
```

## HTTP API 接口

### 获取在线用户数量
- **URL**: `GET /api/ws/online-count`
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "count": 5
    }
}
```

### 获取在线用户列表
- **URL**: `GET /api/ws/online-users`
- **认证**: 需要 JWT Token
- **响应**:
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "users": [1, 2, 3, 4, 5]
    }
}
```

### 发送消息 (HTTP)
- **URL**: `POST /api/ws/send-message`
- **认证**: 需要 JWT Token
- **请求体**:
```json
{
    "room_id": "room_123",
    "to_user_id": 2,
    "message_type": 1,
    "content": "消息内容"
}
```

### 群发消息 (HTTP)
- **URL**: `POST /api/ws/broadcast`
- **认证**: 需要 JWT Token
- **请求体**:
```json
{
    "message": "群发消息内容"
}
```

## 错误码说明

- `1000`: 未知命令
- `1001`: 请求参数错误
- `2001`: 房间参数解析失败
- `2002`: 房间ID不能为空
- `2003`: 房间参数解析失败
- `2004`: 房间ID不能为空
- `3001`: 群聊参数解析失败
- `3002`: 群聊名称不能为空
- `3003`: 群成员列表不能为空
- `3004`: 群聊消息解析失败
- `3005`: 群聊ID和消息内容不能为空
- `3006`: 群聊ID格式无效
- `3008`: 消息包含敏感内容
- `3009`: 群聊消息保存失败
- `10001`: 消息解析失败
- `10002`: 目标用户ID和消息内容不能为空
- `10003`: 目标用户不在线
- `10004`: 消息发送失败
- `10005`: 群发消息内容不能为空
- `10006`: 群发消息失败
- `10007`: 消息包含敏感内容
- `10008`: 消息保存失败

## 重要改进

### 🔐 安全性提升
- **自动Token解析**: 用户ID通过JWT Token自动获取，无需手动传入
- **无需手动上线**: 连接建立后自动上线，提升用户体验
- **敏感词过滤**: 集成百度敏感词API，自动过滤不当内容

### 💬 智能一对一聊天
- **自动房间ID生成**: 根据用户ID自动生成唯一的聊天房间ID (格式: `private_{小ID}_{大ID}`)
- **消息持久化存储**: 自动将聊天消息存储到MongoDB数据库
- **实时消息转发**: 支持用户间的实时消息传递
- **在线状态检测**: 自动检测目标用户是否在线
- **消息状态反馈**: 实时反馈消息发送状态和消息ID

### 🏠 房间管理
- **房间加入/离开**: 支持用户加入和离开聊天房间
- **房间消息**: 消息支持房间ID，便于群聊管理

### 💬 群聊消息增强
- **群聊ID验证**: 支持群聊ID格式验证，确保消息发送安全
- **消息持久化**: 群聊消息自动存储到MongoDB数据库
- **实时广播**: 支持群聊消息的实时转发给所有在线用户
- **消息统计**: 提供发送成功统计和消息ID追踪
- **简化架构**: WebSocket专注实时通信，群聊管理由Chat微服务负责

### 📱 消息类型
- **消息类型字段**: 支持不同类型的消息（文本、图片、文件等）
- **结构化数据**: 更完整的消息数据结构

## 部署说明

1. 确保 MySQL、MongoDB、Redis 服务正常运行
2. 配置 `config.yaml` 文件
3. 运行服务：
```bash
go run websocket_srv_main.go
```

服务将在 `:8009` 端口启动。
