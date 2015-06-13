package gopoc

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", pong)

	return router
}

func pong(c *gin.Context) {
	c.JSON(200, Ping{"pong"})
}
