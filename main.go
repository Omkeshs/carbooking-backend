package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")
	server := gin.Default()
	server.GET("/", func(c *gin.context) {
		c.JSON(200, gin.H{
			"msg":"hello from server"
		})
	})

	r.Run(":6006")
}
