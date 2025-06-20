# ZuLMe 支付流程说明

## 概述
ZuLMe 租车平台采用两步式支付流程：先创建预订，再从预订生成订单并完成支付。

## 支付流程

### 1. 预订阶段
- 用户在车辆详情页选择租赁时间和取车地点
- 系统创建预订记录，状态为 `pending_payment`（等待付款）
- 用户可以在个人中心的"预订管理"中查看预订

### 2. 支付阶段
- 用户在个人中心点击"立即支付"
- 弹出还车地点选择窗口
- 选择还车地点后，系统从预订创建订单
- 生成支付链接，跳转到支付页面

### 3. 支付完成
- 支付成功后，订单状态更新为 `已支付`
- 预订状态更新为 `confirmed`（预订成功）
- 用户可以在"订单管理"中查看订单详情

## 页面结构

### 个人中心 (`/personal-center`)
- **预订管理**: 显示所有预订记录，可以进行支付操作
- **订单管理**: 显示所有订单记录，查看支付状态

### 支付相关页面
- **支付成功页面** (`/payment/success`): 支付完成后的确认页面
- **支付测试页面** (`/payment/test`): 开发测试用的支付流程测试

## 状态说明

### 预订状态
- `processing`: 处理中
- `pending_payment`: 等待付款
- `confirmed`: 预订成功
- `in_use`: 租赁中
- `completed`: 已完成
- `cancelled`: 已取消

### 订单状态
- `1`: 待支付
- `2`: 已支付
- `3`: 使用中
- `4`: 已完成
- `5`: 已取消

## API 接口

### 预订相关
- `POST /reservation/create` - 创建预订
- `GET /reservation/user-list` - 获取用户预订列表
- `PUT /reservation/cancel/:id` - 取消预订

### 订单相关
- `POST /order/create-from-reservation` - 从预订创建订单
- `GET /order/list` - 获取用户订单列表
- `GET /order/detail/:id` - 获取订单详情
- `GET /order/detail-by-sn/:orderSn` - 根据订单号获取详情

### 支付相关
- `POST /payment/alipay/notify` - 支付宝回调通知

## 组件说明

### ReservationList
- 位置: `src/components/ReservationList/index.jsx`
- 功能: 显示预订列表，处理支付操作
- 特性: 支付时可选择还车地点

### OrderList
- 位置: `src/components/OrderList/index.jsx`
- 功能: 显示订单列表，查看订单状态
- 特性: 支持继续支付未完成的订单

### PaymentSuccess
- 位置: `src/pages/PaymentSuccess/index.jsx`
- 功能: 支付成功确认页面
- 特性: 自动获取订单信息并显示

## 测试说明

### 使用支付测试页面
1. 访问 `http://localhost:3000/payment/test`
2. 点击"测试完整支付流程"
3. 观察测试结果，确认各个步骤正常

### 手动测试流程
1. 登录系统
2. 选择车辆并创建预订
3. 在个人中心查看预订
4. 点击"立即支付"并选择还车地点
5. 完成支付宝支付
6. 查看订单状态变化

## 注意事项

1. **登录状态**: 所有支付相关操作都需要用户登录
2. **还车地点**: 支付时必须选择还车地点，支付后无法更改
3. **状态同步**: 支付成功后状态更新可能有延迟，建议刷新页面
4. **错误处理**: 支付失败时会显示错误信息，用户可以重试

## 开发配置

### 前端配置
- API 基础地址: `http://localhost:8888`

### 后端配置
- 支付回调地址: `http://localhost:8888/payment/alipay/notify`
- 支付成功跳转: `http://localhost:3000/payment/success`

## 扩展功能

### 已实现
- ✅ 预订创建和管理
- ✅ 订单生成和支付
- ✅ 还车地点选择
- ✅ 支付状态回调
- ✅ 支付宝支付集成

### 待扩展
- 🔄 微信支付支持
- 🔄 订单退款功能
- 🔄 支付失败重试机制
- 🔄 订单状态推送通知
