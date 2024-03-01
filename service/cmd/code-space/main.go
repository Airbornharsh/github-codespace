package main

import (
	"fmt"

	"github.com/airbornharsh/github-codespace/service/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Code Space",
		})
	})

	routes.Init(r)

	fmt.Println("Server Started at http://localhost:5000")
	r.Run(":5000")
}
