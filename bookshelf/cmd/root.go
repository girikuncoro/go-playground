package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "shelf",
	Short: "Shelf that contains books",
	Long:  "Shelf that contains books from various writers and titles",
	// Uncomment if bare app has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
