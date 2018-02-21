package cmd

import (
	"github.com/spf13/cobra"
)

type listCmd struct{}

func registerListCmds(cli *Cli) {
	cli.listCmd = &listCmd{}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "a namespace for listing books",
		Long:  `From here you can list books`,
	}

	listBookCmd := &cobra.Command{
		Use:   "book",
		Short: "list books",
		Long:  `List books from various title and author`,
		RunE:  cli.runner(listBookRun),
	}

	listCmd.AddCommand(listBookCmd)
	cli.rootCmd.AddCommand(listCmd)
}
