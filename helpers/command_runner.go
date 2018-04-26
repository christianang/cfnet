package helpers

import (
	"io"
	"os/exec"
)

type CommandRunner struct {
	stdout io.Writer
	stderr io.Writer
}

func NewCommandRunner(stdout io.Writer, stderr io.Writer) CommandRunner {
	return CommandRunner{
		stdout: stdout,
		stderr: stderr,
	}
}

func (c CommandRunner) Run(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
