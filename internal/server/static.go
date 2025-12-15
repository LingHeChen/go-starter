package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/linghechen/go-starter/ui" // 引入刚才创建的 ui 包
)

// RegisterStatic 注册静态资源代理/服务
func RegisterStatic(r *gin.Engine) {
	// 判断是否为开发模式 (这里简单通过环境变量判断，你可以根据自己需求改)
	// 比如: APP_ENV=dev
	isDev := os.Getenv("APP_ENV") == "dev"

	if isDev {
		setupDevProxy(r)
	} else {
		setupProdServer(r)
	}
}

// 开发模式：反向代理到 Vite
func setupDevProxy(r *gin.Engine) {
	target := "http://localhost:5173"
	u, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(u)

	// 修改 Director 以处理 Host 头，避免某些情况下 Vite 报错
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = u.Host
	}

	// 捕获所有未定义的路由 (NoRoute)
	r.NoRoute(func(c *gin.Context) {
		// 如果是 API 请求但 404 了，直接返回 JSON 错误，不要返回 HTML
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"code": 404, "msg": "API not found"})
			return
		}

		fmt.Printf("🔄 Proxying %s to Vite...\n", c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

// 生产模式：使用 Embed 文件系统
func setupProdServer(r *gin.Engine) {
	distFS := http.FS(ui.GetDistFS())

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"code": 404, "msg": "API not found"})
			return
		}

		// 尝试查找文件
		path := c.Request.URL.Path
		// 默认行为：Gin 的 StaticFile 逻辑比较简单，这里我们自己处理 SPA 的 Fallback
		// 如果文件存在(比如 /assets/logo.png)，就直接返回
		// 如果文件不存在(比如 /user/1)，就返回 index.html
		c.FileFromFS(path, distFS)
	})
}
