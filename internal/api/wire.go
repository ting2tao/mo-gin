//go:build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/api/controller"
	"github.com/slovty/mo-gin/internal/api/middleware"
	"github.com/slovty/mo-gin/internal/api/router"
	"github.com/slovty/mo-gin/pkg/config"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		config.InitViper,
		controller.ControllerSet,
		InjectorSet,
		middleware.MiddlewareSet,
		router.Set,
	)
	return new(Injector), nil, nil
}
