package handlers

import (
	"fmt"
	"os"

	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

func GitClone(c *gin.Context) {
	rootDir := c.Query("root-dir")
	stack := c.Query("stack")
	newRepo := uuid.New().String()

	git.PlainClone("./tmp/"+newRepo, false, &git.CloneOptions{
		URL:      c.Query("git-link"),
		Progress: os.Stdout,
	})
	id, port, err := helpers.CreateContainer(rootDir, stack, newRepo)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating container",
			"error":   err.Error(),
		})
		return
	}

	_, err = helpers.CreateContainerData(newRepo, id, port)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating container",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(id, port)

	c.JSON(200, gin.H{
		"message": "Clone",
		"repo":    newRepo,
		"id":      id,
	})
}
