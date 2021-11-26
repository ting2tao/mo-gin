package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/slovty/mo-gin/internal/app/controller"
)

type Router struct {
	UserCtl controller.UserCtl
}

var Set = wire.NewSet(wire.Struct(new(Router), "*"))

func (r Router) RegisterRouter(router *gin.Engine) {
	router.GET("/hello", r.UserCtl.Users)

	user := router.Group("user")
	{
		user.POST("/user", r.UserCtl.AddUser)
	}
}
