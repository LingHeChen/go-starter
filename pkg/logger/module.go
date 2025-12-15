package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module 暴露给 Fx 的模块
// 只要引入它，Fx 就知道如何创建 *zap.Logger 了
var Module = fx.Provide(NewLogger)

// NewLogger 构造函数
func NewLogger() (*zap.Logger, error) {
	// 开发阶段先用 NewDevelopment，日志会有颜色，且是人类可读格式
	// 生产环境以后我们改成 NewProduction (JSON格式)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	return config.Build()
}
