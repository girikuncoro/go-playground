package cmd

import (
	"github.com/spf13/cobra"
)

type addCmd struct{}

func registerAddCmds(cli *Cli) {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "a namespace for adding books",
		Long:  `From here you can add books`,
	}

	addBookCmd := &cobra.Command{
		Use:   "book",
		Short: "add a book",
		Long:  `Add a book, you can specify the title and author`,
		RunE:  cli.runner(addBookRun),
	}

	addCmd.AddCommand(addBookCmd)
	cli.rootCmd.AddCommand(addCmd)

	fls := addBookCmd.Flags()
	fls.StringVarP(&cli.bookInfo.Title, "title", "t", "", "title for the book")
	fls.StringVarP(&cli.bookInfo.Author, "author", "a", "", "author of the book")
	fls.StringVarP(&cli.bookInfo.PublishedDate, "date", "d", "", "published date of the book")
}
