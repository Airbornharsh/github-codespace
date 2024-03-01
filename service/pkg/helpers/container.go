package helpers

import (
	"bytes"
	"strconv"

	// "context"
	// "fmt"
	// "io"
	// "os"
	"os/exec"
	"path/filepath"
	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/client"
)

// func buildDockerImage(contextDir, dockerfilePath, imageName string) error {
// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithVersion("1.43"))
// 	if err != nil {
// 		return err
// 	}

// 	absContextDir, err := filepath.Abs(contextDir)
// 	if err != nil {
// 		fmt.Println("Error getting absolute path")
// 		return err
// 	}

// 	fmt.Println(absContextDir)

// 	buildContext, err := os.Open(absContextDir)
// 	if err != nil {
// 		fmt.Println("Error opening context dir")
// 		return err
// 	}
// 	defer buildContext.Close()

// 	absDockerfilePath, err := filepath.Abs(dockerfilePath)
// 	if err != nil {
// 		fmt.Println("Error getting absolute path")
// 		return err
// 	}

// 	fmt.Println(absDockerfilePath)

// 	buildOptions := types.ImageBuildOptions{
// 		Dockerfile: filepath.ToSlash(absDockerfilePath),
// 		Tags:       []string{imageName},
// 	}

// 	buildResponse, err := cli.ImageBuild(context.Background(), buildContext, buildOptions)
// 	if err != nil {
// 		fmt.Println("Error building image")
// 		return err
// 	}
// 	defer buildResponse.Body.Close()

// 	_, err = io.Copy(os.Stdout, buildResponse.Body)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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
