package commands

import "github.com/spf13/cobra"

var RunCRoot string

var RootCmd = &cobra.Command{
	Use:   "cfnet",
	Short: "A cli to interact with container-to-container components in CF.",
	Long:  `A cli to interact with container-to-container networking components in CloudFoundry.`,
}

func init() {
	RootCmd.PersistentFlags().StringVar(&RunCRoot, "runc-root", "/var/run/runc", "runc root directory")
}
