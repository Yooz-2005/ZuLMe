import { collectVehicle, getCollectVehicleList } from './api';

/**
 * 收藏服务
 */
class FavoriteService {
  /**
   * 收藏/取消收藏车辆
   * @param {number} vehicleId - 车辆ID
   * @returns {Promise} API响应
   */
  async toggleCollect(vehicleId) {
    try {
      const response = await collectVehicle(vehicleId);
      return response;
    } catch (error) {
      console.error('收藏操作失败:', error);
      throw error;
    }
  }

  /**
   * 获取用户收藏列表
   * @returns {Promise} 收藏列表
   */
  async getFavoriteList() {
    try {
      const response = await getCollectVehicleList();
      if (response.code === 200) {
        return response.data?.VehicleList || [];
      } else {
        throw new Error(response.message || '获取收藏列表失败');
      }
    } catch (error) {
      console.error('获取收藏列表失败:', error);
      throw error;
    }
  }

  /**
   * 检查车辆是否已收藏
   * @param {number} vehicleId - 车辆ID
   * @param {Array} favoriteList - 收藏列表
   * @returns {boolean} 是否已收藏
   */
  isVehicleFavorited(vehicleId, favoriteList) {
    if (!favoriteList || !Array.isArray(favoriteList)) {
      return false;
    }
    return favoriteList.some(item => item.VehicleId === vehicleId);
  }

  /**
   * 格式化收藏车辆数据
   * @param {Object} vehicle - 原始车辆数据
   * @returns {Object} 格式化后的车辆数据
   */
  formatFavoriteVehicle(vehicle) {
    return {
      id: vehicle.VehicleId,
      vehicleId: vehicle.VehicleId,
      name: vehicle.VehicleName,
      image: vehicle.Image,
      // 可以根据需要添加更多字段
    };
  }
}

// 创建单例实例
const favoriteService = new FavoriteService();

export default favoriteService;
