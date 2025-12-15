package domain

import (
	"github.com/linghechen/go-starter/internal/domain/hello"
	"go.uber.org/fx"
)

var Module = fx.Options(
	hello.Module,
)
