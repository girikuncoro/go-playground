package bookshelf

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS library DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE library;`,
	`CREATE TABLE IF NOT EXISTS books(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NULL,
		author VARCHAR(255) NULL,
		publishedDate VARCHAR(255) NULL,
		createdBy VARCHAR(255) NULL,
		createdById VARCHAR(255) NULL,
		PRIMARY KEY (id)
	)`,
}

type mysqlDB struct {
	conn   *sql.DB
	list   *sql.Stmt
	insert *sql.Stmt
}

var _ BookDatabase = &mysqlDB{}

type MySQLConfig struct {
	Username, Password string
	Host               string
	Port               int
}

func newMySQLDB(config MySQLConfig) (BookDatabase, error) {
	if err := config.ensureTableExists(); err != nil {
		return nil, err
	}

	conn, err := sql.Open("mysql", config.datastoreName("library"))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get connection: %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish good connection: %v", err)
	}

	db := &mysqlDB{
		conn: conn,
	}

	// Prepare statements.
	if db.list, err = conn.Prepare(listStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare insert: %v", err)
	}
	return db, nil
}

func (c MySQLConfig) datastoreName(dbName string) string {
	var cred string
	if c.Username != "" {
		cred = c.Username
		if c.Password != "" {
			cred = cred + ":" + c.Password
		}
		cred = cred + "@"
	}
	return fmt.Sprintf("%stcp([%s]:%d)/%s", cred, c.Host, c.Port, dbName)
}

func (c MySQLConfig) ensureTableExists() error {
	conn, err := sql.Open("mysql", c.datastoreName(""))
	if err != nil {
		return fmt.Errorf("mysql: could not get connection")
	}
	defer conn.Close()

	if conn.Ping() == driver.ErrBadConn {
		return fmt.Errorf("mysql: could not connect to database")
	}

	if _, err := conn.Exec("USE library"); err != nil {
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return createTable(conn)
		}
	}

	if _, err := conn.Exec("DESCRIBE books"); err != nil {
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			return createTable(conn)
		}
		return fmt.Errorf("mysql: could not connect to database")
	}
	return nil
}

func createTable(conn *sql.DB) error {
	for _, stmt := range createTableStatements {
		_, err := conn.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *mysqlDB) Close() {
	db.conn.Close()
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func scanBook(s rowScanner) (*Book, error) {
	var (
		id            int64
		title         sql.NullString
		author        sql.NullString
		publishedDate sql.NullString
		createdBy     sql.NullString
		createdByID   sql.NullString
	)
	if err := s.Scan(&id, &title, &author, &publishedDate,
		&createdBy, &createdByID); err != nil {
		return nil, err
	}

	book := &Book{
		ID:            id,
		Title:         title.String,
		Author:        author.String,
		PublishedDate: publishedDate.String,
		CreatedBy:     createdBy.String,
		CreatedByID:   createdByID.String,
	}
	return book, nil
}

const listStatement = `SELECT * FROM books ORDER BY title`

func (db *mysqlDB) ListBooks() ([]*Book, error) {
	rows, err := db.list.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}
