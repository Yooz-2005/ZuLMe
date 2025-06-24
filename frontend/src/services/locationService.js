import api from './api';

/**
 * 地理位置服务
 */
class LocationService {
  /**
   * 根据地址获取经纬度坐标
   * @param {string} address - 地址字符串
   * @returns {Promise<Object>} 坐标信息
   */
  async getCoordinatesByAddress(address) {
    try {
      const response = await api.post('/geocode/coordinates', {
        address: address
      });

      if (response.code === 200) {
        return {
          success: true,
          data: {
            address: response.address,
            longitude: response.longitude,
            latitude: response.latitude
          }
        };
      } else {
        return {
          success: false,
          message: response.message || '地址解析失败'
        };
      }
    } catch (error) {
      console.error('获取坐标失败:', error);
      return {
        success: false,
        message: '网络请求失败'
      };
    }
  }

  /**
   * 计算用户地址到商家的距离
   * @param {string} userAddress - 用户地址
   * @param {number} merchantId - 商家ID
   * @returns {Promise<Object>} 距离信息
   */
  async calculateDistance(userAddress, merchantId) {
    try {
      const response = await api.post('/user/calculateDistance', {
        location: userAddress,
        merchant_id: merchantId
      });

      if (response.code === 200) {
        return {
          success: true,
          data: {
            distance: response.data.distance,
            distanceMeters: response.data.distance_meters,
            merchantId: merchantId
          }
        };
      } else {
        return {
          success: false,
          message: response.message || '距离计算失败'
        };
      }
    } catch (error) {
      console.error('计算距离失败:', error);
      return {
        success: false,
        message: '距离计算请求失败'
      };
    }
  }

  /**
   * 获取用户附近的商家
   * @param {string} userAddress - 用户地址
   * @param {number} radiusKm - 搜索半径(公里)
   * @param {number} limit - 结果数量限制
   * @returns {Promise<Object>} 附近商家列表
   */
  async getNearbyMerchants(userAddress, radiusKm = 50, limit = 10) {
    try {
      // 首先获取用户地址的坐标
      const coordsResult = await this.getCoordinatesByAddress(userAddress);
      if (!coordsResult.success) {
        return coordsResult;
      }

      // 这里可以调用后端的附近商家查询接口
      // 目前项目中可能还没有这个接口，可以根据需要添加
      const response = await api.post('/merchant/nearby', {
        longitude: coordsResult.data.longitude,
        latitude: coordsResult.data.latitude,
        radius_km: radiusKm,
        limit: limit
      });

      if (response.code === 200) {
        return {
          success: true,
          data: response.data
        };
      } else {
        return {
          success: false,
          message: response.message || '查询附近商家失败'
        };
      }
    } catch (error) {
      console.error('查询附近商家失败:', error);
      return {
        success: false,
        message: '查询附近商家请求失败'
      };
    }
  }

  /**
   * 获取浏览器当前位置
   * @returns {Promise<Object>} 当前位置坐标
   */
  async getCurrentPosition() {
    return new Promise((resolve, reject) => {
      if (!navigator.geolocation) {
        reject(new Error('浏览器不支持地理定位'));
        return;
      }

      navigator.geolocation.getCurrentPosition(
        (position) => {
          resolve({
            success: true,
            data: {
              longitude: position.coords.longitude,
              latitude: position.coords.latitude,
              accuracy: position.coords.accuracy
            }
          });
        },
        (error) => {
          let message = '获取位置失败';
          switch (error.code) {
            case error.PERMISSION_DENIED:
              message = '用户拒绝了地理定位请求';
              break;
            case error.POSITION_UNAVAILABLE:
              message = '位置信息不可用';
              break;
            case error.TIMEOUT:
              message = '获取位置超时';
              break;
          }
          reject(new Error(message));
        },
        {
          enableHighAccuracy: true,
          timeout: 10000,
          maximumAge: 300000
        }
      );
    });
  }

  /**
   * 验证坐标是否有效
   * @param {number} longitude - 经度
   * @param {number} latitude - 纬度
   * @returns {boolean} 是否有效
   */
  validateCoordinates(longitude, latitude) {
    // 中国境内经纬度范围验证
    if (longitude < 73.33 || longitude > 135.05) {
      return false;
    }
    if (latitude < 3.51 || latitude > 53.33) {
      return false;
    }
    return true;
  }

  /**
   * 格式化距离显示
   * @param {number} distanceMeters - 距离(米)
   * @returns {string} 格式化后的距离字符串
   */
  formatDistance(distanceMeters) {
    if (distanceMeters < 1000) {
      return `${Math.round(distanceMeters)}米`;
    } else {
      return `${(distanceMeters / 1000).toFixed(2)}公里`;
    }
  }

  /**
   * 计算两点之间的直线距离 (Haversine公式)
   * @param {number} lat1 - 点1纬度
   * @param {number} lon1 - 点1经度
   * @param {number} lat2 - 点2纬度
   * @param {number} lon2 - 点2经度
   * @returns {number} 距离(米)
   */
  calculateHaversineDistance(lat1, lon1, lat2, lon2) {
    const R = 6371000; // 地球半径(米)
    const dLat = this.toRadians(lat2 - lat1);
    const dLon = this.toRadians(lon2 - lon1);
    const a = 
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(this.toRadians(lat1)) * Math.cos(this.toRadians(lat2)) *
      Math.sin(dLon / 2) * Math.sin(dLon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    return R * c;
  }

  /**
   * 角度转弧度
   * @param {number} degrees - 角度
   * @returns {number} 弧度
   */
  toRadians(degrees) {
    return degrees * (Math.PI / 180);
  }
}

// 创建单例实例
const locationService = new LocationService();

export default locationService;
