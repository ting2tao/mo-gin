package app

import (
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/app/router"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Router *router.Router
}
