package helpers

import (
	"bytes"
	"os"
	"strconv"

	"os/exec"
	"path/filepath"
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

	absDockerfilePath, err := filepath.Abs("./pkg/dockerFiles/React/Dockerfile")
	if err != nil {
		return "", 0, err
	}

	err = exec.Command("cp", absDockerfilePath, codeLocation).Run()
	if err != nil {
		return "", 0, err
	}

	err = exec.Command("docker", "build", "-t", image, codeLocation).Run()
	if err != nil {
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

	err := os.WriteFile(newFilePath, []byte(data), 0644)
	if err != nil {
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
