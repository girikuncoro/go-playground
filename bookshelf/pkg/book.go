package bookshelf

import "fmt"

type Book struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate string
}

type BookDatabase interface {
	ListBooks() ([]*Book, error)
	AddBook(b *Book) (id int64, err error)
	Close()
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return fmt.Errorf("Title must be provided")
	}
	if b.Author == "" {
		return fmt.Errorf("Author must be provided")
	}
	if b.PublishedDate == "" {
		return fmt.Errorf("Published date must be provided")
	}
	return nil
}
