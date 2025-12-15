package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linghechen/go-starter/pkg/xerr"
)

// Standard Response Structure
// 约定大于配置：所有接口都返回这个结构
type Response struct {
	Code int         `json:"code"`           // 业务码 (0=成功, 非0=错误)
	Data interface{} `json:"data,omitempty"` // 数据
	Msg  string      `json:"msg"`            // 提示信息
}

// HandlerFunc 是我们定义的业务函数签名
// 类似于 FastAPI: 接收 Context 和 请求体，返回 响应体 和 错误
type HandlerFunc[Req any, Resp any] func(ctx context.Context, req *Req) (*Resp, xerr.Error)

// Wrap 是核心魔法函数
// 它将业务函数 handler 转换为 Gin 认识的 gin.HandlerFunc
func Wrap[Req any, Resp any](handler HandlerFunc[Req, Resp]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Req

		// 1. 自动绑定参数 (JSON, Query, Form 等)
		// ShouldBind 会根据 Content-Type 自动选择绑定器
		if err := c.ShouldBind(&req); err != nil {
			// 参数绑定失败 (比如 JSON 格式不对，或者字段类型错误)
			c.JSON(http.StatusBadRequest, Response{
				Code: 400,
				Msg:  "Invalid Request: " + err.Error(),
			})
			return
		}

		// 2. 这里的 ShouldBind 其实已经包含了 validator 的校验
		// 如果 struct tag 里有 binding:"required"，上面就会报错返回

		// 3. 调用业务逻辑
		// 我们把 Gin 的 Context 转换成标准 Context，方便透传 TraceID 或 Timeout
		resp, err := handler(c.Request.Context(), &req)

		if err != nil {
			// 4. 统一错误处理
			// 这里以后可以扩展：根据自定义 Error 类型判断是 400 还是 500
			c.JSON(http.StatusInternalServerError, Response{
				Code: err.GetCode(),
				Msg:  err.GetMsg(),
			})
			return
		}

		// 5. 统一成功响应
		c.JSON(http.StatusOK, Response{
			Code: xerr.OK,
			Msg:  "success",
			Data: resp,
		})
	}
}

// Empty 是一个空结构体，用于不需要参数或不需要返回值的场景
type Empty struct{}
