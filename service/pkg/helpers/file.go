package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type ContainerInfo struct {
	ContainerID   string `json:"containerId"`
	ContainerName string `json:"containerName"`
	Port          int    `json:"port"`
}

type ContainerMap map[string]ContainerInfo

func ReadContainersData() (ContainerMap, error) {
	absContextDir, err := filepath.Abs("./data/containers.info.json")
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.ReadFile(absContextDir)
	if err != nil {
		return nil, err
	}

	var containerMap ContainerMap

	err = json.Unmarshal(jsonFile, &containerMap)
	if err != nil {
		return nil, err
	}
	return containerMap, nil
}

func WriteFile(data ContainerMap) error {
	absContextDir, err := filepath.Abs("./data/containers.info.json")
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(absContextDir, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateContainerData(imageName string, containerId string, port int) (string, error) {
	data, err := ReadContainersData()
	if err != nil {
		return "", err
	}

	data[imageName] = ContainerInfo{
		ContainerID:   containerId,
		ContainerName: imageName,
		Port:          port,
	}

	err = WriteFile(data)

	if err != nil {
		return "", err
	}

	return containerId, nil
}

func GetPort(port int) int {
	for {
		port++
		cmd := exec.Command("sudo", "lsof", "-i", "tcp:"+strconv.Itoa(port))
		err := cmd.Run()
		if err != nil {
			break
		}
		fmt.Println("Port in use", port)
	}

	return port
}
