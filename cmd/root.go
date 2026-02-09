package cmd

import (
	"github.com/spf13/cobra"
)

const (
	ROOT_PATH         = "/"
	ORGANIZATION_PATH = "/organization"

	ORG_FILE_PATH = "./employees.json"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "app",
		Short:   "Root command for the application",
		Version: version,
	}
	rootCmd.AddCommand(NewApiCmd(version))
	rootCmd.AddCommand(NewSyncCmd(version))
	return rootCmd
}
