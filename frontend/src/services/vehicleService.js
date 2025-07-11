import api from './api';

// 车辆服务类
class VehicleService {
  
  // 获取车辆列表
  async getVehicleList(params = {}) {
    try {
      const response = await api.get('/vehicle/list', { params });
      return response;
    } catch (error) {
      throw new Error('获取车辆列表失败');
    }
  }

  // 获取车辆详情
  async getVehicleDetail(id) {
    try {
      const response = await api.get(`/vehicle/${id}`);
      return response;
    } catch (error) {
      throw new Error('获取车辆详情失败');
    }
  }

  // 使用ES搜索车辆（高级搜索）
  async searchVehiclesWithES(searchParams) {
    try {
      const params = {
        keyword: searchParams.keyword || searchParams.location,
        type_id: searchParams.carType,
        brand_id: searchParams.brandId,
        price_min: searchParams.priceMin,
        price_max: searchParams.priceMax,
        year_min: searchParams.yearMin,
        year_max: searchParams.yearMax,
        page: searchParams.page || 1,
        page_size: searchParams.pageSize || 12,
        status: searchParams.status !== undefined ? searchParams.status : 1
      };

      // 过滤掉空值
      Object.keys(params).forEach(key => {
        if (params[key] === undefined || params[key] === null || params[key] === '') {
          delete params[key];
        }
      });

      const response = await api.get('/vehicle/search', { params });
      return response;
    } catch (error) {
      console.error('ES搜索车辆失败:', error);
      throw new Error('ES搜索车辆失败');
    }
  }

  // 搜索车辆
  async searchVehicles(searchParams) {
    try {
      // 如果有日期范围，使用可用车辆接口
      const hasDateRange = searchParams.dates &&
                           searchParams.dates.length === 2 &&
                           searchParams.dates[0] &&
                           searchParams.dates[1];

      if (hasDateRange) {
        const params = {
          start_date: searchParams.dates[0].format('YYYY-MM-DD'),
          end_date: searchParams.dates[1].format('YYYY-MM-DD'),
          type_id: searchParams.carType,
          brand_id: searchParams.brandId,
          page: searchParams.page || 1,
          page_size: searchParams.pageSize || 12
        };

        // 过滤掉空值
        Object.keys(params).forEach(key => {
          if (params[key] === undefined || params[key] === null || params[key] === '') {
            delete params[key];
          }
        });

        const response = await api.post('/vehicle-inventory/available-vehicles', params);
        return response;
      } else {
        // 使用普通列表接口进行搜索
        const params = {
          keyword: searchParams.location,
          type_id: searchParams.carType,
          brand_id: searchParams.brandId,
          page: searchParams.page || 1,
          page_size: searchParams.pageSize || 12,
          status: 1 // 只搜索可用状态的车辆
        };

        // 过滤掉空值
        Object.keys(params).forEach(key => {
          if (params[key] === undefined || params[key] === null || params[key] === '') {
            delete params[key];
          }
        });

        const response = await api.get('/vehicle/list', { params });
        return response;
      }
    } catch (error) {
      console.error('搜索车辆失败:', error);
      throw new Error('搜索车辆失败');
    }
  }

  // 获取车辆类型列表
  async getVehicleTypes() {
    try {
      const response = await api.get('/vehicle-type/list');
      return response;
    } catch (error) {
      throw new Error('获取车辆类型失败');
    }
  }

  // 获取车辆类型详情
  async getVehicleTypeDetail(id) {
    try {
      const response = await api.get(`/vehicle-type/${id}`);
      return response;
    } catch (error) {
      throw new Error('获取车辆类型详情失败');
    }
  }

  // ==================== 车辆品牌相关方法 ====================

  // 获取车辆品牌列表
  async getBrandList(params = {}) {
    try {
      const response = await api.get('/vehicle-brand/list', { params });
      return response;
    } catch (error) {
      throw new Error('获取品牌列表失败');
    }
  }

  // 获取车辆品牌详情
  async getBrandDetail(id) {
    try {
      const response = await api.get(`/vehicle-brand/${id}`);
      return response;
    } catch (error) {
      throw new Error('获取品牌详情失败');
    }
  }

  // 创建车辆品牌
  async createBrand(data) {
    try {
      const response = await api.post('/vehicle-brand/create', data);
      return response;
    } catch (error) {
      throw new Error('创建品牌失败');
    }
  }

  // 更新车辆品牌
  async updateBrand(data) {
    try {
      const response = await api.put('/vehicle-brand/update', data);
      return response;
    } catch (error) {
      throw new Error('更新品牌失败');
    }
  }

  // 删除车辆品牌
  async deleteBrand(data) {
    try {
      const response = await api.post('/vehicle-brand/delete', data);
      return response;
    } catch (error) {
      throw new Error('删除品牌失败');
    }
  }

  // 获取热门品牌
  async getHotBrands(limit = 10) {
    try {
      const response = await api.get('/vehicle-brand/list', {
        params: {
          is_hot: 1,
          page_size: limit,
          status: 1
        }
      });
      return response;
    } catch (error) {
      throw new Error('获取热门品牌失败');
    }
  }

  // 根据品牌搜索车辆
  async searchVehiclesByBrand(brandId, params = {}) {
    try {
      const searchParams = {
        brand_id: brandId,
        page: params.page || 1,
        page_size: params.pageSize || 12,
        ...params
      };

      const response = await api.get('/vehicle/list', { params: searchParams });
      return response;
    } catch (error) {
      throw new Error('根据品牌搜索车辆失败');
    }
  }

  // ==================== 车辆库存管理方法 ====================

  // 检查车辆可用性
  async checkAvailability(params) {
    try {
      const response = await api.post('/vehicle-inventory/check-availability', params);
      return response;
    } catch (error) {
      throw new Error('检查车辆可用性失败');
    }
  }

  // ==================== 网点管理方法 ====================

  // 获取网点列表（商户列表）
  async getLocationList(params = {}) {
    try {
      const queryParams = {
        page: params.page || 1,
        page_size: params.pageSize || 100, // 获取所有网点
        status_filter: 1, // 只获取审核通过的商户
        ...params
      };

      const response = await api.get('/merchant/locations', { params: queryParams });
      return response;
    } catch (error) {
      console.error('获取网点列表失败:', error);
      throw new Error('获取网点列表失败');
    }
  }

  // 创建预订
  async createReservation(data) {
    try {
      const response = await api.post('/vehicle-inventory/reservation/create', data);
      return response;
    } catch (error) {
      throw new Error('创建预订失败');
    }
  }

  // 更新预订状态
  async updateReservationStatus(data) {
    try {
      const response = await api.post('/vehicle/update-reservation-status', data);
      return response;
    } catch (error) {
      throw new Error('更新预订状态失败');
    }
  }

  // 获取可用车辆列表
  async getAvailableVehicles(params) {
    try {
      const response = await api.post('/vehicle-inventory/available-vehicles', params);
      return response;
    } catch (error) {
      throw new Error('获取可用车辆失败');
    }
  }

  // 获取库存统计
  async getInventoryStats(params) {
    try {
      const response = await api.get('/vehicle/inventory-stats', { params });
      return response;
    } catch (error) {
      throw new Error('获取库存统计失败');
    }
  }

  // 获取车辆库存日历
  async getVehicleInventory(params) {
    try {
      const response = await api.get('/vehicle-inventory/calendar', { params });
      return response;
    } catch (error) {
      throw new Error('获取车辆库存日历失败');
    }
  }

  // 设置车辆维护状态
  async setMaintenanceStatus(data) {
    try {
      const response = await api.post('/vehicle/set-maintenance', data);
      return response;
    } catch (error) {
      throw new Error('设置维护状态失败');
    }
  }

  // 取消预订
  async cancelReservation(orderId) {
    try {
      const response = await api.post('/vehicle/cancel-reservation', { order_id: orderId });
      return response;
    } catch (error) {
      throw new Error('取消预订失败');
    }
  }

  // 完成租用
  async completeRental(orderId) {
    try {
      const response = await api.post('/vehicle/complete-rental', { order_id: orderId });
      return response;
    } catch (error) {
      throw new Error('完成租用失败');
    }
  }

  // 获取商家车辆库存概览
  async getMerchantInventoryOverview(merchantId) {
    try {
      const response = await api.get(`/vehicle/merchant-inventory/${merchantId}`);
      return response;
    } catch (error) {
      throw new Error('获取商家库存概览失败');
    }
  }
}

// 导出单例
export default new VehicleService();
