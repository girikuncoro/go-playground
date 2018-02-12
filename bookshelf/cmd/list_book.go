package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listBookCmd = &cobra.Command{
	Use:   "book",
	Short: "list books",
	Long:  `List books from various title and author`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listing books")
	},
}

func init() {
	listCmd.AddCommand(listBookCmd)
}
