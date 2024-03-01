package handlers

import (
	"fmt"

	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func DeleteContainer(c *gin.Context) {
	imageId := c.Query("image-id")

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

	err = helpers.DeleteDockerImageAndContainer(imageId+":latest", containerInfo.ContainerID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting container",
			"error":   err.Error(),
		})
		return
	}

	delete(containersData, imageId)
	err = helpers.WriteFile(containersData)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error writing containers data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleted",
	})
}

func GetContainer(c *gin.Context) {
	imageId := c.Query("image-id")

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
	
	cPath := c.Param("path")
	if cPath == "" {
		cPath = "/"
	}

	targetURL := fmt.Sprintf("http://localhost:%d%s", containerInfo.Port, cPath)

	fmt.Println(targetURL)

	proxy := helpers.NewReverseProxy(targetURL)

	proxy.ServeHTTP(c.Writer, c.Request)

	// c.JSON(200, gin.H{
	// 	"message": "Container Found",
	// 	"port":    containerInfo.Port,
	// })
}
