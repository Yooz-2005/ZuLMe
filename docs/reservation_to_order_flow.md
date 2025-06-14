# 预订到订单的完整业务流程

## 概述

本文档详细说明了新的两步式租车流程：**先创建预订，再基于预订创建订单**。

## 🎯 **新流程设计**

### **传统流程 vs 新流程**

#### 传统流程（一步式）
```
用户选择车辆 → 直接创建订单 → 支付 → 完成
```

#### 新流程（两步式）
```
用户选择车辆 → 创建预订 → 基于预订创建订单 → 支付 → 完成
```

### **新流程的优势**

1. **用户体验更好**
   - 用户可以先锁定车辆，再慢慢填写订单详情
   - 避免填写订单时车辆被其他用户预订

2. **业务逻辑更清晰**
   - 预订：锁定车辆和时间段
   - 订单：处理支付和具体安排

3. **系统更灵活**
   - 可以支持预订后修改取还车地点
   - 可以支持预订后的价格调整

## 📋 **完整API流程**

### **第一步：用户创建预订**

#### 1.1 检查车辆可用性（公开）
```http
POST /vehicle-inventory/check-availability
Content-Type: application/json

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "检查成功",
  "data": {
    "is_available": true
  }
}
```

#### 1.2 用户登录
```http
POST /user/login
Content-Type: application/json

{
  "username": "user123",
  "password": "password123"
}
```

#### 1.3 创建预订（需要用户认证）
```http
POST /vehicle-inventory/reservation/create
Content-Type: application/json
Authorization: Bearer <user-jwt-token>

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20",
  "notes": "需要儿童座椅"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "预订创建成功",
  "data": {
    "message": "预订创建成功",
    "reservation_id": 123
  }
}
```

### **第二步：基于预订创建订单**

#### 2.1 基于预订创建订单（需要用户认证）
```http
POST /order/create-from-reservation
Content-Type: application/json
Authorization: Bearer <user-jwt-token>

{
  "reservation_id": 123,
  "pickup_location_id": 1,
  "return_location_id": 2,
  "notes": "需要儿童座椅，下午3点取车",
  "payment_method": 1,
  "expected_total_amount": 1500.00
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "订单创建成功",
  "data": {
    "message": "订单创建成功",
    "order_id": 1001,
    "order_sn": "ORD20240115001",
    "total_amount": 1500.00,
    "status": "待支付",
    "payment_url": "https://pay.example.com/pay/1001"
  }
}
```

### **第三步：支付和状态更新**

#### 3.1 用户支付（外部支付系统）
```
用户通过支付宝/微信完成支付
```

#### 3.2 支付成功后系统自动更新状态
```http
PUT /vehicle-inventory/reservation/status
Content-Type: application/json

{
  "order_id": 1001,
  "status": "rented"
}
```

## 🗄️ **数据库设计变更**

### **预订表 (vehicle_inventory) 变更**

#### 修改前
```sql
CREATE TABLE vehicle_inventory (
  id bigint PRIMARY KEY,
  vehicle_id bigint NOT NULL,
  order_id bigint NOT NULL,  -- 必须字段
  start_date date NOT NULL,
  end_date date NOT NULL,
  status int NOT NULL,
  created_by bigint NOT NULL
);
```

#### 修改后
```sql
CREATE TABLE vehicle_inventory (
  id bigint PRIMARY KEY,
  vehicle_id bigint NOT NULL,
  order_id bigint DEFAULT 0,  -- 可选字段，初始为0
  start_date date NOT NULL,
  end_date date NOT NULL,
  status int NOT NULL,
  notes varchar(500),         -- 新增：预订备注
  created_by bigint NOT NULL
);
```

### **新增方法**

```go
// 创建预订（不需要order_id）
func (vi *VehicleInventory) CreateReservation(vehicleID uint, startDate, endDate time.Time, createdBy uint, notes string) (*VehicleInventory, error)

// 更新预订的订单ID
func (vi *VehicleInventory) UpdateReservationOrderID(reservationID uint, orderID uint) error
```

## 🔄 **状态流转图**

```
预订状态流转:
1. 用户创建预订 → 状态: 已预订 (order_id = 0)
2. 创建订单成功 → 状态: 已预订 (order_id = 实际订单ID)
3. 支付成功     → 状态: 租用中
4. 归还车辆     → 状态: 已完成
```

## 📝 **请求参数说明**

### **CreateReservationRequest（新）**
```go
type CreateReservationRequest struct {
    VehicleID int64  `json:"vehicle_id" binding:"required"`
    StartDate string `json:"start_date" binding:"required"` // YYYY-MM-DD
    EndDate   string `json:"end_date" binding:"required"`   // YYYY-MM-DD
    Notes     string `json:"notes"`                         // 预订备注
    // UserID 从JWT token中获取
}
```

### **CreateOrderFromReservationRequest（新）**
```go
type CreateOrderFromReservationRequest struct {
    ReservationID       int64   `json:"reservation_id" binding:"required"`
    PickupLocationID    int64   `json:"pickup_location_id" binding:"required"`
    ReturnLocationID    int64   `json:"return_location_id" binding:"required"`
    Notes               string  `json:"notes"`
    PaymentMethod       int32   `json:"payment_method"`       // 1:支付宝 2:微信
    ExpectedTotalAmount float64 `json:"expected_total_amount"` // 前端计算的预期金额
}
```

## 🧪 **测试用例**

### **完整流程测试**

```http
### 1. 检查车辆可用性
POST http://localhost:8888/vehicle-inventory/check-availability
Content-Type: application/json

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20"
}

### 2. 用户登录
POST http://localhost:8888/user/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}

### 3. 创建预订
POST http://localhost:8888/vehicle-inventory/reservation/create
Content-Type: application/json
Authorization: Bearer {{userToken}}

{
  "vehicle_id": 1,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20",
  "notes": "需要儿童座椅"
}

### 4. 基于预订创建订单
POST http://localhost:8888/order/create-from-reservation
Content-Type: application/json
Authorization: Bearer {{userToken}}

{
  "reservation_id": 123,
  "pickup_location_id": 1,
  "return_location_id": 2,
  "notes": "下午3点取车",
  "payment_method": 1,
  "expected_total_amount": 1500.00
}
```

## ⚠️ **注意事项**

### **1. 预订超时处理**
- 建议设置预订有效期（如30分钟）
- 超时未创建订单的预订自动取消

### **2. 并发控制**
- 创建预订时需要加锁防止重复预订
- 使用数据库事务确保数据一致性

### **3. 错误处理**
- 预订不存在
- 预订已关联订单
- 预订不属于当前用户
- 车辆已被其他用户预订

### **4. 业务规则**
- 一个预订只能创建一个订单
- 预订创建后车辆立即锁定
- 支持预订取消（未创建订单前）

## 🚀 **后续扩展**

### **可能的功能扩展**
1. **预订修改**：允许用户修改预订时间（在创建订单前）
2. **预订转让**：允许用户将预订转给其他用户
3. **批量预订**：支持一次预订多辆车
4. **预订提醒**：预订即将过期时提醒用户
5. **预订历史**：用户查看自己的预订历史

这种设计为未来的功能扩展提供了良好的基础架构。
