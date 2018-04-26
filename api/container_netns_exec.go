package api

import (
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
)

//go:generate counterfeiter -o ./fakes/ns_adapter.go --fake-name NSAdapter . nsAdapter
type nsAdapter interface {
	GetNS(nspath string) (ns.NetNS, error)
}

//go:generate counterfeiter -o ./fakes/runc.go --fake-name RunC . runc
type runc interface {
	GetNetNSPath(containerID string) (string, error)
}

//go:generate counterfeiter -o ./fakes/command_runner.go --fake-name CommandRunner . commandRunner
type commandRunner interface {
	Run(command string, commandArgs []string) error
}

type ContainerNetNSExec struct {
	ns            nsAdapter
	runc          runc
	commandRunner commandRunner
}

func NewContainerNetNSExec(runc runc, nsAdapter nsAdapter, commandRunner commandRunner) ContainerNetNSExec {
	return ContainerNetNSExec{
		ns:            nsAdapter,
		runc:          runc,
		commandRunner: commandRunner,
	}
}

func (c ContainerNetNSExec) Execute(containerID, command string, commandArgs []string) error {
	nspath, err := c.runc.GetNetNSPath(containerID)
	if err != nil {
		return fmt.Errorf("failed to get net ns path: %s", err)
	}

	netNs, err := c.ns.GetNS(nspath)
	if err != nil {
		return fmt.Errorf("failed to get net ns: %s", err)
	}

	err = netNs.Do(func(_ ns.NetNS) error {
		err := c.commandRunner.Run(command, commandArgs)
		if err != nil {
			return fmt.Errorf("failed to run command in net ns: %s", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
