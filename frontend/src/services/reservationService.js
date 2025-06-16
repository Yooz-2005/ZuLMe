import api from './api';

const reservationService = {
  // 创建预订
  createReservation: async (reservationData) => {
    try {
      const response = await api.post('/vehicle-inventory/reservation/create', reservationData);
      return response;
    } catch (error) {
      console.error('创建预订失败:', error);
      throw error;
    }
  },

  // 获取用户的预订列表
  getUserReservations: async (params = {}) => {
    try {
      const response = await api.get('/vehicle-inventory/reservation/list', { params });
      return response;
    } catch (error) {
      console.error('获取预订列表失败:', error);
      throw error;
    }
  },

  // 获取预订详情
  getReservationDetail: async (reservationId) => {
    try {
      const response = await api.get(`/api/reservations/${reservationId}`);
      return response;
    } catch (error) {
      console.error('获取预订详情失败:', error);
      throw error;
    }
  },

  // 取消预订
  cancelReservation: async (reservationId) => {
    try {
      const response = await api.put(`/api/reservations/${reservationId}/cancel`);
      return response;
    } catch (error) {
      console.error('取消预订失败:', error);
      throw error;
    }
  },

  // 更新预订状态
  updateReservationStatus: async (reservationId, status) => {
    try {
      const response = await api.put(`/api/reservations/${reservationId}/status`, { status });
      return response;
    } catch (error) {
      console.error('更新预订状态失败:', error);
      throw error;
    }
  }
};

export default reservationService;
