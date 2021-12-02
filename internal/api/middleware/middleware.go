package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/slovty/mo-gin/pkg/response"
)

type Middleware struct {
}

var MiddlewareSet = wire.NewSet(wire.Struct(new(Middleware), "*"))

func (m Middleware) RegisterMiddleware(router *gin.Engine) {
	router.Use(response.ResponseHeaders())
}
