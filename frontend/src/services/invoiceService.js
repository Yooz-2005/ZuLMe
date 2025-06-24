import api from './api';

const invoiceService = {
  // 用户申请开发票（推荐使用）
  applyInvoice: async (orderId) => {
    return api.post('/invoice/apply', { order_id: orderId });
  },

  // 商家直接开发票（保留原接口）
  generateInvoice: async (orderId) => {
    return api.post('/merchant/invoice/generate', { orderID: orderId });
  }
};

export default invoiceService;