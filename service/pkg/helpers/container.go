package helpers

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"os/exec"
)

func CreateContainer(id string, name string, giturl string) (string, int, error) {
	imageName := os.Getenv("IMAGE_NAME")
	hostPort, err := GetPort(id, name, giturl)
	if err != nil {
		return "", 0, err
	}

	var containerName string
	prefix := "https://github.com/"
	suffix := ".git"

	if strings.HasPrefix(giturl, prefix) {
		containerName = strings.TrimPrefix(giturl, prefix)
	}

	if strings.HasSuffix(giturl, suffix) {
		containerName = strings.TrimSuffix(containerName, suffix)
	}

	containerName = strings.ReplaceAll(containerName, "/", "-")

	println(containerName)

	createCmd := exec.Command("docker", "create", "--name", containerName, "-p", strconv.Itoa(hostPort)+":8080", imageName)
	output, err := createCmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	containerId := string(bytes.TrimSpace(output))

	startCmd := exec.Command("docker", "start", containerId)
	_, err = startCmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	AddContainerId(containerId, id)
	command := fmt.Sprintf("cd /home/coder && mkdir -p Code && cd Code && git clone %s", giturl)
	cmd := exec.Command("docker", "exec", containerId, "sh", "-c", command)
	if err := cmd.Start(); err != nil {
		SetUnActive(id)
		println(err.Error())
		return "", 0, err
	}
	if err := cmd.Wait(); err != nil {
		SetUnActive(id)
		println(err.Error())
		return "", 0, err
	}

	// go func() {
	// 	cmd := exec.Command("docker", "kill", containerID).Run()
	// 	if cmd != nil {
	// 		fmt.Println("Error killing container", cmd)
	// 	}
	// }()

	return containerId, hostPort, nil
}

func DeleteDockerContainer(containerId string) error {
	err := exec.Command("docker", "stop", containerId).Run()
	if err != nil {
		return err
	}
	err = exec.Command("docker", "rm", containerId).Run()
	if err != nil {
		return err
	}
	RemoveContainer(containerId)
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
