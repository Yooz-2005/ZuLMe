# ZuLMe 租车平台前端

## 项目概述

ZuLMe 是一个现代化的租车平台前端应用，基于 React 18 + Ant Design 5 构建。

## 技术栈

- **React 18** - 前端框架
- **Ant Design 5** - UI 组件库
- **React Router 6** - 路由管理
- **Styled Components** - CSS-in-JS 样式方案
- **Axios** - HTTP 客户端
- **Day.js** - 日期处理库

## 功能特性

### 🏠 首页 (/)
- 轮播图展示
- 车辆搜索功能
- 服务特色介绍
- 热门车型展示

### 🚗 车辆相关页面
- **车辆列表页** (`/vehicles`) - 展示所有可租车辆
- **车辆详情页** (`/vehicle/:id`) - 显示单个车辆的详细信息
- **搜索结果页** (`/search`) - 显示搜索结果

### 👤 用户页面
- **登录注册页** (`/login-register`) - 用户认证界面

## 组件结构

### 页面组件 (Pages)
```
src/pages/
├── Home/                 # 首页
├── VehicleList/         # 车辆列表页
├── VehicleDetail/       # 车辆详情页
├── SearchResults/       # 搜索结果页
└── LoginRegister/       # 登录注册页
```

### 通用组件 (Components)
```
src/components/
├── VehicleCard/         # 车辆卡片组件
└── SearchForm/          # 搜索表单组件
```

### 服务层 (Services)
```
src/services/
├── api.js              # 基础 API 配置
└── vehicleService.js   # 车辆相关 API 服务
```

### 工具类 (Utils)
```
src/utils/
└── constants.js        # 常量定义
```

## API 接口

### 车辆相关接口
- `GET /vehicle/list` - 获取车辆列表
- `GET /vehicle/:id` - 获取车辆详情
- `GET /vehicle-type/list` - 获取车辆类型列表

## 启动项目

1. 安装依赖
```bash
npm install
```

2. 启动开发服务器
```bash
npm start
```

3. 访问应用
```
http://localhost:3000
```

## 项目配置

### API 配置
在 `src/utils/constants.js` 中配置后端 API 地址：
```javascript
export const API_BASE_URL = 'http://localhost:8080';
```

### 路由配置
所有路由在 `src/App.jsx` 中定义：
- `/` - 首页
- `/login-register` - 登录注册
- `/vehicles` - 车辆列表
- `/vehicle/:id` - 车辆详情
- `/search` - 搜索结果

## 开发说明

### 搜索功能
- 支持按地点、日期、车型搜索
- 搜索参数通过 URL 查询参数传递
- 支持分页显示结果

### 车辆展示
- 响应式网格布局
- 支持车辆状态显示（可租用/已租出/维护中）
- 车辆类型标签化显示

### 错误处理
- 统一的 API 错误处理
- 友好的错误提示界面
- 网络异常重试机制

## 注意事项

1. 确保后端 API 服务正常运行
2. 图片资源放置在 `public/images/` 目录下
3. 默认车辆图片为 `default-car.jpg`
4. 所有日期格式使用 `YYYY-MM-DD`
