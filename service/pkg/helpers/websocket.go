package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Command struct {
	Dir     string `json:"dir"`
	Command string `json:"command"`
	Type    string `json:"type"`
	Data    string `json:"data"`
	IsFile  string `json:"isFile"`
}

type Output struct {
	OldDir  string `json:"oldDir"`
	Dir     string `json:"dir"`
	Out     string `json:"out"`
	Error   string `json:"error"`
	Type    string `json:"type"`
	IsFile  string `json:"isFile"`
	Command string `json:"command"`
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

	fmt.Println("WebSocket connection established")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return false
		}

		var command Command
		var output Output

		execCommand := string(p)
		json.Unmarshal([]byte(execCommand), &command)
		var cmd *exec.Cmd
		if command.Data != "" {
			err := WriteFileToContainer(containerInfo.ContainerID, imageId, filepath.Join(command.Dir, command.IsFile), command.Data)
			if err != nil {
				output.Error = fmt.Sprintf("Error writing file to container: %v\n", err)
			}
		} else {
			cmd = exec.Command("docker", "exec", containerInfo.ContainerID, "sh", "-c", "cd "+command.Dir+"&& "+command.Command)
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
			output.Type = command.Type
			if command.Type == "command" {
				pwd, err := getPwd(
					containerInfo.ContainerID,
					command.Dir,
					command.Command,
				)
				if err != nil {
					output.Error = fmt.Sprintf("Error getting pwd: %v\n", err)
				}
				p := strings.Split(pwd, "/app")
				if len(p) > 1 {
					output.Dir = "/app" + p[1]
				} else {
					output.Dir = "/"
				}
				output.OldDir = command.Dir
				output.Command = command.Command
			} else {
				output.Dir = command.Dir
			}
			output.IsFile = command.IsFile
			st, _ := json.Marshal(output)
			if err := conn.WriteMessage(messageType, st); err != nil {
				return false
			}
		}
	}
}

func getPwd(containerId string, dir string, cmd string) (string, error) {
	pwd, err := exec.Command("docker", "exec", containerId, "sh", "-c", "cd "+dir+"&& "+cmd+" && pwd").Output()
	return string(pwd), err
}
