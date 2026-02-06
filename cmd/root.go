package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "app",
		Short:   "Root command for the application",
		Version: version,
	}
	rootCmd.AddCommand(NewApiCmd(version))
	return rootCmd
}
