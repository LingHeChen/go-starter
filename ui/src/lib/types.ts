// src/lib/types.ts

// 对应 Go 后端的 Response 结构体
export interface ApiResponse<T = any> {
    code: number; // 0 代表成功
    msg: string;  // 错误信息
    data: T;      // 具体的业务数据
}