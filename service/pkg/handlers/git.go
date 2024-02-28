package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

func GitClone(c *gin.Context) {
	newRepo := uuid.New().String()

	git.PlainClone("./tmp/"+newRepo, false, &git.CloneOptions{
		URL:      c.Query("git-link"),
		Progress: os.Stdout,
	})

	c.JSON(200, gin.H{
		"message": "Clone",
		"repo":    newRepo,
	})
}
