package helpers

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"os/exec"
)

func CreateContainer(rootDir string, stack string, imageName string) (string, int, error) {
	codeLocation := ""
	hostPort := GetPort(8080)

	if rootDir == "" {
		codeLocation = "./tmp/" + imageName
	} else {
		codeLocation = "./tmp/" + imageName + "/" + rootDir
	}

	image := imageName + ":latest"

	err := exec.Command("cp", "./pkg/dockerFiles/"+stack+"/Dockerfile", codeLocation).Run()
	if err != nil {
		return "", 0, err
	}

	err = exec.Command("docker", "build", "-t", image, codeLocation).Run()
	if err != nil {
		fmt.Println("building ", err)
		return "", 0, err
	}

	createCmd := exec.Command("docker", "create", "--name", imageName, "-p", strconv.Itoa(hostPort)+":3000", image)
	output, err := createCmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	containerID := string(bytes.TrimSpace(output))

	startCmd := exec.Command("docker", "start", containerID)
	_, err = startCmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	go func() {
		cmd := exec.Command("docker", "kill", containerID).Run()
		if cmd != nil {
			fmt.Println("Error killing container", cmd)
		}
	}()

	return containerID, hostPort, nil
}

func DeleteDockerImageAndContainer(imageId string, containerId string) error {
	err := exec.Command("docker", "stop", containerId).Run()
	if err != nil {
		return err
	}

	err = exec.Command("docker", "rm", containerId).Run()
	if err != nil {
		return err
	}
	err = exec.Command("docker", "rmi", imageId).Run()
	if err != nil {
		return err
	}
	return nil
}

func WriteFileToContainer(containerId string, imageId string, filePath string, data string) error {
	newFilePath := "./fileTmp/" + imageId + filePath
	elements := strings.Split(newFilePath, "/")

	if len(elements) > 0 {
		lastIndex := len(elements) - 1
		elements = elements[:lastIndex]
	}

	dirPath := strings.Join(elements, "/")

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory", err)
		return err
	}

	err = os.WriteFile(newFilePath, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return err
	}

	err = exec.Command("docker", "cp", newFilePath, containerId+":"+filePath).Run()

	newErr := os.Remove(newFilePath)
	if newErr != nil {
		return newErr
	}

	if err != nil {
		return err
	}
	return nil
}
