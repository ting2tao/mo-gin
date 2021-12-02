//go:build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/api/controller"
	"github.com/slovty/mo-gin/internal/api/middleware"
	"github.com/slovty/mo-gin/internal/api/router"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		controller.ControllerSet,
		InjectorSet,
		middleware.MiddlewareSet,
		router.Set,
	)
	return new(Injector), nil, nil
}
