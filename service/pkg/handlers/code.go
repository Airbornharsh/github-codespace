package handlers

import (
	"fmt"
	// "os/exec"
	"strings"

	// "github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type File struct {
	FilePath string `json:"filePath"`
	Data     string `json:"data"`
}

func CodeUpdate(c *gin.Context) {
	imageId := strings.Split(c.Request.Host, ".")[0]

	var files []File
	err := c.ShouldBindJSON(&files)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	containersData, err := helpers.ReadContainersData()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error reading containers data",
			"error":   err.Error(),
		})
		return
	}

	containerInfo, ok := containersData[imageId]
	if !ok {
		c.JSON(404, gin.H{
			"message": "Container not found",
		})
		return
	}

	fmt.Println(files)
	for _, file := range files {
		err = helpers.WriteFileToContainer(containerInfo.ContainerID, imageId, file.FilePath, file.Data)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error writing file to container",
				"error":   err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Updated",
	})
}
