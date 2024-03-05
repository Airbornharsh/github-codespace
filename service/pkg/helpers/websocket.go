package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Command struct {
	Dir     string `json:"dir"`
	Command string `json:"command"`
}

type Output struct {
	Out   string `json:"out"`
	Error string `json:"error"`
}

func StartWebSocket(c *gin.Context, upgrader *websocket.Upgrader) bool {
	imageId := strings.Split(c.Request.Host, ".")[0]

	containersData, err := ReadContainersData()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error reading containers data",
			"error":   err.Error(),
		})
		return false
	}

	containerInfo, ok := containersData[imageId]
	if !ok {
		c.JSON(404, gin.H{
			"message": "Container not found",
		})
		return false
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, http.Header{
		"Access-Control-Allow-Origin": []string{"*"},
	})
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return false
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return false
		}

		var command Command
		var output Output

		execCommand := string(p)
		json.Unmarshal([]byte(execCommand), &command)
		cmd := exec.Command("docker", "exec", containerInfo.ContainerID, "sh", "-c", "cd "+command.Dir+"&& "+command.Command)
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			output.Error = fmt.Sprintf("Error starting container: %v\n", err)
		}
		if err := cmd.Wait(); err != nil {
			output.Error = fmt.Sprintf("Error waiting for container: %v\n", err)
		}
		if stdout.Len() > 0 {
			output.Out = stdout.String()
		}
		
		st, _ := json.Marshal(output)
		if err := conn.WriteMessage(messageType, st); err != nil {
			return false
		}
	}
}
