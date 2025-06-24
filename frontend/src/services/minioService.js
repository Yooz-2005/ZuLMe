import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8888';

const minioService = {
  // 获取预签名上传URL（公共接口，不需要认证）
  getPresignedUrl: async (params) => {
    console.log('MinIO服务请求参数:', params);
    console.log('API_BASE_URL:', API_BASE_URL);

    try {
      const requestUrl = `${API_BASE_URL}/minio/public/presigned-url`;
      console.log('请求URL:', requestUrl);

      const response = await axios.get(requestUrl, {
        params: {
          bucket: params.bucket,
          object_name: params.objectName,
          expires: params.expires || 3600
        },
        headers: {
          'Content-Type': 'application/json'
        }
      });

      console.log('MinIO API响应:', response.data);

      return {
        success: true,
        data: response.data.data
      };
    } catch (error) {
      console.error('获取预签名URL失败详细信息:', error);
      console.error('错误响应:', error.response?.data);
      console.error('错误状态:', error.response?.status);
      return {
        success: false,
        message: error.response?.data?.message || '获取上传链接失败'
      };
    }
  },

  // 获取预签名上传URL（需要认证）
  getPresignedUrlAuth: async (params) => {
    console.log('MinIO服务请求参数:', params);
    console.log('API_BASE_URL:', API_BASE_URL);

    try {
      const token = localStorage.getItem('token');
      console.log('使用的token:', token ? '已设置' : '未设置');

      const requestUrl = `${API_BASE_URL}/minio/presigned-url`;
      console.log('请求URL:', requestUrl);

      const response = await axios.get(requestUrl, {
        params: {
          bucket: params.bucket,
          object_name: params.objectName,
          expires: params.expires || 3600
        },
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      });

      console.log('MinIO API响应:', response.data);

      return {
        success: true,
        data: response.data.data
      };
    } catch (error) {
      console.error('获取预签名URL失败详细信息:', error);
      console.error('错误响应:', error.response?.data);
      console.error('错误状态:', error.response?.status);
      return {
        success: false,
        message: error.response?.data?.message || '获取上传链接失败'
      };
    }
  }
};

export default minioService;
