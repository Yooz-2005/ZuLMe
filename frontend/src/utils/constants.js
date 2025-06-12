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
