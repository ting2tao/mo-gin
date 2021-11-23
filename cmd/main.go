package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slovty/mo-gin/pkg/route"
)

func main() {

	r := route.InitRouter()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongSct",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
