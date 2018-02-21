package cmd

import (
	"fmt"
	"log"
	"os"

	bk "github.com/girikuncoro/go-playground/bookshelf/pkg"
	"github.com/spf13/cobra"
)

var bookInfo bk.Book

var addBookCmd = &cobra.Command{
	Use:   "book",
	Short: "add a book",
	Long:  `Add a book, you can specify the title and author`,
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
		newBook := &bk.Book{
			Title:         bookInfo.Title,
			Author:        bookInfo.Author,
			PublishedDate: bookInfo.PublishedDate,
		}
		if err := newBook.Validate(); err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		db.AddBook(newBook)
		fmt.Printf("Successfully added new book: %v", newBook)
	},
}

func addBookRun(cli *Cli) error {
	db, err := cli.dbClient(cli.dbUser, cli.dbPassword, cli.dbHost, cli.dbPort)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	newBook := &bk.Book{
		Title:         bookInfo.Title,
		Author:        bookInfo.Author,
		PublishedDate: bookInfo.PublishedDate,
	}
	if err := newBook.Validate(); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	db.AddBook(newBook)
	fmt.Printf("Successfully added new book: %v", newBook)
	return nil
}
