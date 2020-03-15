package da

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/blackironj/bookchive-server/model"
)

func GetBookByID(db *sqlx.DB, id string) (*model.Books, error) {
	var book model.Books

	stmt := "SELECT * FROM books WHERE id = ?"
	err := db.Get(&book, stmt, id)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func InsertBooks(tx *sqlx.Tx, books []model.Books) error {
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
