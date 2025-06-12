import axios from 'axios';

// 創建一個 axios 實例
const instance = axios.create({
    baseURL: '/', // 您的 API 基礎 URL，如果與前端在同一域，可以留空或設置為 '/'
    timeout: 10000, // 請求超時時間
    headers: {
        'Content-Type': 'application/json',
    },
});

// 請求攔截器：在每個請求發送之前，檢查是否存在 token 並添加到請求頭
instance.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        console.log('Axios 請求攔截器中獲取的 token:', token); // 添加日誌
        if (token) {
            config.headers['X-Token'] = token; // 將 Authorization 改為 X-Token，並移除 Bearer 前綴
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 響應攔截器 (可選): 處理響應錯誤，例如 401 未授權
instance.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        if (error.response && error.response.status === 401) {
            // 處理 401 未授權錯誤，例如導航到登入頁面
            console.error('未授權，請重新登入。');
            // 可以根據您的路由庫進行跳轉，例如：
            // history.push('/login');
        }
        return Promise.reject(error);
    }
);

export default instance; 