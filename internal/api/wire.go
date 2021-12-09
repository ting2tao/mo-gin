//go:build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/api/controller"
	"github.com/slovty/mo-gin/internal/api/middleware"
	"github.com/slovty/mo-gin/internal/api/router"
	"github.com/slovty/mo-gin/pkg/cache"
	"github.com/slovty/mo-gin/pkg/config"
	"github.com/slovty/mo-gin/pkg/db"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		config.InitViper,
		controller.ControllerSet,
		InjectorSet,
		middleware.MiddlewareSet,
		router.Set,
		db.DBSet,
		cache.InitRedis,
	)
	return new(Injector), nil, nil
}
