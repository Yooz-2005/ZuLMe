# 订单微服务完整指南

## 🏗️ **微服务架构概览**

### **服务架构图**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   前端 (React)   │    │  API Gateway    │    │  订单微服务      │
│   Port: 3000    │◄──►│   Port: 8888    │◄──►│   Port: 9093    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌─────────────────┐    ┌─────────────────┐
                       │  车辆微服务      │    │   MySQL数据库    │
                       │   Port: 9092    │    │   Port: 3306    │
                       └─────────────────┘    └─────────────────┘
```

### **服务职责分工**

#### **API Gateway (Port: 8888)**
- 用户认证和授权
- 请求路由和负载均衡
- 统一的错误处理和响应格式
- 支付宝回调处理

#### **订单微服务 (Port: 9093)**
- 订单创建和管理
- 支付状态处理
- 订单状态流转
- 支付宝异步通知处理

#### **车辆微服务 (Port: 9092)**
- 车辆信息管理
- 库存和预订管理
- 可用性检查

## 🚀 **启动服务**

### **1. 启动顺序**
```bash
# 1. 启动MySQL数据库
# 确保MySQL服务正在运行

# 2. 启动车辆微服务
cd ZuLMe/Srv/vehicle_srv
go run vehicle_srv_main.go

# 3. 启动订单微服务
cd ZuLMe/Srv/order_srv
go run order_srv_main.go

# 4. 启动API Gateway
cd ZuLMe/Api
go run api_main.go
```

### **2. 验证服务状态**
```bash
# 检查车辆微服务
grpcurl -plaintext localhost:9092 vehicle.Vehicle/Ping

# 检查订单微服务
grpcurl -plaintext localhost:9093 order.Order/Ping

# 检查API Gateway
curl http://localhost:8888/health
```

## 📋 **完整业务流程测试**

### **第1步：用户登录获取Token**
```http
POST http://localhost:8888/user/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

### **第2步：检查车辆可用性**
```http
POST http://localhost:8888/vehicle-inventory/check-availability
Content-Type: application/json

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20"
}
```

### **第3步：创建预订**
```http
POST http://localhost:8888/vehicle-inventory/reservation/create
Content-Type: application/json
Authorization: Bearer <user-token>

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20",
  "notes": "需要儿童座椅"
}
```

### **第4步：基于预订创建订单（调用订单微服务）**
```http
POST http://localhost:8888/order/create-from-reservation
Content-Type: application/json
Authorization: Bearer <user-token>

{
  "reservation_id": 123,
  "pickup_location_id": 1,
  "return_location_id": 2,
  "notes": "下午3点取车",
  "payment_method": 1,
  "expected_total_amount": 1500.00
}
```

### **第5步：模拟支付宝支付成功**
```http
POST http://localhost:8888/payment/alipay/notify
Content-Type: application/x-www-form-urlencoded

out_trade_no=ORD1703123456&trade_no=2024011522001234567890&trade_status=TRADE_SUCCESS&total_amount=1500.00&gmt_payment=2024-01-15 15:30:00
```

## 🔄 **微服务调用链路**

### **创建订单的调用链路**
```
1. 前端 → API Gateway
   POST /order/create-from-reservation

2. API Gateway → 订单微服务
   gRPC: CreateOrderFromReservation()

3. 订单微服务 → 数据库
   - 验证预订信息
   - 获取车辆信息
   - 创建订单记录
   - 更新预订关联
   - 生成支付链接

4. 订单微服务 → API Gateway
   返回订单信息和支付链接

5. API Gateway → 前端
   返回JSON响应
```

### **支付回调的调用链路**
```
1. 支付宝 → API Gateway
   POST /payment/alipay/notify

2. API Gateway → 订单微服务
   gRPC: AlipayNotify()

3. 订单微服务 → 数据库
   - 更新订单支付状态
   - 更新预订状态
   - 记录支付宝交易号

4. 订单微服务 → API Gateway
   返回处理结果

5. API Gateway → 支付宝
   返回 "success"
```

## 🗄️ **数据库设计**

### **订单表 (orders)**
```sql
CREATE TABLE orders (
  id bigint PRIMARY KEY AUTO_INCREMENT,
  user_id int NOT NULL,
  vehicle_id int NOT NULL,
  reservation_id int NOT NULL,  -- 关联预订ID
  order_sn varchar(50) UNIQUE NOT NULL,
  pickup_location_id int NOT NULL,
  return_location_id int NOT NULL,
  pickup_time datetime NOT NULL,
  return_time datetime NOT NULL,
  rental_days int NOT NULL,
  daily_rate decimal(10,2) NOT NULL,
  total_amount decimal(10,2) NOT NULL,
  status int DEFAULT 1,  -- 订单状态
  payment int DEFAULT 1,  -- 支付方式
  payment_status int DEFAULT 1,  -- 支付状态
  payment_url varchar(500),  -- 支付链接
  alipay_trade_no varchar(100),  -- 支付宝交易号
  notes varchar(500),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);
```

### **预订表 (vehicle_inventory)**
```sql
-- 新增字段
ALTER TABLE vehicle_inventory ADD COLUMN order_id bigint DEFAULT 0;
ALTER TABLE vehicle_inventory ADD COLUMN notes varchar(500);
```

## 🧪 **微服务单独测试**

### **直接测试订单微服务**
```bash
# 使用grpcurl测试
grpcurl -plaintext -d '{
  "reservation_id": 1,
  "user_id": 1,
  "pickup_location_id": 1,
  "return_location_id": 2,
  "notes": "测试订单",
  "payment_method": 1,
  "expected_total_amount": 1500.0
}' localhost:9093 order.Order/CreateOrderFromReservation
```

### **测试支付宝通知处理**
```bash
grpcurl -plaintext -d '{
  "out_trade_no": "ORD1703123456",
  "trade_no": "2024011522001234567890",
  "trade_status": "TRADE_SUCCESS",
  "total_amount": "1500.00",
  "gmt_payment": "2024-01-15 15:30:00"
}' localhost:9093 order.Order/AlipayNotify
```

## 🔧 **配置说明**

### **订单微服务配置**
- **端口**: 9093
- **数据库**: 共享MySQL实例
- **依赖**: 车辆微服务、支付宝服务

### **支付宝配置**
```go
// ZuLMe/Common/payment/alipay.go
config := &AlipayConfig{
    AppID:      "您的支付宝应用ID",
    PrivateKey: "您的应用私钥",
    PublicKey:  "支付宝公钥",
    NotifyURL:  "http://localhost:8888/payment/alipay/notify",
    ReturnURL:  "http://localhost:3000/payment/success",
}
```

## 📊 **监控和日志**

### **服务健康检查**
- 订单微服务: `grpcurl -plaintext localhost:9093 order.Order/Ping`
- 车辆微服务: `grpcurl -plaintext localhost:9092 vehicle.Vehicle/Ping`
- API Gateway: `curl http://localhost:8888/health`

### **日志位置**
- 订单微服务日志: 控制台输出
- API Gateway日志: 控制台输出
- 数据库日志: MySQL日志文件

## ⚠️ **注意事项**

### **1. 服务启动顺序**
必须按照依赖关系启动：数据库 → 车辆微服务 → 订单微服务 → API Gateway

### **2. 端口冲突**
确保以下端口未被占用：
- 3306: MySQL
- 8888: API Gateway
- 9092: 车辆微服务
- 9093: 订单微服务

### **3. 数据一致性**
订单微服务使用数据库事务确保数据一致性，支付成功时会同时更新订单和预订状态。

### **4. 错误处理**
- 网络错误: 自动重试机制
- 业务错误: 返回明确的错误码和消息
- 系统错误: 记录日志并返回通用错误

## 🚀 **生产环境部署**

### **Docker化部署**
```dockerfile
# 订单微服务 Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o order_srv order_srv_main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/order_srv .
EXPOSE 9093
CMD ["./order_srv"]
```

### **Kubernetes部署**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
    spec:
      containers:
      - name: order-service
        image: zulme/order-service:latest
        ports:
        - containerPort: 9093
```

现在您的订单微服务已经完全独立，具备了完整的微服务架构！🎉
