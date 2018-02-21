package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "a namespace for listing books",
	Long:  `From here you can list books`,
}

func init() {
	RootCmd.AddCommand(listCmd)
}
