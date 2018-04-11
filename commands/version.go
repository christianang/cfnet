package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of cfnet.",
	Long:  "Version of cfnet.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
