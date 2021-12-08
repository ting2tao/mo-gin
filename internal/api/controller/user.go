package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slovty/mo-gin/pkg/config"
	"github.com/slovty/mo-gin/pkg/response"
	"net/http"
)

type UserCtl struct {
	*config.Config
}

func (ctl UserCtl) Users(ctx *gin.Context) {

	uri := ctl.Config.GetString("MYSQL_URI")
	fmt.Println(uri)

	response.Success(ctx, 0, "ok", gin.H{"hello": "users111", "sct": ctl.Config.GetString("SCT")})
}

func (ctl UserCtl) MyUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "users"})
	fmt.Println("hello Users")
}

func (ctl UserCtl) AddUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"add": "add user 1"})
	fmt.Println("hello AddUser")
}

func (ctl UserCtl) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"update": "update user 1"})
	fmt.Println("hello AddUser")
}
