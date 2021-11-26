package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCtl struct {
}

func (ctl UserCtl) Users(ctx *gin.Context) {
	ctl.MyUsers(ctx)
	ctx.JSON(1, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}

func (ctl UserCtl) MyUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}

func (ctl UserCtl) AddUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"add": "add user 1"})
	fmt.Println("hello AddUser")
}
