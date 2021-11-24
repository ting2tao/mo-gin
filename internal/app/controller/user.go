package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCtl struct {
}

func (ctl UserCtl) Users(ctx *gin.Context) {
	ctl.MyUsers(ctx)
	ctx.JSON(1, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}

func (ctl UserCtl) MyUsers(ctx *gin.Context) {
	ctx.JSON(1, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}
