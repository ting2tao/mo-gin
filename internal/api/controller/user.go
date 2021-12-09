package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slovty/mo-gin/pkg/cache"
	"github.com/slovty/mo-gin/pkg/config"
	"github.com/slovty/mo-gin/pkg/response"
	"gorm.io/gorm"
	"net/http"
)

type UserCtl struct {
	*config.Config
	*gorm.DB
	*cache.RedisClient
}

func (ctl UserCtl) Users(ctx *gin.Context) {
	var users []map[string]interface{}
	ctl.DB.Table("user").Find(&users)
	var s string
	var err error
	if s, err = ctl.RedisClient.Get("sct"); err != nil {
		fmt.Println(err.Error())
	}
	response.Success(ctx, 0, "ok", gin.H{
		"users": users,
		"sct":   ctl.Config.GetString("SCT"),
		"cache": s,
	})
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
