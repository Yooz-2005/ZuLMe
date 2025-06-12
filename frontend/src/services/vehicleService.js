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

  // 搜索车辆
  async searchVehicles(searchParams) {
    try {
      const params = {
        location: searchParams.location,
        start_date: searchParams.dates?.[0]?.format('YYYY-MM-DD'),
        end_date: searchParams.dates?.[1]?.format('YYYY-MM-DD'),
        vehicle_type: searchParams.carType,
        page: searchParams.page || 1,
        page_size: searchParams.pageSize || 12
      };
      
      // 过滤掉空值
      Object.keys(params).forEach(key => {
        if (params[key] === undefined || params[key] === null || params[key] === '') {
          delete params[key];
        }
      });

      const response = await api.get('/vehicle/list', { params });
      return response;
    } catch (error) {
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
}

// 导出单例
export default new VehicleService();
