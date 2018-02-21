package cmd

import (
	"fmt"
	"log"

	bk "github.com/girikuncoro/go-playground/bookshelf/pkg"
	"github.com/spf13/cobra"
)

var listBookCmd = &cobra.Command{
	Use:   "book",
	Short: "list books",
	Long:  `List books from various title and author`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bk.NewMySQLDB(bk.MySQLConfig{
			Username: "root",
			Password: "my-secret-pw",
			Host:     "0.0.0.0",
			Port:     9000,
		})
		if err != nil {
			log.Fatalln(err)
		}
		books, err := db.ListBooks()
		for _, book := range books {
			fmt.Printf("%#v", book)
		}
	},
}

func init() {
	listCmd.AddCommand(listBookCmd)
}
