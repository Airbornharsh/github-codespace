package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strconv"
	"time"

	"path/filepath"
)

type ContainerInfo struct {
	Id          string `json:"id"`
	ContainerID string `json:"containerId"`
	Name        string `json:"name"`
	Port        int    `json:"port"`
	Active      bool   `json:"active"`
	GitUrl      string `json:"gitUrl"`
}

type DATA struct {
	Currentport int             `json:"currentPort"`
	MaxActive   int             `json:"maxActive"`
	Active      int             `json:"active"`
	List        []ContainerInfo `json:"list"`
	Ids         []string        `json:"ids"`
}

var LOCKED = false

func GetData() (DATA, error) {
	var err error
	if LOCKED {
		try := 10
		for try > 0 && !LOCKED {
			println("Waiting to be unlocked", 10-try)
			time.Sleep(time.Second * 5)
			try--
		}
		if !LOCKED {
			return DATA{}, err
		}
	}

	absContextDir, err := filepath.Abs("./data/containers.info.json")
	if err != nil {
		return DATA{}, err
	}

	jsonFile, err := os.ReadFile(absContextDir)
	if err != nil {
		return DATA{}, err
	}

	var data DATA

	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return DATA{}, err
	}
	return data, nil
}

func ReadContainersData(id string) (ContainerInfo, error) {
	data, err := GetData()
	if err != nil {
		return ContainerInfo{}, err
	}
	LOCKED = true
	for _, container := range data.List {
		if container.Id == id {
			LOCKED = false
			return container, nil
		}
	}
	LOCKED = false
	return ContainerInfo{}, nil
}

func WriteFile(data DATA, lock bool) error {
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
	LOCKED = lock
	return nil
}

func GetPort(id string, name string, gitUrl string) (int, error) {
	data, err := GetData()
	if err != nil {
		return 0, err
	}
	LOCKED = true
	port := data.Currentport + 1
	for {
		port++
		cmd := exec.Command("sudo", "lsof", "-i", "tcp:"+strconv.Itoa(port))
		err := cmd.Run()
		if err != nil {
			break
		}
	}
	data.List = append(data.List, ContainerInfo{
		Id:          id,
		ContainerID: "",
		Name:        name,
		Port:        port,
		Active:      false,
		GitUrl:      gitUrl,
	})
	data.Currentport = port
	WriteFile(data, false)
	return port, nil
}

func AddContainerId(containerId string, id string) error {
	data, err := GetData()
	if err != nil {
		return err
	}
	LOCKED = true
	if data.Active > data.MaxActive {
		return errors.New("have enough containers running")
	}
	for i, container := range data.List {
		if container.Id == id {
			container.ContainerID = containerId
			container.Active = true
			data.List[i] = container
			break
		}
	}
	data.Active++
	WriteFile(data, false)
	return nil
}

func SetUnActive(id string) error {
	data, err := GetData()
	if err != nil {
		return err
	}
	LOCKED = true
	for i, container := range data.List {
		if container.Id == id {
			if !container.Active {
				return nil
			}
			container.Active = true
			data.Active--
			data.List[i] = container
			break
		}
	}
	err = WriteFile(data, false)
	if err != nil {
		return err
	}
	return nil
}

func RemoveContainer(containerId string) error {
	data, err := GetData()
	if err != nil {
		return err
	}
	LOCKED = true
	data.Active--
	for i, container := range data.List {
		if container.ContainerID == containerId {
			container.Active = false
			data.List[i] = container
			break
		}
	}
	WriteFile(data, false)
	return nil
}
