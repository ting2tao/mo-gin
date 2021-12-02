package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slovty/mo-gin/pkg/response"
	"net/http"
)

type UserCtl struct {
}

func (ctl UserCtl) Users(ctx *gin.Context) {
	response.Success(ctx, 0, "ok", gin.H{"hello": "users"})
}

func (ctl UserCtl) MyUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}

func (ctl UserCtl) AddUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"add": "add user 1"})
	fmt.Println("hello AddUser")
}
