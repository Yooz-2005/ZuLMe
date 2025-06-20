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
      // 直接返回订单创建时生成的支付链接
      // 这里orderId实际上是从createOrderFromReservation返回的数据
      const orderData = orderId; // 这里传入的是整个订单数据

      return {
        code: 200,
        data: {
          payment_url: orderData.payment_url,
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
