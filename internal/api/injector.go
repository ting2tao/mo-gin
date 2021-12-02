package api

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/api/middleware"
	"github.com/slovty/mo-gin/internal/api/router"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Router     *router.Router
	Middleware *middleware.Middleware
}
