package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		fmt.Println("Hello World")
		c.JSON(200, gin.H{
			"message": "Welcome to Code Space",
		})
	})
	r.Run(":5000")
	fmt.Println("Server Started at http://localhost:5000")
}
