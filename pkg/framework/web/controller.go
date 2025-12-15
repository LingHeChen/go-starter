package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Controller 接口
// 任何实现了此接口的 Struct，都可以定义自己的路由规则
type Controller interface {
	// Registry 方法里写：r.GET("/users", web.Wrap(h.GetUsers))
	Registry(r *gin.Engine)
}

// Controllers 是核心封装
// 它接收一组构造函数 (ctors...)，把它们都标记为 controller 组，并注册到 Fx
func Controllers(ctors ...any) fx.Option {
	var opts []fx.Option

	for _, ctor := range ctors {
		opts = append(opts, fx.Provide(
			fx.Annotate(
				ctor,
				fx.As(new(Controller)),               // 自动断言接口
				fx.ResultTags(`group:"controllers"`), // 自动打标签
			),
		))
	}

	return fx.Options(opts...)
}
