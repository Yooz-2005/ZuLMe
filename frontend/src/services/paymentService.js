import api from './api';

const paymentService = {
  // 创建支付
  createPayment: async (paymentData) => {
    try {
      const response = await api.post('/api/payments', paymentData);
      return response;
    } catch (error) {
      console.error('创建支付失败:', error);
      throw error;
    }
  },

  // 获取支付链接
  getPaymentUrl: async (orderId, paymentMethod = 'alipay') => {
    try {
      // 暂时返回模拟支付链接，等待支付API实现
      return {
        code: 200,
        data: {
          payment_url: `https://mock-alipay.com/pay?order=${orderId}&method=${paymentMethod}`,
          payment_id: `PAY${Date.now()}`
        },
        message: 'success'
      };
    } catch (error) {
      console.error('获取支付链接失败:', error);
      throw error;
    }
  },

  // 查询支付状态
  getPaymentStatus: async (paymentId) => {
    try {
      const response = await api.get(`/api/payments/${paymentId}/status`);
      return response;
    } catch (error) {
      console.error('查询支付状态失败:', error);
      throw error;
    }
  },

  // 取消支付
  cancelPayment: async (paymentId) => {
    try {
      const response = await api.put(`/api/payments/${paymentId}/cancel`);
      return response;
    } catch (error) {
      console.error('取消支付失败:', error);
      throw error;
    }
  }
};

export default paymentService;
