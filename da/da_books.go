package da

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/blackironj/bookchive-server/model"
)

func GetBookByID(db *sqlx.DB, id string) (*model.Book, error) {
	var book model.Book

	stmt := "SELECT * FROM books WHERE id = ?"
	err := db.Get(&book, stmt, id)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func InsertBooks(tx *sqlx.Tx, books []*model.Book) error {
	valStrings := []string{}
	valArgs := []interface{}{}
	for _, book := range books {
		valStrings = append(valStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")

		valArgs = append(valArgs, book.ID)
		valArgs = append(valArgs, book.Title)
		valArgs = append(valArgs, book.Subtitle)
		valArgs = append(valArgs, book.Authors)
		valArgs = append(valArgs, book.Publisher)
		valArgs = append(valArgs, book.Categories)
		valArgs = append(valArgs, book.Thumbnail)
		valArgs = append(valArgs, book.Pages)
	}

	stmt := `INSERT IGNORE INTO books(id, title, subtitle, authors, publisher, categories, thumbnail, pages) VALUES %s`
	stmt = fmt.Sprintf(stmt, strings.Join(valStrings, ","))

	_, err := tx.Exec(stmt, valArgs...)
	if err != nil {
		return err
	}

	return nil
}

func GetBooksInLibrary(db *sqlx.DB, userUUID string) ([]*model.BookInLibrary, error) {
	var books []*model.BookInLibrary

	stmt :=
		`SELECT 
			libraries.uk,
			libraries.book_id,
			libraries.added_dt,
			books.title, 
			books.authors, 
			books.categories, 
			books.thumbnail 
		FROM libraries
			LEFT JOIN books ON libraries.book_id = books.id
		WHERE libraries.user_uuid = ?`

	err := db.Select(&books, stmt, userUUID)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.New("Cannot find books in your libraries")
	}

	return books, nil
}
