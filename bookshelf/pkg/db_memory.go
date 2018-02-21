package bookshelf

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var _ BookDatabase = &memoryDB{}

type memoryDB struct {
	mu     sync.Mutex
	nextID int64
	books  map[int64]*Book
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		books:  make(map[int64]*Book),
		nextID: 1,
	}
}

func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.books = nil
}

func (db *memoryDB) GetBook(id int64) (*Book, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	book, ok := db.books[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: book not found with ID %d", id)
	}
	return book, nil
}

func (db *memoryDB) AddBook(b *Book) (id int64, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	b.ID = db.nextID
	db.books[b.ID] = b

	db.nextID++

	return b.ID, nil
}

func (db *memoryDB) DeleteBook(id int64) error {
	if id == 0 {
		return errors.New("memorydb: book with invalid ID is passed into deleteBook")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.books[id]; !ok {
		return fmt.Errorf("memorydb: could not delete with ID %d, doesn't exist", id)
	}
	delete(db.books, id)
	return nil
}

func (db *memoryDB) UpdateBook(b *Book) error {
	if b.ID == 0 {
		return errors.New("memorydb: book with invalid ID is passed into updateBook")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.books[b.ID] = b
	return nil
}

type booksByTitle []*Book

func (s booksByTitle) Less(i, j int) bool { return s[i].Title < s[j].Title }
func (s booksByTitle) Len() int           { return len(s) }
func (s booksByTitle) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (db *memoryDB) ListBooks() ([]*Book, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var books []*Book
	for _, b := range db.books {
		books = append(books, b)
	}

	sort.Sort(booksByTitle(books))
	return books, nil
}
