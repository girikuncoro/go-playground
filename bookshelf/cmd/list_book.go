package cmd

import (
	"fmt"
)

func listBookRun(cli *Cli) error {
	db, err := cli.dbClient(cli.dbUser, cli.dbPassword, cli.dbHost, cli.dbPort)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	books, err := db.ListBooks()
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	// Print out book sorted on Title
	for _, book := range books {
		fmt.Printf("%v\n", book)
	}
	return nil
}
