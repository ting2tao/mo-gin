package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Middleware struct {
}

var MiddlewareSet = wire.NewSet(new(Middleware), "*")

func (m Middleware) RegisterMiddleware(router *gin.Engine) {

}
