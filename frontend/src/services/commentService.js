import api from './api';

const commentService = {
  // 创建评论
  createComment: async (commentData) => {
    try {
      const response = await api.post('/comment/create', commentData);
      return response;
    } catch (error) {
      console.error('创建评论失败:', error);
      throw error;
    }
  },

  // 检查订单是否已评论
  checkOrderCommented: async (orderId) => {
    try {
      const response = await api.get(`/comment/check/${orderId}`);
      return response;
    } catch (error) {
      console.error('检查订单评论状态失败:', error);
      throw error;
    }
  },

  // 获取订单评论
  getOrderComment: async (orderId) => {
    try {
      const response = await api.get(`/comment/order/${orderId}`);
      return response;
    } catch (error) {
      console.error('获取订单评论失败:', error);
      throw error;
    }
  },

  // 获取车辆评论列表
  getVehicleComments: async (vehicleId, params = {}) => {
    try {
      const response = await api.get(`/comment/vehicle/${vehicleId}`, { params });
      return response;
    } catch (error) {
      console.error('获取车辆评论失败:', error);
      throw error;
    }
  },

  // 获取车辆评论统计
  getVehicleStats: async (vehicleId) => {
    try {
      const response = await api.get(`/comment/stats/${vehicleId}`);
      return response;
    } catch (error) {
      console.error('获取车辆评论统计失败:', error);
      throw error;
    }
  },

  // 更新评论
  updateComment: async (commentId, commentData) => {
    try {
      const response = await api.put(`/comment/${commentId}`, commentData);
      return response;
    } catch (error) {
      console.error('更新评论失败:', error);
      throw error;
    }
  },

  // 删除评论
  deleteComment: async (commentId) => {
    try {
      const response = await api.delete(`/comment/${commentId}`);
      return response;
    } catch (error) {
      console.error('删除评论失败:', error);
      throw error;
    }
  },

  // 获取用户评论列表
  getUserComments: async (params = {}) => {
    try {
      const response = await api.get('/comment/user/my-comments', { params });
      return response;
    } catch (error) {
      console.error('获取用户评论失败:', error);
      throw error;
    }
  }
};

export default commentService;
