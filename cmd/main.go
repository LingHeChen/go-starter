package main

import (

	// 引入我们定义的包
	"github.com/linghechen/go-starter/internal/domain"
	"github.com/linghechen/go-starter/internal/server"
	"github.com/linghechen/go-starter/pkg/framework/boot"
)

func main() {
	boot.Main(
		// 1. 引入 HTTP Server 模块
		server.Module,
		domain.Module,
	)
}
