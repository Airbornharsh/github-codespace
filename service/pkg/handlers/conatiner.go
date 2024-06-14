package handlers

// import (
// 	"os/exec"
// 	"path/filepath"

// 	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
// 	"github.com/gin-gonic/gin"
// )

// func DeleteContainer(c *gin.Context) {
// 	imageId := c.Query("image-id")

// 	containersData, err := helpers.ReadContainersData()
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error reading containers data",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	containerInfo, ok := containersData[imageId]
// 	if !ok {
// 		c.JSON(404, gin.H{
// 			"message": "Container not found",
// 		})
// 		return
// 	}

// 	err = helpers.DeleteDockerContainer()(imageId+":latest", containerInfo.ContainerID)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error deleting container",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	delete(containersData, imageId)
// 	err = helpers.WriteFile(containersData)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error writing containers data",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	absPath, err := filepath.Abs("./tmp/" + imageId)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error getting absolute path",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	err = exec.Command("rm", "-r", absPath).Run()
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error deleting directory",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "Deleted",
// 	})
// }
