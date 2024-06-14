package handlers

import (
	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GitClone(c *gin.Context) {
	id := uuid.New().String()
	name := c.Query("name")
	gitUrl := c.Query("gitUrl")

	if name == "" || gitUrl == "" {
		c.JSON(400, gin.H{
			"message": "Should Provide Name and Git Url",
			"error":   true,
		})
		return
	}

	// git.PlainClone("./tmp/"+newRepo, false, &git.CloneOptions{
	// 	URL:      ,
	// 	Progress: os.Stdout,
	// })
	conatinerId, port, err := helpers.CreateContainer(id, name, gitUrl)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating container",
			"error":   err.Error(),
		})
		return
	}

	println(conatinerId, port)

	c.JSON(200, gin.H{
		"message":     "Clone",
		"containerId": conatinerId,
		"id":          id,
		"port":        port,
	})
}
