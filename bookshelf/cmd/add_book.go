package cmd

import (
	"fmt"

	bk "github.com/girikuncoro/go-playground/bookshelf/pkg"
)

func addBookRun(cli *Cli) error {
	db, err := cli.dbClient(cli.dbUser, cli.dbPassword, cli.dbHost, cli.dbPort)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	newBook := &bk.Book{
		Title:         cli.bookInfo.Title,
		Author:        cli.bookInfo.Author,
		PublishedDate: cli.bookInfo.PublishedDate,
	}
	if err := newBook.Validate(); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	db.AddBook(newBook)
	fmt.Printf("Successfully added new book: %v", newBook)
	return nil
}
