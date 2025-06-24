import axios from 'axios';
import { API_BASE_URL } from '../utils/constants';

// 创建axios实例
const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 添加token认证信息
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['x-token'] = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    // 统一错误处理
    console.error('API Error:', error);
    
    if (error.response) {
      // 服务器返回错误状态码
      const { status, data } = error.response;
      switch (status) {
        case 401:
          // 未授权，可以跳转到登录页
          console.error('未授权访问');
          break;
        case 403:
          console.error('禁止访问');
          break;
        case 404:
          console.error('资源不存在');
          break;
        case 500:
          console.error('服务器内部错误');
          break;
        default:
          console.error('请求失败:', data?.message || '未知错误');
      }
    } else if (error.request) {
      // 网络错误
      console.error('网络连接失败');
    } else {
      console.error('请求配置错误');
    }
    
    return Promise.reject(error);
  }
);

// ==================== 收藏相关API ====================

// 收藏/取消收藏车辆
export const collectVehicle = async (vehicleId) => {
  try {
    const response = await api.post('/user/collect', {
      vehicle_id: vehicleId
    });
    return response;
  } catch (error) {
    console.error('收藏操作失败:', error);
    throw error;
  }
};

// 获取用户收藏列表
export const getCollectVehicleList = async () => {
  try {
    const response = await api.get('/user/collectList');
    return response;
  } catch (error) {
    console.error('获取收藏列表失败:', error);
    throw error;
  }
};

// ==================== 地理位置相关API ====================

// 根据地址获取经纬度坐标
export const getCoordinatesByAddress = async (address) => {
  try {
    const response = await api.post('/geocode/coordinates', {
      address: address
    });
    return response;
  } catch (error) {
    console.error('获取坐标失败:', error);
    throw error;
  }
};

// 计算用户到商家的距离
export const calculateDistance = async (userAddress, merchantId) => {
  try {
    const response = await api.post('/user/calculateDistance', {
      location: userAddress,
      merchant_id: merchantId
    });
    return response;
  } catch (error) {
    console.error('计算距离失败:', error);
    throw error;
  }
};

export default api;
