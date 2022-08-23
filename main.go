package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run(":8080")
}
