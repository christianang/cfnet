package commands

import (
	"fmt"
	"os"

	"github.com/christianang/cfnet/adapter"
	"github.com/christianang/cfnet/api"
	"github.com/christianang/cfnet/helpers"
	"github.com/christianang/cfnet/runc"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: `exec <container-id> <command>

  <container-id> is the runc container id of the instance.
  <command> is the command to be executed in the net namespace of the container.`,
	Short: "Execute a command within the net namespace of a runc container.",
	Long:  "Execute a command within the net namespace of a runc container.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		containerID := args[0]
		command := args[1]
		commandArgs := args[2:]

		runc := runc.NewRunC(RunCRoot)
		nsAdapter := adapter.NewNS()
		commandRunner := helpers.NewCommandRunner(os.Stdout, os.Stderr)
		exec := api.NewContainerNetNSExec(runc, nsAdapter, commandRunner)

		err := exec.Execute(containerID, command, commandArgs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	execCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(execCmd)
}
