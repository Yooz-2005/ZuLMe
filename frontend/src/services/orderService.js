import api from './api';

const orderService = {
  // 从预订创建订单
  createOrderFromReservation: async (orderData) => {
    try {
      // 如果传入的是数字，说明是旧的调用方式，只传了reservationId
      if (typeof orderData === 'number') {
        orderData = { reservation_id: orderData };
      }

      const response = await api.post('/order/create-from-reservation', orderData);
      return response;
    } catch (error) {
      console.error('从预订创建订单失败:', error);
      throw error;
    }
  },

  // 获取用户的订单列表
  getUserOrders: async (params = {}) => {
    try {
      const response = await api.get('/order/list', { params });
      return response;
    } catch (error) {
      console.error('获取订单列表失败:', error);
      throw error;
    }
  },

  // 获取订单详情
  getOrderDetail: async (orderId) => {
    try {
      const response = await api.get(`/order/detail/${orderId}`);
      return response;
    } catch (error) {
      console.error('获取订单详情失败:', error);
      throw error;
    }
  },

  // 根据订单号获取订单详情
  getOrderDetailBySn: async (orderSn) => {
    try {
      const response = await api.get(`/order/detail-by-sn/${orderSn}`);
      return response;
    } catch (error) {
      console.error('根据订单号获取订单详情失败:', error);
      throw error;
    }
  },

  // 取消订单
  cancelOrder: async (orderId) => {
    try {
      const response = await api.put(`/api/orders/${orderId}/cancel`);
      return response;
    } catch (error) {
      console.error('取消订单失败:', error);
      throw error;
    }
  },

  // 更新订单状态
  updateOrderStatus: async (orderId, status) => {
    try {
      const response = await api.put(`/api/orders/${orderId}/status`, { status });
      return response;
    } catch (error) {
      console.error('更新订单状态失败:', error);
      throw error;
    }
  }
};

export default orderService;
