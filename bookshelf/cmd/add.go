package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "a namespace for adding books",
	Long:  `From here you can add books`,
}

func init() {
	RootCmd.AddCommand(addCmd)
}
