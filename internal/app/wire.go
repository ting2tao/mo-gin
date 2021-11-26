//go:build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/app/controller"
	"github.com/slovty/mo-gin/internal/app/router"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		controller.ControllerSet,
		InjectorSet,
		router.Set,
	)
	return new(Injector), nil, nil
}
