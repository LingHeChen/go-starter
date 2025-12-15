package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linghechen/go-starter/pkg/framework/web"
	"go.uber.org/fx"
)

// Module Fx 模块
var Module = fx.Options(
	fx.Provide(NewGinEngine),
	fx.Invoke(StartServer),
)

// Params 用于接收依赖注入参数
type Params struct {
	fx.In

	// 魔法时刻：Fx 会自动把所有标记为 group:"controllers" 的实例注入到这个切片里
	Controllers []web.Controller `group:"controllers"`
}

func NewGinEngine(p Params) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// 自动循环注册所有路由
	for _, c := range p.Controllers {
		c.Registry(r)
	}

	RegisterStatic(r)

	return r
}

// StartServer 启动服务 (支持端口自动递增)
func StartServer(lc fx.Lifecycle, r *gin.Engine) {
	// 1. 寻找可用端口
	listener, port, err := findAvailablePort(8080)
	if err != nil {
		// 如果连端口都找不到，直接 Panic 终止应用
		panic(fmt.Errorf("failed to find available port: %w", err))
	}

	// 2. 创建 HTTP Server
	srv := &http.Server{
		Handler: r,
	}

	// 3. 注册生命周期
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("\n🚀 HTTP Server is running on: http://localhost:%d\n\n", port)

			go func() {
				// 注意：这里用 Serve 而不是 ListenAndServe
				// 因为我们已经手动拿到了 Listener
				if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
					fmt.Printf("❌ Server startup failed: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("🛑 Shutting down HTTP Server...")
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		},
	})
}

// findAvailablePort 尝试寻找可用端口
// startPort: 起始端口 (例如 8080)
// maxAttempts: 最大尝试次数 (默认尝试 10 个端口)
func findAvailablePort(startPort int) (net.Listener, int, error) {
	for port := startPort; port < startPort+10; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			// 成功抢占到端口
			return listener, port, nil
		}
		// 如果是被占用 (bind error)，则继续循环尝试下一个
		fmt.Printf("⚠️  Port %d is in use, trying %d...\n", port, port+1)
	}
	return nil, 0, fmt.Errorf("no available ports found between %d and %d", startPort, startPort+10)
}
