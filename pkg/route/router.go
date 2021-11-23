package route

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()
	router := gin.Default()

	return router
}
