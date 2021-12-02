package controller

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	wire.Struct(new(UserCtl), "*"),
)
