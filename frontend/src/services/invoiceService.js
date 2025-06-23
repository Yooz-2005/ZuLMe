import api from './api';

const invoiceService = {
  // 申请开发票
  generateInvoice: async (orderId) => {
    return api.post('/invoice/generate', { orderID: orderId });
  }
};

export default invoiceService; 