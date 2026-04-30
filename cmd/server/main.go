package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/doctors", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.Run(":8080")
}
