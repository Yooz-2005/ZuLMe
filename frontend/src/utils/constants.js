// API 基础配置
export const API_BASE_URL = 'http://localhost:8888'; // 根据您的后端端口调整

// 车辆类型常量
export const VEHICLE_TYPES = {
  ECONOMY: 'economy',
  COMFORT: 'comfort', 
  LUXURY: 'luxury',
  SUV: 'suv'
};

// 车辆类型中文映射
export const VEHICLE_TYPE_LABELS = {
  [VEHICLE_TYPES.ECONOMY]: '经济型',
  [VEHICLE_TYPES.COMFORT]: '舒适型',
  [VEHICLE_TYPES.LUXURY]: '豪华型',
  [VEHICLE_TYPES.SUV]: 'SUV'
};

// 车辆状态
export const VEHICLE_STATUS = {
  AVAILABLE: 'available',
  RENTED: 'rented',
  MAINTENANCE: 'maintenance'
};

// 车辆状态中文映射
export const VEHICLE_STATUS_LABELS = {
  [VEHICLE_STATUS.AVAILABLE]: '可租用',
  [VEHICLE_STATUS.RENTED]: '已租出',
  [VEHICLE_STATUS.MAINTENANCE]: '维护中'
};

// 分页默认配置
export const PAGINATION_CONFIG = {
  DEFAULT_PAGE_SIZE: 12,
  DEFAULT_PAGE: 1
};

// 订单状态
export const ORDER_STATUS = {
  PENDING: 1,    // 待支付
  PAID: 2,       // 已支付
  CANCELLED: 3,  // 已取消
  COMPLETED: 4,  // 已完成
  PICKED_UP: 5,  // 已取车
  RETURNED: 6    // 已还车
};

// 订单状态中文映射
export const ORDER_STATUS_LABELS = {
  [ORDER_STATUS.PENDING]: '待支付',
  [ORDER_STATUS.PAID]: '已支付',
  [ORDER_STATUS.CANCELLED]: '已取消',
  [ORDER_STATUS.COMPLETED]: '已完成',
  [ORDER_STATUS.PICKED_UP]: '已取车',
  [ORDER_STATUS.RETURNED]: '已还车'
};

// 支付状态
export const PAYMENT_STATUS = {
  PENDING: 1,    // 待支付
  PAID: 2,       // 已支付
  CANCELLED: 3,  // 已取消
  COMPLETED: 4   // 已完成
};

// 支付状态中文映射
export const PAYMENT_STATUS_LABELS = {
  [PAYMENT_STATUS.PENDING]: '待支付',
  [PAYMENT_STATUS.PAID]: '已支付',
  [PAYMENT_STATUS.CANCELLED]: '已取消',
  [PAYMENT_STATUS.COMPLETED]: '已完成'
};
