package api

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/api/middleware"
	"github.com/slovty/mo-gin/internal/api/router"
	"github.com/slovty/mo-gin/pkg/config"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Router     *router.Router
	Middleware *middleware.Middleware
	Config     *config.Config
}
