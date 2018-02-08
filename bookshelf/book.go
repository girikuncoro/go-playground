package bookshelf

type Book struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate string
	CreatedBy     string
	CreatedByID   string
}

func (b *Book) CreatedByDisplayName() string {
	if b.CreatedBy == "anonymous" {
		return "Anonymous"
	}
	return b.CreatedBy
}

func (b *Book) SetCreatorAnonymous() {
	b.CreatedBy = ""
	b.CreatedBy = "anonymous"
}

type BookDatabase interface {
	ListBooks() ([]*Book, error)
	ListBooksCreatedBy(userID string) ([]*Book, error)
	GetBook(id int64) (*Book, error)
	AddBook(b *Book) (id int64, err error)
	DeleteBook(id int64) error
	UpdateBook(b *Book) error
	Close()
}
