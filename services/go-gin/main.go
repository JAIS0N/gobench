package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "service": "go-gin"})
	})

	r.GET("/compute", func(c *gin.Context) {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		c.JSON(http.StatusOK, gin.H{"result": sum, "service": "go-gin"})
	})

	r.Run(":8080")
}