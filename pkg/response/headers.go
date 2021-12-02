package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ResponseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("server-time", time.Now().Format("2006-01-02 15:04:05"))
		c.Header("api-env", "test-sct")
		c.Header("Access-Control-Expose-Headers", "server-time,api-env")
		c.Next()
	}
}
