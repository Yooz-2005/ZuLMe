# 车辆库存管理 API 文档

## 概述

车辆库存管理API提供了完整的车辆库存管理功能，包括可用性检查、预订管理、维护管理、统计报表等。

## 基础URL

```
http://localhost:8888
```

## 认证

库存管理接口根据功能分为四类：

1. **公开接口**: 无需认证（查看可用性、可用车辆）
2. **用户接口**: 需要用户JWT认证（创建预订）
3. **系统内部接口**: 无需认证，供其他微服务调用（更新预订状态）
4. **商家接口**: 需要商家JWT认证（维护管理、统计报表）

在请求头中包含相应的JWT token：

```
Authorization: Bearer <your-jwt-token>
```

## API接口列表

### 1. 公开接口（无需认证）

#### 1.1 检查车辆可用性

**POST** `/vehicle-inventory/check-availability`

检查指定车辆在特定时间段内是否可用。

**请求参数：**
```json
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
    "message": "检查成功",
    "is_available": true
  }
}
```

#### 1.2 获取可用车辆列表

**POST** `/vehicle-inventory/available-vehicles`

获取在指定时间段内可用的车辆列表。

**请求参数：**
```json
{
  "start_date": "2024-01-15",
  "end_date": "2024-01-20",
  "merchant_id": 1,
  "type_id": 2,
  "brand_id": 3,
  "price_min": 100.0,
  "price_max": 500.0
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "message": "获取成功",
    "vehicles": [
      {
        "id": 1,
        "merchant_id": 1,
        "brand": "奔驰",
        "style": "S级",
        "price": 300.0,
        "status": 1
      }
    ],
    "total": 1
  }
}
```

#### 1.3 获取车辆库存日历

**GET** `/vehicle-inventory/calendar`

获取指定车辆的库存日历，用户可以查看车辆在特定时间段的可用性，便于选择合适的租用日期。

**请求参数：**
```json
{
  "vehicle_id": 1,
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "message": "获取成功",
    "calendar": [
      {
        "date": "2024-01-15",
        "status": 2,
        "count": 1
      },
      {
        "date": "2024-01-16",
        "status": 1,
        "count": 0
      }
    ]
  }
}
```

**状态说明：**
- `1`: 可用
- `2`: 已预订
- `3`: 租用中
- `4`: 维护中

### 2. 用户认证接口

#### 2.1 创建预订

**POST** `/vehicle-inventory/reservation/create`

用户为指定车辆创建预订记录。用户ID从JWT token中自动获取。

**请求参数：**
```json
{
  "vehicle_id": 1,
  "order_id": 1001,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "预订创建成功",
  "data": {
    "message": "预订创建成功"
  }
}
```

### 3. 系统内部接口

#### 3.1 更新预订状态

**PUT** `/vehicle-inventory/reservation/status`

系统内部调用，用于更新预订状态（如从预订变为租用中）。通常由订单服务在订单状态变更时调用。

**请求参数：**
```json
{
  "order_id": 1001,
  "status": "rented"
}
```

**状态值说明：**
- `rented`: 租用中
- `completed`: 已完成
- `cancelled`: 已取消

**响应示例：**
```json
{
  "code": 200,
  "message": "状态更新成功",
  "data": {
    "message": "状态更新成功"
  }
}
```

### 4. 商家认证接口

#### 4.1 获取库存统计

**GET** `/vehicle-inventory/stats`

获取当前商家的库存统计信息。

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "message": "获取成功",
    "total": 10,
    "available": 6,
    "reserved": 2,
    "rented": 1,
    "maintenance": 1
  }
}
```

#### 4.2 设置维护状态

**POST** `/vehicle-inventory/maintenance/set`

商家为车辆设置维护状态。

**请求参数：**
```json
{
  "vehicle_id": 1,
  "start_date": "2024-01-21",
  "end_date": "2024-01-23",
  "notes": "定期保养",
  "created_by": 1
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "维护状态设置成功",
  "data": {
    "message": "维护状态设置成功"
  }
}
```

#### 4.3 获取维护计划

**GET** `/vehicle-inventory/maintenance/schedule`

商家获取指定车辆的维护计划。

**请求参数：**
```json
{
  "vehicle_id": 1
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "message": "获取成功",
    "maintenances": [
      {
        "id": 1,
        "vehicle_id": 1,
        "start_date": "2024-01-21",
        "end_date": "2024-01-23",
        "notes": "定期保养",
        "created_at": "2024-01-15T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

#### 4.4 获取库存报表

**GET** `/vehicle-inventory/report`

商家获取指定时间段的库存利用率报表。

**请求参数：**
```json
{
  "merchant_id": 1,
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "message": "获取成功",
    "total_vehicles": 10,
    "total_days": 31,
    "total_capacity": 310,
    "reservations": 50,
    "rentals": 30,
    "maintenances": 10,
    "used_capacity": 90,
    "utilization_rate": 29.03
  }
}
```

## 错误处理

所有API接口都遵循统一的错误响应格式：

```json
{
  "code": 400,
  "message": "错误描述",
  "data": null
}
```

常见错误码：
- `400`: 请求参数错误
- `401`: 未授权访问
- `403`: 权限不足
- `404`: 资源不存在
- `500`: 服务器内部错误

## 使用示例

### 完整的预订流程

1. **检查可用性**
```bash
curl -X POST http://localhost:8888/vehicle-inventory/check-availability \
  -H "Content-Type: application/json" \
  -d '{
    "vehicle_id": 1,
    "start_date": "2024-01-15",
    "end_date": "2024-01-20"
  }'
```

2. **创建预订**
```bash
curl -X POST http://localhost:8888/vehicle-inventory/reservation/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "vehicle_id": 1,
    "order_id": 1001,
    "user_id": 123,
    "start_date": "2024-01-15",
    "end_date": "2024-01-20"
  }'
```

3. **更新为租用状态**
```bash
curl -X PUT http://localhost:8888/vehicle-inventory/reservation/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "order_id": 1001,
    "status": "rented"
  }'
```

## 注意事项

1. 所有日期格式必须为 `YYYY-MM-DD`
2. 商家认证接口会自动从JWT token中获取商家ID
3. 预订创建需要确保订单已存在
4. 状态更新操作具有原子性，失败时会自动回滚
5. 库存统计数据实时计算，反映当前状态
