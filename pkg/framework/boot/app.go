package boot

import (
	"github.com/linghechen/go-starter/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Main 是整个应用的入口函数
// modules: 传入所有的业务模块 (Database, UserDomain, OrderDomain...)
func Main(modules ...fx.Option) {
	// 基础模块：日志、配置等基础设施
	baseModules := []fx.Option{
		// 这里稍后会加入 ConfigModule, LoggerModule
		// 目前先用 Fx 自带的 logger 演示
		logger.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	}

	// 合并基础模块和用户模块
	allModules := append(baseModules, modules...)

	// 创建并运行 Fx 应用
	app := fx.New(allModules...)

	// 启动应用
	app.Run()
}
