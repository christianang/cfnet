package runc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type RunC struct {
	runcRoot string
}

type State struct {
	NamespacePaths StateNamespacePaths `json:"namespace_paths"`
}

type StateNamespacePaths struct {
	NewNet string `json:"NEWNET"`
}

func NewRunC(runcRoot string) RunC {
	return RunC{
		runcRoot: runcRoot,
	}
}

func (r RunC) GetNetNSPath(containerID string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(r.runcRoot, containerID, "state.json"))
	if err != nil {
		return "", fmt.Errorf("failed to open runc state.json: %s", err)
	}

	var state State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %s", err)
	}

	return state.NamespacePaths.NewNet, nil
}
