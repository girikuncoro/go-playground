package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addBookCmd = &cobra.Command{
	Use:   "book",
	Short: "add a book",
	Long:  `Add a book, you can specify the title and author`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("book is added")
	},
}

func init() {
	addCmd.AddCommand(addBookCmd)
}
