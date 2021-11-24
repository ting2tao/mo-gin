//go:build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/app/router"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InjectorSet,
		router.Set,
	)
	return new(Injector), nil, nil
}
