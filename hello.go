package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"value": "pong"})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
