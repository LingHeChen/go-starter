package hello

import (
	"github.com/linghechen/go-starter/pkg/framework/web"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// 只需要把构造函数扔进去，多少个都可以
	web.Controllers(
		NewHelloHandler,
	),

	// 如果有 Service 或 Repo，依然用普通的 fx.Provide
	// fx.Provide(NewDemoService),
)
