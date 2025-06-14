# 车辆库存管理正确业务流程

## 概述

本文档说明了车辆库存管理系统的正确业务流程和权限设计，特别是预订状态更新的处理方式。

## 角色和权限

### 👤 用户（User）
- **权限**: 用户JWT认证
- **操作**: 创建预订

### 🏪 商家（Merchant）  
- **权限**: 商家JWT认证
- **操作**: 维护管理、统计查看

### 🔧 系统内部（System）
- **权限**: 无需认证（内部调用）
- **操作**: 更新预订状态

### 🌐 公开访问（Public）
- **权限**: 无需认证
- **操作**: 查看可用性、浏览车辆

## 完整业务流程

### 阶段1: 用户浏览和预订

#### 1.1 用户浏览车辆（公开）
```http
GET /vehicle/list
GET /vehicle-inventory/available-vehicles
```
- 用户无需登录即可浏览车辆
- 查看车辆详情和价格信息

#### 1.2 检查车辆可用性（公开）
```http
POST /vehicle-inventory/check-availability
{
  "vehicle_id": 1,
  "start_date": "2024-01-15", 
  "end_date": "2024-01-20"
}
```
- 用户选择心仪车辆和时间段
- 系统实时检查可用性

#### 1.3 用户登录
```http
POST /user/login
```
- 用户必须登录才能预订
- 获得用户JWT token

#### 1.4 用户创建预订（用户认证）
```http
POST /vehicle-inventory/reservation/create
Authorization: Bearer <user-jwt-token>
{
  "vehicle_id": 1,
  "order_id": 1001,
  "start_date": "2024-01-15",
  "end_date": "2024-01-20"
}
```
- 用户ID从JWT自动获取
- 创建预订记录，状态为"已预订"
- 锁定该时间段的车辆

### 阶段2: 订单处理流程

#### 2.1 订单服务处理
```
订单服务内部逻辑:
1. 验证用户信息
2. 计算费用
3. 处理支付
4. 更新订单状态
```

#### 2.2 订单状态变更触发库存更新（系统内部调用）
```http
PUT /vehicle-inventory/reservation/status
Content-Type: application/json
{
  "order_id": 1001,
  "status": "rented"
}
```
- **调用方**: 订单服务
- **触发时机**: 订单支付成功后
- **无需认证**: 系统内部调用
- **作用**: 将预订状态从"已预订"更新为"租用中"

#### 2.3 租用完成（系统内部调用）
```http
PUT /vehicle-inventory/reservation/status
Content-Type: application/json
{
  "order_id": 1001,
  "status": "completed"
}
```
- **调用方**: 订单服务
- **触发时机**: 用户归还车辆后
- **作用**: 将状态更新为"已完成"，车辆重新可用

### 阶段3: 商家管理

#### 3.1 商家查看统计（商家认证）
```http
GET /vehicle-inventory/stats
Authorization: Bearer <merchant-jwt-token>
```
- 商家查看自己车辆的预订统计
- 了解车辆利用率

#### 3.2 商家设置维护（商家认证）
```http
POST /vehicle-inventory/maintenance/set
Authorization: Bearer <merchant-jwt-token>
{
  "vehicle_id": 1,
  "start_date": "2024-01-25",
  "end_date": "2024-01-27",
  "notes": "定期保养"
}
```
- 商家安排车辆维护
- 维护期间车辆不可预订

## 关键设计原则

### 1. 为什么预订状态更新不需要认证？

#### 业务原因
- **自动化流程**: 订单支付成功后自动更新预订状态
- **系统集成**: 订单服务需要调用库存服务更新状态
- **用户体验**: 用户支付后立即生效，无需人工干预

#### 技术原因
- **微服务架构**: 服务间内部调用不需要用户认证
- **性能考虑**: 避免复杂的服务间认证流程
- **可靠性**: 减少认证失败导致的业务中断

### 2. 安全性如何保证？

#### 网络层安全
```
订单服务 → 库存服务
- 内网调用，外网无法直接访问
- 可以添加服务间认证（如API Key）
```

#### 业务层验证
```go
func UpdateReservationStatus(orderID uint, status string) error {
    // 1. 验证订单ID是否存在
    // 2. 验证状态转换是否合法
    // 3. 验证操作权限（如订单所有者）
    // 4. 执行状态更新
}
```

#### 审计日志
```
记录所有状态更新操作:
- 操作时间
- 订单ID
- 状态变更
- 调用来源
```

## 状态流转图

```
预订状态流转:
用户创建预订 → 已预订 → 订单支付成功 → 租用中 → 归还车辆 → 已完成
                ↓
              取消预订 → 已取消
```

## API权限矩阵（修正版）

| 操作 | 游客 | 用户 | 系统内部 | 商家 | 说明 |
|------|------|------|----------|------|------|
| 浏览车辆 | ✅ | ✅ | ✅ | ✅ | 公开信息 |
| 检查可用性 | ✅ | ✅ | ✅ | ✅ | 公开查询 |
| 创建预订 | ❌ | ✅ | ❌ | ❌ | 只有用户可以预订 |
| 更新预订状态 | ❌ | ❌ | ✅ | ❌ | 只有系统内部调用 |
| 查看库存统计 | ❌ | ❌ | ❌ | ✅ | 商家业务数据 |
| 维护管理 | ❌ | ❌ | ❌ | ✅ | 商家运营管理 |

## 实际调用示例

### 完整的租车流程

#### 步骤1: 用户预订
```bash
# 用户检查可用性
curl -X POST http://localhost:8888/vehicle-inventory/check-availability \
  -H "Content-Type: application/json" \
  -d '{"vehicle_id": 1, "start_date": "2024-01-15", "end_date": "2024-01-20"}'

# 用户创建预订
curl -X POST http://localhost:8888/vehicle-inventory/reservation/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <user-token>" \
  -d '{"vehicle_id": 1, "order_id": 1001, "start_date": "2024-01-15", "end_date": "2024-01-20"}'
```

#### 步骤2: 订单服务内部调用
```bash
# 订单支付成功后，订单服务调用
curl -X PUT http://localhost:8888/vehicle-inventory/reservation/status \
  -H "Content-Type: application/json" \
  -d '{"order_id": 1001, "status": "rented"}'

# 租用完成后，订单服务调用
curl -X PUT http://localhost:8888/vehicle-inventory/reservation/status \
  -H "Content-Type: application/json" \
  -d '{"order_id": 1001, "status": "completed"}'
```

#### 步骤3: 商家查看统计
```bash
# 商家查看库存统计
curl -X GET http://localhost:8888/vehicle-inventory/stats \
  -H "Authorization: Bearer <merchant-token>"
```

## 错误处理

### 常见错误场景

1. **用户尝试更新预订状态**
   ```json
   {
     "code": 403,
     "message": "此操作只能由系统内部调用"
   }
   ```

2. **无效的状态转换**
   ```json
   {
     "code": 400,
     "message": "无效的状态转换：从 completed 到 rented"
   }
   ```

3. **订单不存在**
   ```json
   {
     "code": 404,
     "message": "订单不存在"
   }
   ```

## 总结

这种设计的优势：

1. **职责清晰**: 用户负责预订，系统负责状态管理，商家负责运营
2. **自动化程度高**: 减少人工干预，提高效率
3. **安全可控**: 通过业务逻辑验证保证安全性
4. **扩展性好**: 便于添加新的状态和流程
5. **符合微服务架构**: 服务间松耦合，职责分离

这种权限设计既保证了业务流程的自动化，又确保了系统的安全性和可维护性。
