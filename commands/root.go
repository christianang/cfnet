package commands

import "github.com/spf13/cobra"

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "cfnet",
	Short: "A cli to interact with container-to-container components in CF.",
	Long: `A cli to interact with container-to-container components
	in CloudFoundry.`,
}
