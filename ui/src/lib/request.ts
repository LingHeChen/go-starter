// src/lib/request.ts
import { ofetch } from 'ofetch';
import type { ApiResponse } from './types';

// 1. 创建实例
const api = ofetch.create({
    baseURL: '/api', // 配合 Go 的反向代理，请求会被转发到后端

    // 请求拦截器
    onRequest({ options }) {
        // TODO: 这里以后可以从 localStorage 获取 token
        const token = localStorage.getItem('token');
        if (token) {
            // 1. 确保 headers 是一个 Headers 对象实例 (解决了类型兼容问题，也解决了展开为空的问题)
            options.headers = new Headers(options.headers);

            // 2. 使用标准的 set 方法设置 header
            options.headers.set('Authorization', `Bearer ${token}`);
        }
    },

    // 响应拦截器
    async onResponse({ response }) {
        // ofetch 会自动处理 JSON 解析
        // 如果 HTTP 状态码不是 2xx，ofetch 会自动抛错，我们在 onResponseError 处理

        // 这里主要处理 "HTTP 200 但业务失败 (code != 0)" 的情况
        if (response.ok && response._data) {
            const res = response._data as ApiResponse;

            // 约定：code !== 0 视为业务错误
            if (res.code !== 0) {
                // 可以在这里触发全局 Toast 报错，例如：toast.error(res.msg)
                console.error('Business Error:', res.msg);

                // 抛出自定义错误，打断后续逻辑
                throw new Error(res.msg || 'Unknown Business Error');
            }

            // 🟢 核心魔法：直接把 data 拿出来，替换掉原本的 response
            // 这样你在业务代码里拿到的直接就是 UserStruct，而不是 {code:0, data: UserStruct}
            response._data = res.data;
        }
    },

    // 错误处理 (网络错误、404、500 等)
    onResponseError({ response }) {
        console.error('HTTP Error:', response.status, response.statusText);
        // TODO: 触发全局错误提示
        // toast.error(`网络请求失败: ${response.status}`);
    }
});

// 2. 导出泛型封装的方法
// <T> 表示你期望返回的数据类型
export const request = {
    get: <T>(url: string, params?: any) => api<T>(url, { method: 'GET', query: params }),
    post: <T>(url: string, body?: any) => api<T>(url, { method: 'POST', body }),
    put: <T>(url: string, body?: any) => api<T>(url, { method: 'PUT', body }),
    delete: <T>(url: string, body?: any) => api<T>(url, { method: 'DELETE', body }),
};