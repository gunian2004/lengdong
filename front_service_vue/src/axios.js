import axios from 'axios';
import { serviceEndpoints } from '@/services';

const http = axios.create();

http.interceptors.request.use(config => {
    const token = sessionStorage.getItem('token');
    if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
    }

    // 根据请求的路径动态选择对应的 baseURL
    const path = config.url.split('/')[1]; // 假设路径的第一部分是服务名称
    if (serviceEndpoints[path]) {
        config.baseURL = serviceEndpoints[path];
    } else {
        config.baseURL = serviceEndpoints.auth; // 默认使用认证服务
    }

    return config;
});

http.interceptors.response.use(response => {
    return response;
}, error => {
    if (error.response && error.response.status === 401) {
        ElMessage.error('登录已过期，请重新登录');
        sessionStorage.removeItem('token');
        sessionStorage.removeItem('role_type');
        window.location.href = '/'; // 跳转到登录页面
    }
    return Promise.reject(error);
});

export default http;